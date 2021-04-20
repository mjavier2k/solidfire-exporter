package testutils

import (
	"github.com/mjavier2k/solidfire-exporter/pkg/solidfire"
	"github.com/stretchr/testify/mock"
)

type MockSolidfireClient struct {
	mock.Mock
}

func (m *MockSolidfireClient) GetClusterCapacity() (solidfire.GetClusterCapacityResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.GetClusterCapacityResponse), args.Error(1)
}
func (m *MockSolidfireClient) GetClusterFullThreshold() (solidfire.GetClusterFullThresholdResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.GetClusterFullThresholdResponse), args.Error(1)
}
func (m *MockSolidfireClient) GetClusterStats() (solidfire.GetClusterStatsResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.GetClusterStatsResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListAllNodes() (solidfire.ListAllNodesResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListAllNodesResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListClusterFaults() (solidfire.ListClusterFaultsResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListClusterFaultsResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListDrives() (solidfire.ListDrivesResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListDrivesResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListISCSISessions() (solidfire.ListISCSISessionsResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListISCSISessionsResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListNodeStats() (solidfire.ListNodeStatsResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListNodeStatsResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListVolumeQoSHistograms() (solidfire.ListVolumeQoSHistogramsResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListVolumeQoSHistogramsResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListVolumes() (solidfire.ListVolumesResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListVolumesResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListVolumeStats() (solidfire.ListVolumeStatsResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListVolumeStatsResponse), args.Error(1)
}
