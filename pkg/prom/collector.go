package prom

import (
	"math"
	"strconv"
	"strings"
	"sync"

	log "github.com/amoghe/distillog"
	"github.com/mjavier2k/solidfire-exporter/pkg/solidfire"

	"github.com/prometheus/client_golang/prometheus"
)

type solidfireCollector struct {
	client *solidfire.Client
}

var (
	MetricDescriptions = NewMetricDescriptions("solidfire")
	volumeNamesByID    = make(map[int]string)
	nodesNamesByID     = make(map[int]string)
	volumeNamesMux     sync.Mutex
	nodeNamesMux       sync.Mutex
)

func sumHistogram(m map[float64]uint64) (r uint64) {
	r = 0
	for _, val := range m {
		r += val
	}
	return
}

func strCompare(str1 string, str2 string) int {
	if strings.Compare(strings.ToLower(str1), strings.ToLower(str2)) == 0 {
		return 1
	}
	return 0
}

func (c *solidfireCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- MetricDescriptions.ScrapeSuccessDesc

	ch <- MetricDescriptions.VolumeStatsActualIOPSDesc
	ch <- MetricDescriptions.VolumeStatsAverageIOPSizeDesc
	ch <- MetricDescriptions.VolumeStatsBurstIOPSCreditDesc
	ch <- MetricDescriptions.VolumeStatsClientQueueDepthDesc
	ch <- MetricDescriptions.VolumeStatsLatencyUSecDesc
	ch <- MetricDescriptions.VolumeStatsNonZeroBlocksDesc
	ch <- MetricDescriptions.VolumeStatsReadBytesDesc
	ch <- MetricDescriptions.VolumeStatsReadBytesLastSampleDesc
	ch <- MetricDescriptions.VolumeStatsReadLatencyUSecDesc
	ch <- MetricDescriptions.VolumeStatsReadLatencyUSecTotalDesc
	ch <- MetricDescriptions.VolumeStatsReadOpsDesc
	ch <- MetricDescriptions.VolumeStatsReadOpsLastSampleDesc
	ch <- MetricDescriptions.VolumeStatsThrottleDesc
	ch <- MetricDescriptions.VolumeStatsUnalignedReadsDesc
	ch <- MetricDescriptions.VolumeStatsUnalignedWritesDesc
	ch <- MetricDescriptions.VolumeStatsVolumeSizeDesc
	ch <- MetricDescriptions.VolumeStatsVolumeUtilizationDesc
	ch <- MetricDescriptions.VolumeStatsWriteBytesDesc
	ch <- MetricDescriptions.VolumeStatsWriteBytesLastSampleDesc
	ch <- MetricDescriptions.VolumeStatsWriteLatencyUSecDesc
	ch <- MetricDescriptions.VolumeStatsWriteLatencyUSecTotalDesc
	ch <- MetricDescriptions.VolumeStatsWriteOpsDesc
	ch <- MetricDescriptions.VolumeStatsWriteOpsLastSampleDesc
	ch <- MetricDescriptions.VolumeStatsZeroBlocksDesc

	ch <- MetricDescriptions.ClusterCapacityActiveBlockSpace
	ch <- MetricDescriptions.ClusterCapacityActiveSessions
	ch <- MetricDescriptions.ClusterCapacityAverageIOPS
	ch <- MetricDescriptions.ClusterCapacityClusterRecentIOSize
	ch <- MetricDescriptions.ClusterCapacityCurrentIOPS
	ch <- MetricDescriptions.ClusterCapacityMaxIOPS
	ch <- MetricDescriptions.ClusterCapacityMaxOverProvisionableSpace
	ch <- MetricDescriptions.ClusterCapacityMaxProvisionedSpace
	ch <- MetricDescriptions.ClusterCapacityMaxUsedMetadataSpace
	ch <- MetricDescriptions.ClusterCapacityMaxUsedSpace
	ch <- MetricDescriptions.ClusterCapacityNonZeroBlocks
	ch <- MetricDescriptions.ClusterCapacityPeakActiveSessions
	ch <- MetricDescriptions.ClusterCapacityPeakIOPS
	ch <- MetricDescriptions.ClusterCapacityProvisionedSpace
	ch <- MetricDescriptions.ClusterCapacitySnapshotNonZeroBlocks
	ch <- MetricDescriptions.ClusterCapacityTotalOps
	ch <- MetricDescriptions.ClusterCapacityUniqueBlocks
	ch <- MetricDescriptions.ClusterCapacityUniqueBlocksUsedSpace
	ch <- MetricDescriptions.ClusterCapacityUsedMetadataSpace
	ch <- MetricDescriptions.ClusterCapacityUsedMetadataSpaceInSnapshots
	ch <- MetricDescriptions.ClusterCapacityUsedSpace
	ch <- MetricDescriptions.ClusterCapacityZeroBlocks
	ch <- MetricDescriptions.ClusterCapacityThinProvisioningFactor
	ch <- MetricDescriptions.ClusterCapacityDeDuplicationFactor
	ch <- MetricDescriptions.ClusterCapacityCompressionFactor
	ch <- MetricDescriptions.ClusterCapacityEfficiencyFactor

	ch <- MetricDescriptions.ClusterActiveFaultsBestPractice
	ch <- MetricDescriptions.ClusterActiveFaultsWarning
	ch <- MetricDescriptions.ClusterActiveFaultsError
	ch <- MetricDescriptions.ClusterActiveFaultsCritical

	ch <- MetricDescriptions.NodeStatsCBytesIn
	ch <- MetricDescriptions.NodeStatsCBytesOut
	ch <- MetricDescriptions.NodeStatsCount
	ch <- MetricDescriptions.NodeStatsCPUPercentage
	ch <- MetricDescriptions.NodeStatsCPUTotalSeconds
	ch <- MetricDescriptions.NodeStatsMBytesIn
	ch <- MetricDescriptions.NodeStatsMBytesOut
	ch <- MetricDescriptions.NodeStatsNetworkUtilizationCluster
	ch <- MetricDescriptions.NodeStatsNetworkUtilizationStorage
	ch <- MetricDescriptions.NodeStatsReadLatencyUSecTotal
	ch <- MetricDescriptions.NodeStatsReadOps
	ch <- MetricDescriptions.NodeStatsSBytesIn
	ch <- MetricDescriptions.NodeStatsSBytesOut
	ch <- MetricDescriptions.NodeStatsUsedMemory
	ch <- MetricDescriptions.NodeStatsWriteLatencyUSecTotal
	ch <- MetricDescriptions.NodeStatsWriteOps
	ch <- MetricDescriptions.NodeStatsLoadHistogram

	ch <- MetricDescriptions.Node

	ch <- MetricDescriptions.ClusterStatsActualIOPS
	ch <- MetricDescriptions.ClusterStatsAverageIOPSize
	ch <- MetricDescriptions.ClusterStatsClientQueueDepth
	ch <- MetricDescriptions.ClusterStatsClusterUtilization
	ch <- MetricDescriptions.ClusterStatsLatencyUSec
	ch <- MetricDescriptions.ClusterStatsNormalizedIOPS
	ch <- MetricDescriptions.ClusterStatsReadBytes
	ch <- MetricDescriptions.ClusterStatsReadBytesLastSample
	ch <- MetricDescriptions.ClusterStatsReadLatencyUSec
	ch <- MetricDescriptions.ClusterStatsReadLatencyUSecTotal
	ch <- MetricDescriptions.ClusterStatsReadOps
	ch <- MetricDescriptions.ClusterStatsReadOpsLastSample
	ch <- MetricDescriptions.ClusterStatsSamplePeriodMsec
	ch <- MetricDescriptions.ClusterStatsServicesCount
	ch <- MetricDescriptions.ClusterStatsServicesTotal
	ch <- MetricDescriptions.ClusterStatsUnalignedReads
	ch <- MetricDescriptions.ClusterStatsUnalignedWrites
	ch <- MetricDescriptions.ClusterStatsWriteBytes
	ch <- MetricDescriptions.ClusterStatsWriteBytesLastSample
	ch <- MetricDescriptions.ClusterStatsWriteLatencyUSec
	ch <- MetricDescriptions.ClusterStatsWriteLatencyUSecTotal
	ch <- MetricDescriptions.ClusterStatsWriteOps
	ch <- MetricDescriptions.ClusterStatsWriteOpsLastSample

	ch <- MetricDescriptions.ClusterThresholdBlockFullness
	ch <- MetricDescriptions.ClusterThresholdFullness
	ch <- MetricDescriptions.ClusterThresholdMaxMetadataOverProvisionFactor
	ch <- MetricDescriptions.ClusterThresholdMetadataFullness
	ch <- MetricDescriptions.ClusterThresholdSliceReserveUsedThresholdPct
	ch <- MetricDescriptions.ClusterThresholdStage2AwareThreshold
	ch <- MetricDescriptions.ClusterThresholdStage2BlockThresholdBytes
	ch <- MetricDescriptions.ClusterThresholdStage3BlockThresholdBytes
	ch <- MetricDescriptions.ClusterThresholdStage3BlockThresholdPercent
	ch <- MetricDescriptions.ClusterThresholdStage3LowThreshold
	ch <- MetricDescriptions.ClusterThresholdStage4BlockThresholdBytes
	ch <- MetricDescriptions.ClusterThresholdStage4CriticalThreshold
	ch <- MetricDescriptions.ClusterThresholdStage5BlockThresholdBytes
	ch <- MetricDescriptions.ClusterThresholdSumTotalClusterBytes
	ch <- MetricDescriptions.ClusterThresholdSumTotalMetadataClusterBytes
	ch <- MetricDescriptions.ClusterThresholdSumUsedClusterBytes
	ch <- MetricDescriptions.ClusterThresholdSumUsedMetadataClusterBytes

	ch <- MetricDescriptions.ListDrivesStatus
	ch <- MetricDescriptions.ListDrivesType
	ch <- MetricDescriptions.ListDrivesCapacity
}

func (c *solidfireCollector) Collect(ch chan<- prometheus.Metric) {
	var scrapeSuccess float64 = 1

	volumes, err := c.client.ListVolumes()
	if err != nil {
		scrapeSuccess = 0
		log.Errorln(err)
	}

	nodes, err := c.client.ListAllNodes()
	if err != nil {
		scrapeSuccess = 0
		log.Errorln(err)
	}

	for _, vol := range volumes.Result.Volumes {
		volumeNamesMux.Lock()
		volumeNamesByID[vol.VolumeID] = vol.Name
		volumeNamesMux.Unlock()
	}

	for _, node := range nodes.Result.Nodes {
		nodeNamesMux.Lock()
		nodesNamesByID[node.NodeID] = node.Name
		nodeNamesMux.Unlock()

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.Node,
			prometheus.CounterValue,
			1,
			strconv.Itoa(node.NodeID),
			node.Name,
			node.ChassisName,
		)
	}

	volumeStats, err := c.client.ListVolumeStats()
	if err != nil {
		scrapeSuccess = 0
		log.Errorln(err)
	}

	for _, vol := range volumeStats.Result.VolumeStats {
		volumeNamesMux.Lock()
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsActualIOPSDesc,
			prometheus.GaugeValue,
			vol.ActualIOPS,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsAverageIOPSizeDesc,
			prometheus.GaugeValue,
			vol.AverageIOPSize,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsBurstIOPSCreditDesc,
			prometheus.GaugeValue,
			vol.BurstIOPSCredit,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsClientQueueDepthDesc,
			prometheus.GaugeValue,
			vol.ClientQueueDepth,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsLatencyUSecDesc,
			prometheus.GaugeValue,
			vol.LatencyUSec,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsNonZeroBlocksDesc,
			prometheus.GaugeValue,
			vol.NonZeroBlocks,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsReadBytesDesc,
			prometheus.CounterValue,
			vol.ReadBytes,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsReadBytesLastSampleDesc,
			prometheus.GaugeValue,
			vol.ReadBytesLastSample,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsReadLatencyUSecDesc,
			prometheus.GaugeValue,
			vol.ReadLatencyUSec,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsReadLatencyUSecTotalDesc,
			prometheus.CounterValue,
			vol.ReadLatencyUSecTotal,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsReadOpsDesc,
			prometheus.CounterValue,
			vol.ReadOps,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsReadOpsLastSampleDesc,
			prometheus.GaugeValue,
			vol.ReadOpsLastSample,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsThrottleDesc,
			prometheus.GaugeValue,
			vol.Throttle,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsUnalignedReadsDesc,
			prometheus.CounterValue,
			vol.UnalignedReads,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsUnalignedWritesDesc,
			prometheus.CounterValue,
			vol.UnalignedWrites,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsVolumeSizeDesc,
			prometheus.GaugeValue,
			vol.VolumeSize,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsVolumeUtilizationDesc,
			prometheus.GaugeValue,
			vol.VolumeUtilization,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsWriteBytesDesc,
			prometheus.CounterValue,
			vol.WriteBytes,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsWriteBytesLastSampleDesc,
			prometheus.GaugeValue,
			vol.WriteBytesLastSample,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsWriteLatencyUSecDesc,
			prometheus.GaugeValue,
			vol.WriteLatencyUSec,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsWriteLatencyUSecTotalDesc,
			prometheus.CounterValue,
			vol.WriteLatencyUSecTotal,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsWriteOpsDesc,
			prometheus.CounterValue,
			vol.WriteOps,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsWriteOpsLastSampleDesc,
			prometheus.GaugeValue,
			vol.WriteOpsLastSample,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsZeroBlocksDesc,
			prometheus.GaugeValue,
			vol.ZeroBlocks,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		volumeNamesMux.Unlock()
	}

	clusterCapacity, err := c.client.GetClusterCapacity()
	if err != nil {
		scrapeSuccess = 0
		log.Errorln(err)
	}
	cluster := clusterCapacity.Result.ClusterCapacity
	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityActiveBlockSpace,
		prometheus.GaugeValue,
		cluster.ActiveBlockSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityActiveSessions,
		prometheus.GaugeValue,
		cluster.ActiveSessions)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityAverageIOPS,
		prometheus.GaugeValue,
		cluster.AverageIOPS)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityClusterRecentIOSize,
		prometheus.GaugeValue,
		cluster.ClusterRecentIOSize)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityCurrentIOPS,
		prometheus.GaugeValue,
		cluster.CurrentIOPS)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityMaxIOPS,
		prometheus.GaugeValue,
		cluster.MaxIOPS)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityMaxOverProvisionableSpace,
		prometheus.GaugeValue,
		cluster.MaxOverProvisionableSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityMaxProvisionedSpace,
		prometheus.GaugeValue,
		cluster.MaxProvisionedSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityMaxUsedMetadataSpace,
		prometheus.GaugeValue,
		cluster.MaxUsedMetadataSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityMaxUsedSpace,
		prometheus.GaugeValue,
		cluster.MaxUsedSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityNonZeroBlocks,
		prometheus.GaugeValue,
		cluster.NonZeroBlocks)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityPeakActiveSessions,
		prometheus.GaugeValue,
		cluster.PeakActiveSessions)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityPeakIOPS,
		prometheus.GaugeValue,
		cluster.PeakIOPS)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityProvisionedSpace,
		prometheus.GaugeValue,
		cluster.ProvisionedSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacitySnapshotNonZeroBlocks,
		prometheus.GaugeValue,
		cluster.SnapshotNonZeroBlocks)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityTotalOps,
		prometheus.CounterValue,
		cluster.TotalOps)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityUniqueBlocks,
		prometheus.GaugeValue,
		cluster.UniqueBlocks)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityUniqueBlocksUsedSpace,
		prometheus.GaugeValue,
		cluster.UniqueBlocksUsedSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityUsedMetadataSpace,
		prometheus.GaugeValue,
		cluster.UsedMetadataSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityUsedMetadataSpaceInSnapshots,
		prometheus.GaugeValue,
		cluster.UsedMetadataSpaceInSnapshots)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityUsedSpace,
		prometheus.GaugeValue,
		cluster.UsedSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityZeroBlocks,
		prometheus.GaugeValue,
		cluster.ZeroBlocks)

	clusterThinProvisioningFactor := (cluster.NonZeroBlocks + cluster.ZeroBlocks) / cluster.NonZeroBlocks
	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityThinProvisioningFactor,
		prometheus.GaugeValue,
		clusterThinProvisioningFactor)

	clusterDeDuplicationFactor := (cluster.NonZeroBlocks + cluster.SnapshotNonZeroBlocks) / cluster.UniqueBlocks
	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityDeDuplicationFactor,
		prometheus.GaugeValue,
		clusterDeDuplicationFactor)

	clusterCompressionFactor := (cluster.UniqueBlocks * 4096) / (cluster.UniqueBlocksUsedSpace * 0.93)
	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityCompressionFactor,
		prometheus.GaugeValue,
		clusterCompressionFactor)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityEfficiencyFactor,
		prometheus.GaugeValue,
		clusterThinProvisioningFactor*clusterDeDuplicationFactor*clusterCompressionFactor)

	// List Cluster Faults
	ClusterActiveFaults, err := c.client.ListClusterActiveFaults()
	if err != nil {
		scrapeSuccess = 0
		log.Errorln(err)
	}
	severity := map[string]float64{
		solidfire.FaultBestPractice: 0,
		solidfire.FaultWarning:      0,
		solidfire.FaultError:        0,
		solidfire.FaultCritical:     0,
	}

	for _, f := range ClusterActiveFaults.Result.Faults {
		severity[f.Severity]++
	}

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterActiveFaultsBestPractice,
		prometheus.GaugeValue,
		severity[solidfire.FaultBestPractice])

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterActiveFaultsWarning,
		prometheus.GaugeValue,
		severity[solidfire.FaultWarning])

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterActiveFaultsError,
		prometheus.GaugeValue,
		severity[solidfire.FaultError])

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterActiveFaultsCritical,
		prometheus.GaugeValue,
		severity[solidfire.FaultCritical])

	// List Cluster Stats
	ClusterNodeStats, err := c.client.ListNodeStats()
	if err != nil {
		scrapeSuccess = 0
		log.Errorln(err)
	}

	for _, stats := range ClusterNodeStats.Result.NodeStats.Nodes {
		SsLoadHistogram := map[float64]uint64{
			0:   stats.SsLoadHistogram.Bucket0,
			19:  stats.SsLoadHistogram.Bucket1To19,
			39:  stats.SsLoadHistogram.Bucket20To39,
			59:  stats.SsLoadHistogram.Bucket40To59,
			79:  stats.SsLoadHistogram.Bucket60To79,
			100: stats.SsLoadHistogram.Bucket80To100,
		}

		ch <- prometheus.MustNewConstHistogram(
			MetricDescriptions.NodeStatsLoadHistogram,
			stats.Count,
			float64(sumHistogram(SsLoadHistogram)),
			SsLoadHistogram,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsCBytesIn,
			prometheus.GaugeValue,
			stats.CBytesIn,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsCBytesOut,
			prometheus.GaugeValue,
			stats.CBytesOut,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsCount,
			prometheus.GaugeValue,
			float64(stats.Count),
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsCPUPercentage,
			prometheus.GaugeValue,
			stats.CPU,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsCPUTotalSeconds,
			prometheus.GaugeValue,
			stats.CPUTotal,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsMBytesIn,
			prometheus.GaugeValue,
			stats.MBytesIn,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsMBytesOut,
			prometheus.GaugeValue,
			stats.MBytesOut,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsNetworkUtilizationCluster,
			prometheus.GaugeValue,
			stats.NetworkUtilizationCluster,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsNetworkUtilizationStorage,
			prometheus.GaugeValue,
			stats.NetworkUtilizationStorage,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsReadLatencyUSecTotal,
			prometheus.GaugeValue,
			stats.ReadLatencyUSecTotal,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsReadOps,
			prometheus.GaugeValue,
			stats.ReadOps,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsSBytesIn,
			prometheus.GaugeValue,
			stats.SBytesIn,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsSBytesOut,
			prometheus.GaugeValue,
			stats.SBytesOut,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsUsedMemory,
			prometheus.GaugeValue,
			stats.UsedMemory,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsWriteLatencyUSecTotal,
			prometheus.GaugeValue,
			stats.WriteLatencyUSecTotal,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsWriteOps,
			prometheus.GaugeValue,
			stats.WriteOps,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)
	}

	// ListVolumeQoSHistograms
	VolumeQoSHistograms, err := c.client.ListVolumeQoSHistograms()
	if err != nil {
		scrapeSuccess = 0
		log.Errorln(err)
	}

	for _, h := range VolumeQoSHistograms.Result.QosHistograms {

		// Below Min IOPS Percentage
		BelowMinIopsPercentages := map[float64]uint64{
			19:  h.Histograms.BelowMinIopsPercentages.Bucket1To19,
			39:  h.Histograms.BelowMinIopsPercentages.Bucket20To39,
			59:  h.Histograms.BelowMinIopsPercentages.Bucket40To59,
			79:  h.Histograms.BelowMinIopsPercentages.Bucket60To79,
			100: h.Histograms.BelowMinIopsPercentages.Bucket80To100,
		}

		ch <- prometheus.MustNewConstHistogram(
			MetricDescriptions.VolumeQoSBelowMinIopsPercentagesHistogram,
			0,
			float64(sumHistogram(BelowMinIopsPercentages)),
			BelowMinIopsPercentages,
			strconv.Itoa(h.VolumeID),
			volumeNamesByID[h.VolumeID],
		)

		MinToMaxIopsPercentages := map[float64]uint64{
			19:          h.Histograms.MinToMaxIopsPercentages.Bucket1To19,
			39:          h.Histograms.MinToMaxIopsPercentages.Bucket20To39,
			59:          h.Histograms.MinToMaxIopsPercentages.Bucket40To59,
			79:          h.Histograms.MinToMaxIopsPercentages.Bucket60To79,
			100:         h.Histograms.MinToMaxIopsPercentages.Bucket80To100,
			math.Inf(1): h.Histograms.MinToMaxIopsPercentages.Bucket101Plus,
		}

		ch <- prometheus.MustNewConstHistogram(
			MetricDescriptions.VolumeQoSMinToMaxIopsPercentagesHistogram,
			0,
			float64(sumHistogram(MinToMaxIopsPercentages)),
			MinToMaxIopsPercentages,
			strconv.Itoa(h.VolumeID),
			volumeNamesByID[h.VolumeID],
		)

		ReadBlockSizes := map[float64]uint64{
			8191:        h.Histograms.ReadBlockSizes.Bucket4096To8191,
			16383:       h.Histograms.ReadBlockSizes.Bucket8192To16383,
			32767:       h.Histograms.ReadBlockSizes.Bucket16384To32767,
			65535:       h.Histograms.ReadBlockSizes.Bucket32768To65535,
			131071:      h.Histograms.ReadBlockSizes.Bucket65536To131071,
			math.Inf(1): h.Histograms.ReadBlockSizes.Bucket131072Plus,
		}

		ch <- prometheus.MustNewConstHistogram(
			MetricDescriptions.VolumeQoSReadBlockSizesHistogram,
			0,
			float64(sumHistogram(ReadBlockSizes)),
			ReadBlockSizes,
			strconv.Itoa(h.VolumeID),
			volumeNamesByID[h.VolumeID],
		)

		TargetUtilizationPercentages := map[float64]uint64{
			0:           h.Histograms.TargetUtilizationPercentages.Bucket0,
			19:          h.Histograms.TargetUtilizationPercentages.Bucket1To19,
			39:          h.Histograms.TargetUtilizationPercentages.Bucket20To39,
			59:          h.Histograms.TargetUtilizationPercentages.Bucket40To59,
			79:          h.Histograms.TargetUtilizationPercentages.Bucket60To79,
			100:         h.Histograms.TargetUtilizationPercentages.Bucket80To100,
			math.Inf(1): h.Histograms.TargetUtilizationPercentages.Bucket101Plus,
		}

		ch <- prometheus.MustNewConstHistogram(
			MetricDescriptions.VolumeQoSTargetUtilizationPercentages,
			0,
			float64(sumHistogram(TargetUtilizationPercentages)),
			TargetUtilizationPercentages,
			strconv.Itoa(h.VolumeID),
			volumeNamesByID[h.VolumeID],
		)

		ThrottlePercentages := map[float64]uint64{
			0:   h.Histograms.ThrottlePercentages.Bucket0,
			19:  h.Histograms.ThrottlePercentages.Bucket1To19,
			39:  h.Histograms.ThrottlePercentages.Bucket20To39,
			59:  h.Histograms.ThrottlePercentages.Bucket40To59,
			79:  h.Histograms.ThrottlePercentages.Bucket60To79,
			100: h.Histograms.ThrottlePercentages.Bucket80To100,
		}

		ch <- prometheus.MustNewConstHistogram(
			MetricDescriptions.VolumeQoSThrottlePercentages,
			0,
			float64(sumHistogram(ThrottlePercentages)),
			ThrottlePercentages,
			strconv.Itoa(h.VolumeID),
			volumeNamesByID[h.VolumeID],
		)

		WriteBlockSizes := map[float64]uint64{
			8191:        h.Histograms.WriteBlockSizes.Bucket4096To8191,
			16383:       h.Histograms.WriteBlockSizes.Bucket8192To16383,
			32767:       h.Histograms.WriteBlockSizes.Bucket16384To32767,
			65535:       h.Histograms.WriteBlockSizes.Bucket32768To65535,
			131071:      h.Histograms.WriteBlockSizes.Bucket65536To131071,
			math.Inf(1): h.Histograms.WriteBlockSizes.Bucket131072Plus,
		}

		ch <- prometheus.MustNewConstHistogram(
			MetricDescriptions.VolumeQoSWriteBlockSizes,
			0,
			float64(sumHistogram(WriteBlockSizes)),
			WriteBlockSizes,
			strconv.Itoa(h.VolumeID),
			volumeNamesByID[h.VolumeID],
		)
	}

	clusterStats, err := c.client.GetClusterStats()
	if err != nil {
		scrapeSuccess = 0
		log.Errorln(err)
	}

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsActualIOPS,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ActualIOPS,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsAverageIOPSize,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.AverageIOPSize,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsClientQueueDepth,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ClientQueueDepth,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsClusterUtilization,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ClusterUtilization,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsLatencyUSec,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.LatencyUSec,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsNormalizedIOPS,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.NormalizedIOPS,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsReadBytes,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ReadBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsReadBytesLastSample,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ReadBytesLastSample,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsReadLatencyUSec,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ReadLatencyUSec,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsReadLatencyUSecTotal,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ReadLatencyUSecTotal,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsReadOps,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ReadOps,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsReadOpsLastSample,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ReadOpsLastSample,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsSamplePeriodMsec,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.SamplePeriodMsec,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsServicesCount,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ServicesCount,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsServicesTotal,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ServicesTotal,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsUnalignedReads,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.UnalignedReads,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsUnalignedWrites,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.UnalignedWrites,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsWriteBytes,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.WriteBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsWriteBytesLastSample,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.WriteBytesLastSample,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsWriteLatencyUSec,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.WriteLatencyUSec,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsWriteLatencyUSecTotal,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.WriteLatencyUSecTotal,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsWriteOps,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.WriteOps,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsWriteOpsLastSample,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.WriteOpsLastSample,
	)

	clusterFullThreshold, err := c.client.GetClusterFullThreshold()
	if err != nil {
		scrapeSuccess = 0
		log.Errorln(err)
	}

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdBlockFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.BlockFullness, "stage1Happy")),
		"stage1Happy",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdBlockFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.BlockFullness, "stage2Aware")),
		"stage2Aware",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdBlockFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.BlockFullness, "stage3Low")),
		"stage3Low",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdBlockFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.BlockFullness, "stage4Critical")),
		"stage4Critical",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdBlockFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.BlockFullness, "stage5CompletelyConsumed")),
		"stage5CompletelyConsumed",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.Fullness, "blockFullness")),
		"blockFullness",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.Fullness, "metadataFullness")),
		"metadataFullness",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdMaxMetadataOverProvisionFactor,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.MaxMetadataOverProvisionFactor,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdMetadataFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.MetadataFullness, "stage1Happy")),
		"stage1Happy",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdMetadataFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.MetadataFullness, "stage2Aware")),
		"stage2Aware",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdMetadataFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.MetadataFullness, "stage3Low")),
		"stage3Low",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdMetadataFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.MetadataFullness, "stage4Critical")),
		"stage4Critical",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdMetadataFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.MetadataFullness, "stage5CompletelyConsumed")),
		"stage5CompletelyConsumed",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdSliceReserveUsedThresholdPct,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.SliceReserveUsedThresholdPct,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdStage2AwareThreshold,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage2AwareThreshold,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdStage2BlockThresholdBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage2BlockThresholdBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdStage3BlockThresholdBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage3BlockThresholdBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdStage3BlockThresholdPercent,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage3BlockThresholdPercent,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdStage3LowThreshold,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage3LowThreshold,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdStage4BlockThresholdBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage4BlockThresholdBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdStage4CriticalThreshold,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage4CriticalThreshold,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdStage5BlockThresholdBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage5BlockThresholdBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdSumTotalClusterBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.SumTotalClusterBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdSumTotalMetadataClusterBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.SumTotalMetadataClusterBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdSumUsedClusterBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.SumUsedClusterBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterThresholdSumUsedMetadataClusterBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.SumUsedMetadataClusterBytes,
	)

	ListDrives, err := c.client.ListDrives()
	if err != nil {
		scrapeSuccess = 0
		log.Errorln(err)
	}

	for _, d := range ListDrives.Result.Drives {
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ListDrivesStatus,
			prometheus.GaugeValue,
			1,
			strconv.Itoa(d.NodeID),
			nodesNamesByID[d.NodeID],
			strconv.Itoa(d.DriveID),
			d.Serial,
			strconv.Itoa(d.Slot),
			d.Status,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ListDrivesType,
			prometheus.GaugeValue,
			1,
			strconv.Itoa(d.NodeID),
			nodesNamesByID[d.NodeID],
			strconv.Itoa(d.DriveID),
			d.Serial,
			strconv.Itoa(d.Slot),
			d.Type,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ListDrivesCapacity,
			prometheus.GaugeValue,
			d.Capacity,
			strconv.Itoa(d.NodeID),
			nodesNamesByID[d.NodeID],
			strconv.Itoa(d.DriveID),
			d.Serial,
			strconv.Itoa(d.Slot),
		)
	}

	// Set scrape success metric to scrapeSuccess
	ch <- prometheus.MustNewConstMetric(MetricDescriptions.ScrapeSuccessDesc, prometheus.GaugeValue, scrapeSuccess)
}

func NewCollector() (*solidfireCollector, error) {
	log.Infof("initializing new solidfire collector")
	return &solidfireCollector{
		client: solidfire.NewSolidfireClient(),
	}, nil
}
