//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package infiniband

func (i *Infiniband) SampleConfig() string {
	return `# Gets counters from all InfiniBand cards and ports installed
[[inputs.infiniband]]
  # no configuration
`
}
