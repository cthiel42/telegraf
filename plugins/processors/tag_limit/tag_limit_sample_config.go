//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package tag_limit

func (d *TagLimit) SampleConfig() string {
	return `# Restricts the number of tags that can pass through this filter and chooses which tags to preserve when over the limit.
[[processors.tag_limit]]
  ## Maximum number of tags to preserve
  limit = 3

  ## List of tags to preferentially preserve
  keep = ["environment", "region"]
`
}
