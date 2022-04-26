//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package raindrops

func (r *Raindrops) SampleConfig() string {
	return `# Read raindrops stats (raindrops - real-time stats for preforking Rack servers)
[[inputs.raindrops]]
  ## An array of raindrops middleware URI to gather stats.
  urls = ["http://localhost:8080/_raindrops"]
`
}
