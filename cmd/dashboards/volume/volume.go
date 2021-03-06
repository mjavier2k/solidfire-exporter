package volume

import (
	"fmt"
	"strings"

	"github.com/K-Phoen/grabana/axis"
	"github.com/K-Phoen/grabana/dashboard"
	"github.com/K-Phoen/grabana/graph"
	"github.com/K-Phoen/grabana/graph/series"
	"github.com/K-Phoen/grabana/heatmap"
	"github.com/K-Phoen/grabana/row"
	"github.com/K-Phoen/grabana/singlestat"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/mjavier2k/solidfire-exporter/cmd/dashboards/common"
)

func commonQosHeatmap(metric string) row.Option {
	title := strings.TrimPrefix(metric, "solidfire_volume_qos_")
	title = strings.TrimSuffix(title, "_bucket")
	title = strings.ReplaceAll(title, "_", " ")
	title = strings.ToTitle(title)
	return row.WithHeatmap(
		fmt.Sprintf("%s - %s", common.VolumeVar, title),
		heatmap.Transparent(),
		heatmap.DataSource(common.DatasourceVar),
		heatmap.DataFormat(heatmap.TimeSeriesBuckets),
		heatmap.WithPrometheusTarget(
			fmt.Sprintf(`sum by(le, volume_name) (rate(%s{sfcluster=~"%s",volume_name=~"%s"}[%s]))`,
				metric, common.ClusterVar, common.VolumeVar, common.IntervalVar,
			),
			prometheus.Legend(`{{le}}`),
			prometheus.Format(prometheus.FormatHeatmap),
		),
	)
}
func NewVolumeDetailDashboard() dashboard.Builder {
	return dashboard.New(
		"Solidfire Volume Detail",
		common.DashboardTags,
		common.DashboardAutoRefresh,
		common.DatasourceVariable,
		common.ClusterVariable,
		common.VolumeVariable,
		common.IntervalsVariable,
		dashboard.Row(
			fmt.Sprintf("Volume Summary - %s", common.VolumeVar),
			row.RepeatFor(strings.TrimPrefix(common.VolumeVar, "$")),
			row.WithSingleStat(
				"VOLUME SIZE",
				singlestat.Span(2),
				singlestat.Height("120px"),
				singlestat.DataSource(common.DatasourceVar),
				singlestat.ColorBackground(),
				singlestat.Transparent(),
				singlestat.Thresholds([2]string{"1", "1"}),
				singlestat.Colors([3]string{common.ColorGreen, common.ColorGreen, common.ColorGreen}),
				singlestat.Unit("bytes"),
				singlestat.Decimals(2),
				singlestat.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_size_bytes{sfcluster=~"%s", volume_name=~"%s"}`, common.ClusterVar, common.VolumeVar),
					prometheus.Instant(),
				),
			),
			row.WithSingleStat(
				"VOLUME BLOCK UTILIZATION",
				singlestat.Span(2),
				singlestat.Height("120px"),
				singlestat.DataSource(common.DatasourceVar),
				singlestat.Description("Percentage of nonzero blocks (This is the closest measurement we have to 'volume fullness')"),
				singlestat.ColorBackground(),
				singlestat.Transparent(),
				singlestat.Thresholds([2]string{"85", "95"}),
				singlestat.Colors([3]string{common.ColorGreen, common.ColorOrange, common.ColorRed}),
				singlestat.Unit("percent"),
				singlestat.ValueType(singlestat.Current),
				singlestat.SparkLine(),
				singlestat.SparkLineColor(common.ColorBlue),
				singlestat.Decimals(2),
				singlestat.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_non_zero_blocks{sfcluster=~"%s", volume_name=~"%s"} / 
					(solidfire_volume_zero_blocks{sfcluster=~"%s", volume_name=~"%s"} + 
					solidfire_volume_non_zero_blocks{sfcluster=~"%s", volume_name=~"%s"}) * 100`,
						common.ClusterVar, common.VolumeVar,
						common.ClusterVar, common.VolumeVar,
						common.ClusterVar, common.VolumeVar,
					),
				),
			),
			row.WithSingleStat(
				"VOLUME THROUGHPUT UTILIZATION",
				singlestat.Span(2),
				singlestat.Height("120px"),
				singlestat.DataSource(common.DatasourceVar),
				singlestat.ColorBackground(),
				singlestat.Transparent(),
				singlestat.Description("How much of the volume's throughput capacity (IOPS) is being utilized"),
				singlestat.Thresholds([2]string{"80", "90"}),
				singlestat.ValueType(singlestat.Current),
				singlestat.SparkLine(),
				singlestat.SparkLineColor(common.ColorBlue),
				singlestat.Colors([3]string{common.ColorGreen, common.ColorOrange, common.ColorRed}),
				singlestat.Unit("percent"),
				singlestat.Decimals(2),
				singlestat.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_utilization{sfcluster=~"%s", volume_name=~"%s"} * 100`, common.ClusterVar, common.VolumeVar),
				),
			),
			row.WithSingleStat(
				"ISCSI SESSIONS",
				singlestat.Span(2),
				singlestat.Height("120px"),
				singlestat.DataSource(common.DatasourceVar),
				singlestat.ColorBackground(),
				singlestat.Transparent(),
				singlestat.Description("Number of ISCSI sessions to the volume"),
				singlestat.Thresholds([2]string{"1", "1"}),
				singlestat.ValueType(singlestat.Current),
				singlestat.SparkLine(),
				singlestat.SparkLineColor(common.ColorBlue),
				singlestat.Colors([3]string{common.ColorGreen, common.ColorGreen, common.ColorGreen}),
				singlestat.Unit("locale"),
				singlestat.Decimals(2),
				singlestat.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_node_iscsi_sessions{sfcluster=~"%s", volume_name=~"%s"}`,
						common.ClusterVar, common.VolumeVar,
					),
				),
			),
			row.WithSingleStat(
				"OVERALL READ %",
				singlestat.Span(2),
				singlestat.Height("120px"),
				singlestat.DataSource(common.DatasourceVar),
				singlestat.ColorBackground(),
				singlestat.Transparent(),
				singlestat.Description("The total read operations percentage to the volume since the creation of the volume."),
				singlestat.Thresholds([2]string{"1", "1"}),
				singlestat.ValueType(singlestat.Current),
				singlestat.SparkLine(),
				singlestat.SparkLineColor(common.ColorBlue),
				singlestat.Colors([3]string{common.ColorGreen, common.ColorGreen, common.ColorGreen}),
				singlestat.Unit("percent"),
				singlestat.Decimals(2),
				singlestat.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_read_ops_total{sfcluster=~"%s", volume_name=~"%s"} /
					(solidfire_volume_read_ops_total{sfcluster=~"%s", volume_name=~"%s"} + 
						solidfire_volume_write_ops_total{sfcluster=~"%s", volume_name=~"%s"}) * 100`,
						common.ClusterVar, common.VolumeVar,
						common.ClusterVar, common.VolumeVar,
						common.ClusterVar, common.VolumeVar,
					),
				),
			),
			row.WithSingleStat(
				"OVERALL WRITE %",
				singlestat.Span(2),
				singlestat.Height("120px"),
				singlestat.DataSource(common.DatasourceVar),
				singlestat.ColorBackground(),
				singlestat.Transparent(),
				singlestat.Description("The total write operations percentage to the volume since the creation of the volume."),
				singlestat.Thresholds([2]string{"1", "1"}),
				singlestat.ValueType(singlestat.Current),
				singlestat.SparkLine(),
				singlestat.SparkLineColor(common.ColorBlue),
				singlestat.Colors([3]string{common.ColorGreen, common.ColorGreen, common.ColorGreen}),
				singlestat.Unit("percent"),
				singlestat.Decimals(2),
				singlestat.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_write_ops_total{sfcluster=~"%s", volume_name=~"%s"} /
					(solidfire_volume_read_ops_total{sfcluster=~"%s", volume_name=~"%s"} + 
						solidfire_volume_write_ops_total{sfcluster=~"%s", volume_name=~"%s"}) * 100`,
						common.ClusterVar, common.VolumeVar,
						common.ClusterVar, common.VolumeVar,
						common.ClusterVar, common.VolumeVar,
					),
				),
			),
		),
		dashboard.Row(
			"Volume Detail",
			row.WithGraph("VOLUME CAPACITY",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.Description("Volume nonzero blocks / (nonzero blocks + zero blocks)"),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_non_zero_blocks{sfcluster=~"%s", volume_name=~"%s"}`,
						common.ClusterVar, common.VolumeVar,
					),
					prometheus.Legend(`{{volume_name}} - Non-Zero Blocks`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_zero_blocks{sfcluster=~"%s", volume_name=~"%s"}`,
						common.ClusterVar, common.VolumeVar,
					),
					prometheus.Legend(`{{volume_name}} - Zero Blocks`),
				),
				graph.LeftYAxis(axis.Min(0), axis.Label("blocks")),
				graph.SeriesOverride(series.Alias(`/Zero Blocks/`), series.Color(common.ColorBlue)),
				graph.SeriesOverride(series.Alias(`/Non-Zero Blocks/`), series.Color(common.ColorYellow)),
			),
			row.WithGraph("ISCSI SESSIONS",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.Description("Number of ISCSI sessions to the volume"),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_node_iscsi_sessions{sfcluster=~"%s", volume_name=~"%s"}`,
						common.ClusterVar, common.VolumeVar,
					),
					prometheus.Legend(`{{volume_name}} - ISCSI Sessions`),
				),
				graph.SeriesOverride(series.Alias(`/.*/`), series.Color(common.ColorPurple)),
				graph.LeftYAxis(axis.Unit("locale"), axis.Min(0)),
			),
		),
		dashboard.Row(
			"Volume Performance",
			row.WithGraph("VOLUME LATENCY",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_volume_read_latency_seconds_total{sfcluster=~"%s", volume_name=~"%s"}[%s])`,
						common.ClusterVar, common.VolumeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{volume_name}} - read latency`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_volume_write_latency_seconds_total{sfcluster=~"%s", volume_name=~"%s"}[%s])`,
						common.ClusterVar, common.VolumeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{volume_name}} - write latency`),
				),
				graph.SeriesOverride(series.Alias(`/write/`), series.Color(common.ColorOrange)),
				graph.SeriesOverride(series.Alias(`/read/`), series.Color(common.ColorGreen)),
				graph.LeftYAxis(axis.Unit("s"), axis.Min(0)),
			),
			row.WithGraph("VOLUME THROUGHPUT",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_volume_read_bytes_total{sfcluster=~"%s", volume_name=~"%s"}[%s])`,
						common.ClusterVar, common.VolumeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{volume_name}} - read bytes`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_volume_write_bytes_total{sfcluster=~"%s", volume_name=~"%s"}[%s])`,
						common.ClusterVar, common.VolumeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{volume_name}} - write bytes`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_volume_read_bytes_total{sfcluster=~"%s", volume_name=~"%s"}[%s]) +
					rate(solidfire_volume_write_bytes_total{sfcluster=~"%s", volume_name=~"%s"}[%s])`,
						common.ClusterVar, common.VolumeVar, common.IntervalVar,
						common.ClusterVar, common.VolumeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{volume_name}} - total bytes`),
				),
				graph.SeriesOverride(series.Alias(`/write/`), series.Color(common.ColorOrange)),
				graph.SeriesOverride(series.Alias(`/read/`), series.Color(common.ColorGreen)),
				graph.SeriesOverride(series.Alias(`/total/`), series.Color(common.ColorPurple), series.Dashes(true), series.Fill(0), series.LineWidth(3)),
				graph.LeftYAxis(axis.Unit("Bps"), axis.Min(0)),
			),
			row.WithGraph("VOLUME BURST CREDIT",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_burst_iops_credit{sfcluster=~"%s", volume_name=~"%s"}`,
						common.ClusterVar, common.VolumeVar,
					),
					prometheus.Legend(`{{volume_name}} - burst credits`),
				),
				graph.SeriesOverride(series.Alias(`/burst/`), series.Color(common.ColorBlue)),
				graph.LeftYAxis(axis.Min(0)),
			),
			row.WithGraph("VOLUME UTILIZATION",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.Description("Percentage of allotted IOPS consumed by volume. Note: it's possible to burst to above 100% by making use of burst credits the volume has accumulated."),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_utilization{sfcluster=~"%s", volume_name=~"%s"} * 100`,
						common.ClusterVar, common.VolumeVar,
					),
					prometheus.Legend(`{{volume_name}} - utilization`),
				),
				graph.WithPrometheusTarget(
					"vector(100)",
					prometheus.Legend(`max`),
				),
				graph.SeriesOverride(series.Alias(`/utilization/`), series.Color(common.ColorYellow)),
				graph.SeriesOverride(series.Alias(`/max/`), series.Color(common.ColorRed), series.Dashes(true), series.Fill(0), series.LineWidth(3)),
				graph.LeftYAxis(axis.Unit("percent"), axis.Min(0)),
			),
			row.WithGraph("VOLUME ACTUAL IOPS",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.Description("The current actual IOPS for the volume in the last 500 milliseconds"),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_actual_iops{sfcluster=~"%s", volume_name=~"%s"}`,
						common.ClusterVar, common.VolumeVar,
					),
					prometheus.Legend(`{{volume_name}} - actual iops`),
				),
				graph.SeriesOverride(series.Alias(`/actual/`), series.Color(common.ColorGreen)),
				graph.LeftYAxis(axis.Unit("iops"), axis.Min(0)),
			),
			row.WithGraph("VOLUME IOP SIZE",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.Description("The average size in bytes of recent I/O to the volume in the last 500 milliseconds."),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_average_iop_size_bytes{sfcluster=~"%s", volume_name=~"%s"}`,
						common.ClusterVar, common.VolumeVar,
					),
					prometheus.Legend(`{{volume_name}} - iop size`),
				),
				graph.SeriesOverride(series.Alias(`/iop size/`), series.Color(common.ColorPurple)),
				graph.LeftYAxis(axis.Unit("bytes"), axis.Min(0)),
			),
			row.WithGraph("VOLUME CLIENT QUEUE DEPTH",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.Description("The number of outstanding read and write operations to the volume."),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_client_queue_depth{sfcluster=~"%s", volume_name=~"%s"}`,
						common.ClusterVar, common.VolumeVar,
					),
					prometheus.Legend(`{{volume_name}} - queue depth`),
				),
				graph.SeriesOverride(series.Alias(`/queue/`), series.Color(common.ColorOrange)),
				graph.LeftYAxis(axis.Min(0)),
			),
			row.WithGraph("VOLUME THROTTLE",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.Description("Represents how much the system is throttling clients below their maxIOPS because of replication of data, transient errors, and snapshots taken."),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_volume_throttle{sfcluster=~"%s", volume_name=~"%s"} * 100`,
						common.ClusterVar, common.VolumeVar,
					),
					prometheus.Legend(`{{volume_name}} - throttle`),
				),
				graph.SeriesOverride(series.Alias(`/throttle/`), series.Color(common.ColorRed)),
				graph.LeftYAxis(axis.Unit("percent"), axis.Min(0)),
			),
			row.WithGraph("VOLUME UNALIGNED I/O",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_volume_unaligned_reads_total{sfcluster=~"%s", volume_name=~"%s"}[%s])`,
						common.ClusterVar, common.VolumeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{volume_name}} - unaligned reads`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_volume_unaligned_writes_total{sfcluster=~"%s", volume_name=~"%s"}[%s])`,
						common.ClusterVar, common.VolumeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{volume_name}} - unaligned writes`),
				),
				graph.SeriesOverride(series.Alias(`/reads/`), series.Color(common.ColorOrange)),
				graph.SeriesOverride(series.Alias(`/writes/`), series.Color(common.ColorYellow)),
				graph.LeftYAxis(axis.Min(0)),
			),
		),
		dashboard.Row(
			fmt.Sprintf("Volume QoS - %s", common.VolumeVar),
			row.RepeatFor(strings.TrimPrefix(common.VolumeVar, "$")),
			commonQosHeatmap("solidfire_volume_qos_below_min_iops_percentage_bucket"),
			commonQosHeatmap("solidfire_volume_qos_min_to_max_iops_percentage_bucket"),
			commonQosHeatmap("solidfire_volume_qos_read_block_sizes_bytes_bucket"),
			commonQosHeatmap("solidfire_volume_qos_write_block_sizes_bytes_bucket"),
			commonQosHeatmap("solidfire_volume_qos_target_utilization_percentage_bucket"),
			commonQosHeatmap("solidfire_volume_qos_throttle_percentage_bucket"),
		),
	)
}
