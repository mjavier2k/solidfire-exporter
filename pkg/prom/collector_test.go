package prom_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"testing"

	"github.com/mjavier2k/solidfire-exporter/pkg/prom"
	"github.com/mjavier2k/solidfire-exporter/pkg/solidfire"
	"github.com/mjavier2k/solidfire-exporter/pkg/testutils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

var (
	fixtureBasePath = path.Join("..", "..", "test", "fixtures")
)

func Test_GigabytesToBytes(t *testing.T) {
	type args struct {
		gb float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"gb to bytes",
			args{
				gb: 2,
			},
			2e+9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prom.GigabytesToBytes(tt.args.gb); got != tt.want {
				t.Errorf("GigabytesToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MicrosecondsToSeconds(t *testing.T) {
	type args struct {
		microSeconds float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"µs to s",
			args{
				microSeconds: 2e+6,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prom.MicrosecondsToSeconds(tt.args.microSeconds); got != tt.want {
				t.Errorf("MicrosecondsToSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MillisecondsToSeconds(t *testing.T) {
	type args struct {
		milliseconds float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"ms to s",
			args{
				milliseconds: 2e+3,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prom.MillisecondsToSeconds(tt.args.milliseconds); got != tt.want {
				t.Errorf("MillisecondsToSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Collect(t *testing.T) {
	mockSfClient := new(testutils.MockSolidfireClient)

	var getClusterCapacityResponse = solidfire.GetClusterCapacityResponse{}
	var call solidfire.RPC = solidfire.RPCGetClusterCapacity
	bytes, err := ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, call))
	assert.NoError(t, err)
	assert.NoError(t, json.Unmarshal(bytes, &getClusterCapacityResponse))
	mockSfClient.On(string(call)).Return(getClusterCapacityResponse, nil)

	getClusterFullThresholdResponse := solidfire.GetClusterFullThresholdResponse{}
	call = solidfire.RPCGetClusterFullThreshold
	bytes, err = ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, call))
	assert.NoError(t, err)
	assert.NoError(t, json.Unmarshal(bytes, &getClusterFullThresholdResponse))
	mockSfClient.On(string(call)).Return(getClusterFullThresholdResponse, nil)

	getClusterStatsResponse := solidfire.GetClusterStatsResponse{}
	call = solidfire.RPCGetClusterStats
	bytes, err = ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, call))
	assert.NoError(t, err)
	assert.NoError(t, json.Unmarshal(bytes, &getClusterStatsResponse))
	mockSfClient.On(string(call)).Return(getClusterStatsResponse, nil)

	listAllNodesResponse := solidfire.ListAllNodesResponse{}
	call = solidfire.RPCListAllNodes
	bytes, err = ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, call))
	assert.NoError(t, err)
	assert.NoError(t, json.Unmarshal(bytes, &listAllNodesResponse))
	mockSfClient.On(string(call)).Return(listAllNodesResponse, nil)

	listClusterFaultsResponse := solidfire.ListClusterFaultsResponse{}
	call = solidfire.RPCListClusterFaults
	bytes, err = ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, call))
	assert.NoError(t, err)
	assert.NoError(t, json.Unmarshal(bytes, &listClusterFaultsResponse))
	mockSfClient.On(string(call)).Return(listClusterFaultsResponse, nil)

	listDrivesResponse := solidfire.ListDrivesResponse{}
	call = solidfire.RPCListDrives
	bytes, err = ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, call))
	assert.NoError(t, err)
	assert.NoError(t, json.Unmarshal(bytes, &listDrivesResponse))
	mockSfClient.On(string(call)).Return(listDrivesResponse, nil)

	listISCSISessionsResponse := solidfire.ListISCSISessionsResponse{}
	call = solidfire.RPCListISCSISessions
	bytes, err = ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, call))
	assert.NoError(t, err)
	assert.NoError(t, json.Unmarshal(bytes, &listISCSISessionsResponse))
	mockSfClient.On(string(call)).Return(listISCSISessionsResponse, nil)

	listNodeStatsResponse := solidfire.ListNodeStatsResponse{}
	call = solidfire.RPCListNodeStats
	bytes, err = ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, call))
	assert.NoError(t, err)
	assert.NoError(t, json.Unmarshal(bytes, &listNodeStatsResponse))
	mockSfClient.On(string(call)).Return(listNodeStatsResponse, nil)

	listVolumeQoSHistogramsResponse := solidfire.ListVolumeQoSHistogramsResponse{}
	call = solidfire.RPCListVolumeQoSHistograms
	bytes, err = ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, call))
	assert.NoError(t, err)
	assert.NoError(t, json.Unmarshal(bytes, &listVolumeQoSHistogramsResponse))
	mockSfClient.On(string(call)).Return(listVolumeQoSHistogramsResponse, nil)

	listVolumesResponse := solidfire.ListVolumesResponse{}
	call = solidfire.RPCListVolumes
	bytes, err = ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, call))
	assert.NoError(t, err)
	assert.NoError(t, json.Unmarshal(bytes, &listVolumesResponse))
	mockSfClient.On(string(call)).Return(listVolumesResponse, nil)

	listVolumeStatsResponse := solidfire.ListVolumeStatsResponse{}
	call = solidfire.RPCListVolumeStats
	bytes, err = ioutil.ReadFile(testutils.ResolveFixturePath(fixtureBasePath, call))
	assert.NoError(t, err)
	assert.NoError(t, json.Unmarshal(bytes, &listVolumeStatsResponse))
	mockSfClient.On(string(call)).Return(listVolumeStatsResponse, nil)

	collector, err := prom.NewCollector(mockSfClient)
	if err != nil {
		t.Errorf("error initializing collector: %s", err.Error())
	}
	r := prometheus.NewRegistry()
	r.MustRegister(collector)
	got := prometheusOutput(t, r, "solidfire")

	assert.Equal(t, expectedOutputSlice, got, fmt.Sprintf("Here is the full output I got from the collector:\n%s\n", strings.Join(got, "\n")))
}
