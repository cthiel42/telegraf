//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package couchdb

func (*CouchDB) SampleConfig() string {
	return `# Read CouchDB Stats from one or more servers
[[inputs.couchdb]]
  ## Works with CouchDB stats endpoints out of the box
  ## Multiple Hosts from which to read CouchDB stats:
  hosts = ["http://localhost:8086/_stats"]

  ## Use HTTP Basic Authentication.
  # basic_username = "telegraf"
  # basic_password = "p@ssw0rd"
`
}
