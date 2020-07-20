#include "src/stirling/dynamic_trace_connector.h"

#include "src/shared/types/proto/types.pb.h"
#include "src/stirling/dynamic_tracing/dynamic_tracer.h"
#include "src/stirling/utils/linux_headers.h"

namespace pl {
namespace stirling {

using ::pl::stirling::dynamic_tracing::ir::physical::Struct;
using ::pl::stirling::dynamic_tracing::ir::shared::ScalarType;

namespace {

// A generic callback function to be invoked to push a piece of data polled from the perf buffer
// to the DynamicTraceConnector.
void GenericHandleEvent(void* cb_cookie, void* data, int data_size) {
  DCHECK_NE(cb_cookie, nullptr);
  DCHECK_EQ(data_size % 4, 0)
      << "Perf buffer data items are aligned with 8 bytes. "
         "The first 4 bytes are size, therefore data size must be a multiple of 4.";

  auto* parser = static_cast<DynamicTraceConnector*>(cb_cookie);
  std::string buf(static_cast<const char*>(data), data_size);

  parser->AcceptDataEvents(std::move(buf));
}

// A generic callback function to be invoked to process data item loss.
// The input cb_cookie has to be DynamicTraceConnector*.
void GenericHandleEventLoss(void* cb_cookie, uint64_t lost) {
  DCHECK_NE(cb_cookie, nullptr);
  VLOG(1) << absl::Substitute("Lost $0 events", lost);
}

}  // namespace

Status DynamicTraceConnector::InitImpl() {
  PL_RETURN_IF_ERROR(utils::FindOrInstallLinuxHeaders(utils::kDefaultHeaderSearchOrder));

  PL_RETURN_IF_ERROR(InitBPFProgram(bcc_program_.code));

  for (const auto& uprobe : bcc_program_.uprobes) {
    PL_RETURN_IF_ERROR(AttachUProbe(uprobe));
  }

  // TODO(yzhao/oazizi): Might need to change this if we need to support multiple perf buffers.
  bpf_tools::PerfBufferSpec spec = {
      .name = bcc_program_.perf_buffer_specs.front().name,
      .probe_output_fn = &GenericHandleEvent,
      .probe_loss_fn = &GenericHandleEventLoss,
  };

  PL_RETURN_IF_ERROR(OpenPerfBuffer(spec, this));

  return Status::OK();
}

namespace {

// Reads a byte sequence representing a packed C/C++ struct, and extract the values of the fields.
class StructDecoder {
 public:
  explicit StructDecoder(std::string_view buf) : buf_(buf) {}

  template <typename NativeScalarType>
  StatusOr<NativeScalarType> ExtractField() {
    if (buf_.size() < sizeof(NativeScalarType)) {
      return error::ResourceUnavailable("Insufficient number of bytes.");
    }
    NativeScalarType val = {};
    std::memcpy(&val, buf_.data(), sizeof(NativeScalarType));
    buf_.remove_prefix(sizeof(NativeScalarType));
    return val;
  }

 private:
  std::string_view buf_;
};

Status FillColumn(StructDecoder* struct_decoder, DataTable::DynamicRecordBuilder* r, size_t col_idx,
                  ScalarType type) {
#define WRITE_COLUMN(field_type, column_type)                                      \
  PL_ASSIGN_OR_RETURN(field_type val, struct_decoder->ExtractField<field_type>()); \
  r->Append(col_idx++, column_type(val));

  // TODO(yzhao): Right now only support scalar types. We should replace type with ScalarType
  // in Struct::Field.
  switch (type) {
    case ScalarType::BOOL: {
      WRITE_COLUMN(bool, types::BoolValue);
      break;
    }
    case ScalarType::INT: {
      WRITE_COLUMN(int, types::Int64Value);
      break;
    }
    case ScalarType::INT8: {
      WRITE_COLUMN(int8_t, types::Int64Value);
      break;
    }
    case ScalarType::INT16: {
      WRITE_COLUMN(int16_t, types::Int64Value);
      break;
    }
    case ScalarType::INT32: {
      WRITE_COLUMN(int32_t, types::Int64Value);
      break;
    }
    case ScalarType::INT64: {
      WRITE_COLUMN(int64_t, types::Int64Value);
      break;
    }
    case ScalarType::UINT: {
      WRITE_COLUMN(unsigned int, types::Int64Value);
      break;
    }
    case ScalarType::UINT8: {
      WRITE_COLUMN(uint8_t, types::Int64Value);
      break;
    }
    case ScalarType::UINT16: {
      WRITE_COLUMN(uint16_t, types::Int64Value);
      break;
    }
    case ScalarType::UINT32: {
      WRITE_COLUMN(uint32_t, types::Int64Value);
      break;
    }
    case ScalarType::UINT64: {
      WRITE_COLUMN(uint64_t, types::Int64Value);
      break;
    }
    case ScalarType::FLOAT: {
      WRITE_COLUMN(float, types::Float64Value);
      break;
    }
    case ScalarType::DOUBLE: {
      WRITE_COLUMN(double, types::Float64Value);
      break;
    }
    case ScalarType::VOID_POINTER: {
      WRITE_COLUMN(uint64_t, types::Int64Value);
      break;
    }
    case ScalarType::ScalarType_INT_MIN_SENTINEL_DO_NOT_USE_:
    case ScalarType::ScalarType_INT_MAX_SENTINEL_DO_NOT_USE_:
      LOG(DFATAL) << "Impossible enum value";
      break;
  }
#undef WRITE_COLUMN

  return Status::OK();
}

Status AppendRecord(const Struct& st, uint32_t asid, std::string_view buf, DataTable* data_table) {
  StructDecoder struct_decoder(buf);
  DataTable::DynamicRecordBuilder r(data_table);

  // TODO(yzhao): Come up more principled approach to process upid, such that explicit checks
  // can be applied to avoid these fields being misused.
  PL_ASSIGN_OR_RETURN(uint32_t tgid, struct_decoder.ExtractField<uint32_t>());
  PL_ASSIGN_OR_RETURN(uint64_t tgid_start_time, struct_decoder.ExtractField<uint64_t>());

  md::UPID upid(asid, tgid, tgid_start_time);
  r.Append(0, types::UInt128Value(upid.value()));

  // Skip the first 2 fields which are tgid & tgid_start_time, which are combined into upid.
  for (int i = 2, col_idx = 1; i < st.fields_size(); ++i, ++col_idx) {
    PL_RETURN_IF_ERROR(FillColumn(&struct_decoder, &r, col_idx, st.fields(i).type().scalar()));
  }

  return Status::OK();
}

}  // namespace

void DynamicTraceConnector::TransferDataImpl(ConnectorContext* ctx, uint32_t table_num,
                                             DataTable* data_table) {
  DCHECK_EQ(table_num, 0) << "Now only support having exactly one table per DynamicTraceConnector";

  PollPerfBuffers();

  for (const auto& item : data_items_) {
    // TODO(yzhao): Right now only support scalar types. We should replace type with ScalarType
    // in Struct::Field.
    ECHECK_OK(AppendRecord(bcc_program_.perf_buffer_specs.front().output, ctx->GetASID(), item,
                           data_table));
  }

  data_items_.clear();
}

}  // namespace stirling
}  // namespace pl