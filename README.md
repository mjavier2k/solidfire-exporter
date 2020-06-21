# NetApp Solidfire Exporter

![Go](https://github.com/mjavier2k/solidfire-exporter/workflows/Go/badge.svg?event=push)

Prometheus NetApp Solidfire-Exporter queries the Solidfire API and exports the results as Prometheus metrics

Implementation is based on [Solidfire 11.3 API](https://library.netapp.com/ecm/ecm_download_file/ECMLP2856155)


![Volume Metrics](examples/solidfire-volume.jpg?raw=true)

### Installation

Binaries can be downloaded from [Github releases](https://github.com/mjavier2k/solidfire-exporter/releases) page. 

### Usage

```
./solidfire_exporter -u $SOLIDFIRE_USER -p $SOLIDFIRE_PASSWORD -e $SOLIDFIRE_ENDPOINT
```

```
Usage of solidfire-exporter:
  -l, --listenPort int    Port for the exporter to listen on. May also be set by environment variable SOLIDFIRE_PORT. (default 9987)
  -u, --username string   User with which to authenticate to the Solidfire API. May also be set by environment variable SOLIDFIRE_USER. (default "my_solidfire_user")
  -p, --password string   Password with which to authenticate to the Solidfire API. May also be set by environment variable SOLIDFIRE_PASS. (default "my_solidfire_password")
  -e, --endpoint string   Endpoint for the Solidfire API. May also be set by environment variable SOLIDFIRE_RPC_ENDPOINT. (default "https://192.168.1.2/json-rpc/11.3")
  -i, --insecure          Whether to disable TLS validation when calling the Solidfire API. May also be set by environment variable INSECURE_SKIP_VERIFY.
  -t, --timeout int       HTTP Client timeout (in seconds) per call to Solidfire API. (default 30)
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

Please note that we will not merge any changes that encourage insecure behaviour. If in doubt please open an Issue first to discuss your proposal.

### License
Code is licensed under the Apache License 2.0.