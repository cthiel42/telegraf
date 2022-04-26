//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package dovecot

func (d *Dovecot) SampleConfig() string {
	return `# Read metrics about dovecot servers
[[inputs.dovecot]]
  ## specify dovecot servers via an address:port list
  ##  e.g.
  ##    localhost:24242
  ## or as an UDS socket
  ##  e.g.
  ##    /var/run/dovecot/old-stats
  ##
  ## If no servers are specified, then localhost is used as the host.
  servers = ["localhost:24242"]

  ## Type is one of "user", "domain", "ip", or "global"
  type = "global"

  ## Wildcard matches like "*.com". An empty string "" is same as "*"
  ## If type = "ip" filters should be <IP/network>
  filters = [""]
`
}
