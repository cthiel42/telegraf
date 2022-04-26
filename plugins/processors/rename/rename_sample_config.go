//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package rename

func (r *Rename) SampleConfig() string {
	return `# Rename measurements, tags, and fields that pass through this filter.
[[processors.rename]]
  ## Specify one sub-table per rename operation.
  [[processors.rename.replace]]
    measurement = "network_interface_throughput"
    dest = "throughput"

  [[processors.rename.replace]]
    tag = "hostname"
    dest = "host"

  [[processors.rename.replace]]
    field = "lower"
    dest = "min"

  [[processors.rename.replace]]
    field = "upper"
    dest = "max"
`
}
