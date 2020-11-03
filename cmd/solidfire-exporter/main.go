package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

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
	flag.IntP(solidfire.ListenPortFlag, "l", 9987, fmt.Sprintf("Port for the exporter to listen on. May also be set by environment variable %v.", solidfire.ListenPortFlagEnv))
	flag.StringP(solidfire.EndpointFlag, "e", "https://192.168.1.2/json-rpc/11.3", fmt.Sprintf("Endpoint for the Solidfire API. May also be set by environment variable %v.", solidfire.EndpointFlagEnv))
	flag.BoolP(solidfire.InsecureSSLFlag, "i", false, fmt.Sprintf("Whether to disable TLS validation when calling the Solidfire API. May also be set by environment variable %v.", solidfire.InsecureSSLFlagEnv))
	flag.StringP(solidfire.ConfigFileFlag, "c", "config.yaml", fmt.Sprintf("Specify configuration filename."))
	flag.Int64P(solidfire.HTTPClientTimeoutFlag, "t", 30, fmt.Sprintf("HTTP Client timeout (in seconds) per call to Solidfire API."))
	flag.Parse()

	viper.BindPFlags(flag.CommandLine)

	// PORT environment variable takes precedence in order to be backwards-compatible
	if legacyPort, legacyPortFlagExists := os.LookupEnv("PORT"); legacyPortFlagExists {
		viper.BindEnv(solidfire.ListenPortFlag, "PORT")
		log.Warningf("Found environment variable PORT=%v, skipping lookup of %v", legacyPort, solidfire.ListenPortFlagEnv)
	} else {
		viper.BindEnv(solidfire.ListenPortFlag, solidfire.ListenPortFlagEnv)
	}

	// set configname based on --config flag e.g /etc/solidfire-exporter/config.yaml
	viper.SetConfigName(filepath.Base(viper.GetString(solidfire.ConfigFileFlag)))
	viper.SetConfigType("yaml")
	// this adds the dir path based on --config flag if it resides in a different directory
	viper.AddConfigPath(filepath.Dir(viper.GetString(solidfire.ConfigFileFlag)))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Infof("No config file found.")
		}
	} else {
		log.Infof("Found configuration file on %v ", viper.GetViper().ConfigFileUsed())
	}

	viper.BindEnv(solidfire.UsernameFlag, solidfire.UsernameFlagEnv)
	viper.BindEnv(solidfire.PasswordFlag, solidfire.PasswordFlagEnv)
	viper.BindEnv(solidfire.EndpointFlag, solidfire.EndpointFlagEnv)
	viper.BindEnv(solidfire.InsecureSSLFlag, solidfire.InsecureSSLFlagEnv)

}
func main() {
	log.Infof("Version: %v", sha1ver)
	log.Infof("Built: %v", buildTime)
	listenAddr := fmt.Sprintf("127.0.0.1:%v", viper.GetInt(solidfire.ListenPortFlag))
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
