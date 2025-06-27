package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Descriptions struct {
	// Solidfire Metric Descriptions
	upDesc *prometheus.Desc

	// Volume Stats
	VolumeActualIOPS              *prometheus.Desc
	VolumeAverageIOPSizeBytes     *prometheus.Desc
	VolumeBurstIOPSCredit         *prometheus.Desc
	VolumeClientQueueDepth        *prometheus.Desc
	VolumeLastSampleReadBytes     *prometheus.Desc
	VolumeLastSampleReadOps       *prometheus.Desc
	VolumeLastSampleWriteBytes    *prometheus.Desc
	VolumeLatencySeconds          *prometheus.Desc
	VolumeNonZeroBlocks           *prometheus.Desc
	VolumeReadBytesTotal          *prometheus.Desc
	VolumeReadLatencySeconds      *prometheus.Desc
	VolumeReadLatencySecondsTotal *prometheus.Desc
	VolumeReadOpsTotal            *prometheus.Desc
	VolumeThrottle                *prometheus.Desc
	VolumeUnalignedReadsTotal     *prometheus.Desc
	VolumeUnalignedWritesTotal    *prometheus.Desc
	VolumeSizeBytes               *prometheus.Desc
	VolumeUtilization             *prometheus.Desc
	VolumeWriteBytesTotal         *prometheus.Desc
	VolumeWriteLatencySeconds     *prometheus.Desc
	VolumeWriteLatencyTotal       *prometheus.Desc
	VolumeWriteOpsLastSample      *prometheus.Desc
	VolumeWriteOpsTotal           *prometheus.Desc
	VolumeStatsZeroBlocks         *prometheus.Desc

	// ListVolumeQoSHistograms
	VolumeQoSBelowMinIopsPercentagesHistogram      *prometheus.Desc
	VolumeQoSMinToMaxIopsPercentagesHistogram      *prometheus.Desc
	VolumeQoSReadBlockSizesHistogram               *prometheus.Desc
	VolumeQoSTargetUtilizationPercentagesHistogram *prometheus.Desc
	VolumeQoSThrottlePercentagesHistogram          *prometheus.Desc
	VolumeQoSWriteBlockSizesHistogram              *prometheus.Desc

	// Cluster Capacity
	ClusterActiveBlockSpaceBytes             *prometheus.Desc
	ClusterActiveSessions                    *prometheus.Desc
	ClusterAverageIOPS                       *prometheus.Desc
	ClusterClusterRecentIOSizeBytes          *prometheus.Desc
	ClusterCurrentIOPS                       *prometheus.Desc
	ClusterIOPSTotal                         *prometheus.Desc
	ClusterMaxIOPS                           *prometheus.Desc
	ClusterMaxOverProvisionableSpaceBytes    *prometheus.Desc
	ClusterMaxProvisionedSpaceBytes          *prometheus.Desc
	ClusterMaxUsedMetadataSpaceBytes         *prometheus.Desc
	ClusterMaxUsedSpaceBytes                 *prometheus.Desc
	ClusterNonZeroBlocks                     *prometheus.Desc
	ClusterPeakActiveSessions                *prometheus.Desc
	ClusterPeakIOPS                          *prometheus.Desc
	ClusterProvisionedSpaceBytes             *prometheus.Desc
	ClusterSnapshotNonZeroBlocks             *prometheus.Desc
	ClusterUniqueBlocks                      *prometheus.Desc
	ClusterUniqueBlocksUsedSpaceBytes        *prometheus.Desc
	ClusterUsedMetadataSpaceBytes            *prometheus.Desc
	ClusterUsedMetadataSpaceInSnapshotsBytes *prometheus.Desc
	ClusterUsedSpaceBytes                    *prometheus.Desc
	ClusterZeroBlocks                        *prometheus.Desc
	//The following metrics are Calculated by us:
	ClusterCompressionFactor      *prometheus.Desc
	ClusterDeDuplicationFactor    *prometheus.Desc
	ClusterEfficiencyFactor       *prometheus.Desc
	ClusterThinProvisioningFactor *prometheus.Desc

	// ListClusterFaults
	ClusterActiveFaults *prometheus.Desc

	// ListNodeStats
	NodeCPUPercentage                  *prometheus.Desc
	NodeCPUSecondsTotal                *prometheus.Desc
	NodeInterfaceInBytesTotal          *prometheus.Desc
	NodeInterfaceOutBytesTotal         *prometheus.Desc
	NodeInterfaceUtilizationPercentage *prometheus.Desc
	NodeLoadHistogram                  *prometheus.Desc
	NodeReadLatencyTotal               *prometheus.Desc
	NodeReadOpsTotal                   *prometheus.Desc
	NodeSamples                        *prometheus.Desc
	NodeTotalMemoryBytes               *prometheus.Desc
	NodeUsedMemoryBytes                *prometheus.Desc
	NodeWriteLatencyTotal              *prometheus.Desc
	NodeWriteOpsTotal                  *prometheus.Desc

	// ListAllNodes
	NodeInfo *prometheus.Desc

	// GetClusterStats
	ClusterActualIOPS            *prometheus.Desc
	ClusterAverageIOBytes        *prometheus.Desc
	ClusterClientQueueDepth      *prometheus.Desc
	ClusterExpectedServices      *prometheus.Desc
	ClusterLastSampleReadBytes   *prometheus.Desc
	ClusterLastSampleReadOps     *prometheus.Desc
	ClusterLastSampleWriteBytes  *prometheus.Desc
	ClusterLastSampleWriteOps    *prometheus.Desc
	ClusterLatencySeconds        *prometheus.Desc
	ClusterNormalizedIOPS        *prometheus.Desc
	ClusterReadBytesTotal        *prometheus.Desc
	ClusterReadLatencySeconds    *prometheus.Desc
	ClusterReadLatencyTotal      *prometheus.Desc
	ClusterReadOpsTotal          *prometheus.Desc
	ClusterSamplePeriodSeconds   *prometheus.Desc
	ClusterServices              *prometheus.Desc
	ClusterThroughputUtilization *prometheus.Desc
	ClusterUnalignedReadsTotal   *prometheus.Desc
	ClusterUnalignedWritesTotal  *prometheus.Desc
	ClusterWriteBytesTotal       *prometheus.Desc
	ClusterWriteLatency          *prometheus.Desc
	ClusterWriteLatencyTotal     *prometheus.Desc
	ClusterWriteOpsTotal         *prometheus.Desc

	// GetClusterFullThreshold
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

	//ListBulkVolumeJob
	BulkVolumeJobStatus        *prometheus.Desc
	BulkVolumeJobPercentage    *prometheus.Desc
	BulkVolumeJobRemainingTime *prometheus.Desc
}

func NewMetricDescriptions(namespace string) *Descriptions {
	var d Descriptions

	d.upDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Whether last scrape against Solidfire API was successful",
		nil,
		nil)

	d.VolumeActualIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_actual_iops"),
		"The current actual IOPS to the volume in the last 500 milliseconds",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeAverageIOPSizeBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_average_iop_size_bytes"),
		"The average size in bytes of recent I/O to the volume in the last 500 milliseconds",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeBurstIOPSCredit = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_burst_iops_credit"),
		"The total number of IOP credits available to the user. When volumes are not using up to the configured maxIOPS, credits are accrued.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeClientQueueDepth = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_client_queue_depth"),
		"The number of outstanding read and write operations to the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_latency_seconds"),
		"The average time, in seconds, to complete operations to the volume in the last 500 milliseconds. A '0' (zero) value means there is no I/O to the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeNonZeroBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_non_zero_blocks"),
		"The total number of 4KiB blocks that contain data after the last garbage collection operation has completed.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeReadBytesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_read_bytes_total"),
		"The total cumulative bytes read from the volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeLastSampleReadBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_last_sample_read_bytes"),
		"The total number of bytes read from the volume during the last sample period.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeReadLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_read_latency_seconds"),
		"The average time, in seconds, to complete read operations to the volume in the last 500 milliseconds.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeReadLatencySecondsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_read_latency_seconds_total"),
		"The total time spent performing read operations from the volume",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeReadOpsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_read_ops_total"),
		"The total read operations to the volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeLastSampleReadOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_last_sample_read_ops"),
		"The total number of read operations during the last sample period",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeThrottle = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_throttle"),
		"A floating value between 0 and 1 that represents how much the system is throttling clients below their maxIOPS because of rereplication of data, transient errors, and snapshots taken.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeUnalignedReadsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_unaligned_reads_total"),
		"The total cumulative unaligned read operations to a volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeUnalignedWritesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_unaligned_writes_total"),
		"The total cumulative unaligned write operations to a volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeSizeBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_size_bytes"),
		"Total provisioned capacity in bytes.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeUtilization = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_utilization"),
		"A floating value that describes how much the client is using the volume. Value 0: The client is not using the volume. Value 1: The client is using their maximum. Value 1+: The client is using their burst.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeWriteBytesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_write_bytes_total"),
		"The total cumulative bytes written to the volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeLastSampleWriteBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_last_sample_write_bytes"),
		"The total number of bytes written to the volume during the last sample period.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeWriteLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_write_latency_seconds"),
		"The average time, in seconds, to complete write operations to a volume in the last 500 milliseconds.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeWriteLatencyTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_write_latency_seconds_total"),
		"The total time spent performing write operations to the volume",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeWriteOpsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_write_ops_total"),
		"The total cumulative write operations to the volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeWriteOpsLastSample = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_write_ops_last_sample"),
		"The total number of write operations during the last sample period.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsZeroBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_zero_blocks"),
		"The total number of empty 4KiB blocks without data after the last round of garbage collection operation has completed.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.ClusterActiveBlockSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_active_block_space_bytes"),
		"The amount of space on the block drives. This includes additional information such as metadata entries and space which can be cleaned up.",
		nil,
		nil,
	)
	d.ClusterActiveSessions = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_active_sessions"),
		"The number of active iSCSI sessions communicating with the cluster.",
		nil,
		nil,
	)
	d.ClusterAverageIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_average_iops"),
		"The average IOPS for the cluster since midnight Coordinated Universal Time (UTC)",
		nil,
		nil,
	)
	d.ClusterClusterRecentIOSizeBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_recent_io_size_bytes"),
		"The average size of IOPS to all volumes in the cluster",
		nil,
		nil,
	)
	d.ClusterCurrentIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_current_iops"),
		"The average IOPS for all volumes in the cluster over the last 5 seconds",
		nil,
		nil,
	)
	d.ClusterMaxIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_max_iops"),
		"The estimated maximum IOPS capability of the current cluster",
		nil,
		nil,
	)
	d.ClusterMaxOverProvisionableSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_max_over_provisionable_space_bytes"),
		"The maximum amount of provisionable space. This is a computed value. You cannot create new volumes if the current provisioned space plus the new volume size would exceed this number. The value is calculated as follows: maxOverProvisionableSpace = maxProvisionedSpace * maxMetadataOverProvisionFactor",
		nil,
		nil,
	)
	d.ClusterMaxProvisionedSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_max_provisioned_space_bytes"),
		"The total amount of provisionable space if all volumes are 100% filled (no thin provisioned metadata)",
		nil,
		nil,
	)
	d.ClusterMaxUsedMetadataSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_max_used_metadata_space_bytes"),
		"The number of bytes on volume drives used to store metadata",
		nil,
		nil,
	)
	d.ClusterMaxUsedSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_max_used_space_bytes"),
		" The total amount of space on all active block drives",
		nil,
		nil,
	)
	d.ClusterNonZeroBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_non_zero_blocks"),
		"The total number of 4KiB blocks that contain data after the last garbage collection operation has completed",
		nil,
		nil,
	)
	d.ClusterPeakActiveSessions = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_peak_active_sessions"),
		"The peak number of iSCSI connections since midnight UTC",
		nil,
		nil,
	)
	d.ClusterPeakIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_peak_iops"),
		"The highest value for currentIOPS since midnight UTC",
		nil,
		nil,
	)
	d.ClusterProvisionedSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_provisioned_space_bytes"),
		"The total amount of space provisioned in all volumes on the cluster",
		nil,
		nil,
	)
	d.ClusterSnapshotNonZeroBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_snapshot_non_zero_blocks"),
		"The total number of 4KiB blocks that contain data after the last garbage collection operation has completed",
		nil,
		nil,
	)

	d.ClusterIOPSTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_iops_total"),
		"The total number of I/O operations performed throughout the lifetime of the cluster.",
		nil,
		nil,
	)
	d.ClusterUniqueBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_unique_blocks"),
		"The total number of blocks stored on the block drives The value includes replicated blocks",
		nil,
		nil,
	)
	d.ClusterUniqueBlocksUsedSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_unique_blocks_used_space_bytes"),
		"The total amount of data the uniqueBlocks take up on the block drives",
		nil,
		nil,
	)
	d.ClusterUsedMetadataSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_used_metadata_space_bytes"),
		"The total number of bytes on volume drives used to store metadata",
		nil,
		nil,
	)
	d.ClusterUsedMetadataSpaceInSnapshotsBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_used_metadata_space_in_snapshots_bytes"),
		"The number of bytes on volume drives used for storing unique data in snapshots. This number provides an estimate of how much metadata space would be regained by deleting all snapshots on the system",
		nil,
		nil,
	)
	d.ClusterUsedSpaceBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_used_space_bytes"),
		"The total amount of space used by all block drives in the system",
		nil,
		nil,
	)
	d.ClusterZeroBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_zero_blocks"),
		"The total number of empty 4KiB blocks without data after the last round of garbage collection operation has completed",
		nil,
		nil,
	)
	d.ClusterThinProvisioningFactor = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_thin_provisioning_factor"),
		"The cluster thin provisioning factor. thinProvisioningFactor = (nonZeroBlocks + zeroBlocks) / nonZeroBlocks",
		nil,
		nil,
	)
	d.ClusterDeDuplicationFactor = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_de_duplication_factor"),
		"The cluster deDuplication factor. deDuplicationFactor = (nonZeroBlocks + snapshotNonZeroBlocks) / uniqueBlocks",
		nil,
		nil,
	)
	d.ClusterCompressionFactor = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_compression_factor"),
		"The cluster compression factor. compressionFactor = (uniqueBlocks * 4096) / (uniqueBlocksUsedSpace * 0.93)",
		nil,
		nil,
	)
	d.ClusterEfficiencyFactor = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_efficiency_factor"),
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

	d.NodeSamples = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_samples"),
		"Node stat sample count", // Undocumented metric
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeCPUPercentage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_cpu_percentage"),
		"CPU usage in percent.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeCPUSecondsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_cpu_seconds_total"),
		"CPU usage in seconds since last boot.", // undocumented metric
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeInterfaceInBytesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_interface_in_bytes_total"),
		"Bytes in on network interface.",
		[]string{"node_id", "node_name", "interface"},
		nil,
	)

	d.NodeInterfaceOutBytesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_interface_out_bytes_total"),
		"Bytes out on network interface.",
		[]string{"node_id", "node_name", "interface"},
		nil,
	)

	d.NodeInterfaceUtilizationPercentage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_interface_utilization_percentage"),
		"Network interface utilization (in percent) of network interface.",
		[]string{"node_id", "node_name", "interface"},
		nil,
	)

	d.NodeReadLatencyTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_read_latency_seconds_total"),
		"The total time spent performing read operations since the creation of the cluster.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeReadOpsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_read_ops_total"),
		"Total read operations to a node.", // undocumented metric
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeWriteOpsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_write_ops_total"),
		"Total write operations to a node", // undocumented metric
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeTotalMemoryBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_total_memory_bytes"),
		"Total node memory in bytes.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeUsedMemoryBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_used_memory_bytes"),
		"Total node memory used in bytes.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeWriteLatencyTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_write_latency_seconds_total"),
		"The total time spent performing write operations since the creation of the cluster.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeLoadHistogram = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_load"),
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
		prometheus.BuildFQName(namespace, "", "volume_qos_read_block_sizes_bytes"),
		"Volume QoS read block sizes",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeQoSTargetUtilizationPercentagesHistogram = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_qos_target_utilization_percentage"),
		"Volume QoS target utilization percentage",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeQoSThrottlePercentagesHistogram = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_qos_throttle_percentage"),
		"Volume QoS throttle percentage",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeQoSWriteBlockSizesHistogram = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_qos_write_block_sizes_bytes"),
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

	d.ClusterActualIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_iops"),
		"Current actual IOPS for the entire cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterAverageIOBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_average_io_bytes"),
		"Average size in bytes of recent I/O to the cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterClientQueueDepth = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_client_queue_depth"),
		"The number of outstanding read and write operations to the cluster.",
		nil,
		nil,
	)

	d.ClusterThroughputUtilization = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_throughput_utilization"),
		"The cluster capacity being utilized. 0 - not utilized. 1 - 100% utilized.",
		nil,
		nil,
	)

	d.ClusterLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_latency_seconds"),
		"The average time, in seconds, to complete operations to a cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterNormalizedIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_normalized_iops"),
		"Average number of IOPS for the entire cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterReadBytesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_read_bytes_total"),
		"The total cumulative bytes read from the cluster since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterLastSampleReadBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_last_sample_read_bytes"),
		"The total number of bytes read from the cluster during the last sample period.",
		nil,
		nil,
	)

	d.ClusterReadLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_read_latency_seconds"),
		"The average time, in seconds, to complete read operations to the cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterReadLatencyTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_read_latency_seconds_total"),
		"The total time spent performing read operations since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterReadOpsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_read_ops_total"),
		"The total cumulative read operations to the cluster since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterLastSampleReadOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_last_sample_read_ops"),
		"The total number of read operations during the last sample period.",
		nil,
		nil,
	)

	d.ClusterSamplePeriodSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_sample_period_seconds"),
		"The length of the sample period, in seconds.",
		nil,
		nil,
	)

	d.ClusterServices = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_services_running"),
		"The number of services running on the cluster. If equal to the servicesTotal, this indicates that valid statistics were collected from all nodes.",
		nil,
		nil,
	)

	d.ClusterExpectedServices = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_services_expected"),
		"The total number of expected services running on the cluster.",
		nil,
		nil,
	)

	d.ClusterUnalignedReadsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_unaligned_reads_total"),
		"The total cumulative unaligned read operations to a cluster since the creation of the cluster",
		nil,
		nil,
	)

	d.ClusterUnalignedWritesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_unaligned_writes_total"),
		"The total cumulative unaligned write operations to a cluster since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterWriteBytesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_write_bytes_total"),
		"The total cumulative bytes written to the cluster since the creation of the cluster",
		nil,
		nil,
	)

	d.ClusterLastSampleWriteBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_last_sample_write_bytes"),
		"The total number of bytes written to the cluster during the last sample period.",
		nil,
		nil,
	)

	d.ClusterWriteLatency = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_write_latency_seconds"),
		"The average time, in seconds, to complete write operations to a cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterWriteLatencyTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_write_latency_seconds_total"),
		"The total time spent performing write operations since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterWriteOpsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_write_ops_total"),
		"The total cumulative write operations to the cluster since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterLastSampleWriteOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_last_sample_write_ops"),
		"The total number of write operations during the last sample period.",
		nil,
		nil,
	)

	d.ClusterBlockFullness = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_block_fullness"),
		"The current computed level of block fullness of the cluster. See https://library.netapp.com/ecm/ecm_download_file/ECMLP2856155 for more details.",
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
		"The current computed level of metadata fullness of the cluster. See https://library.netapp.com/ecm/ecm_download_file/ECMLP2856155 for more details.",
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

	d.BulkVolumeJobStatus = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "bulk_volume_job_status"),
		"The status of bulk volume read or write operation that is occurring in the system.  0 = Preparing, 1 = Running, 2 = Complete, 3 = Failed.",
		[]string{"bulk_volume_id", "created_time", "format", "key", "script", "volume_id", "status", "type", "snapshot_id"},
		nil,
	)

	d.BulkVolumeJobPercentage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "bulk_volume_job_percentage"),
		"The percentage completed of bulk volume read or write operation that is occurring in the system.",
		[]string{"bulk_volume_id", "created_time", "format", "key", "script", "volume_id", "status", "type", "snapshot_id"},
		nil,
	)

	d.BulkVolumeJobRemainingTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "bulk_volume_job_remaining_time_seconds"),
		"The remaining time in seconds of bulk volume read or write operation that is occurring in the system.",
		[]string{"bulk_volume_id", "created_time", "format", "key", "script", "volume_id", "status", "type", "snapshot_id"},
		nil,
	)

	return &d
}
