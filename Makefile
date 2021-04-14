.PHONY: dashboards
# now=$()
# go build -ldflags "-X main.sha1ver=`git rev-parse HEAD` -X main.buildTime=${now}"

build:
	go build -ldflags "-X main.sha1ver=`git rev-parse HEAD` -X main.buildTime=`date +'%Y-%m-%d_%T'`" -o ./bin/solidfire-exporter ./cmd/solidfire-exporter

build_static:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.sha1ver=`git rev-parse HEAD` -X main.buildTime=`date +'%Y-%m-%d_%T'`" -o ./bin/solidfire-exporter -a -tags netgo -ldflags '-w' ./cmd/solidfire-exporter

test:
	go test ./...

dash: dashboards
dashboard: dashboards
dashboards:
	go run ./cmd/dashboards