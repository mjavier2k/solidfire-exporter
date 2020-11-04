# NetApp Solidfire Exporter

![Go](https://github.com/mjavier2k/solidfire-exporter/workflows/Go/badge.svg?event=push)

NetApp Solidfire-Exporter queries the Solidfire API and exports the results as Prometheus metrics

Implementation is based on [Solidfire 11.3 API](https://library.netapp.com/ecm/ecm_download_file/ECMLP2856155)


![Volume Metrics](examples/solidfire-volume.jpg?raw=true)

### Installation

Binaries can be downloaded from [Github releases](https://github.com/mjavier2k/solidfire-exporter/releases) page. 

### Usage

```
./solidfire_exporter
```

```
Usage of solidfire-exporter:
  -c, --config string     Specify configuration filename. (default: config.yaml)
```

There are 2 ways to specify the configuration for solifire_exporter.

1) Using config.yaml

```
endpoint: https://10.10.10.10/json-rpc/11.3
listenPort: 9987
insecure: true
timeout: 60
username: mySolidfireUsername
password: mySolidfirePassword
```

2) Specify Environment variables. Values specified here takes precedences to config.yaml

```
export SOLIDFIRE_USER="mySolidfireUsername"
export SOLIDFIRE_USER="mySolidfirePassword"
unset SOLIDFIRE_PORT=9987
unset SOLIDFIRE_RPC_ENDPOINT="https://10.10.10.10/json-rpc/11.3"
unset INSECURE_SKIP_VERIFY=true
unset HTTP_CLIENT_TIMEOUT=30
```

__NOTE__: The account for __SOLIDFIRE_USER__ must have administrator access to the solidfire cluster so that QOS data will show up.

### Prometheus Configuration

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

### Using Docker

Create an file with the environment variables set and pass it to docker run. 

```
docker run --env-file=.env_file  --rm -p 8080:8080 mjavier/solidfire-exporter:latest
```

### Grafana Dashboards

Sample Grafana dasboards available on the [example](https://github.com/mjavier2k/solidfire-exporter/tree/master/examples) folder of this repo.

### Alerts

TO DO


### Contributing
We welcome any contributions. Please fork the project on GitHub and open Pull Requests for any proposed changes.

### License
Code is licensed under the Apache License 2.0.
