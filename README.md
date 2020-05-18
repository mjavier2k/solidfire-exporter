# NetApp Solidfire Exporter

Prometheus NetApp solidfire-exporter queries the Solidfire API and exports the results as Prometheus metrics

![Go](https://github.com/mjavier2k/solidfire-exporter/workflows/Go/badge.svg?event=push)

Implementation is based on [Solidfire 11.3 API](https://library.netapp.com/ecm/ecm_download_file/ECMLP2856155)

## Goals
- Volume QoS stats
- Volume stats
- Node stats
- Fault stats
- Cluster stats

## Usage

```
Usage of solidfire-exporter:
  -l, --listenPort int    Port for the exporter to listen on. May also be set by environment variable SOLIDFIRE_PORT. (default 9987)
  -u, --username string   User with which to authenticate to the Solidfire API. May also be set by environment variable SOLIDFIRE_USER. (default "my_solidfire_user")
  -p, --password string   Password with which to authenticate to the Solidfire API. May also be set by environment variable SOLIDFIRE_PASS. (default "my_solidfire_password")
  -e, --endpoint string   Endpoint for the Solidfire API. May also be set by environment variable SOLIDFIRE_RPC_ENDPOINT. (default "https://192.168.1.2/json-rpc/11.3")
  -i, --insecure          Whether to disable TLS validation when calling the Solidfire API. May also be set by environment variable INSECURE_SKIP_VERIFY.
```

## Volume QoS

SOLIDFIRE_USER must have administrator access to be able to gather volume qos data


## Using Docker

Create an file with the environment variables set and pass it to docker run. 

```
docker run --env-file=.env_file  --rm -p 8080:8080 mjavier/solidfire-exporter:latest
```
