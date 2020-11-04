package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	log "github.com/amoghe/distillog"
	"github.com/mjavier2k/solidfire-exporter/pkg/prom"
	"github.com/mjavier2k/solidfire-exporter/pkg/solidfire"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	sha1ver   string // sha1 revision used to build the program
	buildTime string // when the executable was built
)

func init() {
	flag.CommandLine.SortFlags = true
	flag.StringP(solidfire.ConfigFile, "c", solidfire.DefaultConfigFile, fmt.Sprintf("Specify configuration filename."))
	flag.Parse()
	viper.BindPFlags(flag.CommandLine)

	// extracts the filename from the config filename passed on --config flag (e.g /etc/solidfire-exporter/config.yaml)
	viper.SetConfigName(filepath.Base(viper.GetString(solidfire.ConfigFile)))
	viper.SetConfigType("yaml")
	// extracts the directory path from the config filename passed on --config flag (e.g /etc/solidfire-exporter/config.yaml)
	viper.AddConfigPath(filepath.Dir(viper.GetString(solidfire.ConfigFile)))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Infof("No config file found.")
		}
	} else {
		log.Infof("Found configuration file on %v ", viper.GetViper().ConfigFileUsed())
	}

	viper.SetDefault(solidfire.ListenAddress, solidfire.DefaultListenAddress)
	viper.SetDefault(solidfire.Endpoint, solidfire.DefaultEndpoint)
	viper.SetDefault(solidfire.HTTPClientTimeout, solidfire.DefaultHTTPClientTimeout)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("SOLIDFIRE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
func main() {
	log.Infof("Version: %v", sha1ver)
	log.Infof("Built: %v", buildTime)
	listenAddress := fmt.Sprintf("%v", viper.GetString(solidfire.ListenAddress))
	solidfireExporter, _ := prom.NewCollector()
	prometheus.MustRegister(solidfireExporter)
	http.Handle("/metrics", promhttp.Handler())
	for _, key := range viper.AllKeys() {
		value := viper.Get(key)
		if key == solidfire.Password {
			value = "[REDACTED]"
		}
		log.Infof("Config setting found for %s: %v", key, value)
	}
	log.Infof("Booted and listening on %v/metrics\n", listenAddress)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "UP")
	})

	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Errorln(err)
	}
}
