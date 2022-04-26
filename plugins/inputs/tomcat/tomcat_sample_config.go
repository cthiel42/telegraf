//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package tomcat

func (s *Tomcat) SampleConfig() string {
	return `# Gather metrics from the Tomcat server status page.
[[inputs.tomcat]]
  ## URL of the Tomcat server status
  # url = "http://127.0.0.1:8080/manager/status/all?XML=true"

  ## HTTP Basic Auth Credentials
  # username = "tomcat"
  # password = "s3cret"

  ## Request timeout
  # timeout = "5s"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false
`
}
