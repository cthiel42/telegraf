//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package tengine

func (n *Tengine) SampleConfig() string {
	return `# Read Tengine's basic status information (ngx_http_reqstat_module)
[[inputs.tengine]]
  ## An array of Tengine reqstat module URI to gather stats.
  urls = ["http://127.0.0.1/us"]

  ## HTTP response timeout (default: 5s)
  # response_timeout = "5s"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false
`
}
