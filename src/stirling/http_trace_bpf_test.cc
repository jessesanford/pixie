#include <gmock/gmock.h>
#include <gtest/gtest.h>

#include <filesystem>

#include "src/common/base/test_utils.h"
#include "src/common/exec/subprocess.h"
#include "src/stirling/data_table.h"
#include "src/stirling/socket_trace_connector.h"
#include "src/stirling/testing/common.h"
#include "src/stirling/testing/socket_trace_bpf_test_fixture.h"

DEFINE_string(go_grpc_client_path, "", "The path to the go greeter client executable.");
DEFINE_string(go_grpc_server_path, "", "The path to the go greeter server executable.");

namespace pl {
namespace stirling {

using ::pl::stirling::testing::ColWrapperSizeIs;
using ::pl::types::ColumnWrapperRecordBatch;
using ::testing::AllOf;
using ::testing::AnyOf;
using ::testing::ContainsRegex;
using ::testing::Each;
using ::testing::ElementsAre;
using ::testing::HasSubstr;
using ::testing::IsEmpty;
using ::testing::MatchesRegex;
using ::testing::SizeIs;
using ::testing::StrEq;

class GoHTTPTraceTest : public testing::SocketTraceBPFTest {
 protected:
  GoHTTPTraceTest() : SocketTraceBPFTest() {}

  void SetUp() override {
    SocketTraceBPFTest::SetUp();

    CHECK(!FLAGS_go_grpc_client_path.empty())
        << "--go_grpc_client_path cannot be empty. You should run this test with bazel.";
    CHECK(std::filesystem::exists(std::filesystem::path(FLAGS_go_grpc_client_path)))
        << FLAGS_go_grpc_client_path;

    CHECK(!FLAGS_go_grpc_server_path.empty())
        << "--go_grpc_server_path cannot be empty. You should run this test with bazel.";
    CHECK(std::filesystem::exists(std::filesystem::path(FLAGS_go_grpc_server_path)))
        << FLAGS_go_grpc_server_path;

    server_path_ = FLAGS_go_grpc_server_path;
    client_path_ = FLAGS_go_grpc_client_path;

    ASSERT_OK(s_.Start({server_path_}));

    // Give some time for the server to start up.
    sleep(2);

    const std::string port_str = s_.Stdout();
    ASSERT_TRUE(absl::SimpleAtoi(port_str, &s_port_));
    ASSERT_NE(0, s_port_);
  }

  void TearDown() override {
    SocketTraceBPFTest::TearDown();

    s_.Kill();
    EXPECT_EQ(9, s_.Wait()) << "Server should have been killed.";
  }

  std::string server_path_;
  std::string client_path_;

  // Create a context to pass into each TransferData() in the test, using a dummy ASID.
  static constexpr uint32_t kASID = 1;

  DataTable data_table_{kHTTPTable};
  SubProcess c_;
  SubProcess s_;
  int s_port_ = -1;
};

TEST_F(GoHTTPTraceTest, RequestAndResponse) {
  ASSERT_OK(
      c_.Start({client_path_, "-name=PixieLabs", absl::StrCat("-address=localhost:", s_port_)}));
  EXPECT_EQ(0, c_.Wait()) << "Client should exit normally.";

  source_->TransferData(ctx_.get(), kHTTPTableNum, &data_table_);

  const types::ColumnWrapperRecordBatch& record_batch = *data_table_.ActiveRecordBatch();
  const std::vector<size_t> target_record_indices =
      testing::FindRecordIdxMatchesPid(record_batch, kHTTPUPIDIdx, c_.child_pid());
  ASSERT_THAT(target_record_indices, SizeIs(1));

  const size_t target_record_idx = target_record_indices.front();

  EXPECT_THAT(
      std::string(record_batch[kHTTPReqHeadersIdx]->Get<types::StringValue>(target_record_idx)),
      AllOf(HasSubstr(R"("Accept-Encoding":"gzip")"),
            HasSubstr(absl::Substitute(R"(Host":"localhost:$0")", s_port_)),
            ContainsRegex(R"(User-Agent":"Go-http-client/.+")")));
  EXPECT_THAT(
      std::string(record_batch[kHTTPRespHeadersIdx]->Get<types::StringValue>(target_record_idx)),
      AllOf(HasSubstr(R"("Content-Length":"31")"), HasSubstr(R"(Content-Type":"json)")));
  EXPECT_THAT(
      std::string(record_batch[kHTTPRemoteAddrIdx]->Get<types::StringValue>(target_record_idx)),
      HasSubstr("127.0.0.1"));
  EXPECT_EQ(s_port_,
            record_batch[kHTTPRemotePortIdx]->Get<types::Int64Value>(target_record_idx).val);
  EXPECT_THAT(record_batch[kHTTPRespBodyIdx]->Get<types::StringValue>(target_record_idx),
              StrEq(absl::StrCat(R"({"greeter":"Hello PixieLabs!"})", "\n")));
}

struct TraceRoleTestParam {
  EndpointRole role;
  size_t client_records_count;
  size_t server_records_count;
};

class TraceRoleTest : public GoHTTPTraceTest,
                      public ::testing::WithParamInterface<TraceRoleTestParam> {};

TEST_P(TraceRoleTest, VerifyRecordsCount) {
  const TraceRoleTestParam& param = GetParam();

  auto* socket_trace_connector = static_cast<SocketTraceConnector*>(source_.get());
  ASSERT_NE(nullptr, socket_trace_connector);
  EXPECT_OK(socket_trace_connector->UpdateProtocolTraceRole(kProtocolHTTP, param.role));

  ASSERT_OK(
      c_.Start({client_path_, "-name=PixieLabs", absl::StrCat("-address=localhost:", s_port_)}));
  EXPECT_EQ(0, c_.Wait()) << "Client should exit normally.";

  source_->TransferData(ctx_.get(), kHTTPTableNum, &data_table_);

  const types::ColumnWrapperRecordBatch& record_batch = *data_table_.ActiveRecordBatch();
  const std::vector<size_t> client_record_ids =
      testing::FindRecordIdxMatchesPid(record_batch, kHTTPUPIDIdx, c_.child_pid());
  EXPECT_THAT(client_record_ids, SizeIs(param.client_records_count));
  const std::vector<size_t> server_record_ids =
      testing::FindRecordIdxMatchesPid(record_batch, kHTTPUPIDIdx, s_.child_pid());
  EXPECT_THAT(server_record_ids, SizeIs(param.server_records_count));
}

INSTANTIATE_TEST_SUITE_P(AllTraceRoles, TraceRoleTest,
                         ::testing::Values(TraceRoleTestParam{kRoleUnknown, 0, 0},
                                           TraceRoleTestParam{kRoleClient, 1, 0},
                                           TraceRoleTestParam{kRoleServer, 0, 1},
                                           TraceRoleTestParam{kRoleAll, 1, 1}));

}  // namespace stirling
}  // namespace pl
