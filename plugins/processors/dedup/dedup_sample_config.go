//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package dedup

func (d *Dedup) SampleConfig() string {
	return `# Filter metrics with repeating field values
[[processors.dedup]]
  ## Maximum time to suppress output
  dedup_interval = "600s"
`
}
