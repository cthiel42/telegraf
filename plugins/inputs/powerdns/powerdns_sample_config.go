//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package powerdns

func (p *Powerdns) SampleConfig() string {
	return `# Read metrics from one or many PowerDNS servers
[[inputs.powerdns]]
  # An array of sockets to gather stats about.
  # Specify a path to unix socket.
  #
  # If no servers are specified, then '/var/run/pdns.controlsocket' is used as the path.
  unix_sockets = ["/var/run/pdns.controlsocket"]
`
}
