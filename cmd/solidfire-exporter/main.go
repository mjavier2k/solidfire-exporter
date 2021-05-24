package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/apex/log"
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
	logger    *log.Entry
)

func init() {
	logger = log.WithField("app", "solidfire-exporter")

	flag.CommandLine.SortFlags = true
	flag.StringP(solidfire.ConfigFile, "c", solidfire.DefaultConfigFile, fmt.Sprintf("Specify configuration filename."))
	flag.String(solidfire.LogLevel, solidfire.DefaultLogLevel, fmt.Sprintf("Log level."))
	flag.String(solidfire.ListenAddress, solidfire.DefaultListenAddress, fmt.Sprintf("Listen address."))
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
			logger.Warn("no config file found")
		}
	} else {
		logger.Infof("found configuration file on %v ", viper.GetViper().ConfigFileUsed())
	}

	// Set sensible defaults
	viper.SetDefault(solidfire.ListenAddress, solidfire.DefaultListenAddress)
	viper.SetDefault(solidfire.Endpoint, solidfire.DefaultEndpoint)
	viper.SetDefault(solidfire.HTTPClientTimeout, solidfire.DefaultHTTPClientTimeout)
	viper.SetDefault(solidfire.Username, solidfire.DefaultUsername)
	viper.SetDefault(solidfire.Password, solidfire.DefaultPassword)
	viper.SetDefault(solidfire.ConfigFile, solidfire.DefaultConfigFile)
	viper.SetDefault(solidfire.LogLevel, solidfire.DefaultLogLevel)

	// Bind the viper flags to ENV variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("SOLIDFIRE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.BindEnv(solidfire.Username)
	viper.BindEnv(solidfire.Password)
	viper.BindEnv(solidfire.Endpoint)
	viper.BindEnv(solidfire.InsecureSSL)
	viper.BindEnv(solidfire.HTTPClientTimeout)
	viper.BindEnv(solidfire.LogLevel)
}
func main() {
	log.SetLevel(log.MustParseLevel(viper.GetString(solidfire.LogLevel)))
	logger.WithField("version", sha1ver).Infof("version: %v", sha1ver)
	logger.WithField("buildDate", buildTime).Infof("built: %v", buildTime)
	listenAddress := fmt.Sprintf("%v", viper.GetString(solidfire.ListenAddress))

	sfClient, err := solidfire.NewSolidfireClient(logger)
	if err != nil {
		logger.Errorf("error initializing solidfire client: %s\n", err.Error())
		os.Exit(1)
	}

	solidfireExporter, err := prom.NewCollector(sfClient, logger)
	if err != nil {
		logger.Errorf("error initializing collector: %s\n", err.Error())
		os.Exit(1)
	}
	prometheus.MustRegister(solidfireExporter)
	http.Handle("/metrics", promhttp.Handler())

	for _, key := range viper.AllKeys() {
		value := viper.Get(key)
		if key == solidfire.Password {
			value = "[REDACTED]"
		}
		logger.WithField(key, value).Debugf("booted with setting %v", key)
	}
	logger.Infof("Booted and listening on %v/metrics\n", listenAddress)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Solidfire Exporter</title></head>
			<body>
			<h1>Solidfire Exporter</h1>
			<p><a href="/metrics">Metrics</a></p>
			</body>
			</html>`))
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "UP")
	})

	err = http.ListenAndServe(listenAddress, nil)
	if err != nil {
		logger.WithError(err).Errorf("error listening on address %v", listenAddress)
	}
}
