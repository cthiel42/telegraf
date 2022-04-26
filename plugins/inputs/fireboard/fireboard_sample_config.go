//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package fireboard

func (r *Fireboard) SampleConfig() string {
	return `# Read real time temps from fireboard.io servers
[[inputs.fireboard]]
  ## Specify auth token for your account
  auth_token = "invalidAuthToken"
  ## You can override the fireboard server URL if necessary
  # url = https://fireboard.io/api/v1/devices.json
  ## You can set a different http_timeout if you need to
  ## You should set a string using an number and time indicator
  ## for example "12s" for 12 seconds.
  # http_timeout = "4s"
`
}
