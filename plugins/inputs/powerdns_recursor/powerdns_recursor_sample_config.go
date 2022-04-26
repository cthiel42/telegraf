//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package powerdns_recursor

func (p *PowerdnsRecursor) SampleConfig() string {
	return `# Read metrics from one or many PowerDNS Recursor servers
[[inputs.powerdns_recursor]]
  ## Path to the Recursor control socket.
  unix_sockets = ["/var/run/pdns_recursor.controlsocket"]

  ## Directory to create receive socket.  This default is likely not writable,
  ## please reference the full plugin documentation for a recommended setup.
  # socket_dir = "/var/run/"
  ## Socket permissions for the receive socket.
  # socket_mode = "0666"
`
}
