# Official Dashboards

These dashboards are generated using the [grabana](https://github.com/K-Phoen/grabana) library.

- [SolidFire Cluster Overview](https://grafana.com/grafana/dashboards/14025)
- [SolidFire Volume Detail](https://grafana.com/grafana/dashboards/14030)
- [SolidFire Node Detail](https://grafana.com/grafana/dashboards/14026)

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
# Release Process

To update the dashboard releases on grafana.com, a few extra steps are required to generate the JSON files that are compatible for sharing to others:

- Boot local grafana `docker run -d -p 3000:3000 -v foo:/var/lib/grafana grafana/grafana:6.5.1`
- Setup a Prometheus datasource in local grafana, and make it the default
- `make dashboards` to generate the JSON files and write them to local grafana
- For each dashboard in your grafana, click 'share', select `Export for sharing externally`, then upload the new JSON version on grafana.com
- Update the README for each dash on grafana.com if necessary