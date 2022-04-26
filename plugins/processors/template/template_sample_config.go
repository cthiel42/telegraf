//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package template

func (r *TemplateProcessor) SampleConfig() string {
	return `# Uses a Go template to create a new tag
[[processors.template]]
  ## Tag to set with the output of the template.
  tag = "topic"

  ## Go template used to create the tag value.  In order to ease TOML
  ## escaping requirements, you may wish to use single quotes around the
  ## template string.
  template = '