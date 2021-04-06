module github.com/mjavier2k/solidfire-exporter

go 1.15

replace github.com/K-Phoen/grabana => github.com/neufeldtech/grabana v0.0.0-20210406013955-7013907ae24a

require (
	github.com/K-Phoen/grabana v0.16.0
	github.com/amoghe/distillog v0.0.0-20180726233512-ae382b35b717
	github.com/prometheus/client_golang v1.10.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	gopkg.in/h2non/gock.v1 v1.0.16
)
