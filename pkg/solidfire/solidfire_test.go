package solidfire_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"reflect"
	"testing"

	"github.com/mjavier2k/solidfire-exporter/pkg/solidfire"
	"github.com/mjavier2k/solidfire-exporter/pkg/testutils"
	"gopkg.in/h2non/gock.v1"
)

var (
	sfHost        = "https://192.168.1.1"
	sfRPCEndpoint = "/json-rpc/11.3"
	sfClient      = solidfire.Client{
		RPCEndpoint: fmt.Sprintf("%v%v", sfHost, sfRPCEndpoint),
		HttpClient:  &http.Client{},
	}
	fixtureBasePath = path.Join("..", "..", "test", "fixtures")
)

func TestClient_ListVolumeStats(t *testing.T) {
	fixture, err := ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, solidfire.RPCListVolumeStats))
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    float64
		wantErr bool
	}{
		{
			name: "VolumeSize of first volume should match fixture",
			want: 2000683008,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			//gock.Observe(gock.DumpRequest)
			gock.New(sfHost).
				Post(sfRPCEndpoint).
				MatchType("json").
				JSON(solidfire.RPCBody{
					ID:     1,
					Method: solidfire.RPCListVolumeStats,
					Params: solidfire.ListVolumeStatsRPCParams{
						VolumeIDs:             []int{},
						IncludeVirtualVolumes: true,
					}}).
				Reply(200).
				BodyString(string(fixture))
			gotRaw, err := sfClient.ListVolumeStats(context.Background())
			got := gotRaw.Result.VolumeStats[0].VolumeSize
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListVolumeStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ListVolumeStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ListVolumes(t *testing.T) {
	fixture, err := ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, solidfire.RPCListVolumes))
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    string
		wantErr bool
	}{
		{
			name: "Volume Name of first volume should match fixture",
			want: "test-volume1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			//gock.Observe(gock.DumpRequest)
			gock.New(sfHost).
				Post(sfRPCEndpoint).
				MatchType("json").
				JSON(solidfire.RPCBody{
					ID:     1,
					Method: solidfire.RPCListVolumes,
					Params: solidfire.ListVolumesRPCParams{
						IncludeVirtualVolumes: true,
					}}).
				Reply(200).
				BodyString(string(fixture))
			gotRaw, err := sfClient.ListVolumes(context.Background())
			got := gotRaw.Result.Volumes[0].Name
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListVolumes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ListVolumes() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestClient_GetClusterCapacity(t *testing.T) {
	fixture, err := ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, solidfire.RPCGetClusterCapacity))
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    float64
		wantErr bool
	}{
		{
			name: "Cluster Capacity TotalOps should match fixture",
			want: 24181537,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			//gock.Observe(gock.DumpRequest)
			gock.New(sfHost).
				Post(sfRPCEndpoint).
				MatchType("json").
				JSON(solidfire.RPCBody{
					ID:     1,
					Method: solidfire.RPCGetClusterCapacity,
					Params: solidfire.GetClusterCapacityRPCParams{}}).
				Reply(200).
				BodyString(string(fixture))
			gotRaw, err := sfClient.GetClusterCapacity(context.Background())
			got := gotRaw.Result.ClusterCapacity.TotalOps
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetClusterCapacity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetClusterCapacity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ListClusterFaults(t *testing.T) {
	fixture, err := ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, solidfire.RPCListClusterFaults))
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    float64
		wantErr bool
	}{
		{
			name: "Cluster Faults Response should match fixture",
			want: 18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			//			gock.Observe(gock.DumpRequest)
			gock.New(sfHost).
				Post(sfRPCEndpoint).
				MatchType("json").
				JSON(solidfire.RPCBody{
					ID:     1,
					Method: solidfire.RPCListClusterFaults,
					Params: solidfire.ListClusterFaultsRPCParams{
						FaultTypes:    "current",
						BestPractices: true,
					}}).
				Reply(200).
				BodyString(string(fixture))

			gotRaw, err := sfClient.ListClusterFaults(context.Background())
			got := gotRaw.Result.Faults[0].ClusterFaultID

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListClusterFaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ListClusterFaults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ListNodeStats(t *testing.T) {
	fixture, err := ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, solidfire.RPCListNodeStats))
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    float64
		wantErr bool
	}{
		{
			name: "ListNodeStats Response should match fixture",
			want: 282366,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			//			gock.Observe(gock.DumpRequest)
			gock.New(sfHost).
				Post(sfRPCEndpoint).
				MatchType("json").
				JSON(solidfire.RPCBody{
					ID:     1,
					Method: solidfire.RPCListNodeStats,
					Params: solidfire.ListNodeStatsRPCParams{}}).
				Reply(200).
				BodyString(string(fixture))

			gotRaw, err := sfClient.ListNodeStats(context.Background())
			got := gotRaw.Result.NodeStats.Nodes[0].CBytesIn

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListNodeStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ListNodeStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ListVolumeQoSHistograms(t *testing.T) {
	fixture, err := ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, solidfire.RPCListVolumeQoSHistograms))
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    int
		wantErr bool
	}{
		{
			name: "ListVolumeQoSHistograms Response should match fixture",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			//			gock.Observe(gock.DumpRequest)
			gock.New(sfHost).
				Post(sfRPCEndpoint).
				MatchType("json").
				JSON(solidfire.RPCBody{
					ID:     1,
					Method: solidfire.RPCListVolumeQoSHistograms,
					Params: solidfire.ListVolumeQoSHistogramsRPCParams{
						VolumeIDs: []int{}, // blank gives us all of them
					}}).
				Reply(200).
				BodyString(string(fixture))

			gotRaw, err := sfClient.ListVolumeQoSHistograms(context.Background())
			got := gotRaw.Result.QosHistograms[0].VolumeID

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListVolumeQoSHistograms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ListVolumeQoSHistograms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ListAllNodes(t *testing.T) {
	fixture, err := ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, solidfire.RPCListAllNodes))
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    int
		wantErr bool
	}{
		{
			name: "ListAllNodes Response should match fixture",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			//			gock.Observe(gock.DumpRequest)
			gock.New(sfHost).
				Post(sfRPCEndpoint).
				MatchType("json").
				JSON(solidfire.RPCBody{
					ID:     1,
					Method: solidfire.RPCListAllNodes,
					Params: solidfire.ListAllNodesRPCParams{}}).
				Reply(200).
				BodyString(string(fixture))

			gotRaw, err := sfClient.ListAllNodes(context.Background())
			got := gotRaw.Result.Nodes[0].NodeID

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListAllNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ListAllNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetClusterStats(t *testing.T) {
	fixture, err := ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, solidfire.RPCGetClusterStats))
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    float64
		wantErr bool
	}{
		{
			name: "GetClusterStats Response should match fixture",
			want: 11092150,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			//			gock.Observe(gock.DumpRequest)
			gock.New(sfHost).
				Post(sfRPCEndpoint).
				MatchType("json").
				JSON(solidfire.RPCBody{
					ID:     1,
					Method: solidfire.RPCGetClusterStats,
					Params: solidfire.GetClusterStatsRPCParams{}}).
				Reply(200).
				BodyString(string(fixture))

			gotRaw, err := sfClient.GetClusterStats(context.Background())
			got := gotRaw.Result.ClusterStats.ReadOps

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetClusterStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetClusterStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetClusterFullThreshold(t *testing.T) {
	fixture, err := ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, solidfire.RPCGetClusterFullThreshold))
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    float64
		wantErr bool
	}{
		{
			name: "GetClusterFullThreshold Response should match fixture",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			//			gock.Observe(gock.DumpRequest)
			gock.New(sfHost).
				Post(sfRPCEndpoint).
				MatchType("json").
				JSON(solidfire.RPCBody{
					ID:     1,
					Method: solidfire.RPCGetClusterFullThreshold,
					Params: solidfire.GetClusterFullThresholdParams{}}).
				Reply(200).
				BodyString(string(fixture))

			gotRaw, err := sfClient.GetClusterFullThreshold(context.Background())
			got := gotRaw.Result.Stage2AwareThreshold

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetClusterFullThreshold() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetClusterFullThreshold() = %v, want %v", got, tt.want)
			}
		})
	}
}
