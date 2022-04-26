//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package dmcache

func (c *DMCache) SampleConfig() string {
	return `# Provide a native collection for dmsetup based statistics for dm-cache
[[inputs.dmcache]]
  ## Whether to report per-device stats or not
  per_device = true
`
}
