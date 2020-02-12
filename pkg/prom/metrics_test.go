package prom_test

import (
	"testing"

	"github.com/mjavier2k/solidfire-exporter/pkg/prom"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	namespace = "namespace"
	registry  = prometheus.NewRegistry()
)

func TestRegistration(t *testing.T) {
	t.Skip("TODO: make sure the metric descriptions are legit when registering then with a registry...")
	{
		_ = prom.NewMetricDescriptions(namespace)
	}
}
