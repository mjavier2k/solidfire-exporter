package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/K-Phoen/grabana"
	"github.com/K-Phoen/grabana/dashboard"
	"github.com/mjavier2k/solidfire-exporter/cmd/dashboards/cluster"
	"github.com/mjavier2k/solidfire-exporter/cmd/dashboards/node"
	"github.com/mjavier2k/solidfire-exporter/cmd/dashboards/volume"
)

func main() {
	var dashes []dashboard.Builder
	dashes = append(dashes, cluster.NewClusterOverviewDashboard())
	dashes = append(dashes, node.NewNodeDetailDashboard())
	dashes = append(dashes, volume.NewVolumeDetailDashboard())

	for _, builder := range dashes {
		dashBytes, err := json.MarshalIndent(builder.Internal(), "", "  ")
		if err != nil {
			fmt.Printf("Could not marshal JSON: %s\n", err)
			continue
		}
		fileName := "dashboards/" + strings.ToLower(strings.ReplaceAll(builder.Internal().Title, " ", "-")) + ".json"
		err = ioutil.WriteFile(fileName, dashBytes, 0644)
		if err != nil {
			fmt.Printf("error writing dashboard json to file: %s\n", err)
			continue
		}
		fmt.Printf("Wrote dashboard to %s \n", fileName)
	}

	fmt.Println("Attempting to write dashboards to local grafana...")
	ctx := context.Background()
	client := grabana.NewClient(&http.Client{}, "http://localhost:3000", grabana.WithBasicAuth("admin", "admin"))
	folder, err := client.FindOrCreateFolder(ctx, "Solidfire")
	if err != nil {
		fmt.Printf("Could not find or create folder in local grafana: %s\n", err)
		return
	}

	for _, builder := range dashes {
		if _, err := client.UpsertDashboard(ctx, folder, builder); err != nil {
			fmt.Printf("Could not create dashboard: %s\n", err)
		}
		fmt.Printf("Wrote dashboard %s to local grafana\n", builder.Internal().Title)
	}

}
