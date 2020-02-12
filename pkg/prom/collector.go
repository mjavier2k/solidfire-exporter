package prom

import (
	"strconv"
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
	volumeNamesMux     sync.Mutex
)

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

}

func (c *solidfireCollector) Collect(ch chan<- prometheus.Metric) {
	var scrapeSuccess float64 = 1

	volumes, err := c.client.ListVolumes()
	if err != nil {
		scrapeSuccess = 0
		log.Errorln(err)
	}

	for _, vol := range volumes.Result.Volumes {
		volumeNamesMux.Lock()
		volumeNamesByID[vol.VolumeID] = vol.Name
		volumeNamesMux.Unlock()
	}

	// log.Infof("%+v", volumeNamesByID)

	volumeStats, err := c.client.ListVolumeStats()
	if err != nil {
		scrapeSuccess = 0
		log.Errorln(err)
	}
	// log.Infof("%+v", details)

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
		node_ssload_histogram_data := map[float64]uint64{
			0:   stats.SsLoadHistogram.Bucket0,
			19:  stats.SsLoadHistogram.Bucket1To19,
			39:  stats.SsLoadHistogram.Bucket20To39,
			59:  stats.SsLoadHistogram.Bucket40To59,
			79:  stats.SsLoadHistogram.Bucket60To79,
			100: stats.SsLoadHistogram.Bucket80To100,
		}

		var sum uint64 = 0
		for _, val := range node_ssload_histogram_data {
			sum += val
		}
		ch <- prometheus.MustNewConstHistogram(
			MetricDescriptions.NodeStatsLoadHistogram,
			stats.Count,
			float64(sum),
			node_ssload_histogram_data,
			strconv.Itoa(stats.NodeID),
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsCBytesIn,
			prometheus.GaugeValue,
			stats.CBytesIn,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsCBytesOut,
			prometheus.GaugeValue,
			stats.CBytesOut,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsCount,
			prometheus.GaugeValue,
			float64(stats.Count),
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsCPUPercentage,
			prometheus.GaugeValue,
			stats.CPU,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsCPUTotalSeconds,
			prometheus.GaugeValue,
			stats.CPUTotal,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsMBytesIn,
			prometheus.GaugeValue,
			stats.MBytesIn,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsMBytesOut,
			prometheus.GaugeValue,
			stats.MBytesOut,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsNetworkUtilizationCluster,
			prometheus.GaugeValue,
			stats.NetworkUtilizationCluster,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsNetworkUtilizationStorage,
			prometheus.GaugeValue,
			stats.NetworkUtilizationStorage,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsReadLatencyUSecTotal,
			prometheus.GaugeValue,
			stats.ReadLatencyUSecTotal,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsReadOps,
			prometheus.GaugeValue,
			stats.ReadOps,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsSBytesIn,
			prometheus.GaugeValue,
			stats.SBytesIn,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsSBytesOut,
			prometheus.GaugeValue,
			stats.SBytesOut,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsUsedMemory,
			prometheus.GaugeValue,
			stats.UsedMemory,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsWriteLatencyUSecTotal,
			prometheus.GaugeValue,
			stats.WriteLatencyUSecTotal,
			strconv.Itoa(stats.NodeID))

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsWriteOps,
			prometheus.GaugeValue,
			stats.WriteOps,
			strconv.Itoa(stats.NodeID))
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
