package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/amoghe/distillog"
	"github.com/mjavier2k/solidfire-exporter/pkg/prom"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	sha1ver   string // sha1 revision used to build the program
	buildTime string // when the executable was built
)

func resolvePort() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9987"
	}
	return port
}

func main() {
	log.Infof("Version: %v", sha1ver)
	log.Infof("Built: %v", buildTime)
	listenAddr := fmt.Sprintf("0.0.0.0:%v", resolvePort())
	solidfireExporter, _ := prom.NewCollector()
	prometheus.MustRegister(solidfireExporter)
	http.Handle("/metrics", promhttp.Handler())
	log.Infof("Booted and listening on %v/metrics\n", listenAddr)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "UP")
	})

	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Errorln(err)
	}
}
