//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package udp_listener

func (u *UDPListener) SampleConfig() string {
	return `# Generic UDP listener
[[inputs.udp_listener]]
  # see https://github.com/influxdata/telegraf/tree/master/plugins/inputs/socket_listener
`
}
