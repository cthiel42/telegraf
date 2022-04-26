//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package instrumental

func (i *Instrumental) SampleConfig() string {
	return `# Configuration for sending metrics to an Instrumental project
[[outputs.instrumental]]
  ## Project API Token (required)
  api_token = "API Token"  # required
  ## Prefix the metrics with a given name
  prefix = ""
  ## Stats output template (Graphite formatting)
  ## see https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_OUTPUT.md#graphite
  template = "host.tags.measurement.field"
  ## Timeout in seconds to connect
  timeout = "2s"
  ## Debug true - Print communication to Instrumental
  debug = false
`
}
