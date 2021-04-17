package cluster

import (
	"fmt"
	"strings"

	"github.com/K-Phoen/grabana/axis"
	"github.com/K-Phoen/grabana/dashboard"
	"github.com/K-Phoen/grabana/graph"
	"github.com/K-Phoen/grabana/graph/series"
	"github.com/K-Phoen/grabana/row"
	"github.com/K-Phoen/grabana/singlestat"
	"github.com/K-Phoen/grabana/table"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/mjavier2k/solidfire-exporter/cmd/dashboards/common"
)

func faultSingleStat(severity string, thresholds [2]string, colors [3]string) row.Option {
	return row.WithSingleStat(
		strings.ToUpper(severity),
		singlestat.Span(3),
		singlestat.Height("100px"),
		singlestat.DataSource(common.DatasourceVar),
		singlestat.ColorBackground(),
		singlestat.Thresholds(thresholds),
		singlestat.Colors(colors),
		singlestat.WithPrometheusTarget(
			fmt.Sprintf(`sum(solidfire_cluster_active_faults{sfcluster=~"%s",severity="%s"}) or vector(0)`, common.ClusterVar, severity),
			prometheus.Instant(),
		),
	)
}

func efficiencySingleStat(metric string, title string, thresholds [2]string, colors [3]string) row.Option {
	return row.WithSingleStat(
		strings.ToUpper(title),
		singlestat.Span(3),
		singlestat.Height("100px"),
		singlestat.Decimals(2),
		singlestat.Postfix("x"),
		singlestat.PostfixFontSize("75%"),
		singlestat.DataSource(common.DatasourceVar),
		singlestat.ColorBackground(),
		singlestat.Thresholds(thresholds),
		singlestat.Colors(colors),
		singlestat.WithPrometheusTarget(
			fmt.Sprintf(`%s{sfcluster=~"%s"}`, metric, common.ClusterVar),
			prometheus.Instant(),
		),
	)
}

func NewClusterOverviewDashboard() dashboard.Builder {
	return dashboard.New(
		"Solidfire Cluster Overview",
		common.DashboardTags,
		common.DashboardAutoRefresh,
		common.DatasourceVariable,
		common.ClusterVariable,
		common.IntervalsVariable,
		dashboard.Row(
			"Fault Summary",
			faultSingleStat("bestpractice", [2]string{"1", "1"}, [3]string{common.ColorGreen, common.ColorBlue, common.ColorBlue}),
			faultSingleStat("warning", [2]string{"1", "1"}, [3]string{common.ColorGreen, common.ColorOrange, common.ColorOrange}),
			faultSingleStat("errors", [2]string{"1", "1"}, [3]string{common.ColorGreen, common.ColorRed, common.ColorRed}),
			faultSingleStat("critical", [2]string{"1", "1"}, [3]string{common.ColorGreen, common.ColorRed, common.ColorRed}),
		),
		dashboard.Row(
			"Fault Detail",
			row.WithTable("Faults",
				table.Span(12),
				table.Height("150px"),
				table.DataSource(common.DatasourceVar),
				table.Transparent(),
				table.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_cluster_active_faults{sfcluster=~"%s"} > 0`, common.ClusterVar),
					prometheus.Format(prometheus.FormatTable),
					prometheus.Instant(),
				),

				table.HideColumn(`Time|__name__|Value`),
			),
		),
		dashboard.Row(
			"Storage Capacity",
			efficiencySingleStat("solidfire_cluster_thin_provisioning_factor", "Thin Provisioning Factor", [2]string{"0.5", "1"}, [3]string{common.ColorRed, common.ColorOrange, common.ColorGreen}),
			efficiencySingleStat("solidfire_cluster_efficiency_factor", "Efficiency Factor", [2]string{"0.5", "1"}, [3]string{common.ColorRed, common.ColorOrange, common.ColorGreen}),
			efficiencySingleStat("solidfire_cluster_compression_factor", "Compression Factor", [2]string{"0.5", "1"}, [3]string{common.ColorRed, common.ColorOrange, common.ColorGreen}),
			efficiencySingleStat("solidfire_cluster_de_duplication_factor", "Deduplication Factor", [2]string{"0.5", "1"}, [3]string{common.ColorRed, common.ColorOrange, common.ColorGreen}),
			row.WithGraph("Capacity",
				graph.DataSource(common.DatasourceVar),
				graph.Span(12),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_cluster_max_used_space_bytes{sfcluster=~"%s"}`, common.ClusterVar),
					prometheus.Legend(`{{sfcluster}} - Max Usable Capacity`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_cluster_unique_blocks_used_space_bytes{sfcluster=~"%s"}`, common.ClusterVar),
					prometheus.Legend(`{{sfcluster}} - Unique Blocks Space Used`),
				),
				graph.SeriesOverride(series.Alias("/Max Usable Capacity/"), series.Dashes(true), series.Color(common.ColorRed), series.Fill(0), series.LineWidth(2)),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_cluster_provisioned_space_bytes{sfcluster=~"%s"}`, common.ClusterVar),
					prometheus.Legend(`{{sfcluster}} - Provisioned Space`),
				),
				graph.Legend(graph.AsTable, graph.Min, graph.Max, graph.Current, graph.ToTheRight),
				graph.LeftYAxis(axis.Unit("bytes")),
			),
		),
		dashboard.Row(
			"Cluster Performance",
			row.WithGraph("Cpu Usage",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_node_cpu_percentage{sfcluster=~"%s"}`, common.ClusterVar),
					prometheus.Legend(`{{sfcluster}} - {{node_name}}`),
				),
				graph.SeriesOverride(series.Alias("/.*/"), series.Color(common.ColorGreen), series.Fill(7)),
				graph.LeftYAxis(axis.Unit("percent")),
			),
			row.WithGraph("Performance Utilization",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_cluster_throughput_utilization{sfcluster=~"%s"} * 100`, common.ClusterVar),
					prometheus.Legend(`{{sfcluster}}`),
				),
				graph.SeriesOverride(series.Alias("/.*/"), series.Color(common.ColorBlue), series.Fill(7)),
				graph.LeftYAxis(axis.Unit("percent")),
			),
			row.WithGraph("IOPS",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`sum by (sfcluster) (rate(solidfire_node_read_ops_total{sfcluster=~"%s"}[%s]))`, common.ClusterVar, common.IntervalVar),
					prometheus.Legend(`{{sfcluster}} read iops`),
				),
				graph.SeriesOverride(series.Alias("/write/"), series.Color(common.ColorGreen)),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`sum by (sfcluster) (rate(solidfire_node_write_ops_total{sfcluster=~"%s"}[%s]))`, common.ClusterVar, common.IntervalVar),
					prometheus.Legend(`{{sfcluster}} write iops`),
				),
				graph.SeriesOverride(series.Alias("/write/"), series.Color(common.ColorBlue)),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`sum by (sfcluster) (rate(solidfire_node_read_ops_total{sfcluster=~"%s"}[%s])) + sum by (sfcluster) (rate(solidfire_node_write_ops_total{sfcluster=~"%s"}[%s]))`, common.ClusterVar, common.IntervalVar, common.ClusterVar, common.IntervalVar),
					prometheus.Legend(`{{sfcluster}} total iops`),
				),
				graph.SeriesOverride(series.Alias("/total/"), series.Color(common.ColorRed), series.Dashes(true), series.LineWidth(2), series.Fill(0)),
				graph.LeftYAxis(axis.Unit("iops")),
			),
			row.WithGraph("Throughput/s",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`sum by (sfcluster) (rate(solidfire_cluster_read_bytes_total{sfcluster=~"%s"}[%s]))`, common.ClusterVar, common.IntervalVar),
					prometheus.Legend(`{{sfcluster}} read bytes`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`sum by (sfcluster) (rate(solidfire_cluster_write_bytes_total{sfcluster=~"%s"}[%s]))`, common.ClusterVar, common.IntervalVar),
					prometheus.Legend(`{{sfcluster}} write bytes`),
				),
				graph.SeriesOverride(series.Alias("/write/"), series.Color(common.ColorBlue)),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`sum by (sfcluster) (rate(solidfire_cluster_read_bytes_total{sfcluster=~"%s"}[%s])) + sum by (sfcluster) (rate(solidfire_cluster_write_bytes_total{sfcluster=~"%s"}[%s]))`, common.ClusterVar, common.IntervalVar, common.ClusterVar, common.IntervalVar),
					prometheus.Legend(`{{sfcluster}} total bytes`),
				),
				graph.SeriesOverride(series.Alias("/total/"), series.Color(common.ColorRed), series.Dashes(true), series.LineWidth(2), series.Fill(0)),
				graph.LeftYAxis(axis.Unit("Bps")),
			),

			row.WithGraph("iSCSI Sessions",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`sum by (sfcluster) (solidfire_cluster_active_sessions{sfcluster=~"%s"})`, common.ClusterVar),
					prometheus.Legend(`{{sfcluster }} sessions`),
				),
				graph.SeriesOverride(series.Alias("/sessions/"), series.Color(common.ColorPurple), series.Fill(7)),
				graph.LeftYAxis(axis.Unit("locale")),
			),
			row.WithGraph("Queue Depth",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`sum by (sfcluster) (solidfire_cluster_client_queue_depth{sfcluster=~"%s"})`, common.ClusterVar),
					prometheus.Legend(`{{sfcluster}} queue depth`),
				),
				graph.SeriesOverride(series.Alias("/queue depth/"), series.Color(common.ColorOrange), series.Fill(7)),
				graph.LeftYAxis(axis.Unit("locale")),
			),
		),
	)
}
