#pragma once
#include <memory>
#include <string>
#include <vector>

#include "src/carnot/planner/compiler_state/compiler_state.h"
#include "src/carnot/planner/objects/funcobject.h"
#include "src/common/uuid/uuid.h"
#include "src/shared/metadata/base_types.h"
#include "src/stirling/dynamic_tracing/ir/logical.pb.h"

namespace pl {
namespace carnot {
namespace planner {
namespace compiler {

class ProbeOutput {
 public:
  ProbeOutput() = delete;
  ProbeOutput(const std::string& output_name, const std::vector<std::string>& col_names,
              const std::vector<std::string>& var_names)
      : output_name_(output_name), col_names_(col_names), var_names_(var_names) {}
  /**
   * @brief Returns the schema as a string. Just joins column_names together. If two outputActions
   * write to the same columns w/ the same schemas, they can still disagree on type which will be
   * piped up in a later Stirling IR processing stage.
   *
   * @return std::string_view the string representation of the schema.
   */
  // std::string_view Schema() const { return absl::StrJoin(col_names_, ","); }

  Status ToActionProto(stirling::dynamic_tracing::ir::logical::OutputAction* pb);
  Status ToOutputProto(stirling::dynamic_tracing::ir::shared::Output* pb);
  const std::string& name() const { return output_name_; }

 private:
  // Output name is the name of the output to write.
  std::string output_name_;
  // Column names are the names to write to.
  std::vector<std::string> col_names_;
  // Var names are the ids of the variables to write to those columns.
  std::vector<std::string> var_names_;
};

class ProbeIR {
 public:
  ProbeIR(stirling::dynamic_tracing::ir::shared::BinarySpec::Language language,
          const std::string& function_name)
      : language_(language), symbol_(function_name) {}

  /**
   * @brief Serializes this probe definition as a protobuf.
   *
   * @param pb
   * @return Any error that occured during serialization.
   */
  Status ToProto(stirling::dynamic_tracing::ir::logical::Probe* pb);

  /**
   * @return true a latency column has been defined for the probe
   * @return false otherwise
   */
  bool HasLatencyCol() { return latency_col_id_.size() > 0; }

  /**
   * @brief Set the ID of a function latency object. There can only be one, so this is structured
   * differently than arguments and return values.
   *
   * @param latency_col_id
   */
  void SetFunctionLatencyID(const std::string& latency_col_id) { latency_col_id_ = latency_col_id; }

  /**
   * @brief Add Argument expression to the probe.
   *
   * @param id the id of the argument expression.
   * @param expr the expression to parse for the argument.
   */
  void AddArgument(const std::string& id, const std::string& expr);

  /**
   * @brief Add Return expression to the probe.
   *
   * @param id the id of the argument expression.
   * @param expr the expression to parse for the argument.
   */
  void AddReturnValue(const std::string& id, const std::string& expr);

  /**
   * @brief Adds a latency function to this probe.
   *
   * @param id the id of the returned value.
   */
  void AddLatencyFunc(const std::string& id);

  /**
   * @brief Convenience function that returns the next string to set as an id of an argument. In the
   * future we might name these.
   *
   * @return std::string
   */
  std::string NextArgName() { return absl::Substitute("arg$0", args_.size()); }

  /**
   * @brief Convenience function that returns the next string to set as an id of return value. In
   * the future we might name these.
   *
   * @return std::string
   */
  std::string NextReturnName() { return absl::Substitute("ret$0", ret_vals_.size()); }

  /**
   * @brief Convenience function that returns the next string to set as an id of a latency function
   * call. We only support 1 latency call for an entire function so we always return the same value.
   *
   * @return std::string
   */
  std::string NextLatencyName() { return "lat0"; }

  /**
   * @brief Create a New Output definition with the given col_names and var_names. For now we only
   * support one output.
   *
   * @param output_name
   * @param col_names
   * @param var_names
   */
  void CreateNewOutput(const std::string& output_name, const std::vector<std::string>& col_names,
                       const std::vector<std::string>& var_names);
  std::shared_ptr<ProbeOutput> output() const { return output_; }
  stirling::dynamic_tracing::ir::shared::BinarySpec::Language language() const { return language_; }

 private:
  stirling::dynamic_tracing::ir::shared::BinarySpec::Language language_;
  std::string symbol_;
  std::string latency_col_id_;
  std::vector<stirling::dynamic_tracing::ir::logical::Argument> args_;
  std::vector<stirling::dynamic_tracing::ir::logical::ReturnValue> ret_vals_;
  std::shared_ptr<ProbeOutput> output_ = nullptr;
};

class TracingProgram {
 public:
  /**
   * @brief Converts Program to the proto representation.
   *
   * @param pb
   * @return Status errors if they happen.
   */
  Status ToProto(stirling::dynamic_tracing::ir::logical::Program* pb) const;

  /**
   * @brief Add a Probe to the current program being traced.
   *
   * @param name
   * @param probe_ir
   * @param ttl_ns
   * @return Status
   */
  Status AddProbe(const std::string& name, ProbeIR* probe_ir, int64_t ttl_ns);

 private:
  std::string binary_path_;
  std::vector<stirling::dynamic_tracing::ir::logical::Probe> probes_;
  std::vector<stirling::dynamic_tracing::ir::shared::Output> outputs_;
  absl::flat_hash_map<std::string, stirling::dynamic_tracing::ir::shared::Output*> output_map_;
  stirling::dynamic_tracing::ir::shared::BinarySpec::Language language_;
};

class DynamicTraceIR {
 public:
  /**
   * @brief Creates a new probe definition and stores it in the current_probe() of the Builder.
   *
   * @param function_name
   * @return std::shared_ptr<ProbeIR>
   */
  std::shared_ptr<ProbeIR> StartProbe(
      stirling::dynamic_tracing::ir::shared::BinarySpec::Language language,
      const std::string& function_name);

  /**
   * @brief Upserts a probe definition into a program, as defined by the pod_name and binary_path.
   *
   * @param probe_ir the probe to upsert
   * @param probe_name the name ot give the probe
   * @param pod_name the name of the pod to trace on
   * @param binary_path the path of the binary in the pod?
   * @return Status
   */
  Status UpsertProbe(std::shared_ptr<ProbeIR> probe_ir, const std::string& probe_name,
                     const std::string& pod_name, const std::string& container_name,
                     const std::string& binary_path, int64_t ttl_ns);

  /**
   * @brief Upserts a probe definition into a program defined by the UPID>
   *
   * @param probe_ir the probe to upsert
   * @param probe_name the name ot give the probe
   * @param upid the upid definition
   * @param ttl_ns the duration of the live probe
   * @return Status
   */
  Status UpsertUPIDProbe(std::shared_ptr<ProbeIR> probe_ir, const std::string& probe_name,
                         const md::UPID& upid, int64_t ttl_ns);

  /**
   * @brief Get the CurrentProbe or return an error. Nice shorthand to support a clean error
   * message that points to the write position in the code.
   *
   * @param ast the ast object to use as a pointer to the line of code.
   * @return StatusOr<ProbeIR*> the current probe or an error if it doesn't exist.
   */
  StatusOr<ProbeIR*> GetCurrentProbeOrError(const pypa::AstPtr& ast);

  /**
   * @brief Converts the Probe definitions into the proto definition that can be used by the
   * Stirling probe tracer. Note: If there are multiple binaries defined then this will error out
   * for now. In the future we might support multiple binaries, but for now it will error out.
   *
   * @param pb the protobuf object to write to.
   * @return Status
   */
  Status ToProto(stirling::dynamic_tracing::ir::logical::Program* pb);

  /**
   * @brief Stops recording changes to the current_probe_ and removes it from the current_probe_
   * position.
   */
  void EndProbe();

  ProbeIR* current_probe() { return current_probe_.get(); }

 private:
  absl::flat_hash_map<std::string, TracingProgram> binary_to_program_map_;
  absl::flat_hash_map<md::UPID, TracingProgram> upid_to_program_map_;
  std::vector<std::shared_ptr<ProbeIR>> probes_pool_;
  std::shared_ptr<ProbeIR> current_probe_;
};

}  // namespace compiler
}  // namespace planner
}  // namespace carnot
}  // namespace pl