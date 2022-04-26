//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package fluentd

func (h *Fluentd) SampleConfig() string {
	return `# Read metrics exposed by fluentd in_monitor plugin
[[inputs.fluentd]]
  ## This plugin reads information exposed by fluentd (using /api/plugins.json endpoint).
  ##
  ## Endpoint:
  ## - only one URI is allowed
  ## - https is not supported
  endpoint = "http://localhost:24220/api/plugins.json"

  ## Define which plugins have to be excluded (based on "type" field - e.g. monitor_agent)
  exclude = [
    "monitor_agent",
    "dummy",
  ]
`
}
