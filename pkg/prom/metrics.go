package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Descriptions struct {
	// Solidfire Metric Descriptions
	ScrapeSuccessDesc *prometheus.Desc

	// Volume Stats
	VolumeStatsActualIOPSDesc            *prometheus.Desc
	VolumeStatsAverageIOPSizeDesc        *prometheus.Desc
	VolumeStatsBurstIOPSCreditDesc       *prometheus.Desc
	VolumeStatsClientQueueDepthDesc      *prometheus.Desc
	VolumeStatsLatencyUSecDesc           *prometheus.Desc
	VolumeStatsNonZeroBlocksDesc         *prometheus.Desc
	VolumeStatsReadBytesDesc             *prometheus.Desc
	VolumeStatsReadBytesLastSampleDesc   *prometheus.Desc
	VolumeStatsReadLatencyUSecDesc       *prometheus.Desc
	VolumeStatsReadLatencyUSecTotalDesc  *prometheus.Desc
	VolumeStatsReadOpsDesc               *prometheus.Desc
	VolumeStatsReadOpsLastSampleDesc     *prometheus.Desc
	VolumeStatsThrottleDesc              *prometheus.Desc
	VolumeStatsUnalignedReadsDesc        *prometheus.Desc
	VolumeStatsUnalignedWritesDesc       *prometheus.Desc
	VolumeStatsVolumeSizeDesc            *prometheus.Desc
	VolumeStatsVolumeUtilizationDesc     *prometheus.Desc
	VolumeStatsWriteBytesDesc            *prometheus.Desc
	VolumeStatsWriteBytesLastSampleDesc  *prometheus.Desc
	VolumeStatsWriteLatencyUSecDesc      *prometheus.Desc
	VolumeStatsWriteLatencyUSecTotalDesc *prometheus.Desc
	VolumeStatsWriteOpsDesc              *prometheus.Desc
	VolumeStatsWriteOpsLastSampleDesc    *prometheus.Desc
	VolumeStatsZeroBlocksDesc            *prometheus.Desc

	// Cluster Capacity
	ClusterCapacityActiveBlockSpace             *prometheus.Desc
	ClusterCapacityActiveSessions               *prometheus.Desc
	ClusterCapacityAverageIOPS                  *prometheus.Desc
	ClusterCapacityClusterRecentIOSize          *prometheus.Desc
	ClusterCapacityCurrentIOPS                  *prometheus.Desc
	ClusterCapacityMaxIOPS                      *prometheus.Desc
	ClusterCapacityMaxOverProvisionableSpace    *prometheus.Desc
	ClusterCapacityMaxProvisionedSpace          *prometheus.Desc
	ClusterCapacityMaxUsedMetadataSpace         *prometheus.Desc
	ClusterCapacityMaxUsedSpace                 *prometheus.Desc
	ClusterCapacityNonZeroBlocks                *prometheus.Desc
	ClusterCapacityPeakActiveSessions           *prometheus.Desc
	ClusterCapacityPeakIOPS                     *prometheus.Desc
	ClusterCapacityProvisionedSpace             *prometheus.Desc
	ClusterCapacitySnapshotNonZeroBlocks        *prometheus.Desc
	ClusterCapacityTotalOps                     *prometheus.Desc
	ClusterCapacityUniqueBlocks                 *prometheus.Desc
	ClusterCapacityUniqueBlocksUsedSpace        *prometheus.Desc
	ClusterCapacityUsedMetadataSpace            *prometheus.Desc
	ClusterCapacityUsedMetadataSpaceInSnapshots *prometheus.Desc
	ClusterCapacityUsedSpace                    *prometheus.Desc
	ClusterCapacityZeroBlocks                   *prometheus.Desc
	//The following metrics are Calculated by us:
	ClusterCapacityThinProvisioningFactor *prometheus.Desc
	ClusterCapacityDeDuplicationFactor    *prometheus.Desc
	ClusterCapacityCompressionFactor      *prometheus.Desc
	ClusterCapacityEfficiencyFactor       *prometheus.Desc

	// ListClusterFaults
	ClusterActiveFaultsWarning      *prometheus.Desc
	ClusterActiveFaultsError        *prometheus.Desc
	ClusterActiveFaultsCritical     *prometheus.Desc
	ClusterActiveFaultsBestPractice *prometheus.Desc

	// ListNodeStats
	NodeStatsCBytesIn                  *prometheus.Desc
	NodeStatsCBytesOut                 *prometheus.Desc
	NodeStatsCount                     *prometheus.Desc
	NodeStatsCPUPercentage             *prometheus.Desc
	NodeStatsCPUTotalSeconds           *prometheus.Desc
	NodeStatsMBytesIn                  *prometheus.Desc
	NodeStatsMBytesOut                 *prometheus.Desc
	NodeStatsNetworkUtilizationCluster *prometheus.Desc
	NodeStatsNetworkUtilizationStorage *prometheus.Desc
	NodeStatsReadLatencyUSecTotal      *prometheus.Desc
	NodeStatsReadOps                   *prometheus.Desc
	NodeStatsSBytesIn                  *prometheus.Desc
	NodeStatsSBytesOut                 *prometheus.Desc
	NodeStatsUsedMemory                *prometheus.Desc
	NodeStatsWriteLatencyUSecTotal     *prometheus.Desc
	NodeStatsWriteOps                  *prometheus.Desc
	NodeStatsLoadHistogram             *prometheus.Desc

	// ListVolumeQoSHistograms

	VolumeQoSBelowMinIopsPercentagesHistogram *prometheus.Desc
	VolumeQoSMinToMaxIopsPercentagesHistogram *prometheus.Desc
	VolumeQoSReadBlockSizesHistogram          *prometheus.Desc
	VolumeQoSTargetUtilizationPercentages     *prometheus.Desc
	VolumeQoSThrottlePercentages              *prometheus.Desc
	VolumeQoSWriteBlockSizes                  *prometheus.Desc

	// ListAllNodes
	Node *prometheus.Desc

	// GetClusterStats
	ClusterStatsActualIOPS            *prometheus.Desc
	ClusterStatsAverageIOPSize        *prometheus.Desc
	ClusterStatsClientQueueDepth      *prometheus.Desc
	ClusterStatsClusterUtilization    *prometheus.Desc
	ClusterStatsLatencyUSec           *prometheus.Desc
	ClusterStatsNormalizedIOPS        *prometheus.Desc
	ClusterStatsReadBytes             *prometheus.Desc
	ClusterStatsReadBytesLastSample   *prometheus.Desc
	ClusterStatsReadLatencyUSec       *prometheus.Desc
	ClusterStatsReadLatencyUSecTotal  *prometheus.Desc
	ClusterStatsReadOps               *prometheus.Desc
	ClusterStatsReadOpsLastSample     *prometheus.Desc
	ClusterStatsSamplePeriodMsec      *prometheus.Desc
	ClusterStatsServicesCount         *prometheus.Desc
	ClusterStatsServicesTotal         *prometheus.Desc
	ClusterStatsUnalignedReads        *prometheus.Desc
	ClusterStatsUnalignedWrites       *prometheus.Desc
	ClusterStatsWriteBytes            *prometheus.Desc
	ClusterStatsWriteBytesLastSample  *prometheus.Desc
	ClusterStatsWriteLatencyUSec      *prometheus.Desc
	ClusterStatsWriteLatencyUSecTotal *prometheus.Desc
	ClusterStatsWriteOps              *prometheus.Desc
	ClusterStatsWriteOpsLastSample    *prometheus.Desc

	ClusterThresholdBlockFullness                  *prometheus.Desc
	ClusterThresholdFullness                       *prometheus.Desc
	ClusterThresholdMaxMetadataOverProvisionFactor *prometheus.Desc
	ClusterThresholdMetadataFullness               *prometheus.Desc
	ClusterThresholdSliceReserveUsedThresholdPct   *prometheus.Desc
	ClusterThresholdStage2AwareThreshold           *prometheus.Desc
	ClusterThresholdStage2BlockThresholdBytes      *prometheus.Desc
	ClusterThresholdStage3BlockThresholdBytes      *prometheus.Desc
	ClusterThresholdStage3BlockThresholdPercent    *prometheus.Desc
	ClusterThresholdStage3LowThreshold             *prometheus.Desc
	ClusterThresholdStage4BlockThresholdBytes      *prometheus.Desc
	ClusterThresholdStage4CriticalThreshold        *prometheus.Desc
	ClusterThresholdStage5BlockThresholdBytes      *prometheus.Desc
	ClusterThresholdSumTotalClusterBytes           *prometheus.Desc
	ClusterThresholdSumTotalMetadataClusterBytes   *prometheus.Desc
	ClusterThresholdSumUsedClusterBytes            *prometheus.Desc
	ClusterThresholdSumUsedMetadataClusterBytes    *prometheus.Desc

	ListDrivesStatus   *prometheus.Desc
	ListDrivesType     *prometheus.Desc
	ListDrivesCapacity *prometheus.Desc
}

func NewMetricDescriptions(namespace string) *Descriptions {
	var d Descriptions

	d.ScrapeSuccessDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "scrape_success"),
		"Whether last scrape against Solidfire API was successful",
		nil,
		nil)

	d.VolumeStatsActualIOPSDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_actual_iops"),
		"The current actual IOPS to the volume in the last 500 milliseconds",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsAverageIOPSizeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_average_iop_size"),
		"The average size in bytes of recent I/O to the volume in the last 500 milliseconds",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsBurstIOPSCreditDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_burst_iops_credit"),
		"The total number of IOP credits available to the user. When volumes are not using up to the configured maxIOPS, credits are accrued.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsClientQueueDepthDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_client_queue_depth"),
		"The number of outstanding read and write operations to the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsLatencyUSecDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_latency_usec"),
		"The average time, in microseconds, to complete operations to the volume in the last 500 milliseconds. A '0' (zero) value means there is no I/O to the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsNonZeroBlocksDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_non_zero_blocks"),
		"The total number of 4KiB blocks that contain data after the last garbage collection operation has completed.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsReadBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_read_bytes"),
		"The total cumulative bytes read from the volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsReadBytesLastSampleDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_read_bytes_last_sample"),
		"The total number of bytes read from the volume during the last sample period.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsReadLatencyUSecDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_read_latency_usec"),
		"The average time, in microseconds, to complete read operations to the volume in the last 500 milliseconds.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsReadLatencyUSecTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_read_latency_usec_total"),
		"The total time spent performing read operations from the volume",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsReadOpsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_read_ops"),
		"The total read operations to the volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsReadOpsLastSampleDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_read_ops_last_sample"),
		"The total number of read operations during the last sample period",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsThrottleDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_throttle"),
		"A floating value between 0 and 1 that represents how much the system is throttling clients below their maxIOPS because of rereplication of data, transient errors, and snapshots taken.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsUnalignedReadsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_unaligned_reads"),
		"The total cumulative unaligned read operations to a volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsUnalignedWritesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_unaligned_writes"),
		"The total cumulative unaligned write operations to a volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsVolumeSizeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_size"),
		"Total provisioned capacity in bytes.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsVolumeUtilizationDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_utilization"),
		"A floating value that describes how much the client is using the volume. Value 0: The client is not using the volume. Value 1: The client is using their maximum. Value 1+: The client is using their burst.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsWriteBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_write_bytes"),
		"The total cumulative bytes written to the volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsWriteBytesLastSampleDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_write_bytes_last_sample"),
		"The total number of bytes written to the volume during the last sample period.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsWriteLatencyUSecDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_write_latency_usec"),
		"The average time, in microseconds, to complete write operations to a volume in the last 500 milliseconds.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsWriteLatencyUSecTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_write_latency_usec_total"),
		"The total time spent performing write operations to the volume",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsWriteOpsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_write_ops"),
		"The total cumulative write operations to the volume since the creation of the volume.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsWriteOpsLastSampleDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_write_ops_last_sample"),
		"The total number of write operations during the last sample period.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.VolumeStatsZeroBlocksDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "volume_stats_zero_blocks"),
		"The total number of empty 4KiB blocks without data after the last round of garbage collection operation has completed.",
		[]string{"volume_id", "volume_name"},
		nil,
	)

	d.ClusterCapacityActiveBlockSpace = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_active_block_space"),
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
	d.ClusterCapacityClusterRecentIOSize = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_cluster_recent_io_size"),
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
	d.ClusterCapacityMaxOverProvisionableSpace = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_max_over_provisionable_space"),
		"The maximum amount of provisionable space. This is a computed value. You cannot create new volumes if the current provisioned space plus the new volume size would exceed this number. The value is calculated as follows: maxOverProvisionableSpace = maxProvisionedSpace * maxMetadataOverProvisionFactor",
		nil,
		nil,
	)
	d.ClusterCapacityMaxProvisionedSpace = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_max_provisioned_space"),
		"The total amount of provisionable space if all volumes are 100% filled (no thin provisioned metadata)",
		nil,
		nil,
	)
	d.ClusterCapacityMaxUsedMetadataSpace = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_max_used_metadata_space"),
		"The number of bytes on volume drives used to store metadata",
		nil,
		nil,
	)
	d.ClusterCapacityMaxUsedSpace = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_max_used_space"),
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
	d.ClusterCapacityProvisionedSpace = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_provisioned_space"),
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

	d.ClusterCapacityTotalOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_total_ops"),
		"The total number of I/O operations performed throughout the lifetime of the cluster",
		nil,
		nil,
	)
	d.ClusterCapacityUniqueBlocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_unique_blocks"),
		"The total number of blocks stored on the block drives The value includes replicated blocks",
		nil,
		nil,
	)
	d.ClusterCapacityUniqueBlocksUsedSpace = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_unique_blocks_used_space"),
		"The total amount of data the uniqueBlocks take up on the block drives",
		nil,
		nil,
	)
	d.ClusterCapacityUsedMetadataSpace = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_used_metadata_space"),
		"The total number of bytes on volume drives used to store metadata",
		nil,
		nil,
	)
	d.ClusterCapacityUsedMetadataSpaceInSnapshots = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_used_metadata_space_in_snapshots"),
		"The number of bytes on volume drives used for storing unique data in snapshots. This number provides an estimate of how much metadata space would be regained by deleting all snapshots on the system",
		nil,
		nil,
	)
	d.ClusterCapacityUsedSpace = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_capacity_used_space"),
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

	d.ClusterActiveFaultsWarning = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_active_faults_warning"),
		"The total number of warning faults in the system",
		nil,
		nil,
	)

	d.ClusterActiveFaultsError = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_active_faults_error"),
		"The total number of error faults in the system",
		nil,
		nil,
	)

	d.ClusterActiveFaultsCritical = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_active_faults_critical"),
		"The total number of critical faults in the system",
		nil,
		nil,
	)

	d.ClusterActiveFaultsBestPractice = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_active_faults_best_practice"),
		"The total number of best practice faults in the system",
		nil,
		nil,
	)

	d.NodeStatsCBytesIn = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_cbytes_in"),
		"Bytes in on the cluster interface",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsCBytesOut = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_cbytes_out"),
		"Bytes out on the cluster interface",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_count"),
		"Node stat sample count", // Undocumented metric
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsCPUPercentage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_cpu_percentage"),
		"CPU usage in percent.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsCPUTotalSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_cpu_total_seconds"),
		"CPU usage in seconds since last boot.", // undocumented metric
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsMBytesIn = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_mbytes_in"),
		"Bytes in on the management interface.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsMBytesOut = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_mbytes_out"),
		"Bytes out on the management interface.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsNetworkUtilizationCluster = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_network_utilization_cluster_percentage"),
		"Network interface utilization (in percent) for the cluster network interface.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsNetworkUtilizationStorage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_network_utilization_storage_percentage"),
		"Network interface utilization (in percent) for the storage network interface.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsReadLatencyUSecTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_read_latency_usec_total"),
		"The number, in milliseconds, of read latency between clusters.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsReadOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_read_ops"),
		"Read Operations", // undocumented metric
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsWriteOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_write_ops"),
		"Write Operations", // undocumented metric
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsSBytesIn = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_sbytes_in"),
		"Bytes in on the storage interface.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsSBytesOut = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_sbytes_out"),
		"Bytes out on the storage interface.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsUsedMemory = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_used_memory"),
		"Total memory usage in bytes.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsWriteLatencyUSecTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_write_latency_usec_total"),
		"The number, in milliseconds, of read latency between clusters.",
		[]string{"node_id", "node_name"},
		nil,
	)

	d.NodeStatsLoadHistogram = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_load_histogram"),
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

	d.Node = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "node_name"),
		"Cluster node name",
		[]string{"node_id", "node_name", "chassis_name"},
		nil,
	)

	d.ClusterStatsActualIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_actual_iops"),
		"Current actual IOPS for the entire cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterStatsAverageIOPSize = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_average_iops_size"),
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

	d.ClusterStatsClusterUtilization = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_utilization"),
		"The cluster capacity being utilized.",
		nil,
		nil,
	)

	d.ClusterStatsLatencyUSec = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_latency_usec"),
		"The average time, in microseconds, to complete operations to a cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterStatsNormalizedIOPS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_normalized_iops"),
		"Average number of IOPS for the entire cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterStatsReadBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_read_bytes"),
		"The total cumulative bytes read from the cluster since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsReadBytesLastSample = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_read_bytes_last_sample"),
		"The total number of bytes read from the cluster during the last sample period.",
		nil,
		nil,
	)

	d.ClusterStatsReadLatencyUSec = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_read_latency_usec"),
		"The average time, in microseconds, to complete read operations to the cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterStatsReadLatencyUSecTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_read_latency_usec_total"),
		"The total time spent performing read operations since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsReadOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_read_ops"),
		"The total cumulative read operations to the cluster since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsReadOpsLastSample = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_read_ops_last_sample"),
		"The total number of read operations during the last sample period.",
		nil,
		nil,
	)

	d.ClusterStatsSamplePeriodMsec = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_sample_period_msec"),
		"The length of the sample period, in milliseconds.",
		nil,
		nil,
	)

	d.ClusterStatsServicesCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_services_count"),
		"The number of services running on the cluster. If equal to the servicesTotal, this indicates that valid statistics were collected from all nodes.",
		nil,
		nil,
	)

	d.ClusterStatsServicesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_services_total"),
		"The total number of expected services running on the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsUnalignedReads = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_unaligned_reads"),
		"The total cumulative unaligned read operations to a cluster since the creation of the cluster",
		nil,
		nil,
	)

	d.ClusterStatsUnalignedWrites = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_unaligned_writes"),
		"The total cumulative unaligned write operations to a cluster since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsWriteBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_write_bytes"),
		"The total cumulative bytes written to the cluster since the creation of the cluster",
		nil,
		nil,
	)

	d.ClusterStatsWriteBytesLastSample = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_write_bytes_last_sample"),
		"The total number of bytes written to the cluster during the last sample period.",
		nil,
		nil,
	)

	d.ClusterStatsWriteLatencyUSec = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_write_latency_usec"),
		"The average time, in microseconds, to complete write operations to a cluster in the last 500 milliseconds.",
		nil,
		nil,
	)

	d.ClusterStatsWriteLatencyUSecTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_write_latency_usec_total"),
		"The total time spent performing write operations since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsWriteOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_write_ops"),
		"The total cumulative write operations to the cluster since the creation of the cluster.",
		nil,
		nil,
	)

	d.ClusterStatsWriteOpsLastSample = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_stats_write_ops_last_sample"),
		"The total number of write operations during the last sample period.",
		nil,
		nil,
	)

	d.ClusterThresholdBlockFullness = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_block_fullness"),
		"The current computed level of block fullness of the cluster.",
		[]string{"level"},
		nil,
	)

	d.ClusterThresholdFullness = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_fullness"),
		"Reflects the highest level of fullness between 'blockFullness' and 'metadataFullness'.",
		[]string{"level"},
		nil,
	)

	d.ClusterThresholdMaxMetadataOverProvisionFactor = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_max_metadata_over_provision_factor"),
		"A value representative of the number of times metadata space can be over provisioned relative to the amount of space available.",
		nil,
		nil,
	)

	d.ClusterThresholdMetadataFullness = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_metadata_fullness"),
		"The current computed level of metadata fullness of the cluster.",
		[]string{"level"},
		nil,
	)

	d.ClusterThresholdSliceReserveUsedThresholdPct = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_slice_reserve_used_threshold_percentage"),
		"Error condition. A system alert is triggered if the reserved slice utilization is greater than the sliceReserveUsedThresholdPct value returned.",
		nil,
		nil,
	)

	d.ClusterThresholdStage2AwareThreshold = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_stage2_aware_threshold"),
		"Awareness condition. The value that is set for 'Stage 2' cluster threshold level.",
		nil,
		nil,
	)

	d.ClusterThresholdStage2BlockThresholdBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_stage2_block_threshold_bytes"),
		"Number of bytes being used by the cluster at which a stage2 condition will exist.",
		nil,
		nil,
	)

	d.ClusterThresholdStage3BlockThresholdBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_stage3_block_threshold_bytes"),
		"Number of bytes being used by the cluster at which a stage3 condition will exist.",
		nil,
		nil,
	)

	d.ClusterThresholdStage3BlockThresholdPercent = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_stage3_block_threshold_percentage"),
		"Percent value set for stage3. At this percent full, a warning is posted in the Alerts log.",
		nil,
		nil,
	)

	d.ClusterThresholdStage3LowThreshold = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_stage3_low_threshold"),
		"Error condition. The threshold at which a system alert is created due to low capacity on a cluster",
		nil,
		nil,
	)

	d.ClusterThresholdStage4BlockThresholdBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_stage4_block_threshold_bytes"),
		"Number of bytes being used by the cluster at which a stage4 condition will exist",
		nil,
		nil,
	)

	d.ClusterThresholdStage4CriticalThreshold = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_stage4_critical_threshold"),
		"Error condition. The threshold at which a system alert is created to warn about critically low capacity on a cluster.",
		nil,
		nil,
	)

	d.ClusterThresholdStage5BlockThresholdBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_stage5_block_threshold_bytes"),
		"The number of bytes being used by the cluster at which a stage5 condition will exist.",
		nil,
		nil,
	)

	d.ClusterThresholdSumTotalClusterBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_sum_total_cluster_bytes"),
		"Physical capacity of the cluster, measured in bytes.",
		nil,
		nil,
	)

	d.ClusterThresholdSumTotalMetadataClusterBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_sum_total_metadata_cluster_bytes"),
		"Total amount of space that can be used to store metadata",
		nil,
		nil,
	)

	d.ClusterThresholdSumUsedClusterBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_sum_used_cluster_bytes"),
		"Number of bytes used on the cluster.",
		nil,
		nil,
	)

	d.ClusterThresholdSumUsedMetadataClusterBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_threshold_sum_used_metadata_cluster_bytes"),
		"Amount of space used on volume drives to store metadata.",
		nil,
		nil,
	)

	d.ListDrivesStatus = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "list_drives_status"),
		"The drive staus for each individual drives in the cluster's active nodes",
		[]string{"node_id", "node_name", "drive_id", "serial", "slot", "status"},
		nil,
	)

	d.ListDrivesType = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "list_drives_types"),
		"The drive type for each individual drives in the cluster's active nodes",
		[]string{"node_id", "node_name", "drive_id", "serial", "slot", "status"},
		nil,
	)

	d.ListDrivesCapacity = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "list_drives_capacity"),
		"The drive capacity for each individual drives in the cluster's active nodes",
		[]string{"node_id", "node_name", "drive_id", "serial", "slot"},
		nil,
	)

	return &d
}
