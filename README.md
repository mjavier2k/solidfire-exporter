# NetApp Solidfire Exporter

![Go](https://github.com/mjavier2k/solidfire-exporter/workflows/Go/badge.svg?event=push)

NetApp Solidfire-Exporter queries the Solidfire API and exports the results as Prometheus metrics

Implementation is based on [Solidfire 11.3 API](https://library.netapp.com/ecm/ecm_download_file/ECMLP2856155)


![Volume Metrics](examples/solidfire-volume.jpg?raw=true)

### Installation

Binaries can be downloaded from [Github releases](https://github.com/mjavier2k/solidfire-exporter/releases) page. 

### Usage

```
./solidfire_exporter --config config.yaml
```

```
Usage of solidfire-exporter:
  -c, --config string     Specify configuration filename. (default: config.yaml)
```


### Configuration

The following configuration options are supported by YAML or Environment variable.

```
client.username - User with which to authenticate to the Solidfire API. May also be set by environment variable SOLIDFIRE_CLIENT_USERNAME. This account must have administrator access to the solidfire cluster so that QOS data will show up.
client.password - Password with which to authenticate to the Solidfire API. May also be set by environment variable SOLIDFIRE_CLIENT_PASSWORD.
client.endpoint - Endpoint for the Solidfire API. May also be set by environment variable SOLIDFIRE_CLIENT_ENDPOINT.
client.insecure - Whether to disable TLS validation when calling the Solidfire API. May also be set by environment variable SOLIDFIRE_CLIENT_INSECURE.
client.timeout  - HTTP Client timeout (in seconds) per call to Solidfire API. May also be set by environment variable SOLIDFIRE_CLIENT_TIMEOUT.
listen.address  - IP address and port where the Solidfire exporter is listening. May also be set by environment variable SOLIDFIRE_LISTEN_ADDRESS.
```

1) Using config.yaml

```
listen:
  address: 127.0.0.1:9987
client:
  endpoint: https://192.168.1.2/json-rpc/11.3
  username: mySolidfireUsername
  password: mySolidfirePassword
  insecure: false
  timeout: 130
```

2) Using environment variables. These option takes precedence if config.yaml is also specified.

```
export SOLIDFIRE_CLIENT_USERNAME="mySolidfireUsername"
export SOLIDFIRE_CLIENT_PASSWORD="mySolidfirePassword"
export SOLIDFIRE_LISTEN_ADDRESS="127.0.0.1:9987"
export SOLIDFIRE_CLIENT_ENDPOINT="https://10.10.10.10/json-rpc/11.3"
export SOLIDFIRE_CLIENT_INSECURE=true
export SOLIDFIRE_CLIENT_TIMEOUT=30
```

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
