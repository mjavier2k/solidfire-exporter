package common

import (
	"fmt"
	"strings"

	"github.com/K-Phoen/grabana/dashboard"
	"github.com/K-Phoen/grabana/variable/datasource"
	"github.com/K-Phoen/grabana/variable/interval"
	"github.com/K-Phoen/grabana/variable/query"
)

var (
	DatasourceVar   = "$datasource"
	ClusterVar      = "$sfcluster"
	IntervalVar     = "$interval"
	VolumeVar       = "$volume"
	NodeVar         = "$node"
	ClusterVariable = dashboard.VariableAsQuery(
		strings.TrimLeft(ClusterVar, "$"),
		query.DataSource(DatasourceVar),
		query.Request("label_values(solidfire_cluster_max_iops, sfcluster)"),
		query.Refresh(query.TimeChange),
	)
	VolumeVariable = dashboard.VariableAsQuery(
		strings.TrimLeft(VolumeVar, "$"),
		query.Multi(),
		query.DataSource(DatasourceVar),
		query.Request(fmt.Sprintf(`label_values(solidfire_volume_size_bytes{sfcluster=~"%s"}, volume_name)`, ClusterVar)),
		query.Refresh(query.TimeChange),
	)
	NodeVariable = dashboard.VariableAsQuery(
		strings.TrimLeft(NodeVar, "$"),
		query.Multi(),
		query.IncludeAll(),
		query.DefaultAll(),
		query.DataSource(DatasourceVar),
		query.Request(fmt.Sprintf(`label_values(solidfire_node_cpu_percentage{sfcluster=~"%s"}, node_name)`, ClusterVar)),
		query.Refresh(query.TimeChange),
	)
	IntervalsVariable = dashboard.VariableAsInterval(
		strings.TrimLeft(IntervalVar, "$"),
		interval.Values([]string{"30s", "1m", "5m", "10m", "30m", "1h", "6h", "12h"}),
	)
	DatasourceVariable = dashboard.VariableAsDatasource(
		strings.TrimLeft(DatasourceVar, "$"),
		datasource.Type("prometheus"),
	)
	DashboardAutoRefresh = dashboard.AutoRefresh("30s")
	DashboardTags        = dashboard.Tags([]string{"solidfire", "generated"})

	// NetApp Color pallette https://www-download.netapp.com/edm/email-guideline/
	ColorGreen  = "#118B42"
	ColorRed    = "#AA342C"
	ColorOrange = "#D74822"
	ColorBlue   = "#0077BF"
	ColorPurple = "#804C9D"
	ColorYellow = "#E2AB80"
)
