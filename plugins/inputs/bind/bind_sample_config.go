//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package bind

func (b *Bind) SampleConfig() string {
	return `# Read BIND nameserver XML statistics
[[inputs.bind]]
  ## An array of BIND XML statistics URI to gather stats.
  ## Default is "http://localhost:8053/xml/v3".
  # urls = ["http://localhost:8053/xml/v3"]
  # gather_memory_contexts = false
  # gather_views = false

  ## Timeout for http requests made by bind nameserver
  # timeout = "4s"
`
}
