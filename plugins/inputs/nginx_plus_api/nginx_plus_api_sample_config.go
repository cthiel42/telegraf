//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package nginx_plus_api

func (n *NginxPlusAPI) SampleConfig() string {
	return `# Read Nginx Plus API advanced status information
[[inputs.nginx_plus_api]]
  ## An array of Nginx API URIs to gather stats.
  urls = ["http://localhost/api"]
  # Nginx API version, default: 3
  # api_version = 3

  # HTTP response timeout (default: 5s)
  response_timeout = "5s"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false
`
}
