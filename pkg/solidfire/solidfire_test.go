package solidfire_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/mjavier2k/solidfire-exporter/pkg/solidfire"
	"gopkg.in/h2non/gock.v1"
)

var (
	sfHost        = "https://192.168.1.1"
	sfRPCEndpoint = "/json-rpc/11.3"
	sfClient      = solidfire.Client{
		RPCEndpoint: fmt.Sprintf("%v%v", sfHost, sfRPCEndpoint),
		HttpClient:  &http.Client{},
	}
)

func TestClient_ListVolumeStats(t *testing.T) {
	fixture, err := ioutil.ReadFile("../../test/fixtures/listvolumestats.json")
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    float64
		wantErr bool
	}{
		{
			name: "VolumeSize of first volume should match fixture",
			want: 5000658944,
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
					Method: "ListVolumeStats",
					Params: solidfire.ListVolumeStatsRPCParams{
						VolumeIDs:             []int{},
						IncludeVirtualVolumes: true,
					}}).
				Reply(200).
				BodyString(string(fixture))
			gotRaw, err := sfClient.ListVolumeStats()
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
	fixture, err := ioutil.ReadFile("../../test/fixtures/listvolumes.json")
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    string
		wantErr bool
	}{
		{
			name: "Volume Name of first volume should match fixture",
			want: "testVolume1",
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
					Method: "ListVolumes",
					Params: solidfire.ListVolumesRPCParams{
						IncludeVirtualVolumes: true,
					}}).
				Reply(200).
				BodyString(string(fixture))
			gotRaw, err := sfClient.ListVolumes()
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
	fixture, err := ioutil.ReadFile("../../test/fixtures/getclustercapacity.json")
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    float64
		wantErr bool
	}{
		{
			name: "Cluster Capacity TotalOps should match fixture",
			want: 422890150883,
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
					Method: "GetClusterCapacity",
					Params: solidfire.GetClusterCapacityRPCParams{}}).
				Reply(200).
				BodyString(string(fixture))
			gotRaw, err := sfClient.GetClusterCapacity()
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

func TestClient_ListClusterActiveFaults(t *testing.T) {
	fixture, err := ioutil.ReadFile("../../test/fixtures/listclusterfaults.json")
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    float64
		wantErr bool
	}{
		{
			name: "Cluster Faults Response should match fixture",
			want: 1,
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
					Method: "ListClusterActiveFaults",
					Params: solidfire.ListClusterFaultsRPCParams{
						FaultTypes:    "current",
						BestPractices: true,
					}}).
				Reply(200).
				BodyString(string(fixture))

			gotRaw, err := sfClient.ListClusterActiveFaults()
			got := gotRaw.Result.Faults[0].ClusterFaultID

			severity := map[string]float64{
				solidfire.FaultBestPractice: 0,
				solidfire.FaultWarning:      0,
				solidfire.FaultError:        0,
				solidfire.FaultCritical:     0,
			}

			for _, f := range gotRaw.Result.Faults {
				severity[f.Severity]++
			}

			fmt.Printf("%v", severity)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListClusterActiveFaults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ListClusterActiveFaults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ListNodeStats(t *testing.T) {
	fixture, err := ioutil.ReadFile("../../test/fixtures/listnodestats.json")
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name    string
		s       solidfire.Client
		want    float64
		wantErr bool
	}{
		{
			name: "ListNodeStats Response should match fixture",
			want: 685484495902625,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer gock.Off()
			gock.Observe(gock.DumpRequest)
			gock.New(sfHost).
				Post(sfRPCEndpoint).
				MatchType("json").
				JSON(solidfire.RPCBody{
					ID:     1,
					Method: "ListNodeStats",
					Params: solidfire.ListNodeStatsRPCParams{}}).
				Reply(200).
				BodyString(string(fixture))

			gotRaw, err := sfClient.ListNodeStats()
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
	fixture, err := ioutil.ReadFile("../../test/fixtures/listvolumeqoshistograms.json")
	if err != nil {
		panic(err)
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
			gock.Observe(gock.DumpRequest)
			gock.New(sfHost).
				Post(sfRPCEndpoint).
				MatchType("json").
				JSON(solidfire.RPCBody{
					ID:     1,
					Method: "ListVolumeQoSHistograms",
					Params: solidfire.ListVolumeQoSHistogramsRPCParams{
						VolumeIDs: []int{}, // blank gives us all of them
					}}).
				Reply(200).
				BodyString(string(fixture))

			gotRaw, err := sfClient.ListVolumeQoSHistograms()
			got := gotRaw.Result.QosHistograms[0].VolumeID

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
