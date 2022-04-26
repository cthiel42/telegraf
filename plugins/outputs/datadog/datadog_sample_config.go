//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package datadog

func (d *Datadog) SampleConfig() string {
	return `# Configuration for DataDog API to send metrics to.
[[outputs.datadog]]
  ## Datadog API key
  apikey = "my-secret-key"

  ## Connection timeout.
  # timeout = "5s"

  ## Write URL override; useful for debugging.
  # url = "https://app.datadoghq.com/api/v1/series"

  ## Set http_proxy (telegraf uses the system wide proxy settings if it isn't set)
  # http_proxy_url = "http://localhost:8888"

  ## Override the default (none) compression used to send data.
  ## Supports: "zlib", "none"
  # compression = "none"
`
}
