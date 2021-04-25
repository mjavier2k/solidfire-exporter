package prom

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	log "github.com/amoghe/distillog"
	"github.com/mjavier2k/solidfire-exporter/pkg/solidfire"
	"golang.org/x/sync/errgroup"

	"github.com/prometheus/client_golang/prometheus"
)

type SolidfireCollector struct {
	client solidfire.Interface
}

var (
	MetricDescriptions    = NewMetricDescriptions("solidfire")
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

func (c *SolidfireCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- MetricDescriptions.upDesc

	ch <- MetricDescriptions.VolumeActualIOPS
	ch <- MetricDescriptions.VolumeAverageIOPSizeBytes
	ch <- MetricDescriptions.VolumeBurstIOPSCredit
	ch <- MetricDescriptions.VolumeClientQueueDepth
	ch <- MetricDescriptions.VolumeLatencySeconds
	ch <- MetricDescriptions.VolumeNonZeroBlocks
	ch <- MetricDescriptions.VolumeReadBytesTotal
	ch <- MetricDescriptions.VolumeLastSampleReadBytes
	ch <- MetricDescriptions.VolumeReadLatencySeconds
	ch <- MetricDescriptions.VolumeReadLatencySecondsTotal
	ch <- MetricDescriptions.VolumeReadOpsTotal
	ch <- MetricDescriptions.VolumeLastSampleReadOps
	ch <- MetricDescriptions.VolumeThrottle
	ch <- MetricDescriptions.VolumeUnalignedReadsTotal
	ch <- MetricDescriptions.VolumeUnalignedWritesTotal
	ch <- MetricDescriptions.VolumeSizeBytes
	ch <- MetricDescriptions.VolumeUtilization
	ch <- MetricDescriptions.VolumeWriteBytesTotal
	ch <- MetricDescriptions.VolumeLastSampleWriteBytes
	ch <- MetricDescriptions.VolumeWriteLatencySeconds
	ch <- MetricDescriptions.VolumeWriteLatencyTotal
	ch <- MetricDescriptions.VolumeWriteOpsTotal
	ch <- MetricDescriptions.VolumeWriteOpsLastSample
	ch <- MetricDescriptions.VolumeStatsZeroBlocks

	ch <- MetricDescriptions.ClusterActiveBlockSpaceBytes
	ch <- MetricDescriptions.ClusterActiveSessions
	ch <- MetricDescriptions.ClusterAverageIOPS
	ch <- MetricDescriptions.ClusterClusterRecentIOSizeBytes
	ch <- MetricDescriptions.ClusterCurrentIOPS
	ch <- MetricDescriptions.ClusterMaxIOPS
	ch <- MetricDescriptions.ClusterMaxOverProvisionableSpaceBytes
	ch <- MetricDescriptions.ClusterMaxProvisionedSpaceBytes
	ch <- MetricDescriptions.ClusterMaxUsedMetadataSpaceBytes
	ch <- MetricDescriptions.ClusterMaxUsedSpaceBytes
	ch <- MetricDescriptions.ClusterNonZeroBlocks
	ch <- MetricDescriptions.ClusterPeakActiveSessions
	ch <- MetricDescriptions.ClusterPeakIOPS
	ch <- MetricDescriptions.ClusterProvisionedSpaceBytes
	ch <- MetricDescriptions.ClusterSnapshotNonZeroBlocks
	ch <- MetricDescriptions.ClusterIOPSTotal
	ch <- MetricDescriptions.ClusterUniqueBlocks
	ch <- MetricDescriptions.ClusterUniqueBlocksUsedSpaceBytes
	ch <- MetricDescriptions.ClusterUsedMetadataSpaceBytes
	ch <- MetricDescriptions.ClusterUsedMetadataSpaceInSnapshotsBytes
	ch <- MetricDescriptions.ClusterUsedSpaceBytes
	ch <- MetricDescriptions.ClusterZeroBlocks
	ch <- MetricDescriptions.ClusterThinProvisioningFactor
	ch <- MetricDescriptions.ClusterDeDuplicationFactor
	ch <- MetricDescriptions.ClusterCompressionFactor
	ch <- MetricDescriptions.ClusterEfficiencyFactor

	ch <- MetricDescriptions.ClusterActiveFaults

	ch <- MetricDescriptions.NodeSamples
	ch <- MetricDescriptions.NodeCPUPercentage
	ch <- MetricDescriptions.NodeCPUSecondsTotal
	ch <- MetricDescriptions.NodeInterfaceInBytesTotal
	ch <- MetricDescriptions.NodeInterfaceOutBytesTotal
	ch <- MetricDescriptions.NodeInterfaceUtilizationPercentage
	ch <- MetricDescriptions.NodeReadLatencyTotal
	ch <- MetricDescriptions.NodeReadOpsTotal
	ch <- MetricDescriptions.NodeUsedMemoryBytes
	ch <- MetricDescriptions.NodeWriteLatencyTotal
	ch <- MetricDescriptions.NodeWriteOpsTotal
	ch <- MetricDescriptions.NodeLoadHistogram

	ch <- MetricDescriptions.NodeInfo

	ch <- MetricDescriptions.ClusterActualIOPS
	ch <- MetricDescriptions.ClusterAverageIOBytes
	ch <- MetricDescriptions.ClusterClientQueueDepth
	ch <- MetricDescriptions.ClusterThroughputUtilization
	ch <- MetricDescriptions.ClusterLatencySeconds
	ch <- MetricDescriptions.ClusterNormalizedIOPS
	ch <- MetricDescriptions.ClusterReadBytesTotal
	ch <- MetricDescriptions.ClusterLastSampleReadBytes
	ch <- MetricDescriptions.ClusterReadLatencySeconds
	ch <- MetricDescriptions.ClusterReadLatencyTotal
	ch <- MetricDescriptions.ClusterReadOpsTotal
	ch <- MetricDescriptions.ClusterLastSampleReadOps
	ch <- MetricDescriptions.ClusterSamplePeriodSeconds
	ch <- MetricDescriptions.ClusterServices
	ch <- MetricDescriptions.ClusterExpectedServices
	ch <- MetricDescriptions.ClusterUnalignedReadsTotal
	ch <- MetricDescriptions.ClusterUnalignedWritesTotal
	ch <- MetricDescriptions.ClusterWriteBytesTotal
	ch <- MetricDescriptions.ClusterLastSampleWriteBytes
	ch <- MetricDescriptions.ClusterWriteLatency
	ch <- MetricDescriptions.ClusterWriteLatencyTotal
	ch <- MetricDescriptions.ClusterWriteOpsTotal
	ch <- MetricDescriptions.ClusterLastSampleWriteOps

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

func (c *SolidfireCollector) Collect(ch chan<- prometheus.Metric) {
	var up float64 = 0
	var volumeNamesByID = make(map[int]string)
	var nodesNamesByID = make(map[int]string)
	defer func() { ch <- prometheus.MustNewConstMetric(MetricDescriptions.upDesc, prometheus.GaugeValue, up) }()

	// We must get volume and node details before anything else
	// because we need this metadata for labels in other calls
	metadataGroup := new(errgroup.Group)
	metadataGroup.Go(func() error {
		volumes, err := c.client.ListVolumes()
		if err != nil {
			return err
		}
		for _, vol := range volumes.Result.Volumes {
			volumeNamesByID[vol.VolumeID] = vol.Name
		}
		return nil
	})

	metadataGroup.Go(func() error {
		nodes, err := c.client.ListAllNodes()
		if err != nil {
			return err
		}
		for _, node := range nodes.Result.Nodes {
			nodesNamesByID[node.NodeID] = node.Name
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
				MetricDescriptions.NodeTotalMemoryBytes,
				prometheus.GaugeValue,
				GigabytesToBytes(node.PlatformInfo.NodeMemoryGB),
				strconv.Itoa(node.NodeID),
				node.Name,
			)
		}
		return nil
	})

	if err := metadataGroup.Wait(); err != nil {
		log.Errorln(err)
		return
	}

	metricsGroup := new(errgroup.Group)

	metricsGroup.Go(func() error {
		volumeStats, err := c.client.ListVolumeStats()
		if err != nil {
			return err
		}
		for _, vol := range volumeStats.Result.VolumeStats {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeActualIOPS,
				prometheus.GaugeValue,
				vol.ActualIOPS,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeAverageIOPSizeBytes,
				prometheus.GaugeValue,
				vol.AverageIOPSize,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeBurstIOPSCredit,
				prometheus.GaugeValue,
				vol.BurstIOPSCredit,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeClientQueueDepth,
				prometheus.GaugeValue,
				vol.ClientQueueDepth,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeLatencySeconds,
				prometheus.GaugeValue,
				MicrosecondsToSeconds(vol.LatencyUSec),
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeNonZeroBlocks,
				prometheus.GaugeValue,
				vol.NonZeroBlocks,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeReadBytesTotal,
				prometheus.CounterValue,
				vol.ReadBytes,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeLastSampleReadBytes,
				prometheus.GaugeValue,
				vol.ReadBytesLastSample,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeReadLatencySeconds,
				prometheus.GaugeValue,
				MicrosecondsToSeconds(vol.ReadLatencyUSec),
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeReadLatencySecondsTotal,
				prometheus.CounterValue,
				MicrosecondsToSeconds(vol.ReadLatencyUSecTotal),
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeReadOpsTotal,
				prometheus.CounterValue,
				vol.ReadOps,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeLastSampleReadOps,
				prometheus.GaugeValue,
				vol.ReadOpsLastSample,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeThrottle,
				prometheus.GaugeValue,
				vol.Throttle,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeUnalignedReadsTotal,
				prometheus.CounterValue,
				vol.UnalignedReads,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeUnalignedWritesTotal,
				prometheus.CounterValue,
				vol.UnalignedWrites,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeSizeBytes,
				prometheus.GaugeValue,
				vol.VolumeSize,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeUtilization,
				prometheus.GaugeValue,
				vol.VolumeUtilization,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeWriteBytesTotal,
				prometheus.CounterValue,
				vol.WriteBytes,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeLastSampleWriteBytes,
				prometheus.GaugeValue,
				vol.WriteBytesLastSample,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeWriteLatencySeconds,
				prometheus.GaugeValue,
				MicrosecondsToSeconds(vol.WriteLatencyUSec),
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeWriteLatencyTotal,
				prometheus.CounterValue,
				MicrosecondsToSeconds(vol.WriteLatencyUSecTotal),
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeWriteOpsTotal,
				prometheus.CounterValue,
				vol.WriteOps,
				strconv.Itoa(vol.VolumeID),
				volumeNamesByID[vol.VolumeID])

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.VolumeWriteOpsLastSample,
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
		}
		return nil
	})

	metricsGroup.Go(func() error {
		clusterCapacity, err := c.client.GetClusterCapacity()
		if err != nil {
			return err
		}
		cluster := clusterCapacity.Result.ClusterCapacity
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterActiveBlockSpaceBytes,
			prometheus.GaugeValue,
			cluster.ActiveBlockSpace)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterActiveSessions,
			prometheus.GaugeValue,
			cluster.ActiveSessions)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterAverageIOPS,
			prometheus.GaugeValue,
			cluster.AverageIOPS)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterClusterRecentIOSizeBytes,
			prometheus.GaugeValue,
			cluster.ClusterRecentIOSize)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterCurrentIOPS,
			prometheus.GaugeValue,
			cluster.CurrentIOPS)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterMaxIOPS,
			prometheus.GaugeValue,
			cluster.MaxIOPS)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterMaxOverProvisionableSpaceBytes,
			prometheus.GaugeValue,
			cluster.MaxOverProvisionableSpace)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterMaxProvisionedSpaceBytes,
			prometheus.GaugeValue,
			cluster.MaxProvisionedSpace)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterMaxUsedMetadataSpaceBytes,
			prometheus.GaugeValue,
			cluster.MaxUsedMetadataSpace)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterMaxUsedSpaceBytes,
			prometheus.GaugeValue,
			cluster.MaxUsedSpace)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterNonZeroBlocks,
			prometheus.GaugeValue,
			cluster.NonZeroBlocks)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterPeakActiveSessions,
			prometheus.GaugeValue,
			cluster.PeakActiveSessions)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterPeakIOPS,
			prometheus.GaugeValue,
			cluster.PeakIOPS)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterProvisionedSpaceBytes,
			prometheus.GaugeValue,
			cluster.ProvisionedSpace)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterSnapshotNonZeroBlocks,
			prometheus.GaugeValue,
			cluster.SnapshotNonZeroBlocks)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterIOPSTotal,
			prometheus.CounterValue,
			cluster.TotalOps)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterUniqueBlocks,
			prometheus.GaugeValue,
			cluster.UniqueBlocks)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterUniqueBlocksUsedSpaceBytes,
			prometheus.GaugeValue,
			cluster.UniqueBlocksUsedSpace)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterUsedMetadataSpaceBytes,
			prometheus.GaugeValue,
			cluster.UsedMetadataSpace)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterUsedMetadataSpaceInSnapshotsBytes,
			prometheus.GaugeValue,
			cluster.UsedMetadataSpaceInSnapshots)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterUsedSpaceBytes,
			prometheus.GaugeValue,
			cluster.UsedSpace)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterZeroBlocks,
			prometheus.GaugeValue,
			cluster.ZeroBlocks)

		clusterThinProvisioningFactor := (cluster.NonZeroBlocks + cluster.ZeroBlocks) / cluster.NonZeroBlocks
		if cluster.NonZeroBlocks == 0 {
			clusterThinProvisioningFactor = 1
		}

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterThinProvisioningFactor,
			prometheus.GaugeValue,
			clusterThinProvisioningFactor)

		clusterDeDuplicationFactor := (cluster.NonZeroBlocks + cluster.SnapshotNonZeroBlocks) / cluster.UniqueBlocks
		if cluster.UniqueBlocks == 0 {
			clusterDeDuplicationFactor = 1
		}

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterDeDuplicationFactor,
			prometheus.GaugeValue,
			clusterDeDuplicationFactor)

		clusterCompressionFactor := (cluster.UniqueBlocks * 4096) / (cluster.UniqueBlocksUsedSpace * 0.93)
		if cluster.UniqueBlocksUsedSpace == 0 {
			clusterCompressionFactor = 1
		}

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterCompressionFactor,
			prometheus.GaugeValue,
			clusterCompressionFactor)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterEfficiencyFactor,
			prometheus.GaugeValue,
			clusterThinProvisioningFactor*clusterDeDuplicationFactor*clusterCompressionFactor)
		return nil
	})

	metricsGroup.Go(func() error {
		ClusterActiveFaults, err := c.client.ListClusterFaults()
		if err != nil {
			return err
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
		return nil
	})

	metricsGroup.Go(func() error {
		ClusterNodeStats, err := c.client.ListNodeStats()
		if err != nil {
			return err
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
				MetricDescriptions.NodeLoadHistogram,
				stats.Count,
				float64(sumHistogram(SsLoadHistogram)),
				SsLoadHistogram,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeInterfaceInBytesTotal,
				prometheus.CounterValue,
				stats.CBytesIn,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
				"cluster",
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeInterfaceOutBytesTotal,
				prometheus.CounterValue,
				stats.CBytesOut,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
				"cluster",
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeSamples,
				prometheus.GaugeValue,
				float64(stats.Count),
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeCPUPercentage,
				prometheus.GaugeValue,
				stats.CPU,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeCPUSecondsTotal,
				prometheus.CounterValue,
				stats.CPUTotal,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeInterfaceInBytesTotal,
				prometheus.CounterValue,
				stats.MBytesIn,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
				"management",
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeInterfaceOutBytesTotal,
				prometheus.CounterValue,
				stats.MBytesOut,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
				"management",
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeInterfaceUtilizationPercentage,
				prometheus.GaugeValue,
				stats.NetworkUtilizationCluster,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
				"cluster",
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeInterfaceUtilizationPercentage,
				prometheus.GaugeValue,
				stats.NetworkUtilizationStorage,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
				"storage",
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeReadLatencyTotal,
				prometheus.CounterValue,
				MicrosecondsToSeconds(stats.ReadLatencyUSecTotal),
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeReadOpsTotal,
				prometheus.CounterValue,
				stats.ReadOps,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeInterfaceInBytesTotal,
				prometheus.CounterValue,
				stats.SBytesIn,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
				"storage",
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeInterfaceOutBytesTotal,
				prometheus.CounterValue,
				stats.SBytesOut,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
				"storage",
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeUsedMemoryBytes,
				prometheus.GaugeValue,
				stats.UsedMemory,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeWriteLatencyTotal,
				prometheus.CounterValue,
				MicrosecondsToSeconds(stats.WriteLatencyUSecTotal),
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
			)

			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NodeWriteOpsTotal,
				prometheus.CounterValue,
				stats.WriteOps,
				strconv.Itoa(stats.NodeID),
				nodesNamesByID[stats.NodeID],
			)
		}
		return nil
	})

	metricsGroup.Go(func() error {
		VolumeQoSHistograms, err := c.client.ListVolumeQoSHistograms()
		if err != nil {
			return err
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
				MetricDescriptions.VolumeQoSTargetUtilizationPercentagesHistogram,
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
				MetricDescriptions.VolumeQoSThrottlePercentagesHistogram,
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
				MetricDescriptions.VolumeQoSWriteBlockSizesHistogram,
				0,
				float64(sumHistogram(WriteBlockSizes)),
				WriteBlockSizes,
				strconv.Itoa(h.VolumeID),
				volumeNamesByID[h.VolumeID],
			)
		}
		return nil
	})

	metricsGroup.Go(func() error {
		clusterStats, err := c.client.GetClusterStats()
		if err != nil {
			return err
		}

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterActualIOPS,
			prometheus.GaugeValue,
			clusterStats.Result.ClusterStats.ActualIOPS,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterAverageIOBytes,
			prometheus.GaugeValue,
			clusterStats.Result.ClusterStats.AverageIOPSize,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterClientQueueDepth,
			prometheus.GaugeValue,
			clusterStats.Result.ClusterStats.ClientQueueDepth,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterThroughputUtilization,
			prometheus.GaugeValue,
			clusterStats.Result.ClusterStats.ClusterUtilization,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterLatencySeconds,
			prometheus.GaugeValue,
			MicrosecondsToSeconds(clusterStats.Result.ClusterStats.LatencyUSec),
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterNormalizedIOPS,
			prometheus.GaugeValue,
			clusterStats.Result.ClusterStats.NormalizedIOPS,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterReadBytesTotal,
			prometheus.CounterValue,
			clusterStats.Result.ClusterStats.ReadBytes,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterLastSampleReadBytes,
			prometheus.GaugeValue,
			clusterStats.Result.ClusterStats.ReadBytesLastSample,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterReadLatencySeconds,
			prometheus.GaugeValue,
			MicrosecondsToSeconds(clusterStats.Result.ClusterStats.ReadLatencyUSec),
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterReadLatencyTotal,
			prometheus.CounterValue,
			MicrosecondsToSeconds(clusterStats.Result.ClusterStats.ReadLatencyUSecTotal),
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterReadOpsTotal,
			prometheus.CounterValue,
			clusterStats.Result.ClusterStats.ReadOps,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterLastSampleReadOps,
			prometheus.GaugeValue,
			clusterStats.Result.ClusterStats.ReadOpsLastSample,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterSamplePeriodSeconds,
			prometheus.GaugeValue,
			MillisecondsToSeconds(clusterStats.Result.ClusterStats.SamplePeriodMsec),
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterServices,
			prometheus.GaugeValue,
			clusterStats.Result.ClusterStats.ServicesCount,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterExpectedServices,
			prometheus.GaugeValue,
			clusterStats.Result.ClusterStats.ServicesTotal,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterUnalignedReadsTotal,
			prometheus.CounterValue,
			clusterStats.Result.ClusterStats.UnalignedReads,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterUnalignedWritesTotal,
			prometheus.CounterValue,
			clusterStats.Result.ClusterStats.UnalignedWrites,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterWriteBytesTotal,
			prometheus.CounterValue,
			clusterStats.Result.ClusterStats.WriteBytes,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterLastSampleWriteBytes,
			prometheus.GaugeValue,
			clusterStats.Result.ClusterStats.WriteBytesLastSample,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterWriteLatency,
			prometheus.GaugeValue,
			MicrosecondsToSeconds(clusterStats.Result.ClusterStats.WriteLatencyUSec),
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterWriteLatencyTotal,
			prometheus.CounterValue,
			MicrosecondsToSeconds(clusterStats.Result.ClusterStats.WriteLatencyUSecTotal),
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterWriteOpsTotal,
			prometheus.CounterValue,
			clusterStats.Result.ClusterStats.WriteOps,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ClusterLastSampleWriteOps,
			prometheus.GaugeValue,
			clusterStats.Result.ClusterStats.WriteOpsLastSample,
		)
		return nil
	})

	metricsGroup.Go(func() error {
		clusterFullThreshold, err := c.client.GetClusterFullThreshold()
		if err != nil {
			return err
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
		return nil
	})
	metricsGroup.Go(func() error {
		ListDrives, err := c.client.ListDrives()
		if err != nil {
			return err
		}

		for _, d := range ListDrives.Result.Drives {
			for _, ds := range possibleDriveStatuses {
				var driveStatusValue float64 = 0
				if ds == d.Status {
					driveStatusValue = 1
				}
				ch <- prometheus.MustNewConstMetric(
					MetricDescriptions.DriveStatus,
					prometheus.GaugeValue,
					driveStatusValue,
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
		return nil
	})
	metricsGroup.Go(func() error {
		ListISCSISessions, err := c.client.ListISCSISessions()
		if err != nil {
			return err
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
		return nil
	})

	if err := metricsGroup.Wait(); err != nil {
		log.Errorln(err)
		return
	}
	up = 1
}

func NewCollector(client solidfire.Interface) (*SolidfireCollector, error) {
	var err error
	if client == nil {
		log.Infof("initializing new solidfire client")
		client, err = solidfire.NewSolidfireClient()
		if err != nil {
			return nil, err
		}
	}
	return &SolidfireCollector{
		client: client,
	}, nil
}

func GigabytesToBytes(gb float64) float64 {
	return gb * 1e+9
}

func MicrosecondsToSeconds(microSeconds float64) float64 {
	return microSeconds * 1e-6
}

func MillisecondsToSeconds(milliseconds float64) float64 {
	return milliseconds * 1e-3
}
