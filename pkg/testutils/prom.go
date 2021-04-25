package testutils

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
)

//prometheusOutput was originally copied from https://github.com/peterbourgon/fastly-exporter - thanks Peter
func PrometheusOutput(t *testing.T, registry *prometheus.Registry, prefix string) []string {
	t.Helper()

	server := httptest.NewServer(promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	defer server.Close()
	resp, err := http.Get(server.URL)
	assert.NoError(t, err)

	var selected []string
	s := bufio.NewScanner(resp.Body)
	for s.Scan() {
		if strings.HasPrefix(s.Text(), prefix) {
			// fmt.Println(s.Text())
			selected = append(selected, s.Text())
		}
	}
	return selected
}
