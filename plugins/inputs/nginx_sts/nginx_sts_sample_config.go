//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package nginx_sts

func (n *NginxSTS) SampleConfig() string {
	return `# Read Nginx virtual host traffic status module information (nginx-module-sts)
[[inputs.nginx_sts]]
  ## An array of ngx_http_status_module or status URI to gather stats.
  urls = ["http://localhost/status"]

  ## HTTP response timeout (default: 5s)
  response_timeout = "5s"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false
`
}
