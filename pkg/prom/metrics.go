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
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsCBytesOut = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_cbytes_out"),
		"Bytes out on the cluster interface",
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_count"),
		"Node stat sample count", // Undocumented metric
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsCPUPercentage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_cpu_percentage"),
		"CPU usage in percent.",
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsCPUTotalSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_cpu_total_seconds"),
		"CPU usage in seconds since last boot.", // undocumented metric
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsMBytesIn = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_mbytes_in"),
		"Bytes in on the management interface.",
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsMBytesOut = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_mbytes_out"),
		"Bytes out on the management interface.",
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsNetworkUtilizationCluster = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_network_utilization_cluster_percentage"),
		"Network interface utilization (in percent) for the cluster network interface.",
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsNetworkUtilizationStorage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_network_utilization_storage_percentage"),
		"Network interface utilization (in percent) for the storage network interface.",
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsReadLatencyUSecTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_read_latency_usec_total"),
		"The number, in milliseconds, of read latency between clusters.",
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsReadOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_read_ops"),
		"Read Operations", // undocumented metric
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsWriteOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_write_ops"),
		"Write Operations", // undocumented metric
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsSBytesIn = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_sbytes_in"),
		"Bytes in on the storage interface.",
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsSBytesOut = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_sbytes_out"),
		"Bytes out on the storage interface.",
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsUsedMemory = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_used_memory"),
		"Total memory usage in bytes.",
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsWriteLatencyUSecTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_write_latency_usec_total"),
		"The number, in milliseconds, of read latency between clusters.",
		[]string{"node_id"},
		nil,
	)

	d.NodeStatsLoadHistogram = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cluster_node_stats_load_histogram"),
		"System load histogram",
		[]string{"node_id"},
		nil,
	)

	return &d
}
