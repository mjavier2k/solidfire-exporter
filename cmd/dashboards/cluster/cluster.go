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
	"github.com/K-Phoen/grabana/row"
	"github.com/K-Phoen/grabana/singlestat"
	"github.com/K-Phoen/grabana/table"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/K-Phoen/grabana/variable/datasource"
	"github.com/K-Phoen/grabana/variable/interval"
	"github.com/K-Phoen/grabana/variable/query"
)

var (
	datasourceVar   = "$datasource"
	clusterVar      = "$sfcluster"
	clusterVariable = dashboard.VariableAsQuery(
		strings.TrimLeft(clusterVar, "$"),
		query.DataSource(datasourceVar),
		query.Request("label_values(solidfire_cluster_max_iops, sfcluster)"),
		query.Refresh(query.TimeChange),
	)
	intervalsVariable = dashboard.VariableAsInterval(
		"interval",
		interval.Values([]string{"30s", "1m", "5m", "10m", "30m", "1h", "6h", "12h"}),
	)
	datasourceVariable   = dashboard.VariableAsDatasource(strings.TrimLeft(datasourceVar, "$"), datasource.Type("prometheus"))
	dashboardAutoRefresh = dashboard.AutoRefresh("30s")
	dashboardTags        = dashboard.Tags([]string{"solidfire", "generated"})

	// NetApp Color pallette https://www-download.netapp.com/edm/email-guideline/
	colorGreen  = "#118B42"
	colorRed    = "#CF2128"
	colorOrange = "#EE6023"
	colorBlue   = "#0077BF"
)

func faultSingleStat(severity string, thresholds [2]string, colors [3]string) row.Option {
	return row.WithSingleStat(
		strings.ToUpper(severity),
		singlestat.Span(3),
		singlestat.Height("100px"),
		singlestat.DataSource(datasourceVar),
		singlestat.ColorBackground(),
		singlestat.Thresholds(thresholds),
		singlestat.Colors(colors),
		singlestat.WithPrometheusTarget(
			fmt.Sprintf(`sum(solidfire_cluster_active_faults{sfcluster=~"%s",severity="%s"}) or vector(0)`, clusterVar, severity),
			prometheus.Instant(),
		),
	)
}

func efficiencySingleStat(metric string, title string, thresholds [2]string, colors [3]string) row.Option {
	return row.WithSingleStat(
		strings.ToUpper(title),
		singlestat.Span(3),
		singlestat.Height("100px"),
		singlestat.Decimals(1),
		singlestat.Postfix("x"),
		singlestat.PostfixFontSize("75%"),
		singlestat.DataSource(datasourceVar),
		singlestat.ColorBackground(),
		singlestat.Thresholds(thresholds),
		singlestat.Colors(colors),
		singlestat.WithPrometheusTarget(
			fmt.Sprintf(`%s{sfcluster=~"%s"}`, metric, clusterVar),
			prometheus.Instant(),
		),
	)
}

func main() {
	builder := dashboard.New(
		"Cluster Overview",
		dashboardTags,
		dashboardAutoRefresh,
		datasourceVariable,
		clusterVariable,
		intervalsVariable,
		dashboard.Row(
			"Fault Summary",
			faultSingleStat("bestpractice", [2]string{"1", "1"}, [3]string{colorGreen, colorBlue, colorBlue}),
			faultSingleStat("warning", [2]string{"1", "1"}, [3]string{colorGreen, colorOrange, colorOrange}),
			faultSingleStat("errors", [2]string{"1", "1"}, [3]string{colorGreen, colorRed, colorRed}),
			faultSingleStat("critical", [2]string{"1", "1"}, [3]string{colorGreen, colorRed, colorRed}),
		),
		dashboard.Row(
			"Fault Detail",
			row.WithTable("Faults",
				table.Span(12),
				table.DataSource(datasourceVar),
				table.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_cluster_active_faults{sfcluster=~"%s"} > 0`, clusterVar),
					prometheus.Format(prometheus.FormatTable),
					prometheus.Instant(),
				),
				table.HideColumn(`Time|__name__|Value`),
			),
		),
		dashboard.Row(
			"Storage Capacity",
			efficiencySingleStat("solidfire_cluster_thin_provisioning_factor", "Thin Provisioning Factor", [2]string{"0.5", "1"}, [3]string{colorRed, colorOrange, colorGreen}),
			efficiencySingleStat("solidfire_cluster_efficiency_factor", "Efficiency Factor", [2]string{"0.5", "1"}, [3]string{colorRed, colorOrange, colorGreen}),
			efficiencySingleStat("solidfire_cluster_compression_factor", "Compression Factor", [2]string{"0.5", "1"}, [3]string{colorRed, colorOrange, colorGreen}),
			efficiencySingleStat("solidfire_cluster_de_duplication_factor", "Deduplication Factor", [2]string{"0.5", "1"}, [3]string{colorRed, colorOrange, colorGreen}),
			row.WithGraph("Capacity",
				graph.DataSource(datasourceVar),
				graph.Span(12),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_cluster_max_used_space_bytes{sfcluster=~"%s"}`, clusterVar),
					prometheus.Legend(`{{sfcluster}} - Max Usable Capacity`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_cluster_unique_blocks_used_space_bytes{sfcluster=~"%s"}`, clusterVar),
					prometheus.Legend(`{{sfcluster}} - Unique Blocks Space Used`),
				),
				graph.WithPrometheusTarget(
					fmt.Sprintf(`solidfire_cluster_provisioned_space_bytes{sfcluster=~"%s"}`, clusterVar),
					prometheus.Legend(`{{sfcluster}} - Provisioned Space`),
				),
				graph.Legend(graph.AsTable, graph.Min, graph.Max, graph.Current),
				graph.LeftYAxis(axis.Unit("bytes")),
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
