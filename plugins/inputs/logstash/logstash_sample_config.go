//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package logstash

func (logstash *Logstash) SampleConfig() string {
	return `# Read metrics exposed by Logstash
[[inputs.logstash]]
  ## The URL of the exposed Logstash API endpoint.
  url = "http://127.0.0.1:9600"

  ## Use Logstash 5 single pipeline API, set to true when monitoring
  ## Logstash 5.
  # single_pipeline = false

  ## Enable optional collection components.  Can contain
  ## "pipelines", "process", and "jvm".
  # collect = ["pipelines", "process", "jvm"]

  ## Timeout for HTTP requests.
  # timeout = "5s"

  ## Optional HTTP Basic Auth credentials.
  # username = "username"
  # password = "pa$$word"

  ## Optional TLS Config.
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"

  ## Use TLS but skip chain & host verification.
  # insecure_skip_verify = false

  ## Optional HTTP headers.
  # [inputs.logstash.headers]
  #   "X-Special-Header" = "Special-Value"
`
}
