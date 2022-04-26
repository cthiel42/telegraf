//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package multifile

func (m *MultiFile) SampleConfig() string {
	return `# Aggregates the contents of multiple files into a single point
[[inputs.multifile]]
  ## Base directory where telegraf will look for files.
  ## Omit this option to use absolute paths.
  base_dir = "/sys/bus/i2c/devices/1-0076/iio:device0"

  ## If true discard all data when a single file can't be read.
  ## Else, Telegraf omits the field generated from this file.
  # fail_early = true

  ## Files to parse each interval.
  [[inputs.multifile.file]]
    file = "in_pressure_input"
    dest = "pressure"
    conversion = "float"
  [[inputs.multifile.file]]
    file = "in_temp_input"
    dest = "temperature"
    conversion = "float(3)"
  [[inputs.multifile.file]]
    file = "in_humidityrelative_input"
    dest = "humidityrelative"
    conversion = "float(3)"
`
}
