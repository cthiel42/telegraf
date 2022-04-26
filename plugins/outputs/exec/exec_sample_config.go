//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package exec

func (e *Exec) SampleConfig() string {
	return `# Send metrics to command as input over stdin
[[outputs.exec]]
  ## Command to ingest metrics via stdin.
  command = ["tee", "-a", "/dev/null"]

  ## Timeout for command to complete.
  # timeout = "5s"

  ## Data format to output.
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_OUTPUT.md
  # data_format = "influx"
`
}
