//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package synproxy

func (k *Synproxy) SampleConfig() string {
	return `# Get synproxy counter statistics from procfs
[[inputs.synproxy]]
  # no configuration
`
}
