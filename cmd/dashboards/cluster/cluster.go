package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/K-Phoen/grabana"
	"github.com/K-Phoen/grabana/dashboard"
	"github.com/K-Phoen/grabana/row"
	"github.com/K-Phoen/grabana/singlestat"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/K-Phoen/grabana/variable/datasource"
	"github.com/K-Phoen/grabana/variable/interval"
	"github.com/K-Phoen/grabana/variable/query"
)

var (
	datasourceVar     = "$datasource"
	clusterVar        = "$sfcluster"
	clusterVariable   = dashboard.VariableAsQuery(strings.TrimLeft(clusterVar, "$"), query.DataSource(datasourceVar), query.Request("label_values(solidfire_cluster_max_iops, sfcluster)"))
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

func singleStatAlert(severity string, thresholds [2]string, colors [3]string) row.Option {
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

func main() {
	builder := dashboard.New(
		"Cluster Overview",
		dashboardTags,
		dashboardAutoRefresh,
		datasourceVariable,
		clusterVariable,
		intervalsVariable,
		dashboard.Row(
			"Alerts",
			singleStatAlert("bestpractice", [2]string{"1", "1"}, [3]string{colorGreen, colorBlue, colorBlue}),
			singleStatAlert("warning", [2]string{"1", "1"}, [3]string{colorGreen, colorOrange, colorOrange}),
			singleStatAlert("errors", [2]string{"1", "1"}, [3]string{colorGreen, colorRed, colorRed}),
			singleStatAlert("critical", [2]string{"1", "1"}, [3]string{colorGreen, colorRed, colorRed}),
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
