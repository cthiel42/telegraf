//go:build !windows
// +build !windows

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package bcache

func (b *Bcache) SampleConfig() string {
	return `# Read metrics of bcache from stats_total and dirty_data
[[inputs.bcache]]
  ## Bcache sets path
  ## If not specified, then default is:
  bcachePath = "/sys/fs/bcache"

  ## By default, Telegraf gather stats for all bcache devices
  ## Setting devices will restrict the stats to the specified
  ## bcache devices.
  bcacheDevs = ["bcache0"]
`
}
