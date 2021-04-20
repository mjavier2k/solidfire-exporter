package testutils

import (
	"path"

	"github.com/mjavier2k/solidfire-exporter/pkg/solidfire"
)

func ResolveFixturePath(basePath string, r solidfire.RPC) string {
	return path.Join(basePath, string(r)+".json")
}
