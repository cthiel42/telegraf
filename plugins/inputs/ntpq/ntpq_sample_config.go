//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package ntpq

func (n *NTPQ) SampleConfig() string {
	return `# Get standard NTP query metrics, requires ntpq executable.
[[inputs.ntpq]]
  ## If false, set the -n ntpq flag. Can reduce metric gather time.
  dns_lookup = true
`
}
