//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package csgo

func (*CSGO) SampleConfig() string {
	return `# Fetch metrics from a CSGO SRCDS
[[inputs.csgo]]
  ## Specify servers using the following format:
  ##    servers = [
  ##      ["ip1:port1", "rcon_password1"],
  ##      ["ip2:port2", "rcon_password2"],
  ##    ]
  #
  ## If no servers are specified, no data will be collected
  servers = []
`
}
