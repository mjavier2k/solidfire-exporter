package testutils

import (
	"context"

	"github.com/mjavier2k/solidfire-exporter/pkg/solidfire"
	"github.com/stretchr/testify/mock"
)

type MockSolidfireClient struct {
	mock.Mock
}

func (m *MockSolidfireClient) GetClusterCapacity(ctx context.Context) (solidfire.GetClusterCapacityResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.GetClusterCapacityResponse), args.Error(1)
}
func (m *MockSolidfireClient) GetClusterFullThreshold(ctx context.Context) (solidfire.GetClusterFullThresholdResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.GetClusterFullThresholdResponse), args.Error(1)
}
func (m *MockSolidfireClient) GetClusterStats(ctx context.Context) (solidfire.GetClusterStatsResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.GetClusterStatsResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListAllNodes(ctx context.Context) (solidfire.ListAllNodesResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListAllNodesResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListClusterFaults(ctx context.Context) (solidfire.ListClusterFaultsResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListClusterFaultsResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListDrives(ctx context.Context) (solidfire.ListDrivesResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListDrivesResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListISCSISessions(ctx context.Context) (solidfire.ListISCSISessionsResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListISCSISessionsResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListNodeStats(ctx context.Context) (solidfire.ListNodeStatsResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListNodeStatsResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListVolumeQoSHistograms(ctx context.Context) (solidfire.ListVolumeQoSHistogramsResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListVolumeQoSHistogramsResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListVolumes(ctx context.Context) (solidfire.ListVolumesResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListVolumesResponse), args.Error(1)
}
func (m *MockSolidfireClient) ListVolumeStats(ctx context.Context) (solidfire.ListVolumeStatsResponse, error) {
	args := m.Called()
	return args.Get(0).(solidfire.ListVolumeStatsResponse), args.Error(1)
}
