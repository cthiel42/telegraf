//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package chrony

func (*Chrony) SampleConfig() string {
	return `# Get standard chrony metrics, requires chronyc executable.
[[inputs.chrony]]
  ## If true, chronyc tries to perform a DNS lookup for the time server.
  # dns_lookup = false
`
}
