//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package leofs

func (l *LeoFS) SampleConfig() string {
	return `# Read metrics from a LeoFS Server via SNMP
[[inputs.leofs]]
  ## An array of URLs of the form:
  ##   host [ ":" port]
  servers = ["127.0.0.1:4010"]
`
}
