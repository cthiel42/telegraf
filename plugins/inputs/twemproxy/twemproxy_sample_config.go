//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package twemproxy

func (t *Twemproxy) SampleConfig() string {
	return `# Read Twemproxy stats data
[[inputs.twemproxy]]
  ## Twemproxy stats address and port (no scheme)
  addr = "localhost:22222"
  ## Monitor pool name
  pools = ["redis_pool", "mc_pool"]
`
}
