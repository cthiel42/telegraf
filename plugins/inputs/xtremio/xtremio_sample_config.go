//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package xtremio

func (xio *XtremIO) SampleConfig() string {
	return ` # Gathers Metrics From a Dell EMC XtremIO Storage Array's V3 API
[[inputs.xtremio]]
  ## XtremIO User Interface Endpoint
  url = "https://xtremio.example.com/" # required

  ## Credentials
  username = "user1"
  password = "pass123"

  ## Metrics to collect from the XtremIO
  # collectors = ["bbus","clusters","ssds","volumes","xms"]

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use SSL but skip chain & host verification
  # insecure_skip_verify = false
`
}
