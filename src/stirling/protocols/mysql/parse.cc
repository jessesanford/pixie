#include "src/stirling/protocols/mysql/parse.h"

#include <arpa/inet.h>
#include <deque>
#include <string_view>
#include <utility>

#include "src/common/base/byte_utils.h"
#include "src/stirling/common/parse_state.h"
#include "src/stirling/protocols/mysql/types.h"

namespace pl {
namespace stirling {
namespace protocols {
namespace mysql {

ParseState ParseFrame(MessageType type, std::string_view* buf, Packet* result) {
  if (type != MessageType::kRequest && type != MessageType::kResponse) {
    return ParseState::kInvalid;
  }

  if (buf->size() < kPacketHeaderLength) {
    return ParseState::kNeedsMoreData;
  }

  result->sequence_id = static_cast<uint8_t>((*buf)[3]);

  int packet_length = utils::LEndianBytesToInt<int, kPayloadLengthLength>(*buf);

  // TODO(oazizi): Is pre-checking requests here a good idea? Somewhat out of place.
  // Better fit for stitcher (when analyzing structure of packet bodies).
  if (type == MessageType::kRequest) {
    if (buf->size() < kPacketHeaderLength + 1) {
      return ParseState::kInvalid;
    }
    uint8_t command = (*buf)[kPacketHeaderLength];
    if (!IsValidCommand(command)) {
      return ParseState::kInvalid;
    }

    // We can constrain the expected lengths, by command type.
    NumberRange length_range = mysql::kMySQLCommandLengths[command];
    if (packet_length < length_range.min || packet_length > length_range.max) {
      return ParseState::kInvalid;
    }
  }

  ssize_t buffer_length = buf->length();
  if (buffer_length < kPacketHeaderLength + packet_length) {
    return ParseState::kNeedsMoreData;
  }

  result->msg = buf->substr(kPacketHeaderLength, packet_length);
  buf->remove_prefix(kPacketHeaderLength + packet_length);

  return ParseState::kSuccess;
}

size_t FindFrameBoundary(MessageType type, std::string_view buf, size_t start_pos) {
  if (buf.length() < mysql::kPacketHeaderLength) {
    return std::string::npos;
  }

  if (type == MessageType::kResponse) {
    // No real search implemented for responses.
    // TODO(oazizi): Is there something we can implement here?
    return std::string::npos;
  }

  // Need at least kPacketHeaderLength bytes + 1 command byte in buf.
  for (size_t i = start_pos; i < buf.size() - mysql::kPacketHeaderLength; ++i) {
    std::string_view cur_buf = buf.substr(i);
    int packet_length = utils::LEndianBytesToInt<int, mysql::kPayloadLengthLength>(cur_buf);
    uint8_t sequence_id = static_cast<uint8_t>(cur_buf[3]);
    auto command_byte = magic_enum::enum_cast<mysql::Command>(cur_buf[mysql::kPacketHeaderLength]);

    // Requests must have sequence id of 0.
    if (sequence_id != 0) {
      continue;
    }

    // If the command byte doesn't decode to a valid command, then this can't a message boundary.
    if (!command_byte.has_value()) {
      continue;
    }

    // We can constrain the expected lengths, by command type.
    auto length_range = mysql::kMySQLCommandLengths[static_cast<int>(command_byte.value())];
    if (packet_length < length_range.min || packet_length > length_range.max) {
      continue;
    }

    return i;
  }

  return std::string::npos;
}

}  // namespace mysql

template <>
ParseState ParseFrame(MessageType type, std::string_view* buf, mysql::Packet* result) {
  return mysql::ParseFrame(type, buf, result);
}

template <>
size_t FindFrameBoundary<mysql::Packet>(MessageType type, std::string_view buf, size_t start_pos) {
  return mysql::FindFrameBoundary(type, buf, start_pos);
}

}  // namespace protocols
}  // namespace stirling
}  // namespace pl