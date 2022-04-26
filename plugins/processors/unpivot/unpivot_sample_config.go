//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package unpivot

func (p *Unpivot) SampleConfig() string {
	return `# Rotate multi field metric into several single field metrics
[[processors.unpivot]]
  ## Tag to use for the name.
  tag_key = "name"
  ## Field to use for the name of the value.
  value_key = "value"
`
}
