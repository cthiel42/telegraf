//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package teamspeak

func (ts *Teamspeak) SampleConfig() string {
	return `# Reads metrics from a Teamspeak 3 Server via ServerQuery
[[inputs.teamspeak]]
  ## Server address for Teamspeak 3 ServerQuery
  # server = "127.0.0.1:10011"
  ## Username for ServerQuery
  username = "serverqueryuser"
  ## Password for ServerQuery
  password = "secret"
  ## Nickname of the ServerQuery client
  nickname = "telegraf"
  ## Array of virtual servers
  # virtual_servers = [1]
`
}
