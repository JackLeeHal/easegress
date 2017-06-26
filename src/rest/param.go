package rest

import (
	"time"

	"config"
)

//
// Admin API
//

type pluginCreationRequest struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}

type pluginsRetrieveRequest struct {
	NamePattern string   `json:"name_pattern,omitempty"`
	Types       []string `json:"types"`
}

type pluginsRetrieveResponse struct {
	Plugins []config.PluginSpec `json:"plugins"`
}

type pluginUpdateRequest struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}

////

type pipelineCreationRequest struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}

type pipelinesRetrieveRequest struct {
	NamePattern string   `json:"name_pattern,omitempty"`
	Types       []string `json:"types"`
}

type pipelinesRetrieveResponse struct {
	Pipelines []config.PipelineSpec `json:"pipelines"`
}

type pipelineUpdateRequest struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}

////

type pluginTypesRetrieveResponse struct {
	PluginTypes []string `json:"plugin_types"`
}

type pipelineTypesRetrieveResponse struct {
	PipelineTypes []string `json:"pipeline_types"`
}

//
// Statistics API
//

type indicatorNamesRetrieveResponse struct {
	Names []string `json:"names"`
}

type indicatorValueRetrieveResponse struct {
	Value interface{} `json:"value"`
}

type indicatorDescriptionRetrieveResponse struct {
	Description interface{} `json:"desc"`
}

////

type gatewayUpTimeRetrieveResponse struct {
	UpTime time.Duration `json:"up_nanosec"`
}

//
// Cluster API
//

type clusterRequest struct {
	TimeoutSec uint16 `json:"timeout_sec"` // 10-65535, zero means using default value (30)
}

//
// Cluster admin API
//

type clusterOperationSeqRequest struct {
	clusterRequest
}

type clusterOperationSeqResponse struct {
	Group             string `json:"cluster_group"`
	OperationSequence uint64 `json:"operation_seq"`
}

type clusterOperation struct {
	clusterRequest
	OperationSeqSnapshot uint64 `json:"operation_seq_snapshot"`
	Consistent           bool   `json:"consistent"`
}

type pluginCreationClusterRequest struct {
	clusterOperation
	pluginCreationRequest
}

type pluginsRetrieveClusterRequest struct {
	clusterRequest
	pluginsRetrieveRequest
	Consistent bool `json:"consistent"`
}

type pluginRetrieveClusterRequest struct {
	clusterRequest
	Consistent bool `json:"consistent"`
}

type pluginUpdateClusterRequest struct {
	clusterOperation
	pluginUpdateRequest
}

type pluginDeletionClusterRequest struct {
	clusterOperation
}

////

type pipelineCreationClusterRequest struct {
	clusterOperation
	pipelineCreationRequest
}

type pipelinesRetrieveClusterRequest struct {
	clusterRequest
	pipelinesRetrieveRequest
	Consistent bool `json:"consistent"`
}

type pipelineRetrieveClusterRequest struct {
	clusterRequest
	Consistent bool `json:"consistent"`
}

type pipelineUpdateClusterRequest struct {
	clusterOperation
	pipelineUpdateRequest
}

type pipelineDeletionClusterRequest struct {
	clusterOperation
}

////

type pluginTypesRetrieveClusterRequest struct {
	clusterRequest
	Consistent bool `json:"consistent"`
}

type pipelineTypesRetrieveClusterRequest struct {
	clusterRequest
	Consistent bool `json:"consistent"`
}

//
// Cluster statistics API
//

type statisticsClusterRequest struct {
	clusterRequest
	// TODO: aggregateStatResponses() returns result of each member as the part of aggregation result
	// if the flag is turned on
	Details bool `json:"details"`
}

//
// Health check API
//

type clusterInfo struct {
	Name                  string   `json:"node_name"`
	Mode                  string   `json:"node_mode"`
	Group                 string   `json:"group_name"`
	GroupMaxSeq           string   `json:"group_operation_sequence"`
	LocalMaxSeq           string   `json:"local_operation_sequence"`
	Peers                 []string `json:"alive_peers_in_group"`
	OPLogMaxSeqGapToPull  uint16   `json:"oplog_max_seq_gap_to_pull"`
	OPLogPullMaxCountOnce uint16   `json:"oplog_pull_max_count_once"`
	OPLogPullInterval     int      `json:"oplog_pull_interval_in_second"`
	OPLogPullTimeout      int      `json:"oplog_pull_timeout_in_second"`
}
