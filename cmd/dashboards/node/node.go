package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/K-Phoen/grabana"
	"github.com/K-Phoen/grabana/axis"
	"github.com/K-Phoen/grabana/dashboard"
	"github.com/K-Phoen/grabana/graph"
	"github.com/K-Phoen/grabana/graph/series"
	"github.com/K-Phoen/grabana/row"
	"github.com/K-Phoen/grabana/singlestat"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/mjavier2k/solidfire-exporter/cmd/dashboards/common"
)

func driveSingleStat(status, triggerColor string, span float32) row.Option {
	return row.WithSingleStat(
		strings.ToUpper(status),
		singlestat.Span(span),
		singlestat.Height("120px"),
		singlestat.Decimals(0),
		singlestat.DataSource(common.DatasourceVar),
		singlestat.ColorBackground(),
		singlestat.Thresholds([2]string{"1", "1"}),
		singlestat.Colors([3]string{common.ColorGreen, common.ColorGreen, triggerColor}),
		singlestat.WithPrometheusTarget(
			fmt.Sprintf(`sum(solidfire_drive_status{sfcluster=~"%s",node_name=~"%s",status="%s"}) by (node_name)`, common.ClusterVar, common.NodeVar, status),
			prometheus.Instant(),
		),
	)
}

func main() {
	builder := dashboard.New(
		"Node Detail",
		common.DashboardTags,
		common.DashboardAutoRefresh,
		common.DatasourceVariable,
		common.ClusterVariable,
		common.NodeVariable,
		common.IntervalsVariable,
		dashboard.Row(
			fmt.Sprintf("Node Performance Summary - %s", common.NodeVar),
			row.RepeatFor(strings.TrimPrefix(common.NodeVar, "$")),
			row.WithSingleStat(
				"CPU USAGE",
				singlestat.Span(2),
				singlestat.Height("120px"),
				singlestat.Decimals(0),
				singlestat.DataSource(common.DatasourceVar),
				singlestat.ColorBackground(),
				singlestat.Thresholds([2]string{"80", "90"}),
				singlestat.ValueType(singlestat.Current),
				singlestat.SparkLine(),
				singlestat.Unit("percent"),
				singlestat.Colors([3]string{common.ColorGreen, common.ColorOrange, common.ColorRed}),
				singlestat.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_node_cpu_percentage{sfcluster=~"%s",node_name=~"%s"}`, common.ClusterVar, common.NodeVar),
				),
			),
			row.WithSingleStat(
				"MEMORY USAGE",
				singlestat.Span(2),
				singlestat.Height("120px"),
				singlestat.Decimals(0),
				singlestat.DataSource(common.DatasourceVar),
				singlestat.ColorBackground(),
				singlestat.Thresholds([2]string{"80", "90"}),
				singlestat.ValueType(singlestat.Current),
				singlestat.SparkLine(),
				singlestat.Unit("percent"),
				singlestat.Colors([3]string{common.ColorGreen, common.ColorOrange, common.ColorRed}),
				singlestat.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_node_used_memory_bytes{sfcluster=~"%s",node_name=~"%s"} / 
					solidfire_node_total_memory_bytes{sfcluster=~"%s",node_name=~"%s"} * 100`,
						common.ClusterVar, common.NodeVar,
						common.ClusterVar, common.NodeVar,
					),
				),
			),
			row.WithSingleStat(
				"ISCSI SESSIONS",
				singlestat.Span(2),
				singlestat.Height("120px"),
				singlestat.Decimals(0),
				singlestat.DataSource(common.DatasourceVar),
				singlestat.ColorBackground(),
				singlestat.Thresholds([2]string{"1", "1"}),
				singlestat.ValueType(singlestat.Current),
				singlestat.SparkLine(),
				singlestat.Colors([3]string{common.ColorGreen, common.ColorGreen, common.ColorGreen}),
				singlestat.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_node_iscsi_sessions{sfcluster=~"%s", node_name=~"%s"}`,
						common.ClusterVar, common.NodeVar,
					),
				),
			),
			row.WithSingleStat(
				"MGMT BANDWIDTH",
				singlestat.Span(2),
				singlestat.Description("Bytes i/o on management network interface"),
				singlestat.Height("120px"),
				singlestat.Decimals(0),
				singlestat.DataSource(common.DatasourceVar),
				singlestat.ColorBackground(),
				singlestat.Thresholds([2]string{"1", "1"}),
				singlestat.ValueType(singlestat.Current),
				singlestat.SparkLine(),
				singlestat.Unit("Bps"),
				singlestat.Colors([3]string{common.ColorGreen, common.ColorGreen, common.ColorGreen}),
				singlestat.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_node_interface_in_bytes_total{interface="management",sfcluster=~"%s",node_name=~"%s"}[%s]) +
					rate(solidfire_node_interface_out_bytes_total{interface="management",sfcluster=~"%s",node_name=~"%s"}[%s])`,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
					),
				),
			),
			row.WithSingleStat(
				"STORAGE BANDWIDTH",
				singlestat.Span(2),
				singlestat.Description("Bytes i/o on storage network interface"),
				singlestat.Height("120px"),
				singlestat.Decimals(0),
				singlestat.DataSource(common.DatasourceVar),
				singlestat.ColorBackground(),
				singlestat.Thresholds([2]string{"1", "1"}),
				singlestat.ValueType(singlestat.Current),
				singlestat.SparkLine(),
				singlestat.Unit("Bps"),
				singlestat.Colors([3]string{common.ColorGreen, common.ColorGreen, common.ColorGreen}),
				singlestat.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_node_interface_in_bytes_total{interface="storage",sfcluster=~"%s",node_name=~"%s"}[%s]) +
					rate(solidfire_node_interface_out_bytes_total{interface="storage",sfcluster=~"%s",node_name=~"%s"}[%s])`,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
					),
				),
			),
			row.WithSingleStat(
				"IOPS",
				singlestat.Span(2),
				singlestat.Height("120px"),
				singlestat.Decimals(1),
				singlestat.DataSource(common.DatasourceVar),
				singlestat.ColorBackground(),
				singlestat.Thresholds([2]string{"1", "1"}),
				singlestat.ValueType(singlestat.Current),
				singlestat.SparkLine(),
				singlestat.Unit("iops"),
				singlestat.Colors([3]string{common.ColorGreen, common.ColorGreen, common.ColorGreen}),
				singlestat.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_node_read_ops_total{sfcluster=~"%s", node_name=~"%s"}[%s]) +
					rate(solidfire_node_write_ops_total{sfcluster=~"%s", node_name=~"%s"}[%s])`,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
					),
				),
			),
		),
		dashboard.Row(
			fmt.Sprintf("Node Drive Summary - %s", common.NodeVar),
			row.RepeatFor(strings.TrimPrefix(common.NodeVar, "$")),
			driveSingleStat("active", common.ColorGreen, 3),
			driveSingleStat("available", common.ColorBlue, 3),
			driveSingleStat("removing", common.ColorBlue, 2),
			driveSingleStat("failed", common.ColorRed, 2),
			driveSingleStat("erasing", common.ColorOrange, 2),
		),
		dashboard.Row(
			"Node Performance Detail",
			row.WithGraph("CPU",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_node_cpu_percentage{sfcluster=~"%s", node_name=~"%s"}`, common.ClusterVar, common.NodeVar),
					prometheus.Legend(`{{node_name}} - cpu usage`),
				),
				graph.SeriesOverride(series.Alias("/.*/"), series.Color(common.ColorGreen)),
				graph.LeftYAxis(axis.Min(0), axis.Max(100), axis.Unit("percent")),
			),
			row.WithGraph("MEMORY",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_node_used_memory_bytes{sfcluster=~"%s",node_name=~"%s"} / 
					solidfire_node_total_memory_bytes{sfcluster=~"%s",node_name=~"%s"} * 100`,
						common.ClusterVar, common.NodeVar,
						common.ClusterVar, common.NodeVar,
					),
					prometheus.Legend(`{{node_name}} - memory usage`),
				),
				graph.SeriesOverride(series.Alias("/.*/"), series.Color(common.ColorYellow)),
				graph.LeftYAxis(axis.Min(0), axis.Max(100), axis.Unit("percent")),
			),
			row.WithGraph("IOPS",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_node_read_ops_total{sfcluster=~"%s", node_name=~"%s"}[%s])`,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{node_name}} - read iops`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_node_write_ops_total{sfcluster=~"%s", node_name=~"%s"}[%s])`,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{node_name}} - write iops`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_node_read_ops_total{sfcluster=~"%s", node_name=~"%s"}[%s]) +
					rate(solidfire_node_write_ops_total{sfcluster=~"%s", node_name=~"%s"}[%s])`,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{node_name}} - total iops`),
				),
				graph.SeriesOverride(series.Alias("/read/"), series.Color(common.ColorGreen)),
				graph.SeriesOverride(series.Alias("/write/"), series.Color(common.ColorYellow)),
				graph.SeriesOverride(series.Alias("/total/"), series.Color(common.ColorRed), series.Dashes(true), series.Fill(0), series.LineWidth(3)),
				graph.LeftYAxis(axis.Min(0), axis.Unit("iops")),
			),

			row.WithGraph("ISCSI SESSIONS",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_node_iscsi_sessions{sfcluster=~"%s", node_name=~"%s"}`,
						common.ClusterVar, common.NodeVar,
					),
					prometheus.Legend(`{{node_name}} - {{volume_name}} sessions`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`sum(solidfire_node_iscsi_sessions{sfcluster=~"%s", node_name=~"%s"}) by (node_name)`,
						common.ClusterVar, common.NodeVar,
					),
					prometheus.Legend(`{{node_name}} - total`),
				),
				graph.SeriesOverride(series.Alias("/sessions/"), series.Color(common.ColorBlue)),
				graph.SeriesOverride(series.Alias("/total/"), series.Color(common.ColorRed), series.Dashes(true), series.Fill(0), series.LineWidth(3)),
				graph.LeftYAxis(axis.Min(0)),
			),
		),
		dashboard.Row(
			"Node Throughput Detail",
			row.WithGraph("MGMT IFACE THROUGHPUT ",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_node_interface_in_bytes_total{sfcluster=~"%s", node_name=~"%s", interface="management"}[%s])`,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{node_name}} - in`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_node_interface_out_bytes_total{sfcluster=~"%s", node_name=~"%s", interface="management"}[%s])`,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{node_name}} - out`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_node_interface_in_bytes_total{sfcluster=~"%s", node_name=~"%s", interface="management"}[%s]) +
					rate(solidfire_node_interface_out_bytes_total{sfcluster=~"%s", node_name=~"%s", interface="management"}[%s])`,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{node_name}} - total`),
				),
				graph.SeriesOverride(series.Alias("/in/"), series.Color(common.ColorYellow)),
				graph.SeriesOverride(series.Alias("/out/"), series.Color(common.ColorBlue)),
				graph.SeriesOverride(series.Alias("/total/"), series.Color(common.ColorRed), series.Dashes(true), series.LineWidth(3), series.Fill(0)),
				graph.LeftYAxis(axis.Min(0), axis.Unit("Bps")),
			),
			row.WithGraph("STORAGE IFACE THROUGHPUT ",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_node_interface_in_bytes_total{sfcluster=~"%s", node_name=~"%s", interface="storage"}[%s])`,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{node_name}} - in`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_node_interface_out_bytes_total{sfcluster=~"%s", node_name=~"%s", interface="storage"}[%s])`,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{node_name}} - out`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`rate(solidfire_node_interface_in_bytes_total{sfcluster=~"%s", node_name=~"%s", interface="storage"}[%s]) +
					rate(solidfire_node_interface_out_bytes_total{sfcluster=~"%s", node_name=~"%s", interface="storage"}[%s])`,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
						common.ClusterVar, common.NodeVar, common.IntervalVar,
					),
					prometheus.Legend(`{{node_name}} - total`),
				),
				graph.SeriesOverride(series.Alias("/in/"), series.Color(common.ColorYellow)),
				graph.SeriesOverride(series.Alias("/out/"), series.Color(common.ColorBlue)),
				graph.SeriesOverride(series.Alias("/total/"), series.Color(common.ColorRed), series.Dashes(true), series.LineWidth(3), series.Fill(0)),
				graph.LeftYAxis(axis.Min(0), axis.Unit("Bps")),
			),
		),
	)

	ctx := context.Background()
	client := grabana.NewClient(&http.Client{}, "http://localhost:3000", grabana.WithBasicAuth("admin", "admin"))

	// create the folder holding the dashboard for the service
	folder, err := client.FindOrCreateFolder(ctx, "Solidfire")
	if err != nil {
		fmt.Printf("Could not find or create folder: %s\n", err)
		os.Exit(1)
	}

	if _, err := client.UpsertDashboard(ctx, folder, builder); err != nil {
		fmt.Printf("Could not create dashboard: %s\n", err)
		os.Exit(1)
	}
}
