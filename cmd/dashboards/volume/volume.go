package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

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

func main() {
	builder := dashboard.New(
		"Volume Detail",
		common.DashboardTags,
		common.DashboardAutoRefresh,
		common.DatasourceVariable,
		common.ClusterVariable,
		common.VolumeVariable,
		common.IntervalsVariable,
		dashboard.Row(
			"Volume Summary",
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
					fmt.Sprintf(`sum(solidfire_volume_size_bytes{sfcluster=~"%s", volume_name=~"%s"}) by (volume_name)`, common.ClusterVar, common.VolumeVar),
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
					fmt.Sprintf(`sum by (volume_name) (solidfire_node_iscsi_sessions{sfcluster=~"%s", volume_name=~"%s"})`,
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
					fmt.Sprintf(`sum(solidfire_volume_read_ops_total{sfcluster=~"%s", volume_name=~"%s"}) by (sfcluster, volume_name) /
					sum(solidfire_volume_read_ops_total{sfcluster=~"%s", volume_name=~"%s"} + 
						solidfire_volume_write_ops_total{sfcluster=~"%s", volume_name=~"%s"}) 
					by (sfcluster, volume_name) * 100`,
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
					fmt.Sprintf(`sum(solidfire_volume_write_ops_total{sfcluster=~"%s", volume_name=~"%s"}) by (sfcluster, volume_name) /
					sum(solidfire_volume_read_ops_total{sfcluster=~"%s", volume_name=~"%s"} + 
						solidfire_volume_write_ops_total{sfcluster=~"%s", volume_name=~"%s"}) 
					by (sfcluster, volume_name) * 100`,
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
				graph.SeriesOverride(series.Alias(`/Non-Zero Blocks/`), series.Color(common.Coloryellow)),
			),
			row.WithGraph("ISCSI SESSIONS",
				graph.DataSource(common.DatasourceVar),
				graph.Span(6),
				graph.Transparent(),
				graph.Description("Number of ISCSI sessions to the volume"),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`sum by (volume_name) (solidfire_node_iscsi_sessions{sfcluster=~"%s", volume_name=~"%s"})`,
						common.ClusterVar, common.VolumeVar,
					),
					prometheus.Legend(`{{volume_name}} - ISCSI Sessions`),
				),
				graph.SeriesOverride(series.Alias(`/.*/`), series.Color(common.ColorPurple), series.Fill(4)),
				graph.LeftYAxis(axis.Unit("locale"), axis.Min(0)),
			),
		),
		dashboard.Row(
			"Volume Performance",
			row.WithText(" "),
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
