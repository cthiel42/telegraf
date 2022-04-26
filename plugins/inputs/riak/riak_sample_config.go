//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package riak

func (r *Riak) SampleConfig() string {
	return `# Read metrics one or many Riak servers
[[inputs.riak]]
  # Specify a list of one or more riak http servers
  servers = ["http://localhost:8098"]
`
}
