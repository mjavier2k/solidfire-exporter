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

| Yaml Configuration Option | CLI Flag | Environment Variable      | Default                         | Example                            | Description                                                                                                                   |
| ------------------------- | -------- | ------------------------- | ------------------------------- | ---------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| client.username           | N/A      | SOLIDFIRE_CLIENT_USERNAME | ""                              | myUsername                         | User with which to authenticate to the Solidfire API. NOTE: User must have administrator access to be able to query QOS data. |
| client.password           | N/A      | SOLIDFIRE_CLIENT_PASSWORD | ""                              | myPassword                         | Password with which to authenticate to the Solidfire API.                                                                     |
| client.endpoint           | N/A      | SOLIDFIRE_CLIENT_ENDPOINT | https://127.0.0.1/json-rpc/11.3 | http://192.168.12.10/json-rpc/11.3 | HTTP(s) endpoint of the Solidfire API.                                                                                        |
| client.insecure           | N/A      | SOLIDFIRE_CLIENT_INSECURE | false                           | true                               | Disables TLS validation when calling Solidfire API. Useful for bypassing self-signed certificates in testing.                 |
| client.timeout            | N/A      | SOLIDFIRE_CLIENT_TIMEOUT  | 30                              | 75                                 | Timeout in seconds per call to the Solidfire API.                                                                             |
| listen.address            | N/A      | SOLIDFIRE_LISTEN_ADDRESS  | 0.0.0.0:9987                    | 192.168.4.2:13987                  | IP address and port where the http server of this exporter should listen                                                      |
| N/A                       | -c       | SOLIDFIRE_CONFIG          | config.yaml                     | mySolidfireConfig.yaml             | Path to configuration file                                                                                                    |

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
