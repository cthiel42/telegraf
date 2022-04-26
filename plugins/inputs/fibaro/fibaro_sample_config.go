//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package fibaro

func (f *Fibaro) SampleConfig() string {
	return `# Read devices value(s) from a Fibaro controller
[[inputs.fibaro]]
  ## Required Fibaro controller address/hostname.
  ## Note: at the time of writing this plugin, Fibaro only implemented http - no https available
  url = "http://<controller>:80"

  ## Required credentials to access the API (http://<controller/api/<component>)
  username = "<username>"
  password = "<password>"

  ## Amount of time allowed to complete the HTTP request
  # timeout = "5s"
`
}
