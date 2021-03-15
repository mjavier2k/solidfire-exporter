package prom

import (
	"fmt"
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
	MetricDescriptions    = NewMetricDescriptions("solidfire")
	volumeNamesByID       = make(map[int]string)
	nodesNamesByID        = make(map[int]string)
	volumeNamesMux        sync.Mutex
	nodeNamesMux          sync.Mutex
	possibleDriveStatuses = []string{"active", "available", "erasing", "failed", "removing"}
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
	ch <- MetricDescriptions.upDesc

	ch <- MetricDescriptions.VolumeStatsActualIOPS
	ch <- MetricDescriptions.VolumeStatsAverageIOPSizeBytes
	ch <- MetricDescriptions.VolumeStatsBurstIOPSCredit
	ch <- MetricDescriptions.VolumeStatsClientQueueDepth
	ch <- MetricDescriptions.VolumeStatsLatencySeconds
	ch <- MetricDescriptions.VolumeStatsNonZeroBlocks
	ch <- MetricDescriptions.VolumeStatsReadBytes
	ch <- MetricDescriptions.VolumeStatsLastSampleReadBytes
	ch <- MetricDescriptions.VolumeStatsReadLatencySeconds
	ch <- MetricDescriptions.VolumeStatsReadLatencySecondsTotal
	ch <- MetricDescriptions.VolumeStatsReadOpsTotal
	ch <- MetricDescriptions.VolumeStatsLastSampleReadOps
	ch <- MetricDescriptions.VolumeStatsThrottle
	ch <- MetricDescriptions.VolumeStatsUnalignedReadsTotal
	ch <- MetricDescriptions.VolumeStatsUnalignedWritesTotal
	ch <- MetricDescriptions.VolumeStatsVolumeSizeBytes
	ch <- MetricDescriptions.VolumeStatsVolumeUtilization
	ch <- MetricDescriptions.VolumeStatsWriteBytes
	ch <- MetricDescriptions.VolumeStatsLastSampleWriteBytes
	ch <- MetricDescriptions.VolumeStatsWriteLatencySeconds
	ch <- MetricDescriptions.VolumeStatsWriteLatencyTotal
	ch <- MetricDescriptions.VolumeStatsWriteOpsTotal
	ch <- MetricDescriptions.VolumeStatsWriteOpsLastSample
	ch <- MetricDescriptions.VolumeStatsZeroBlocks

	ch <- MetricDescriptions.ClusterCapacityActiveBlockSpaceBytes
	ch <- MetricDescriptions.ClusterCapacityActiveSessions
	ch <- MetricDescriptions.ClusterCapacityAverageIOPS
	ch <- MetricDescriptions.ClusterCapacityClusterRecentIOSizeBytes
	ch <- MetricDescriptions.ClusterCapacityCurrentIOPS
	ch <- MetricDescriptions.ClusterCapacityMaxIOPS
	ch <- MetricDescriptions.ClusterCapacityMaxOverProvisionableSpaceBytes
	ch <- MetricDescriptions.ClusterCapacityMaxProvisionedSpaceBytes
	ch <- MetricDescriptions.ClusterCapacityMaxUsedMetadataSpaceBytes
	ch <- MetricDescriptions.ClusterCapacityMaxUsedSpaceBytes
	ch <- MetricDescriptions.ClusterCapacityNonZeroBlocks
	ch <- MetricDescriptions.ClusterCapacityPeakActiveSessions
	ch <- MetricDescriptions.ClusterCapacityPeakIOPS
	ch <- MetricDescriptions.ClusterCapacityProvisionedSpaceBytes
	ch <- MetricDescriptions.ClusterCapacitySnapshotNonZeroBlocks
	ch <- MetricDescriptions.ClusterCapacityIOPSTotal
	ch <- MetricDescriptions.ClusterCapacityUniqueBlocks
	ch <- MetricDescriptions.ClusterCapacityUniqueBlocksUsedSpaceBytes
	ch <- MetricDescriptions.ClusterCapacityUsedMetadataSpaceBytes
	ch <- MetricDescriptions.ClusterCapacityUsedMetadataSpaceInSnapshotsBytes
	ch <- MetricDescriptions.ClusterCapacityUsedSpaceBytes
	ch <- MetricDescriptions.ClusterCapacityZeroBlocks
	ch <- MetricDescriptions.ClusterCapacityThinProvisioningFactor
	ch <- MetricDescriptions.ClusterCapacityDeDuplicationFactor
	ch <- MetricDescriptions.ClusterCapacityCompressionFactor
	ch <- MetricDescriptions.ClusterCapacityEfficiencyFactor

	ch <- MetricDescriptions.ClusterActiveFaults

	ch <- MetricDescriptions.NodeStatsCount
	ch <- MetricDescriptions.NodeStatsCPUPercentage
	ch <- MetricDescriptions.NodeStatsCPUTotalSeconds
	ch <- MetricDescriptions.NodeStatsInterfaceInBytesTotal
	ch <- MetricDescriptions.NodeStatsInterfaceOutBytesTotal
	ch <- MetricDescriptions.NodeStatsInterfaceUtilizationPercentage
	ch <- MetricDescriptions.NodeStatsReadLatencyTotal
	ch <- MetricDescriptions.NodeStatsReadOps
	ch <- MetricDescriptions.NodeStatsUsedMemoryBytes
	ch <- MetricDescriptions.NodeStatsWriteLatencyTotal
	ch <- MetricDescriptions.NodeStatsWriteOps
	ch <- MetricDescriptions.NodeStatsLoadHistogram

	ch <- MetricDescriptions.NodeInfo

	ch <- MetricDescriptions.ClusterStatsActualIOPS
	ch <- MetricDescriptions.ClusterStatsAverageIOBytes
	ch <- MetricDescriptions.ClusterStatsClientQueueDepth
	ch <- MetricDescriptions.ClusterStatsThroughputUtilization
	ch <- MetricDescriptions.ClusterStatsLatencySeconds
	ch <- MetricDescriptions.ClusterStatsNormalizedIOPS
	ch <- MetricDescriptions.ClusterStatsReadBytesTotal
	ch <- MetricDescriptions.ClusterStatsLastSampleReadBytes
	ch <- MetricDescriptions.ClusterStatsReadLatencySeconds
	ch <- MetricDescriptions.ClusterStatsReadLatencyTotal
	ch <- MetricDescriptions.ClusterStatsReadOpsTotal
	ch <- MetricDescriptions.ClusterStatsLastSampleReadOps
	ch <- MetricDescriptions.ClusterStatsSamplePeriodSeconds
	ch <- MetricDescriptions.ClusterStatsServices
	ch <- MetricDescriptions.ClusterStatsExpectedServices
	ch <- MetricDescriptions.ClusterStatsUnalignedReadsTotal
	ch <- MetricDescriptions.ClusterStatsUnalignedWrites
	ch <- MetricDescriptions.ClusterStatsWriteBytesTotal
	ch <- MetricDescriptions.ClusterStatsLastSampleWriteBytes
	ch <- MetricDescriptions.ClusterStatsWriteLatency
	ch <- MetricDescriptions.ClusterStatsWriteLatencyTotal
	ch <- MetricDescriptions.ClusterStatsWriteOpsTotal
	ch <- MetricDescriptions.ClusterStatsLastSampleWriteOps

	ch <- MetricDescriptions.ClusterBlockFullness
	ch <- MetricDescriptions.ClusterFullness
	ch <- MetricDescriptions.ClusterMaxMetadataOverProvisionFactor
	ch <- MetricDescriptions.ClusterMetadataFullness
	ch <- MetricDescriptions.ClusterSliceReserveUsedThresholdPercentage
	ch <- MetricDescriptions.ClusterStage2AwareThresholdPercentage
	ch <- MetricDescriptions.ClusterStage2BlockThresholdBytes
	ch <- MetricDescriptions.ClusterStage3BlockThresholdBytes
	ch <- MetricDescriptions.ClusterStage3BlockThresholdPercentage
	ch <- MetricDescriptions.ClusterStage3LowThresholdPercentage
	ch <- MetricDescriptions.ClusterStage4BlockThresholdBytes
	ch <- MetricDescriptions.ClusterStage4CriticalThreshold
	ch <- MetricDescriptions.ClusterStage5BlockThresholdBytes
	ch <- MetricDescriptions.ClusterTotalBytes
	ch <- MetricDescriptions.ClusterTotalMetadataBytes
	ch <- MetricDescriptions.ClusterUsedBytes
	ch <- MetricDescriptions.ClusterUsedMetadataBytes

	ch <- MetricDescriptions.DriveStatus
	ch <- MetricDescriptions.DriveCapacityBytes

	ch <- MetricDescriptions.NodeISCSISessions
}

func (c *solidfireCollector) Collect(ch chan<- prometheus.Metric) {
	var up float64 = 1

	volumes, err := c.client.ListVolumes()
	if err != nil {
		up = 0
		log.Errorln(err)
	}

	nodes, err := c.client.ListAllNodes()
	if err != nil {
		up = 0
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
			MetricDescriptions.NodeInfo,
			prometheus.GaugeValue,
			1,
			strconv.Itoa(node.NodeID),
			node.Name,
			node.ChassisName,
			strconv.Itoa(node.AssociatedFServiceID),
			strconv.Itoa(node.AssociatedMasterServiceID),
			node.PlatformInfo.ChassisType,
			node.PlatformInfo.CPUModel,
			node.PlatformInfo.NodeType,
			node.PlatformInfo.PlatformConfigVersion,
			node.Sip,
			node.Sipi,
			node.SoftwareVersion,
			node.UUID,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsTotalMemoryBytes,
			prometheus.GaugeValue,
			GBToBytes(node.PlatformInfo.NodeMemoryGB),
			strconv.Itoa(node.NodeID),
			node.Name,
		)
	}

	volumeStats, err := c.client.ListVolumeStats()
	if err != nil {
		up = 0
		log.Errorln(err)
	}

	for _, vol := range volumeStats.Result.VolumeStats {
		volumeNamesMux.Lock()
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsActualIOPS,
			prometheus.GaugeValue,
			vol.ActualIOPS,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsAverageIOPSizeBytes,
			prometheus.GaugeValue,
			vol.AverageIOPSize,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsBurstIOPSCredit,
			prometheus.GaugeValue,
			vol.BurstIOPSCredit,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsClientQueueDepth,
			prometheus.GaugeValue,
			vol.ClientQueueDepth,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsLatencySeconds,
			prometheus.GaugeValue,
			UsToSecondss(vol.LatencyUSec),
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsNonZeroBlocks,
			prometheus.GaugeValue,
			vol.NonZeroBlocks,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsReadBytes,
			prometheus.CounterValue,
			vol.ReadBytes,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsLastSampleReadBytes,
			prometheus.GaugeValue,
			vol.ReadBytesLastSample,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsReadLatencySeconds,
			prometheus.GaugeValue,
			UsToSecondss(vol.ReadLatencyUSec),
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsReadLatencySecondsTotal,
			prometheus.CounterValue,
			UsToSecondss(vol.ReadLatencyUSecTotal),
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsReadOpsTotal,
			prometheus.CounterValue,
			vol.ReadOps,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsLastSampleReadOps,
			prometheus.GaugeValue,
			vol.ReadOpsLastSample,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsThrottle,
			prometheus.GaugeValue,
			vol.Throttle,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsUnalignedReadsTotal,
			prometheus.CounterValue,
			vol.UnalignedReads,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsUnalignedWritesTotal,
			prometheus.CounterValue,
			vol.UnalignedWrites,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsVolumeSizeBytes,
			prometheus.GaugeValue,
			vol.VolumeSize,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsVolumeUtilization,
			prometheus.GaugeValue,
			vol.VolumeUtilization,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsWriteBytes,
			prometheus.CounterValue,
			vol.WriteBytes,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsLastSampleWriteBytes,
			prometheus.GaugeValue,
			vol.WriteBytesLastSample,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsWriteLatencySeconds,
			prometheus.GaugeValue,
			UsToSecondss(vol.WriteLatencyUSec),
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsWriteLatencyTotal,
			prometheus.CounterValue,
			UsToSecondss(vol.WriteLatencyUSecTotal),
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsWriteOpsTotal,
			prometheus.CounterValue,
			vol.WriteOps,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsWriteOpsLastSample,
			prometheus.GaugeValue,
			vol.WriteOpsLastSample,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.VolumeStatsZeroBlocks,
			prometheus.GaugeValue,
			vol.ZeroBlocks,
			strconv.Itoa(vol.VolumeID),
			volumeNamesByID[vol.VolumeID])

		volumeNamesMux.Unlock()
	}

	clusterCapacity, err := c.client.GetClusterCapacity()
	if err != nil {
		up = 0
		log.Errorln(err)
	}
	cluster := clusterCapacity.Result.ClusterCapacity
	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityActiveBlockSpaceBytes,
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
		MetricDescriptions.ClusterCapacityClusterRecentIOSizeBytes,
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
		MetricDescriptions.ClusterCapacityMaxOverProvisionableSpaceBytes,
		prometheus.GaugeValue,
		cluster.MaxOverProvisionableSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityMaxProvisionedSpaceBytes,
		prometheus.GaugeValue,
		cluster.MaxProvisionedSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityMaxUsedMetadataSpaceBytes,
		prometheus.GaugeValue,
		cluster.MaxUsedMetadataSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityMaxUsedSpaceBytes,
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
		MetricDescriptions.ClusterCapacityProvisionedSpaceBytes,
		prometheus.GaugeValue,
		cluster.ProvisionedSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacitySnapshotNonZeroBlocks,
		prometheus.GaugeValue,
		cluster.SnapshotNonZeroBlocks)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityIOPSTotal,
		prometheus.CounterValue,
		cluster.TotalOps)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityUniqueBlocks,
		prometheus.GaugeValue,
		cluster.UniqueBlocks)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityUniqueBlocksUsedSpaceBytes,
		prometheus.GaugeValue,
		cluster.UniqueBlocksUsedSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityUsedMetadataSpaceBytes,
		prometheus.GaugeValue,
		cluster.UsedMetadataSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityUsedMetadataSpaceInSnapshotsBytes,
		prometheus.GaugeValue,
		cluster.UsedMetadataSpaceInSnapshots)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityUsedSpaceBytes,
		prometheus.GaugeValue,
		cluster.UsedSpace)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityZeroBlocks,
		prometheus.GaugeValue,
		cluster.ZeroBlocks)

	clusterThinProvisioningFactor := (cluster.NonZeroBlocks + cluster.ZeroBlocks) / cluster.NonZeroBlocks
	if cluster.NonZeroBlocks == 0 {
		clusterThinProvisioningFactor = 1
	}

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityThinProvisioningFactor,
		prometheus.GaugeValue,
		clusterThinProvisioningFactor)

	clusterDeDuplicationFactor := (cluster.NonZeroBlocks + cluster.SnapshotNonZeroBlocks) / cluster.UniqueBlocks
	if cluster.UniqueBlocks == 0 {
		clusterDeDuplicationFactor = 1
	}

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterCapacityDeDuplicationFactor,
		prometheus.GaugeValue,
		clusterDeDuplicationFactor)

	clusterCompressionFactor := (cluster.UniqueBlocks * 4096) / (cluster.UniqueBlocksUsedSpace * 0.93)
	if cluster.UniqueBlocksUsedSpace == 0 {
		clusterCompressionFactor = 1
	}

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
		up = 0
		log.Errorln(err)
	}

	for _, f := range ClusterActiveFaults.Result.Faults {
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterActiveFaults,
			prometheus.GaugeValue,
			1,
			strconv.Itoa(f.NodeID),
			nodesNamesByID[f.NodeID],
			f.Code,
			f.Severity,
			f.Type,
			fmt.Sprintf("%f", f.ServiceID),
			strconv.FormatBool(f.Resolved),
			fmt.Sprintf("%f", f.NodeHardwareFaultID),
			fmt.Sprintf("%f", f.DriveID),
			f.Details,
		)
	}

	// List Cluster Stats
	ClusterNodeStats, err := c.client.ListNodeStats()
	if err != nil {
		up = 0
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
			MetricDescriptions.NodeStatsInterfaceInBytesTotal,
			prometheus.CounterValue,
			stats.CBytesIn,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
			"cluster",
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsInterfaceOutBytesTotal,
			prometheus.CounterValue,
			stats.CBytesOut,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
			"cluster",
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
			MetricDescriptions.NodeStatsInterfaceInBytesTotal,
			prometheus.CounterValue,
			stats.MBytesIn,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
			"management",
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsInterfaceOutBytesTotal,
			prometheus.CounterValue,
			stats.MBytesOut,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
			"management",
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsInterfaceUtilizationPercentage,
			prometheus.GaugeValue,
			stats.NetworkUtilizationCluster,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
			"cluster",
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsInterfaceUtilizationPercentage,
			prometheus.GaugeValue,
			stats.NetworkUtilizationStorage,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
			"storage",
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsReadLatencyTotal,
			prometheus.CounterValue,
			UsToSecondss(stats.ReadLatencyUSecTotal),
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
			MetricDescriptions.NodeStatsInterfaceInBytesTotal,
			prometheus.CounterValue,
			stats.SBytesIn,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
			"storage",
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsInterfaceOutBytesTotal,
			prometheus.CounterValue,
			stats.SBytesOut,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
			"storage",
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsUsedMemoryBytes,
			prometheus.GaugeValue,
			stats.UsedMemory,
			strconv.Itoa(stats.NodeID),
			nodesNamesByID[stats.NodeID],
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NodeStatsWriteLatencyTotal,
			prometheus.CounterValue,
			UsToSecondss(stats.WriteLatencyUSecTotal),
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
		up = 0
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
		up = 0
		log.Errorln(err)
	}

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsActualIOPS,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ActualIOPS,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsAverageIOBytes,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.AverageIOPSize,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsClientQueueDepth,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ClientQueueDepth,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsThroughputUtilization,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ClusterUtilization,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsLatencySeconds,
		prometheus.GaugeValue,
		UsToSecondss(clusterStats.Result.ClusterStats.LatencyUSec),
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsNormalizedIOPS,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.NormalizedIOPS,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsReadBytesTotal,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ReadBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsLastSampleReadBytes,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ReadBytesLastSample,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsReadLatencySeconds,
		prometheus.GaugeValue,
		UsToSecondss(clusterStats.Result.ClusterStats.ReadLatencyUSec),
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsReadLatencyTotal,
		prometheus.GaugeValue,
		UsToSecondss(clusterStats.Result.ClusterStats.ReadLatencyUSecTotal),
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsReadOpsTotal,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ReadOps,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsLastSampleReadOps,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ReadOpsLastSample,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsSamplePeriodSeconds,
		prometheus.GaugeValue,
		MsToSeconds(clusterStats.Result.ClusterStats.SamplePeriodMsec),
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsServices,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ServicesCount,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsExpectedServices,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.ServicesTotal,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsUnalignedReadsTotal,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.UnalignedReads,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsUnalignedWrites,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.UnalignedWrites,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsWriteBytesTotal,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.WriteBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsLastSampleWriteBytes,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.WriteBytesLastSample,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsWriteLatency,
		prometheus.GaugeValue,
		UsToSecondss(clusterStats.Result.ClusterStats.WriteLatencyUSec),
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsWriteLatencyTotal,
		prometheus.GaugeValue,
		UsToSecondss(clusterStats.Result.ClusterStats.WriteLatencyUSecTotal),
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsWriteOpsTotal,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.WriteOps,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStatsLastSampleWriteOps,
		prometheus.GaugeValue,
		clusterStats.Result.ClusterStats.WriteOpsLastSample,
	)

	clusterFullThreshold, err := c.client.GetClusterFullThreshold()
	if err != nil {
		up = 0
		log.Errorln(err)
	}

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterBlockFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.BlockFullness, "stage1Happy")),
		"stage1Happy",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterBlockFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.BlockFullness, "stage2Aware")),
		"stage2Aware",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterBlockFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.BlockFullness, "stage3Low")),
		"stage3Low",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterBlockFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.BlockFullness, "stage4Critical")),
		"stage4Critical",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterBlockFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.BlockFullness, "stage5CompletelyConsumed")),
		"stage5CompletelyConsumed",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.Fullness, "blockFullness")),
		"blockFullness",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.Fullness, "metadataFullness")),
		"metadataFullness",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterMaxMetadataOverProvisionFactor,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.MaxMetadataOverProvisionFactor,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterMetadataFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.MetadataFullness, "stage1Happy")),
		"stage1Happy",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterMetadataFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.MetadataFullness, "stage2Aware")),
		"stage2Aware",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterMetadataFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.MetadataFullness, "stage3Low")),
		"stage3Low",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterMetadataFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.MetadataFullness, "stage4Critical")),
		"stage4Critical",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterMetadataFullness,
		prometheus.GaugeValue,
		float64(strCompare(clusterFullThreshold.Result.MetadataFullness, "stage5CompletelyConsumed")),
		"stage5CompletelyConsumed",
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterSliceReserveUsedThresholdPercentage,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.SliceReserveUsedThresholdPct,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStage2AwareThresholdPercentage,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage2AwareThreshold,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStage2BlockThresholdBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage2BlockThresholdBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStage3BlockThresholdBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage3BlockThresholdBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStage3BlockThresholdPercentage,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage3BlockThresholdPercent,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStage3LowThresholdPercentage,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage3LowThreshold,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStage4BlockThresholdBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage4BlockThresholdBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStage4CriticalThreshold,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage4CriticalThreshold,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterStage5BlockThresholdBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.Stage5BlockThresholdBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterTotalBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.SumTotalClusterBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterTotalMetadataBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.SumTotalMetadataClusterBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterUsedBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.SumUsedClusterBytes,
	)

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ClusterUsedMetadataBytes,
		prometheus.GaugeValue,
		clusterFullThreshold.Result.SumUsedMetadataClusterBytes,
	)

	ListDrives, err := c.client.ListDrives()
	if err != nil {
		up = 0
		log.Errorln(err)
	}

	for _, d := range ListDrives.Result.Drives {
		for _, ds := range possibleDriveStatuses {
			var driveStatus float64 = 0
			if ds == d.Status {
				driveStatus = 1
			}
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.DriveStatus,
				prometheus.GaugeValue,
				driveStatus,
				strconv.Itoa(d.NodeID),
				nodesNamesByID[d.NodeID],
				strconv.Itoa(d.DriveID),
				d.Serial,
				strconv.Itoa(d.Slot),
				ds,
				d.Type,
			)
		}

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.DriveCapacityBytes,
			prometheus.GaugeValue,
			d.Capacity,
			strconv.Itoa(d.NodeID),
			nodesNamesByID[d.NodeID],
			strconv.Itoa(d.DriveID),
			d.Serial,
			strconv.Itoa(d.Slot),
			d.Type,
		)
	}

	ListISCSISessions, err := c.client.ListISCSISessions()
	if err != nil {
		up = 0
		log.Errorln(err)
	}

	sessions := make(map[int]map[int]float64)

	for _, session := range ListISCSISessions.Result.Sessions {
		if sessions[session.NodeID] == nil {
			sessions[session.NodeID] = make(map[int]float64)
		}
		sessions[session.NodeID][session.VolumeID]++
	}

	for node, v := range sessions {
		for vol, val := range v {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeISCSISessions,
				prometheus.GaugeValue,
				val,
				strconv.Itoa(node),
				nodesNamesByID[node],
				strconv.Itoa(vol),
				volumeNamesByID[vol],
			)
		}
	}

	// Set scrape success metric to scrapeSuccess
	ch <- prometheus.MustNewConstMetric(MetricDescriptions.upDesc, prometheus.GaugeValue, up)
}

func NewCollector() (*solidfireCollector, error) {
	log.Infof("initializing new solidfire collector")
	c, err := solidfire.NewSolidfireClient()
	if err != nil {
		return nil, err
	}
	return &solidfireCollector{
		client: c,
	}, nil
}

func GBToBytes(gb float64) float64 {
	return gb * 1e+9
}

func UsToSecondss(microSeconds float64) float64 {
	return microSeconds * 1e-6
}

func MsToSeconds(milliseconds float64) float64 {
	return milliseconds * 1e-3
}
