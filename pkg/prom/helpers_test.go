package prom_test

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
)

func prometheusOutput(t *testing.T, registry *prometheus.Registry, prefix string) []string {
	t.Helper()

	server := httptest.NewServer(promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	defer server.Close()
	resp, err := http.Get(server.URL)
	assert.NoError(t, err)

	var selected []string
	s := bufio.NewScanner(resp.Body)
	for s.Scan() {
		if strings.HasPrefix(s.Text(), prefix) {
			// fmt.Println(s.Text())
			selected = append(selected, s.Text())
		}
	}

	return selected
}

var expectedOutputSlice = strings.Split(strings.TrimSpace(`
solidfire_cluster_active_block_space_bytes 4.977419581e+09
solidfire_cluster_active_faults{code="driveAvailable",details="Node ID 1 has 1 available drive(s).",drive_id="0.000000",node_hardware_fault_id="0.000000",node_id="1",node_name="n01",resolved="false",service_id="0.000000",severity="warning",type="drive"} 1
solidfire_cluster_active_sessions 1
solidfire_cluster_average_io_bytes 0
solidfire_cluster_average_iops 0
solidfire_cluster_block_fullness{level="stage1Happy"} 0
solidfire_cluster_block_fullness{level="stage2Aware"} 1
solidfire_cluster_block_fullness{level="stage3Low"} 0
solidfire_cluster_block_fullness{level="stage4Critical"} 0
solidfire_cluster_block_fullness{level="stage5CompletelyConsumed"} 0
solidfire_cluster_client_queue_depth 0
solidfire_cluster_compression_factor 2.094133784391091
solidfire_cluster_current_iops 0
solidfire_cluster_de_duplication_factor 1.0000545044935927
solidfire_cluster_efficiency_factor 18.580522988214764
solidfire_cluster_fullness{level="blockFullness"} 0
solidfire_cluster_fullness{level="metadataFullness"} 0
solidfire_cluster_iops 0
solidfire_cluster_iops_total 2.4181537e+07
solidfire_cluster_last_sample_read_bytes 0
solidfire_cluster_last_sample_read_ops 0
solidfire_cluster_last_sample_write_bytes 0
solidfire_cluster_last_sample_write_ops 0
solidfire_cluster_latency_seconds 0
solidfire_cluster_max_iops 3000
solidfire_cluster_max_metadata_over_provision_factor 5
solidfire_cluster_max_over_provisionable_space_bytes 1.855425871872e+13
solidfire_cluster_max_provisioned_space_bytes 3.710851743744e+12
solidfire_cluster_max_used_metadata_space_bytes 1.4495514624e+10
solidfire_cluster_max_used_space_bytes 1.073741824e+11
solidfire_cluster_metadata_fullness{level="stage1Happy"} 1
solidfire_cluster_metadata_fullness{level="stage2Aware"} 0
solidfire_cluster_metadata_fullness{level="stage3Low"} 0
solidfire_cluster_metadata_fullness{level="stage4Critical"} 0
solidfire_cluster_metadata_fullness{level="stage5CompletelyConsumed"} 0
solidfire_cluster_non_zero_blocks 165133
solidfire_cluster_normalized_iops 0
solidfire_cluster_peak_active_sessions 1
solidfire_cluster_peak_iops 6
solidfire_cluster_provisioned_space_bytes 6.001000448e+09
solidfire_cluster_read_bytes_total 4.5445102592e+10
solidfire_cluster_read_latency_seconds 0
solidfire_cluster_read_latency_seconds_total 0
solidfire_cluster_read_ops_total 1.109215e+07
solidfire_cluster_recent_io_size_bytes 0
solidfire_cluster_sample_period_seconds 0.5
solidfire_cluster_services_expected 1
solidfire_cluster_services_running 1
solidfire_cluster_slice_reserve_used_threshold_percentage 5
solidfire_cluster_snapshot_non_zero_blocks 0
solidfire_cluster_stage2_aware_threshold_percentage 3
solidfire_cluster_stage2_block_threshold_bytes 0
solidfire_cluster_stage3_block_threshold_bytes 9.8784247808e+10
solidfire_cluster_stage3_block_threshold_percentage 3
solidfire_cluster_stage3_low_threshold_percentage 2
solidfire_cluster_stage4_block_threshold_bytes 1.0200547328e+11
solidfire_cluster_stage4_critical_threshold_percentage 1
solidfire_cluster_stage5_block_threshold_bytes 1.073741824e+11
solidfire_cluster_thin_provisioning_factor 8.872169705631219
solidfire_cluster_throughput_utilization 0
solidfire_cluster_total_bytes 1.073741824e+11
solidfire_cluster_total_metadata_bytes 1.4495514624e+10
solidfire_cluster_unaligned_reads_total 13
solidfire_cluster_unaligned_writes_total 0
solidfire_cluster_unique_blocks 165124
solidfire_cluster_unique_blocks_used_space_bytes 3.47282402e+08
solidfire_cluster_used_bytes 3.47282402e+08
solidfire_cluster_used_metadata_bytes 7.221248e+06
solidfire_cluster_used_metadata_space_bytes 7.221248e+06
solidfire_cluster_used_metadata_space_in_snapshots_bytes 7.221248e+06
solidfire_cluster_used_space_bytes 3.47282402e+08
solidfire_cluster_write_bytes_total 1.21720639488e+11
solidfire_cluster_write_latency_seconds 0
solidfire_cluster_write_latency_seconds_total 0
solidfire_cluster_write_ops_total 1.3089387e+07
solidfire_cluster_zero_blocks 1.299955e+06
solidfire_drive_capacity_bytes{drive_id="1",node_id="1",node_name="n01",serial="sdb",slot="1",type="volume"} 3.221225472e+10
solidfire_drive_capacity_bytes{drive_id="2",node_id="1",node_name="n01",serial="sdc",slot="2",type="block"} 5.36870912e+10
solidfire_drive_capacity_bytes{drive_id="3",node_id="1",node_name="n01",serial="sdd",slot="3",type="block"} 5.36870912e+10
solidfire_drive_capacity_bytes{drive_id="4",node_id="1",node_name="n01",serial="sde",slot="4",type="block"} 5.36870912e+10
solidfire_drive_status{drive_id="1",node_id="1",node_name="n01",serial="sdb",slot="1",status="active",type="volume"} 1
solidfire_drive_status{drive_id="1",node_id="1",node_name="n01",serial="sdb",slot="1",status="available",type="volume"} 0
solidfire_drive_status{drive_id="1",node_id="1",node_name="n01",serial="sdb",slot="1",status="erasing",type="volume"} 0
solidfire_drive_status{drive_id="1",node_id="1",node_name="n01",serial="sdb",slot="1",status="failed",type="volume"} 0
solidfire_drive_status{drive_id="1",node_id="1",node_name="n01",serial="sdb",slot="1",status="removing",type="volume"} 0
solidfire_drive_status{drive_id="2",node_id="1",node_name="n01",serial="sdc",slot="2",status="active",type="block"} 0
solidfire_drive_status{drive_id="2",node_id="1",node_name="n01",serial="sdc",slot="2",status="available",type="block"} 1
solidfire_drive_status{drive_id="2",node_id="1",node_name="n01",serial="sdc",slot="2",status="erasing",type="block"} 0
solidfire_drive_status{drive_id="2",node_id="1",node_name="n01",serial="sdc",slot="2",status="failed",type="block"} 0
solidfire_drive_status{drive_id="2",node_id="1",node_name="n01",serial="sdc",slot="2",status="removing",type="block"} 0
solidfire_drive_status{drive_id="3",node_id="1",node_name="n01",serial="sdd",slot="3",status="active",type="block"} 1
solidfire_drive_status{drive_id="3",node_id="1",node_name="n01",serial="sdd",slot="3",status="available",type="block"} 0
solidfire_drive_status{drive_id="3",node_id="1",node_name="n01",serial="sdd",slot="3",status="erasing",type="block"} 0
solidfire_drive_status{drive_id="3",node_id="1",node_name="n01",serial="sdd",slot="3",status="failed",type="block"} 0
solidfire_drive_status{drive_id="3",node_id="1",node_name="n01",serial="sdd",slot="3",status="removing",type="block"} 0
solidfire_drive_status{drive_id="4",node_id="1",node_name="n01",serial="sde",slot="4",status="active",type="block"} 1
solidfire_drive_status{drive_id="4",node_id="1",node_name="n01",serial="sde",slot="4",status="available",type="block"} 0
solidfire_drive_status{drive_id="4",node_id="1",node_name="n01",serial="sde",slot="4",status="erasing",type="block"} 0
solidfire_drive_status{drive_id="4",node_id="1",node_name="n01",serial="sde",slot="4",status="failed",type="block"} 0
solidfire_drive_status{drive_id="4",node_id="1",node_name="n01",serial="sde",slot="4",status="removing",type="block"} 0
solidfire_node_cpu_percentage{node_id="1",node_name="n01"} 0
solidfire_node_cpu_seconds_total{node_id="1",node_name="n01"} 2247
solidfire_node_info{associated_fservice_id="0",associated_master_service_id="1",chassis_name="",chassis_type="SFVIRT",cpu_model="Intel(R) Xeon(R) CPU E7-8891 v4 @ 2.80GHz\nUnknown Processor",node_id="1",node_name="n01",node_type="SFDEMO-NE",platform_config_version="0.0.0.0",sip="10.0.0.91",sipi="eth1",software_version="11.7.0.76",uuid="5329FE1F-A41F-DC41-8BD9-2016FF2DD8FF"} 1
solidfire_node_interface_in_bytes_total{interface="cluster",node_id="1",node_name="n01"} 282366
solidfire_node_interface_in_bytes_total{interface="management",node_id="1",node_name="n01"} 332883
solidfire_node_interface_in_bytes_total{interface="storage",node_id="1",node_name="n01"} 282366
solidfire_node_interface_out_bytes_total{interface="cluster",node_id="1",node_name="n01"} 59773
solidfire_node_interface_out_bytes_total{interface="management",node_id="1",node_name="n01"} 130104
solidfire_node_interface_out_bytes_total{interface="storage",node_id="1",node_name="n01"} 59773
solidfire_node_interface_utilization_percentage{interface="cluster",node_id="1",node_name="n01"} 0
solidfire_node_interface_utilization_percentage{interface="storage",node_id="1",node_name="n01"} 0
solidfire_node_iscsi_sessions{node_id="1",node_name="n01",volume_id="1",volume_name="test-volume1"} 1
solidfire_node_load_bucket{node_id="1",node_name="n01",le="0"} 1.205996e+06
solidfire_node_load_bucket{node_id="1",node_name="n01",le="19"} 4.606744e+06
solidfire_node_load_bucket{node_id="1",node_name="n01",le="39"} 1.192471e+06
solidfire_node_load_bucket{node_id="1",node_name="n01",le="59"} 89
solidfire_node_load_bucket{node_id="1",node_name="n01",le="79"} 0
solidfire_node_load_bucket{node_id="1",node_name="n01",le="100"} 0
solidfire_node_load_bucket{node_id="1",node_name="n01",le="+Inf"} 294
solidfire_node_load_sum{node_id="1",node_name="n01"} 7.0053e+06
solidfire_node_load_count{node_id="1",node_name="n01"} 294
solidfire_node_read_latency_seconds_total{node_id="1",node_name="n01"} 0
solidfire_node_read_ops_total{node_id="1",node_name="n01"} 1.109215e+07
solidfire_node_samples{node_id="1",node_name="n01"} 294
solidfire_node_total_memory_bytes{node_id="1",node_name="n01"} 1.6e+10
solidfire_node_used_memory_bytes{node_id="1",node_name="n01"} 9.000198144e+09
solidfire_node_write_latency_seconds_total{node_id="1",node_name="n01"} 0
solidfire_node_write_ops_total{node_id="1",node_name="n01"} 1.3089387e+07
solidfire_up 1
solidfire_volume_actual_iops{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_actual_iops{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_average_iop_size_bytes{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_average_iop_size_bytes{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_burst_iops_credit{volume_id="1",volume_name="test-volume1"} 600000
solidfire_volume_burst_iops_credit{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_client_queue_depth{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_client_queue_depth{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_last_sample_read_bytes{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_last_sample_read_bytes{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_last_sample_read_ops{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_last_sample_read_ops{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_last_sample_write_bytes{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_last_sample_write_bytes{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_latency_seconds{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_latency_seconds{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_non_zero_blocks{volume_id="1",volume_name="test-volume1"} 165133
solidfire_volume_non_zero_blocks{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_qos_below_min_iops_percentage_bucket{volume_id="1",volume_name="test-volume1",le="19"} 32
solidfire_volume_qos_below_min_iops_percentage_bucket{volume_id="1",volume_name="test-volume1",le="39"} 6
solidfire_volume_qos_below_min_iops_percentage_bucket{volume_id="1",volume_name="test-volume1",le="59"} 4
solidfire_volume_qos_below_min_iops_percentage_bucket{volume_id="1",volume_name="test-volume1",le="79"} 2
solidfire_volume_qos_below_min_iops_percentage_bucket{volume_id="1",volume_name="test-volume1",le="100"} 4
solidfire_volume_qos_below_min_iops_percentage_bucket{volume_id="1",volume_name="test-volume1",le="+Inf"} 0
solidfire_volume_qos_below_min_iops_percentage_sum{volume_id="1",volume_name="test-volume1"} 48
solidfire_volume_qos_below_min_iops_percentage_count{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_qos_below_min_iops_percentage_bucket{volume_id="2",volume_name="test-volume2",le="19"} 0
solidfire_volume_qos_below_min_iops_percentage_bucket{volume_id="2",volume_name="test-volume2",le="39"} 0
solidfire_volume_qos_below_min_iops_percentage_bucket{volume_id="2",volume_name="test-volume2",le="59"} 0
solidfire_volume_qos_below_min_iops_percentage_bucket{volume_id="2",volume_name="test-volume2",le="79"} 0
solidfire_volume_qos_below_min_iops_percentage_bucket{volume_id="2",volume_name="test-volume2",le="100"} 0
solidfire_volume_qos_below_min_iops_percentage_bucket{volume_id="2",volume_name="test-volume2",le="+Inf"} 0
solidfire_volume_qos_below_min_iops_percentage_sum{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_qos_below_min_iops_percentage_count{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_qos_min_to_max_iops_percentage_bucket{volume_id="1",volume_name="test-volume1",le="19"} 167
solidfire_volume_qos_min_to_max_iops_percentage_bucket{volume_id="1",volume_name="test-volume1",le="39"} 3823
solidfire_volume_qos_min_to_max_iops_percentage_bucket{volume_id="1",volume_name="test-volume1",le="59"} 2304
solidfire_volume_qos_min_to_max_iops_percentage_bucket{volume_id="1",volume_name="test-volume1",le="79"} 5867
solidfire_volume_qos_min_to_max_iops_percentage_bucket{volume_id="1",volume_name="test-volume1",le="100"} 28539
solidfire_volume_qos_min_to_max_iops_percentage_bucket{volume_id="1",volume_name="test-volume1",le="+Inf"} 6162
solidfire_volume_qos_min_to_max_iops_percentage_sum{volume_id="1",volume_name="test-volume1"} 46862
solidfire_volume_qos_min_to_max_iops_percentage_count{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_qos_min_to_max_iops_percentage_bucket{volume_id="2",volume_name="test-volume2",le="19"} 0
solidfire_volume_qos_min_to_max_iops_percentage_bucket{volume_id="2",volume_name="test-volume2",le="39"} 0
solidfire_volume_qos_min_to_max_iops_percentage_bucket{volume_id="2",volume_name="test-volume2",le="59"} 0
solidfire_volume_qos_min_to_max_iops_percentage_bucket{volume_id="2",volume_name="test-volume2",le="79"} 0
solidfire_volume_qos_min_to_max_iops_percentage_bucket{volume_id="2",volume_name="test-volume2",le="100"} 0
solidfire_volume_qos_min_to_max_iops_percentage_bucket{volume_id="2",volume_name="test-volume2",le="+Inf"} 0
solidfire_volume_qos_min_to_max_iops_percentage_sum{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_qos_min_to_max_iops_percentage_count{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_qos_read_block_sizes_bytes_bucket{volume_id="1",volume_name="test-volume1",le="8191"} 1.1091915e+07
solidfire_volume_qos_read_block_sizes_bytes_bucket{volume_id="1",volume_name="test-volume1",le="16383"} 27
solidfire_volume_qos_read_block_sizes_bytes_bucket{volume_id="1",volume_name="test-volume1",le="32767"} 78
solidfire_volume_qos_read_block_sizes_bytes_bucket{volume_id="1",volume_name="test-volume1",le="65535"} 62
solidfire_volume_qos_read_block_sizes_bytes_bucket{volume_id="1",volume_name="test-volume1",le="131071"} 39
solidfire_volume_qos_read_block_sizes_bytes_bucket{volume_id="1",volume_name="test-volume1",le="+Inf"} 16
solidfire_volume_qos_read_block_sizes_bytes_sum{volume_id="1",volume_name="test-volume1"} 1.1092137e+07
solidfire_volume_qos_read_block_sizes_bytes_count{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_qos_read_block_sizes_bytes_bucket{volume_id="2",volume_name="test-volume2",le="8191"} 0
solidfire_volume_qos_read_block_sizes_bytes_bucket{volume_id="2",volume_name="test-volume2",le="16383"} 0
solidfire_volume_qos_read_block_sizes_bytes_bucket{volume_id="2",volume_name="test-volume2",le="32767"} 0
solidfire_volume_qos_read_block_sizes_bytes_bucket{volume_id="2",volume_name="test-volume2",le="65535"} 0
solidfire_volume_qos_read_block_sizes_bytes_bucket{volume_id="2",volume_name="test-volume2",le="131071"} 0
solidfire_volume_qos_read_block_sizes_bytes_bucket{volume_id="2",volume_name="test-volume2",le="+Inf"} 0
solidfire_volume_qos_read_block_sizes_bytes_sum{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_qos_read_block_sizes_bytes_count{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="1",volume_name="test-volume1",le="0"} 5.755624e+06
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="1",volume_name="test-volume1",le="19"} 157
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="1",volume_name="test-volume1",le="39"} 3778
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="1",volume_name="test-volume1",le="59"} 2277
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="1",volume_name="test-volume1",le="79"} 5812
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="1",volume_name="test-volume1",le="100"} 28690
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="1",volume_name="test-volume1",le="+Inf"} 6162
solidfire_volume_qos_target_utilization_percentage_sum{volume_id="1",volume_name="test-volume1"} 5.8025e+06
solidfire_volume_qos_target_utilization_percentage_count{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="2",volume_name="test-volume2",le="0"} 0
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="2",volume_name="test-volume2",le="19"} 0
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="2",volume_name="test-volume2",le="39"} 0
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="2",volume_name="test-volume2",le="59"} 0
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="2",volume_name="test-volume2",le="79"} 0
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="2",volume_name="test-volume2",le="100"} 0
solidfire_volume_qos_target_utilization_percentage_bucket{volume_id="2",volume_name="test-volume2",le="+Inf"} 0
solidfire_volume_qos_target_utilization_percentage_sum{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_qos_target_utilization_percentage_count{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="1",volume_name="test-volume1",le="0"} 5.8025e+06
solidfire_volume_qos_throttle_percentage_bucket{volume_id="1",volume_name="test-volume1",le="19"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="1",volume_name="test-volume1",le="39"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="1",volume_name="test-volume1",le="59"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="1",volume_name="test-volume1",le="79"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="1",volume_name="test-volume1",le="100"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="1",volume_name="test-volume1",le="+Inf"} 0
solidfire_volume_qos_throttle_percentage_sum{volume_id="1",volume_name="test-volume1"} 5.8025e+06
solidfire_volume_qos_throttle_percentage_count{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="2",volume_name="test-volume2",le="0"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="2",volume_name="test-volume2",le="19"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="2",volume_name="test-volume2",le="39"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="2",volume_name="test-volume2",le="59"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="2",volume_name="test-volume2",le="79"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="2",volume_name="test-volume2",le="100"} 0
solidfire_volume_qos_throttle_percentage_bucket{volume_id="2",volume_name="test-volume2",le="+Inf"} 0
solidfire_volume_qos_throttle_percentage_sum{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_qos_throttle_percentage_count{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_qos_write_block_sizes_bytes_bucket{volume_id="1",volume_name="test-volume1",le="8191"} 6.859016e+06
solidfire_volume_qos_write_block_sizes_bytes_bucket{volume_id="1",volume_name="test-volume1",le="16383"} 4.182242e+06
solidfire_volume_qos_write_block_sizes_bytes_bucket{volume_id="1",volume_name="test-volume1",le="32767"} 1.634812e+06
solidfire_volume_qos_write_block_sizes_bytes_bucket{volume_id="1",volume_name="test-volume1",le="65535"} 368092
solidfire_volume_qos_write_block_sizes_bytes_bucket{volume_id="1",volume_name="test-volume1",le="131071"} 43454
solidfire_volume_qos_write_block_sizes_bytes_bucket{volume_id="1",volume_name="test-volume1",le="+Inf"} 1771
solidfire_volume_qos_write_block_sizes_bytes_sum{volume_id="1",volume_name="test-volume1"} 1.3089387e+07
solidfire_volume_qos_write_block_sizes_bytes_count{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_qos_write_block_sizes_bytes_bucket{volume_id="2",volume_name="test-volume2",le="8191"} 0
solidfire_volume_qos_write_block_sizes_bytes_bucket{volume_id="2",volume_name="test-volume2",le="16383"} 0
solidfire_volume_qos_write_block_sizes_bytes_bucket{volume_id="2",volume_name="test-volume2",le="32767"} 0
solidfire_volume_qos_write_block_sizes_bytes_bucket{volume_id="2",volume_name="test-volume2",le="65535"} 0
solidfire_volume_qos_write_block_sizes_bytes_bucket{volume_id="2",volume_name="test-volume2",le="131071"} 0
solidfire_volume_qos_write_block_sizes_bytes_bucket{volume_id="2",volume_name="test-volume2",le="+Inf"} 0
solidfire_volume_qos_write_block_sizes_bytes_sum{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_qos_write_block_sizes_bytes_count{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_read_bytes_total{volume_id="1",volume_name="test-volume1"} 4.5445102592e+10
solidfire_volume_read_bytes_total{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_read_latency_seconds{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_read_latency_seconds{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_read_latency_seconds_total{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_read_latency_seconds_total{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_read_ops_total{volume_id="1",volume_name="test-volume1"} 1.109215e+07
solidfire_volume_read_ops_total{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_size_bytes{volume_id="1",volume_name="test-volume1"} 2.000683008e+09
solidfire_volume_size_bytes{volume_id="2",volume_name="test-volume2"} 4.00031744e+09
solidfire_volume_throttle{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_throttle{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_unaligned_reads_total{volume_id="1",volume_name="test-volume1"} 13
solidfire_volume_unaligned_reads_total{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_unaligned_writes_total{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_unaligned_writes_total{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_utilization{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_utilization{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_write_bytes_total{volume_id="1",volume_name="test-volume1"} 1.21720639488e+11
solidfire_volume_write_bytes_total{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_write_latency_seconds{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_write_latency_seconds{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_write_latency_seconds_total{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_write_latency_seconds_total{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_write_ops_last_sample{volume_id="1",volume_name="test-volume1"} 0
solidfire_volume_write_ops_last_sample{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_write_ops_total{volume_id="1",volume_name="test-volume1"} 1.3089387e+07
solidfire_volume_write_ops_total{volume_id="2",volume_name="test-volume2"} 0
solidfire_volume_zero_blocks{volume_id="1",volume_name="test-volume1"} 323315
solidfire_volume_zero_blocks{volume_id="2",volume_name="test-volume2"} 976640
`), "\n")
