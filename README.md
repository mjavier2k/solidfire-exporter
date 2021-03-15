# NetApp Solidfire Exporter

![Go](https://github.com/mjavier2k/solidfire-exporter/workflows/Go/badge.svg?event=push)

NetApp Solidfire-Exporter queries the Solidfire API and exports the results as Prometheus metrics


![Cluster Metrics](examples/dashboards/solidfire-cluster.png?raw=true)

<!--  Table of Contents Generated by the VS Code "markdown all in one" extension: yzhang.markdown-all-in-one -->
- [NetApp Solidfire Exporter](#netapp-solidfire-exporter)
  - [Supported ElementOS Versions](#supported-elementos-versions)
  - [Installation](#installation)
  - [Metrics](#metrics)
  - [Usage](#usage)
  - [Configuration](#configuration)
      - [Shell-Like Environments](#shell-like-environments)
      - [Docker-Type Environments / SystemD EnvironmentFile Environment](#docker-type-environments--systemd-environmentfile-environment)
  - [Prometheus Configuration](#prometheus-configuration)
  - [Using Docker](#using-docker)
  - [Grafana Dashboards](#grafana-dashboards)
  - [Contributing](#contributing)
  - [License](#license)

## Supported ElementOS Versions

| API Version | Endpoint | Notes |
|-|-|-|
| [11.3](https://library.netapp.com/ecm/ecm_download_file/ECMLP2856155) | https://your-mgmt-vip/json-rpc/11.3 | There is a known bug in v11.3 where the api user requires <br>Administrator access is required to view the QoS data. |
| [12.2](https://docs.netapp.com/sfe-122/index.jsp) | https://your-mgmt-vip/json-rpc/12.2 |  |

## Installation

Binaries can be downloaded from [Github releases](https://github.com/mjavier2k/solidfire-exporter/releases) page. 


## Metrics

| Metric | Type | Description |
|-|-|-|
| solidfire_cluster_capacity_active_block_space_bytes | gauge | The amount of space on the block drives. This includes additional information such as metadata entries and space which can be cleaned up. |
| solidfire_cluster_capacity_active_sessions | gauge | The number of active iSCSI sessions communicating with the cluster. |
| solidfire_cluster_capacity_average_iops | gauge | The average IOPS for the cluster since midnight Coordinated Universal Time (UTC) |
| solidfire_cluster_capacity_cluster_recent_io_size_bytes | gauge | The average size of IOPS to all volumes in the cluster |
| solidfire_cluster_capacity_compression_factor | gauge | The cluster compression factor. compressionFactor = (uniqueBlocks * 4096) / (uniqueBlocksUsedSpace * 0.93) |
| solidfire_cluster_capacity_current_iops | gauge | The average IOPS for all volumes in the cluster over the last 5 seconds |
| solidfire_cluster_capacity_de_duplication_factor | gauge | The cluster deDuplication factor. deDuplicationFactor = (nonZeroBlocks + snapshotNonZeroBlocks) / uniqueBlocks |
| solidfire_cluster_capacity_efficiency_factor | gauge | The cluster efficiency factor. efficiencyFactor = thinProvisioningFactor * deDuplicationFactor * compressionFactor |
| solidfire_cluster_capacity_max_iops | gauge | The estimated maximum IOPS capability of the current cluster |
| solidfire_cluster_capacity_max_over_provisionable_space_bytes | gauge | The maximum amount of provisionable space. This is a computed value. You cannot create new volumes if the current provisioned space plus the new volume size would exceed this number. The value is calculated as follows: maxOverProvisionableSpace = maxProvisionedSpace * maxMetadataOverProvisionFactor |
| solidfire_cluster_capacity_max_provisioned_space_bytes | gauge | The total amount of provisionable space if all volumes are 100% filled (no thin provisioned metadata) |
| solidfire_cluster_capacity_max_used_metadata_space_bytes | gauge | The number of bytes on volume drives used to store metadata |
| solidfire_cluster_capacity_max_used_space_bytes | gauge | The total amount of space on all active block drives |
| solidfire_cluster_capacity_non_zero_blocks | gauge | The total number of 4KiB blocks that contain data after the last garbage collection operation has completed |
| solidfire_cluster_capacity_peak_active_sessions | gauge | The peak number of iSCSI connections since midnight UTC |
| solidfire_cluster_capacity_peak_iops | gauge | The highest value for currentIOPS since midnight UTC |
| solidfire_cluster_capacity_provisioned_space_bytes | gauge | The total amount of space provisioned in all volumes on the cluster |
| solidfire_cluster_capacity_snapshot_non_zero_blocks | gauge | The total number of 4KiB blocks that contain data after the last garbage collection operation has completed |
| solidfire_cluster_capacity_thin_provisioning_factor | gauge | The cluster thin provisioning factor. thinProvisioningFactor = (nonZeroBlocks + zeroBlocks) / nonZeroBlocks |
| solidfire_cluster_capacity_iops_total | counter | The total number of I/O operations performed throughout the lifetime of the cluster |
| solidfire_cluster_capacity_unique_blocks | gauge | The total number of blocks stored on the block drives The value includes replicated blocks |
| solidfire_cluster_capacity_unique_blocks_used_space_bytes | gauge | The total amount of data the uniqueBlocks take up on the block drives |
| solidfire_cluster_capacity_used_metadata_space_bytes | gauge | The total number of bytes on volume drives used to store metadata |
| solidfire_cluster_capacity_used_metadata_space_in_snapshots_bytes | gauge | The number of bytes on volume drives used for storing unique data in snapshots. This number provides an estimate of how much metadata space would be regained by deleting all snapshots on the system |
| solidfire_cluster_capacity_used_space_bytes | gauge | The total amount of space used by all block drives in the system |
| solidfire_cluster_capacity_zero_blocks | gauge | The total number of empty 4KiB blocks without data after the last round of garbage collection operation has completed |
| solidfire_node_stats_cbytes_in_total | counter | Bytes in on the cluster interface |
| solidfire_node_stats_cbytes_out_total | counter | Bytes out on the cluster interface |
| solidfire_node_stats_samples | gauge | Node stat sample count |
| solidfire_node_stats_cpu_percentage | gauge | CPU usage in percent. |
| solidfire_node_stats_cpu_seconds_total | gauge | CPU usage in seconds since last boot. |
| solidfire_node_stats_load | histogram | System load histogram |
| solidfire_node_stats_mbytes_in_total | counter | Bytes in on the management interface. |
| solidfire_node_stats_mbytes_out_total | counter | Bytes out on the management interface. |
| solidfire_node_stats_network_utilization_cluster_percentage | gauge | Network interface utilization (in percent) for the cluster network interface. |
| solidfire_node_stats_network_utilization_storage_percentage | gauge | Network interface utilization (in percent) for the storage network interface. |
| solidfire_node_stats_read_latency_seconds_total | gauge | The number, in milliseconds, of read latency between clusters. |
| solidfire_node_stats_read_ops | gauge | Read Operations |
| solidfire_node_stats_sbytes_in_total | counter | Bytes in on the storage interface. |
| solidfire_node_stats_sbytes_out_total | counter | Bytes out on the storage interface. |
| solidfire_node_stats_used_memory_bytes | gauge | Total memory usage in bytes. |
| solidfire_node_stats_write_latency_seconds_total | gauge | The number, in milliseconds, of read latency between clusters. |
| solidfire_node_stats_write_ops | gauge | Write Operations |
| solidfire_cluster_stats_read_latency_seconds_total | gauge | The total time spent performing read operations since the creation of the cluster. |
| solidfire_cluster_stats_iops | gauge | Current actual IOPS for the entire cluster in the last 500 milliseconds. |
| solidfire_cluster_stats_average_io_bytes | gauge | Average size in bytes of recent I/O to the cluster in the last 500 milliseconds. |
| solidfire_cluster_stats_client_queue_depth | gauge | The number of outstanding read and write operations to the cluster. |
| solidfire_cluster_stats_latency_seconds | gauge | The average time, in microseconds, to complete operations to a cluster in the last 500 milliseconds. |
| solidfire_cluster_stats_normalized_iops | gauge | Average number of IOPS for the entire cluster in the last 500 milliseconds. |
| solidfire_cluster_stats_read_bytes_total | gauge | The total cumulative bytes read from the cluster since the creation of the cluster. |
| solidfire_cluster_stats_last_sample_read_bytes | gauge | The total number of bytes read from the cluster during the last sample period. |
| solidfire_cluster_stats_read_latency_seconds | gauge | The average time, in microseconds, to complete read operations to the cluster in the last 500 milliseconds. |
| solidfire_cluster_stats_read_ops_total | gauge | The total cumulative read operations to the cluster since the creation of the cluster. |
| solidfire_cluster_stats_last_sample_read_ops | gauge | The total number of read operations during the last sample period. |
| solidfire_cluster_stats_sample_period_seconds | gauge | The length of the sample period, in milliseconds. |
| solidfire_cluster_stats_running_services | gauge | The number of services running on the cluster. If equal to the servicesTotal, this indicates that valid statistics were collected from all nodes. |
| solidfire_cluster_stats_expected_services | gauge | The total number of expected services running on the cluster. |
| solidfire_cluster_stats_unaligned_reads_total | gauge | The total cumulative unaligned read operations to a cluster since the creation of the cluster |
| solidfire_cluster_stats_unaligned_writes_total | gauge | The total cumulative unaligned write operations to a cluster since the creation of the cluster. |
| solidfire_cluster_stats_throughput_utilization | gauge | The cluster capacity being utilized. |
| solidfire_cluster_stats_write_bytes_total | gauge | The total cumulative bytes written to the cluster since the creation of the cluster |
| solidfire_cluster_stats_last_sample_write_bytes | gauge | The total number of bytes written to the cluster during the last sample period. |
| solidfire_cluster_stats_write_latency_seconds | gauge | The average time, in microseconds, to complete write operations to a cluster in the last 500 milliseconds. |
| solidfire_cluster_stats_write_latency_seconds_total | gauge | The total time spent performing write operations since the creation of the cluster. |
| solidfire_cluster_stats_write_ops_total | gauge | The total cumulative write operations to the cluster since the creation of the cluster. |
| solidfire_cluster_stats_last_sample_write_ops | gauge | The total number of write operations during the last sample period. |
| solidfire_cluster_block_fullness | gauge | The current computed level of block fullness of the cluster. |
| solidfire_cluster_fullness | gauge | Reflects the highest level of fullness between 'blockFullness' and 'metadataFullness'. |
| solidfire_cluster_max_metadata_over_provision_factor | gauge | A value representative of the number of times metadata space can be over provisioned relative to the amount of space available. |
| solidfire_cluster_metadata_fullness | gauge | The current computed level of metadata fullness of the cluster. |
| solidfire_cluster_slice_reserve_used_threshold_percentage | gauge | Error condition. A system alert is triggered if the reserved slice utilization is greater than the sliceReserveUsedThresholdPct value returned. |
| solidfire_cluster_stage2_aware_threshold_percentage | gauge | Awareness condition. The value that is set for 'Stage 2' cluster threshold level. |
| solidfire_cluster_stage2_block_threshold_bytes | gauge | Number of bytes being used by the cluster at which a stage2 condition will exist. |
| solidfire_cluster_stage3_block_threshold_bytes | gauge | Number of bytes being used by the cluster at which a stage3 condition will exist. |
| solidfire_cluster_stage3_block_threshold_percentage | gauge | Percent value set for stage3. At this percent full, a warning is posted in the Alerts log. |
| solidfire_cluster_stage3_low_threshold_percentage | gauge | Error condition. The threshold at which a system alert is created due to low capacity on a cluster |
| solidfire_cluster_stage4_block_threshold_bytes | gauge | Number of bytes being used by the cluster at which a stage4 condition will exist |
| solidfire_cluster_stage4_critical_threshold_percentage | gauge | Error condition. The threshold at which a system alert is created to warn about critically low capacity on a cluster. |
| solidfire_cluster_stage5_block_threshold_bytes | gauge | The number of bytes being used by the cluster at which a stage5 condition will exist. |
| solidfire_cluster_total_bytes | gauge | Physical capacity of the cluster, measured in bytes. |
| solidfire_cluster_total_metadata_bytes | gauge | Total amount of space that can be used to store metadata |
| solidfire_cluster_used_bytes | gauge | Number of bytes used on the cluster. |
| solidfire_cluster_used_metadata_bytes | gauge | Amount of space used on volume drives to store metadata. |
| solidfire_drive_capacity_bytes | gauge | The drive capacity for each individual drives in the cluster's active nodes |
| solidfire_drive_status | gauge | The drive status for each individual drives in the cluster's active nodes |
| solidfire_node_iscsi_sessions | gauge | The total number of iscsi sessions per node and volume |
| solidfire_node_info | gauge | Cluster node info |
| solidfire_node_stats_total_memory_bytes | counter | Cluster node total memory in GB |
| solidfire_scrape_success | gauge | Whether last scrape against Solidfire API was successful |
| solidfire_volume_qos_below_min_iops_percentage | histogram | Volume QoS Below minimum IOPS percentage |
| solidfire_volume_qos_min_to_max_iops_percentage | histogram | Volume QoS min to max IOPS percentage |
| solidfire_volume_qos_read_block_sizes | histogram | Volume QoS read block sizes |
| solidfire_volume_qos_target_utilization_percentage | histogram | Volume QoS target utilization percentage |
| solidfire_volume_qos_throttle_percentage | histogram | Volume QoS throttle percentage |
| solidfire_volume_qos_write_block_sizes | histogram | Volume QoS write block sizes |
| solidfire_volume_stats_actual_iops | gauge | The current actual IOPS to the volume in the last 500 milliseconds |
| solidfire_volume_stats_average_iop_size_bytes | gauge | The average size in bytes of recent I/O to the volume in the last 500 milliseconds |
| solidfire_volume_stats_burst_iops_credit | gauge | The total number of IOP credits available to the user. When volumes are not using up to the configured maxIOPS, credits are accrued. |
| solidfire_volume_stats_client_queue_depth | gauge | The number of outstanding read and write operations to the volume. |
| solidfire_volume_stats_latency_seconds | gauge | The average time, in microseconds, to complete operations to the volume in the last 500 milliseconds. A '0' (zero) value means there is no I/O to the volume. |
| solidfire_volume_stats_non_zero_blocks | gauge | The total number of 4KiB blocks that contain data after the last garbage collection operation has completed. |
| solidfire_volume_stats_read_bytes | counter | The total cumulative bytes read from the volume since the creation of the volume. |
| solidfire_volume_stats_last_sample_read_bytes | gauge | The total number of bytes read from the volume during the last sample period. |
| solidfire_volume_stats_read_latency_seconds | gauge | The average time, in microseconds, to complete read operations to the volume in the last 500 milliseconds. |
| solidfire_volume_stats_read_latency_seconds_total | counter | The total time spent performing read operations from the volume |
| solidfire_volume_stats_read_ops_total | counter | The total read operations to the volume since the creation of the volume. |
| solidfire_volume_stats_last_sample_read_ops | gauge | The total number of read operations during the last sample period |
| solidfire_volume_stats_size_bytes | gauge | Total provisioned capacity in bytes. |
| solidfire_volume_stats_throttle | gauge | A floating value between 0 and 1 that represents how much the system is throttling clients below their maxIOPS because of rereplication of data, transient errors, and snapshots taken. |
| solidfire_volume_stats_unaligned_reads_total | counter | The total cumulative unaligned read operations to a volume since the creation of the volume. |
| solidfire_volume_stats_unaligned_writes_total | counter | The total cumulative unaligned write operations to a volume since the creation of the volume. |
| solidfire_volume_stats_utilization | gauge | A floating value that describes how much the client is using the volume. Value 0: The client is not using the volume. Value 1: The client is using their maximum. Value 1+: The client is using their burst. |
| solidfire_volume_stats_write_bytes | counter | The total cumulative bytes written to the volume since the creation of the volume. |
| solidfire_volume_stats_last_sample_write_bytes | gauge | The total number of bytes written to the volume during the last sample period. |
| solidfire_volume_stats_write_latency_seconds | gauge | The average time, in microseconds, to complete write operations to a volume in the last 500 milliseconds. |
| solidfire_volume_stats_write_latency_seconds_total | counter | The total time spent performing write operations to the volume |
| solidfire_volume_stats_write_ops_total | counter | The total cumulative write operations to the volume since the creation of the volume. |
| solidfire_volume_stats_write_ops_last_sample | gauge | The total number of write operations during the last sample period. |
| solidfire_volume_stats_zero_blocks | gauge | The total number of empty 4KiB blocks without data after the last round of garbage collection operation has completed. |

## Usage

```
./solidfire_exporter --config config.yaml
```

```
Usage of solidfire-exporter:
  -c, --config string     Specify configuration filename. (default: config.yaml)
```

## Configuration

| Yaml Configuration Option | CLI Flag | Environment Variable      | Default                         | Example                            | Description                                                                                                                   |
| ------------------------- | -------- | ------------------------- | ------------------------------- | ---------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| client.username           | N/A      | SOLIDFIRE_CLIENT_USERNAME | ""                              | myUsername                         | User with which to authenticate to the Solidfire API. NOTE: User must have administrator access to be able to query QOS data. |
| client.password           | N/A      | SOLIDFIRE_CLIENT_PASSWORD | ""                              | myPassword                         | Password with which to authenticate to the Solidfire API.                                                                     |
| client.endpoint           | N/A      | SOLIDFIRE_CLIENT_ENDPOINT | https://127.0.0.1/json-rpc/11.3 | http://192.168.12.10/json-rpc/11.3 | HTTP(s) endpoint of the Solidfire API.                                                                                        |
| client.insecure           | N/A      | SOLIDFIRE_CLIENT_INSECURE | false                           | true                               | Disables TLS validation when calling Solidfire API. Useful for bypassing self-signed certificates in testing.                 |
| client.timeout            | N/A      | SOLIDFIRE_CLIENT_TIMEOUT  | 30                              | 75                                 | Timeout in seconds per call to the Solidfire API.                                                                             |
| listen.address            | N/A      | SOLIDFIRE_LISTEN_ADDRESS  | 0.0.0.0:9987                    | 192.168.4.2:13987                  | IP address and port where the http server of this exporter should listen                                                      |
| N/A                       | -c       | SOLIDFIRE_CONFIG          | config.yaml                     | mySolidfireConfig.yaml             | Path to configuration file                                                                                                    |

There are two different options to configure the solidfire-exporter

1) Using environment variables. These option takes precedence if config.yaml is also specified.

#### Shell-Like Environments

Note that we *are* allowed to double-quote the values here

```bash
export SOLIDFIRE_CLIENT_USERNAME="mySolidfireUsername"
export SOLIDFIRE_CLIENT_PASSWORD="mySolidfirePassword"
export SOLIDFIRE_LISTEN_ADDRESS="127.0.0.1:9987"
export SOLIDFIRE_CLIENT_ENDPOINT="https://10.10.10.10/json-rpc/11.3"
export SOLIDFIRE_CLIENT_INSECURE=true
export SOLIDFIRE_CLIENT_TIMEOUT=30
```
#### Docker-Type Environments / SystemD EnvironmentFile Environment

Note that we **MAY NOT** have any double quotes in the values below!

```bash
SOLIDFIRE_CLIENT_USERNAME=mySolidfireUsername
SOLIDFIRE_CLIENT_PASSWORD=mySolidfirePassword
SOLIDFIRE_LISTEN_ADDRESS=0.0.0.0:9987
SOLIDFIRE_CLIENT_ENDPOINT=https://10.10.10.10/json-rpc/11.3
SOLIDFIRE_CLIENT_INSECURE=true
SOLIDFIRE_CLIENT_TIMEOUT=30
```

2) Using a config.yaml file

```yaml
listen:
  address: 127.0.0.1:9987
client:
  endpoint: https://192.168.1.2/json-rpc/11.3
  username: mySolidfireUsername
  password: mySolidfirePassword
  insecure: false
  timeout: 130
```

## Prometheus Configuration

```
- job_name: solidfire_exporter
  honor_timestamps: true
  scrape_interval: 30s
  scrape_timeout: 20s
  metrics_path: /metrics
  scheme: http
  static_configs:
  - targets:
    - localhost:9987
    labels:
      app: solidfire-exporter
      group: prometheus
      sfcluster: sfcluster01
```

## Using Docker

Create an file with the environment variables set and pass it to docker run. 

```
docker run --env-file=.env_file  --rm -p 8080:8080 mjavier/solidfire-exporter:latest
```

## Grafana Dashboards

Official Dashboards can be found at:
- [SolidFire Cluster Overview](https://grafana.com/grafana/dashboards/14025)
- [SolidFire Volume Details](https://grafana.com/grafana/dashboards/14030)
- [SolidFire Node Details](https://grafana.com/grafana/dashboards/14026)
- [SolidFire TopK Volumes](https://grafana.com/grafana/dashboards/14029)
- [SolidFire Overutilized Volumes](https://grafana.com/grafana/dashboards/14027)

Dashboard sources live in the [examples/dashboards](https://github.com/mjavier2k/solidfire-exporter/tree/master/examples/dashboards) folder of this repo.

## Contributing
We welcome contributions. Please fork the project on GitHub and open Pull Requests for any proposed changes.

## License
Code is licensed under the Apache License 2.0.
