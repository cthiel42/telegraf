//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package zipkin

func (z Zipkin) SampleConfig() string {
	return `# This plugin implements the Zipkin http server to gather trace and timing data needed to troubleshoot latency problems in microservice architectures.
[[inputs.zipkin]]
  # path = "/api/v1/spans" # URL path for span data
  # port = 9411 # Port on which Telegraf listens
`
}
