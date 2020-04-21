package main

import (
	"fmt"
	"net/http"

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
	flag.CommandLine.SortFlags = false
	flag.StringP(solidfire.ListenFlag, "l", "0.0.0.0:9987", fmt.Sprintf("Address + port for the exporter to listen on. May also be set by environment variable %v.", solidfire.ListenFlagEnv))
	flag.StringP(solidfire.UsernameFlag, "u", "my_solidfire_user", fmt.Sprintf("User with which to authenticate to the Solidfire API. May also be set by environment variable %v.", solidfire.UsernameFlagEnv))
	flag.StringP(solidfire.PasswordFlag, "p", "my_solidfire_password", fmt.Sprintf("Password with which to authenticate to the Solidfire API. May also be set by environment variable %v.", solidfire.PasswordFlagEnv))
	flag.StringP(solidfire.EndpointFlag, "e", "https://192.168.1.2/json-rpc/11.3", fmt.Sprintf("Endpoint for the Solidfire API. May also be set by environment variable %v.", solidfire.EndpointFlagEnv))
	flag.BoolP(solidfire.InsecureSSLFlag, "i", false, fmt.Sprintf("Whether to disable TLS validation when calling the Solidfire API. May also be set by environment variable %v.", solidfire.InsecureSSLFlagEnv))
	flag.Parse()

	viper.BindEnv(solidfire.ListenFlag, solidfire.ListenFlagEnv)
	viper.BindEnv(solidfire.UsernameFlag, solidfire.UsernameFlagEnv)
	viper.BindEnv(solidfire.PasswordFlag, solidfire.PasswordFlagEnv)
	viper.BindEnv(solidfire.EndpointFlag, solidfire.EndpointFlagEnv)
	viper.BindEnv(solidfire.InsecureSSLFlag, solidfire.InsecureSSLFlagEnv)
	viper.BindPFlags(flag.CommandLine)
}
func main() {
	log.Infof("Version: %v", sha1ver)
	log.Infof("Built: %v", buildTime)
	listenAddr := viper.GetString(solidfire.ListenFlag)
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
