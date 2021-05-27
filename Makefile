all: test build dashboards

build:
	goreleaser --snapshot --skip-publish --rm-dist

.PHONY: test
test:
	go test ./...

.PHONY: dashboards
dashboards:
	go run ./cmd/dashboards
