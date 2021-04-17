# Official Dashboards

These dashboards are generated using the [grabana](https://github.com/K-Phoen/grabana) library.

- Cluster Overview
- Node Detail
- Volume Detail

## Generating


```bash
# Bootup local grafana (optional)
$ docker run -d -p 3000:3000 -v foo:/var/lib/grafana grafana/grafana:6.5.1

# Generate them
$ make dashboards

# JSON dashboard output:
$ ls -1 dashboards/*.json
dashboards/cluster-overview.json
dashboards/node-detail.json
dashboards/volume-detail.json

# If you booted local Grafana via docker, the dashboards should also be uploaded there as well in the Solidfire folder
```