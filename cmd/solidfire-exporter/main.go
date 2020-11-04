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
	flag.StringP(solidfire.ConfigFile, "c", "config.yaml", fmt.Sprintf("Specify configuration filename."))
	flag.Parse()
	viper.BindPFlags(flag.CommandLine)

	// set configname based on --config flag e.g /etc/solidfire-exporter/config.yaml
	viper.SetConfigName(filepath.Base(viper.GetString(solidfire.ConfigFile)))
	viper.SetConfigType("yaml")
	// this adds the dir path based on --config flag if it resides in a different directory
	viper.AddConfigPath(filepath.Dir(viper.GetString(solidfire.ConfigFile)))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Infof("No config file found.")
		}
	} else {
		log.Infof("Found configuration file on %v ", viper.GetViper().ConfigFileUsed())
	}

	// viper.BindEnv(solidfire.UsernameFlag, solidfire.UsernameFlagEnv)
	// viper.BindEnv(solidfire.PasswordFlag, solidfire.PasswordFlagEnv)
	// viper.BindEnv(solidfire.EndpointFlag, solidfire.EndpointFlagEnv)
	// viper.BindEnv(solidfire.InsecureSSLFlag, solidfire.InsecureSSLFlagEnv)
	// viper.BindEnv(solidfire.HTTPClientTimeoutFlag, solidfire.HTTPClientTimeoutFlagEnv)
	// viper.SetEnvPrefix("SOLIDFIRE")
	viper.SetEnvKeyReplacer(strings.NewReplacer("solidfire", ""))
	viper.AutomaticEnv()
	fmt.Println(viper.GetViper().AllKeys())

}
func main() {
	log.Infof("Version: %v", sha1ver)
	log.Infof("Built: %v", buildTime)
	listenAddr := fmt.Sprintf("0.0.0.0:%v", viper.GetInt(solidfire.ListenPort))
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
