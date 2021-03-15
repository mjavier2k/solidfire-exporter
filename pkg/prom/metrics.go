package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Descriptions struct {
	// Solidfire Metric Descriptions
	upDesc *prometheus.Desc

	// Volume Stats
	VolumeStatsActualIOPS              *prometheus.Desc
	VolumeStatsAverageIOPSizeBytes     *prometheus.Desc
	VolumeStatsBurstIOPSCredit         *prometheus.Desc
	VolumeStatsClientQueueDepth        *prometheus.Desc
	VolumeStatsLatencySeconds          *prometheus.Desc
	VolumeStatsNonZeroBlocks           *prometheus.Desc
	VolumeStatsReadBytes               *prometheus.Desc
	VolumeStatsLastSampleReadBytes     *prometheus.Desc
	VolumeStatsReadLatencySeconds      *prometheus.Desc
	VolumeStatsReadLatencySecondsTotal *prometheus.Desc
	VolumeStatsReadOpsTotal            *prometheus.Desc
	VolumeStatsLastSampleReadOps       *prometheus.Desc
	VolumeStatsThrottle                *prometheus.Desc
	VolumeStatsUnalignedReadsTotal     *prometheus.Desc
	VolumeStatsUnalignedWritesTotal    *prometheus.Desc
	VolumeStatsVolumeSizeBytes         *prometheus.Desc
	VolumeStatsVolumeUtilization       *prometheus.Desc
	VolumeStatsWriteBytes              *prometheus.Desc
	VolumeStatsLastSampleWriteBytes    *prometheus.Desc
	VolumeStatsWriteLatencySeconds     *prometheus.Desc
	VolumeStatsWriteLatencyTotal       *prometheus.Desc
	VolumeStatsWriteOpsTotal           *prometheus.Desc
	VolumeStatsWriteOpsLastSample      *prometheus.Desc
	VolumeStatsZeroBlocks              *prometheus.Desc

	// Cluster Capacity
	ClusterCapacityActiveBlockSpaceBytes             *prometheus.Desc
	ClusterCapacityActiveSessions                    *prometheus.Desc
	ClusterCapacityAverageIOPS                       *prometheus.Desc
	ClusterCapacityClusterRecentIOSizeBytes          *prometheus.Desc
	ClusterCapacityCurrentIOPS                       *prometheus.Desc
	ClusterCapacityMaxIOPS                           *prometheus.Desc
	ClusterCapacityMaxOverProvisionableSpaceBytes    *prometheus.Desc
	ClusterCapacityMaxProvisionedSpaceBytes          *prometheus.Desc
	ClusterCapacityMaxUsedMetadataSpaceBytes         *prometheus.Desc
	ClusterCapacityMaxUsedSpaceBytes                 *prometheus.Desc
	ClusterCapacityNonZeroBlocks                     *prometheus.Desc
	ClusterCapacityPeakActiveSessions                *prometheus.Desc
	ClusterCapacityPeakIOPS                          *prometheus.Desc
	ClusterCapacityProvisionedSpaceBytes             *prometheus.Desc
	ClusterCapacitySnapshotNonZeroBlocks             *prometheus.Desc
	ClusterCapacityIOPSTotal                         *prometheus.Desc
	ClusterCapacityUniqueBlocks                      *prometheus.Desc
	ClusterCapacityUniqueBlocksUsedSpaceBytes        *prometheus.Desc
	ClusterCapacityUsedMetadataSpaceBytes            *prometheus.Desc
	ClusterCapacityUsedMetadataSpaceInSnapshotsBytes *prometheus.Desc
	ClusterCapacityUsedSpaceBytes                    *prometheus.Desc
	ClusterCapacityZeroBlocks                        *prometheus.Desc
	//The following metrics are Calculated by us:
	ClusterCapacityThinProvisioningFactor *prometheus.Desc
	ClusterCapacityDeDuplicationFactor    *prometheus.Desc
	ClusterCapacityCompressionFactor      *prometheus.Desc
	ClusterCapacityEfficiencyFactor       *prometheus.Desc

	// ListClusterFaults
	ClusterActiveFaults *prometheus.Desc

	// ListNodeStats
	// NodeStatsCBytesIn                       *prometheus.Desc
	// NodeStatsCBytesOut                      *prometheus.Desc
	NodeStatsCount                          *prometheus.Desc
	NodeStatsCPUPercentage                  *prometheus.Desc
	NodeStatsCPUTotalSeconds                *prometheus.Desc
	NodeStatsInterfaceInBytesTotal          *prometheus.Desc
	NodeStatsInterfaceOutBytesTotal         *prometheus.Desc
	NodeStatsInterfaceUtilizationPercentage *prometheus.Desc
	// NodeStatsMBytesIn                    *prometheus.Desc
	// NodeStatsMBytesOut                   *prometheus.Desc
	// NodeStatsNetworkUtilizationCluster   *prometheus.Desc
	// NodeStatsNetworkUtilizationStorage   *prometheus.Desc
	NodeStatsReadLatencyTotal *prometheus.Desc
	NodeStatsReadOps          *prometheus.Desc
	// NodeStatsSBytesIn          *prometheus.Desc
	// NodeStatsSBytesOut         *prometheus.Desc
	NodeStatsTotalMemoryBytes  *prometheus.Desc
	NodeStatsUsedMemoryBytes   *prometheus.Desc
	NodeStatsWriteLatencyTotal *prometheus.Desc
	NodeStatsWriteOps          *prometheus.Desc
	NodeStatsLoadHistogram     *prometheus.Desc

	// ListVolumeQoSHistograms

	VolumeQoSBelowMinIopsPercentagesHistogram *prometheus.Desc
	VolumeQoSMinToMaxIopsPercentagesHistogram *prometheus.Desc
	VolumeQoSReadBlockSizesHistogram          *prometheus.Desc
	VolumeQoSTargetUtilizationPercentages     *prometheus.Desc
	VolumeQoSThrottlePercentages              *prometheus.Desc
	VolumeQoSWriteBlockSizes                  *prometheus.Desc

	// ListAllNodes
	NodeInfo *prometheus.Desc

	// GetClusterStats
	ClusterStatsActualIOPS            *prometheus.Desc
	ClusterStatsAverageIOBytes        *prometheus.Desc
	ClusterStatsClientQueueDepth      *prometheus.Desc
	ClusterStatsThroughputUtilization *prometheus.Desc
	ClusterStatsLatencySeconds        *prometheus.Desc
	ClusterStatsNormalizedIOPS        *prometheus.Desc
	ClusterStatsReadBytesTotal        *prometheus.Desc
	ClusterStatsLastSampleReadBytes   *prometheus.Desc
	ClusterStatsReadLatencySeconds    *prometheus.Desc
	ClusterStatsReadLatencyTotal      *prometheus.Desc
	ClusterStatsReadOpsTotal          *prometheus.Desc
	ClusterStatsLastSampleReadOps     *prometheus.Desc
	ClusterStatsSamplePeriodSeconds   *prometheus.Desc
	ClusterStatsServices              *prometheus.Desc
	ClusterStatsExpectedServices      *prometheus.Desc
	ClusterStatsUnalignedReadsTotal   *prometheus.Desc
	ClusterStatsUnalignedWrites       *prometheus.Desc
	ClusterStatsWriteBytesTotal       *prometheus.Desc
	ClusterStatsLastSampleWriteBytes  *prometheus.Desc
	ClusterStatsWriteLatency          *prometheus.Desc
	ClusterStatsWriteLatencyTotal     *prometheus.Desc
	ClusterStatsWriteOpsTotal         *prometheus.Desc
	ClusterStatsLastSampleWriteOps    *prometheus.Desc

	ClusterBlockFullness                       *prometheus.Desc
	ClusterFullness                            *prometheus.Desc
	ClusterMaxMetadataOverProvisionFactor      *prometheus.Desc
	ClusterMetadataFullness                    *prometheus.Desc
	ClusterSliceReserveUsedThresholdPercentage *prometheus.Desc
	ClusterStage2AwareThresholdPercentage      *prometheus.Desc
	ClusterStage2BlockThresholdBytes           *prometheus.Desc
	ClusterStage3BlockThresholdBytes           *prometheus.Desc
	ClusterStage3BlockThresholdPercentage      *prometheus.Desc
	ClusterStage3LowThresholdPercentage        *prometheus.Desc
	ClusterStage4BlockThresholdBytes           *prometheus.Desc
	ClusterStage4CriticalThreshold             *prometheus.Desc
	ClusterStage5BlockThresholdBytes           *prometheus.Desc
	ClusterTotalBytes                          *prometheus.Desc
	ClusterTotalMetadataBytes                  *prometheus.Desc
	ClusterUsedBytes                           *prometheus.Desc
	ClusterUsedMetadataBytes                   *prometheus.Desc

	DriveStatus        *prometheus.Desc
	DriveCapacityBytes *prometheus.Desc

	NodeISCSISessions *prometheus.Desc
	//NodeISCSIVolumes       *prometheus.Desc
}

func NewMetricDescriptions(namespace string) *Descriptions {
	var d Descriptions

	d.upDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Whether last scrape against Solidfire API was successful",
		nil,
		nil)

	d.VolumeStatsActualIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_actual_iops"),
		"The current actual IOPS to the volume in the last 500 milliseconds",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsAverageIOPSizeBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_average_iop_size_bytes"),
		"The average size in bytes of recent I/O to the volume in the last 500 milliseconds",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsBurstIOPSCredit = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_burst_iops_credit"),
		"The total number of IOP credits available to the user. When volumes are not using up to the configured maxIOPS, credits are accrued.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsClientQueueDepth = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_client_queue_depth"),
		"The number of outstanding read and write operations to the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_latency_seconds"),
		"The average time, in seconds, to complete operations to the volume in the last 500 milliseconds. A '0' (zero) value means there is no I/O to the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsNonZeroBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_non_zero_blocks"),
		"The total number of 4KiB blocks that contain data after the last garbage collection operation has completed.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsReadBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_read_bytes"),
		"The total cumulative bytes read from the volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsLastSampleReadBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_last_sample_read_bytes"),
		"The total number of bytes read from the volume during the last sample period.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsReadLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_read_latency_seconds"),
		"The average time, in seconds, to complete read operations to the volume in the last 500 milliseconds.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsReadLatencySecondsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_read_latency_seconds_total"),
		"The total time spent performing read operations from the volume",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsReadOpsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_read_ops_total"),
		"The total read operations to the volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsLastSampleReadOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_last_sample_read_ops"),
		"The total number of read operations during the last sample period",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsThrottle = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_throttle"),
		"A floating value between 0 and 1 that represents how much the system is throttling clients below their maxIOPS because of rereplication of data, transient errors, and snapshots taken.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsUnalignedReadsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_unaligned_reads_total"),
		"The total cumulative unaligned read operations to a volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsUnalignedWritesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_unaligned_writes_total"),
		"The total cumulative unaligned write operations to a volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsVolumeSizeBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_size_bytes"),
		"Total provisioned capacity in bytes.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsVolumeUtilization = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_utilization"),
		"A floating value that describes how much the client is using the volume. Value 0: The client is not using the volume. Value 1: The client is using their maximum. Value 1+: The client is using their burst.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsWriteBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_write_bytes"),
		"The total cumulative bytes written to the volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsLastSampleWriteBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_last_sample_write_bytes"),
		"The total number of bytes written to the volume during the last sample period.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsWriteLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_write_latency_seconds"),
		"The average time, in seconds, to complete write operations to a volume in the last 500 milliseconds.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsWriteLatencyTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_write_latency_seconds_total"),
		"The total time spent performing write operations to the volume",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsWriteOpsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_write_ops_total"),
		"The total cumulative write operations to the volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsWriteOpsLastSample = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_write_ops_last_sample"),
		"The total number of write operations during the last sample period.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsZeroBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_zero_blocks"),
		"The total number of empty 4KiB blocks without data after the last round of garbage collection operation has completed.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.ClusterCapacityActiveBlockSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_active_block_space_bytes"),
		"The amount of space on the block drives. This includes additional information such as metadata entries and space which can be cleaned up.",
		nil,
		nil,
	)
	d.ClusterCapacityActiveSessions = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_active_sessions"),
		"The number of active iSCSI sessions communicating with the cluster.",
		nil,
		nil,
	)
	d.ClusterCapacityAverageIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_average_iops"),
		"The average IOPS for the cluster since midnight Coordinated Universal Time (UTC)",
		nil,
		nil,
	)
	d.ClusterCapacityClusterRecentIOSizeBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_cluster_recent_io_size_bytes"),
		"The average size of IOPS to all volumes in the cluster",
		nil,
		nil,
	)
	d.ClusterCapacityCurrentIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_current_iops"),
		"The average IOPS for all volumes in the cluster over the last 5 seconds",
		nil,
		nil,
	)
	d.ClusterCapacityMaxIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_max_iops"),
		"The estimated maximum IOPS capability of the current cluster",
		nil,
		nil,
	)
	d.ClusterCapacityMaxOverProvisionableSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_max_over_provisionable_space_bytes"),
		"The maximum amount of provisionable space. This is a computed value. You cannot create new volumes if the current provisioned space plus the new volume size would exceed this number. The value is calculated as follows: maxOverProvisionableSpace = maxProvisionedSpace * maxMetadataOverProvisionFactor",
		nil,
		nil,
	)
	d.ClusterCapacityMaxProvisionedSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_max_provisioned_space_bytes"),
		"The total amount of provisionable space if all volumes are 100% filled (no thin provisioned metadata)",
		nil,
		nil,
	)
	d.ClusterCapacityMaxUsedMetadataSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_max_used_metadata_space_bytes"),
		"The number of bytes on volume drives used to store metadata",
		nil,
		nil,
	)
	d.ClusterCapacityMaxUsedSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_max_used_space_bytes"),
		" The total amount of space on all active block drives",
		nil,
		nil,
	)
	d.ClusterCapacityNonZeroBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_non_zero_blocks"),
		"The total number of 4KiB blocks that contain data after the last garbage collection operation has completed",
		nil,
		nil,
	)
	d.ClusterCapacityPeakActiveSessions = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_peak_active_sessions"),
		"The peak number of iSCSI connections since midnight UTC",
		nil,
		nil,
	)
	d.ClusterCapacityPeakIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_peak_iops"),
		"The highest value for currentIOPS since midnight UTC",
		nil,
		nil,
	)
	d.ClusterCapacityProvisionedSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_provisioned_space_bytes"),
		"The total amount of space provisioned in all volumes on the cluster",
		nil,
		nil,
	)
	d.ClusterCapacitySnapshotNonZeroBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_snapshot_non_zero_blocks"),
		"The total number of 4KiB blocks that contain data after the last garbage collection operation has completed",
		nil,
		nil,
	)

	d.ClusterCapacityIOPSTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_iops_total"),
		"The total number of I/O operations performed throughout the lifetime of the cluster.",
		nil,
		nil,
	)
	d.ClusterCapacityUniqueBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_unique_blocks"),
		"The total number of blocks stored on the block drives The value includes replicated blocks",
		nil,
		nil,
	)
	d.ClusterCapacityUniqueBlocksUsedSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_unique_blocks_used_space_bytes"),
		"The total amount of data the uniqueBlocks take up on the block drives",
		nil,
		nil,
	)
	d.ClusterCapacityUsedMetadataSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_used_metadata_space_bytes"),
		"The total number of bytes on volume drives used to store metadata",
		nil,
		nil,
	)
	d.ClusterCapacityUsedMetadataSpaceInSnapshotsBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_used_metadata_space_in_snapshots_bytes"),
		"The number of bytes on volume drives used for storing unique data in snapshots. This number provides an estimate of how much metadata space would be regained by deleting all snapshots on the system",
		nil,
		nil,
	)
	d.ClusterCapacityUsedSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_used_space_bytes"),
		"The total amount of space used by all block drives in the system",
		nil,
		nil,
	)
	d.ClusterCapacityZeroBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_zero_blocks"),
		"The total number of empty 4KiB blocks without data after the last round of garbage collection operation has completed",
		nil,
		nil,
	)
	d.ClusterCapacityThinProvisioningFactor = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_thin_provisioning_factor"),
		"The cluster thin provisioning factor. thinProvisioningFactor = (nonZeroBlocks + zeroBlocks) / nonZeroBlocks",
		nil,
		nil,
	)
	d.ClusterCapacityDeDuplicationFactor = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_de_duplication_factor"),
		"The cluster deDuplication factor. deDuplicationFactor = (nonZeroBlocks + snapshotNonZeroBlocks) / uniqueBlocks",
		nil,
		nil,
	)
	d.ClusterCapacityCompressionFactor = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_compression_factor"),
		"The cluster compression factor. compressionFactor = (uniqueBlocks * 4096) / (uniqueBlocksUsedSpace * 0.93)",
		nil,
		nil,
	)
	d.ClusterCapacityEfficiencyFactor = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_efficiency_factor"),
		"The cluster efficiency factor. efficiencyFactor = thinProvisioningFactor * deDuplicationFactor * compressionFactor",
		nil,
		nil,
	)

	d.ClusterActiveFaults = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_active_faults"),
		"List of any active faults detected on the cluster",
		[]string{"node_id", "node_name", "code", "severity", "type", "service_id", "resolved", "node_hardware_fault_id", "drive_id", "details"},
		nil,
	)

	d.NodeStatsCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_samples"),
		"Node stat sample count", // Undocumented metric
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsCPUPercentage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_cpu_percentage"),
		"CPU usage in percent.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsCPUTotalSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_cpu_seconds_total"),
		"CPU usage in seconds since last boot.", // undocumented metric
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsInterfaceInBytesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_interface_in_bytes_total"),
		"Bytes in on network interface.",
		[]string{"node_id", "node_name", "interface"},
		nil,
	)

	d.NodeStatsInterfaceOutBytesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_interface_out_bytes_total"),
		"Bytes out on network interface.",
		[]string{"node_id", "node_name", "interface"},
		nil,
	)

	d.NodeStatsInterfaceUtilizationPercentage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_interface_utilization_percentage"),
		"Network interface utilization (in percent) of network interface.",
		[]string{"node_id", "node_name", "interface"},
		nil,
	)

	d.NodeStatsReadLatencyTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_read_latency_seconds_total"),
		"The total time spent performing read operations since the creation of the cluster.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsReadOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_read_ops"),
		"Read Operations", // undocumented metric
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsWriteOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_write_ops"),
		"Write Operations", // undocumented metric
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsTotalMemoryBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_total_memory_bytes"),
		"Total node memory in bytes.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsUsedMemoryBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_used_memory_bytes"),
		"Total node memory used in bytes.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsWriteLatencyTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_write_latency_seconds_total"),
		"The total time spent performing write operations since the creation of the cluster.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsLoadHistogram = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_stats_load"),
		"System load histogram",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.VolumeQoSBelowMinIopsPercentagesHistogram = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_qos_below_min_iops_percentage"),
		"Volume QoS Below minimum IOPS percentage",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeQoSMinToMaxIopsPercentagesHistogram = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_qos_min_to_max_iops_percentage"),
		"Volume QoS min to max IOPS percentage",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeQoSReadBlockSizesHistogram = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_qos_read_block_sizes"),
		"Volume QoS read block sizes",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeQoSTargetUtilizationPercentages = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_qos_target_utilization_percentage"),
		"Volume QoS target utilization percentage",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeQoSThrottlePercentages = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_qos_throttle_percentage"),
		"Volume QoS throttle percentage",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeQoSWriteBlockSizes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_qos_write_block_sizes"),
		"Volume QoS write block sizes",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.NodeInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_info"),
		"Cluster node info",
		[]string{"node_id", "node_name", "chassis_name", "associated_fservice_id", "associated_master_service_id", "chassis_type", "cpu_model", "node_type", "platform_config_version", "sip", "sipi", "software_version", "uuid"},
		nil,
	)

	d.ClusterStatsActualIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_iops"),
		"Current actual IOPS for the entire cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterStatsAverageIOBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_average_io_bytes"),
		"Average size in bytes of recent I/O to the cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterStatsClientQueueDepth = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_client_queue_depth"),
		"The number of outstanding read and write operations to the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsThroughputUtilization = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_throughput_utilization"),
		"The cluster capacity being utilized. 0 - not utilized. 1 - 100% utilized.",
		nil,
		nil,
	)

	d.ClusterStatsLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_latency_seconds"),
		"The average time, in seconds, to complete operations to a cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterStatsNormalizedIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_normalized_iops"),
		"Average number of IOPS for the entire cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterStatsReadBytesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_read_bytes_total"),
		"The total cumulative bytes read from the cluster since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsLastSampleReadBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_last_sample_read_bytes"),
		"The total number of bytes read from the cluster during the last sample period.",
		nil,
		nil,
	)

	d.ClusterStatsReadLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_read_latency_seconds"),
		"The average time, in seconds, to complete read operations to the cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterStatsReadLatencyTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_read_latency_seconds_total"),
		"The total time spent performing read operations since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsReadOpsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_read_ops_total"),
		"The total cumulative read operations to the cluster since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsLastSampleReadOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_last_sample_read_ops"),
		"The total number of read operations during the last sample period.",
		nil,
		nil,
	)

	d.ClusterStatsSamplePeriodSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_sample_period_seconds"),
		"The length of the sample period, in seconds.",
		nil,
		nil,
	)

	d.ClusterStatsServices = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_running_services"),
		"The number of services running on the cluster. If equal to the servicesTotal, this indicates that valid statistics were collected from all nodes.",
		nil,
		nil,
	)

	d.ClusterStatsExpectedServices = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_expected_services"),
		"The total number of expected services running on the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsUnalignedReadsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_unaligned_reads_total"),
		"The total cumulative unaligned read operations to a cluster since the creation of the cluster",
		nil,
		nil,
	)

	d.ClusterStatsUnalignedWrites = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_unaligned_writes_total"),
		"The total cumulative unaligned write operations to a cluster since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsWriteBytesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_write_bytes_total"),
		"The total cumulative bytes written to the cluster since the creation of the cluster",
		nil,
		nil,
	)

	d.ClusterStatsLastSampleWriteBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_last_sample_write_bytes"),
		"The total number of bytes written to the cluster during the last sample period.",
		nil,
		nil,
	)

	d.ClusterStatsWriteLatency = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_write_latency_seconds"),
		"The average time, in seconds, to complete write operations to a cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterStatsWriteLatencyTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_write_latency_seconds_total"),
		"The total time spent performing write operations since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsWriteOpsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_write_ops_total"),
		"The total cumulative write operations to the cluster since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsLastSampleWriteOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_last_sample_write_ops"),
		"The total number of write operations during the last sample period.",
		nil,
		nil,
	)

	d.ClusterBlockFullness = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_block_fullness"),
		"The current computed level of block fullness of the cluster.",
		[]string{"level"},
		nil,
	)

	d.ClusterFullness = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_fullness"),
		"Reflects the highest level of fullness between 'blockFullness' and 'metadataFullness'.",
		[]string{"level"},
		nil,
	)

	d.ClusterMaxMetadataOverProvisionFactor = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_max_metadata_over_provision_factor"),
		"A value representative of the number of times metadata space can be over provisioned relative to the amount of space available.",
		nil,
		nil,
	)

	d.ClusterMetadataFullness = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_metadata_fullness"),
		"The current computed level of metadata fullness of the cluster.",
		[]string{"level"},
		nil,
	)

	d.ClusterSliceReserveUsedThresholdPercentage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_slice_reserve_used_threshold_percentage"),
		"Error condition. A system alert is triggered if the reserved slice utilization is greater than the sliceReserveUsedThresholdPct value returned.",
		nil,
		nil,
	)

	d.ClusterStage2AwareThresholdPercentage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stage2_aware_threshold_percentage"),
		"Awareness condition. The value that is set for 'Stage 2' cluster threshold level.",
		nil,
		nil,
	)

	d.ClusterStage2BlockThresholdBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stage2_block_threshold_bytes"),
		"Number of bytes being used by the cluster at which a stage2 condition will exist.",
		nil,
		nil,
	)

	d.ClusterStage3BlockThresholdBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stage3_block_threshold_bytes"),
		"Number of bytes being used by the cluster at which a stage3 condition will exist.",
		nil,
		nil,
	)

	d.ClusterStage3BlockThresholdPercentage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stage3_block_threshold_percentage"),
		"Percent value set for stage3. At this percent full, a warning is posted in the Alerts log.",
		nil,
		nil,
	)

	d.ClusterStage3LowThresholdPercentage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stage3_low_threshold_percentage"),
		"Error condition. The threshold at which a system alert is created due to low capacity on a cluster",
		nil,
		nil,
	)

	d.ClusterStage4BlockThresholdBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stage4_block_threshold_bytes"),
		"Number of bytes being used by the cluster at which a stage4 condition will exist",
		nil,
		nil,
	)

	d.ClusterStage4CriticalThreshold = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stage4_critical_threshold_percentage"),
		"Error condition. The threshold at which a system alert is created to warn about critically low capacity on a cluster.",
		nil,
		nil,
	)

	d.ClusterStage5BlockThresholdBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stage5_block_threshold_bytes"),
		"The number of bytes being used by the cluster at which a stage5 condition will exist.",
		nil,
		nil,
	)

	d.ClusterTotalBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_total_bytes"),
		"Physical capacity of the cluster, measured in bytes.",
		nil,
		nil,
	)

	d.ClusterTotalMetadataBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_total_metadata_bytes"),
		"Total amount of space that can be used to store metadata",
		nil,
		nil,
	)

	d.ClusterUsedBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_used_bytes"),
		"Number of bytes used on the cluster.",
		nil,
		nil,
	)

	d.ClusterUsedMetadataBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_used_metadata_bytes"),
		"Amount of space used on volume drives to store metadata.",
		nil,
		nil,
	)

	d.DriveStatus = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "drive_status"),
		"The drive status for each individual drives in the cluster's active nodes",
		[]string{"node_id", "node_name", "drive_id", "serial", "slot", "status", "type"},
		nil,
	)

	d.DriveCapacityBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "drive_capacity_bytes"),
		"The drive capacity for each individual drives in the cluster's active nodes",
		[]string{"node_id", "node_name", "drive_id", "serial", "slot", "type"},
		nil,
	)

	d.NodeISCSISessions = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_iscsi_sessions"),
		"The total number of iscsi sessions per node and volume",
		[]string{"node_id", "node_name", "volume_id", "volume_name"},
		nil,
	)

	return &d
}
