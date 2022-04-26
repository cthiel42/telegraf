//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package internet_speed

func (is *InternetSpeed) SampleConfig() string {
	return `# Monitors internet speed using speedtest.net service
[[inputs.internet_speed]]
  ## Sets if runs file download test
  # enable_file_download = false

  ## Caches the closest server location
  # cache = false
`
}
