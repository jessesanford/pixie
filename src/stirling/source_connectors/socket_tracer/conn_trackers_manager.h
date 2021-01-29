#pragma once

#include <list>
#include <map>
#include <memory>
#include <set>

#include "src/stirling/source_connectors/socket_tracer/connection_tracker.h"

namespace pl {
namespace stirling {

/**
 * ConnTrackersManager is a container that keeps track of all ConnectionTrackers.
 * Interface designed for two primary operations:
 *  1) Insertion of events indexed by conn_id (PID+FD+TSID) as they arrive from BPF.
 *  2) Iteration through trackers by protocols.
 */
class ConnTrackersManager {
 public:
  /**
   * Get a connection tracker for the specified conn_id. If a tracker does not exist,
   * one will be created and returned.
   */
  ConnectionTracker& GetOrCreateConnTracker(struct conn_id_t conn_id);

  /**
   * A TrackersList consists of a list of trackers.
   * It can only be created via ConnTrackersForProtocol(), such that it returns a list
   * of trackers that have the requested protocol.
   *
   * Usage model example:
   * ConnTrackersManager::TrackersList http_conn_trackers =
   *     conn_trackers_mgr.ConnTrackersForProtocol(kProtocolHTTP);
   *
   * for (auto iter = http_conn_trackers.begin(); iter != http_conn_trackers.end(); ++iter) {
   *   ConnectionTracker* tracker = *iter;
   *
   *   // Relevant actions on tracker go here.
   * }
   */
  class TrackersList {
   public:
    /**
     * A custom iterator for going through the list of trackers for a given protocol.
     * This iterator automatically handles removing trackers whose protocol has changed
     * (currently this should only be possible from kProtocolUnknown), and the removal of
     * trackers that are ReadyForDestruction().
     */
    class TrackersListIterator {
     public:
      bool operator!=(const TrackersListIterator& other) { return other.iter_ != this->iter_; }

      ConnectionTracker* operator*() { return *iter_; }

      // Prefix increment operator.
      TrackersListIterator operator++();

     private:
      TrackersListIterator(std::list<ConnectionTracker*>* trackers,
                           std::list<ConnectionTracker*>::iterator iter, TrafficProtocol protocol,
                           ConnTrackersManager* conn_trackers_manager);

      void AdvanceToValidTracker();

      std::list<ConnectionTracker*>* trackers_;
      std::list<ConnectionTracker*>::iterator iter_;
      TrafficProtocol protocol_;
      ConnTrackersManager* conn_trackers_manager_;

      friend class TrackersList;
    };

    TrackersListIterator begin() {
      return TrackersListIterator(list_, list_->begin(), protocol_, conn_trackers_);
    }

    TrackersListIterator end() {
      return TrackersListIterator(list_, list_->end(), protocol_, conn_trackers_);
    }

   private:
    TrackersList(std::list<ConnectionTracker*>* list, TrafficProtocol protocol,
                 ConnTrackersManager* conn_trackers)
        : list_(list), protocol_(protocol), conn_trackers_(conn_trackers) {}

    std::list<ConnectionTracker*>* list_;
    TrafficProtocol protocol_;
    ConnTrackersManager* conn_trackers_;

    friend class ConnTrackersManager;
  };

  /**
   * Returns a list of all the trackers that belong to a particular protocol.
   */
  TrackersList ConnTrackersForProtocol(TrafficProtocol protocol) {
    return TrackersList(&conn_trackers_by_protocol_[protocol], protocol, this);
  }

  /**
   * Returns the latest generation of a connection tracker for the given pid and fd.
   * If there is no tracker for {pid, fd}, returns error::NotFound.
   */
  StatusOr<const ConnectionTracker*> GetConnectionTracker(uint32_t pid, uint32_t fd) const;

  /**
   * If a connection tracker has its protocol changed, then one must manually call this function.
   * TODO(oazizi): Find a cleaner/more automatic way that can avoid this call altogether.
   */
  void NotifyProtocolChange(ConnectionTracker* tracker);

  /**
   * Deletes trackers that are ReadyForDestruction().
   * Call this only after accumulating enough trackers to clean-up, to avoid the performance
   * impact of scanning through all trackers every iteration.
   */
  void CleanupTrackers();

  /**
   * Runs checks for checking the consistency of the data structure.
   * Useful for catching bugs.
   */
  void CheckConsistency();

  /**
   * Debug utility to dump information about connection trackers.
   */
  void DebugInfo() const;

 private:
  // A map from conn_id (PID+FD+TSID) to tracker. This is for easy update on BPF events.
  // Structured as two nested maps to be explicit about "generations" of trackers per PID+FD.
  // Key is {PID, FD} for outer map, and tsid for inner map.
  // Note that the inner map cannot be a vector, because there is no guaranteed order
  // in which events are read from perf buffers.
  // Inner map could be a priority_queue, but benchmarks showed better performance with a std::map.
  absl::flat_hash_map<uint64_t, std::map<uint64_t, ConnectionTracker> > connection_trackers_;

  // A set of lists of pointers to all the contained trackers, organized by protocol
  // This is for easy access to the trackers during TransferData().
  // Key is protocol.
  // TODO(jps): Convert to vector?
  absl::flat_hash_map<TrafficProtocol, std::list<ConnectionTracker*> > conn_trackers_by_protocol_;

  // Keep track of total number of trackers, and other counts.
  // Used to check for consistency.
  size_t num_trackers_ = 0;
  size_t num_trackers_ready_for_destruction_ = 0;
  size_t num_trackers_in_lists_ = 0;
  size_t num_tracker_dups_ = 0;
};

}  // namespace stirling
}  // namespace pl