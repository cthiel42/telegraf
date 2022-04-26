//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package systemd_units

func (s *SystemdUnits) SampleConfig() string {
	return `# Gather systemd units state
[[inputs.systemd_units]]
  ## Set timeout for systemctl execution
  # timeout = "1s"
  #
  ## Filter for a specific unit type, default is "service", other possible
  ## values are "socket", "target", "device", "mount", "automount", "swap",
  ## "timer", "path", "slice" and "scope ":
  # unittype = "service"
  #
  ## Filter for a specific pattern, default is "" (i.e. all), other possible
  ## values are valid pattern for systemctl, e.g. "a*" for all units with
  ## names starting with "a"
  # pattern = ""
  ## pattern = "telegraf* influxdb*"
  ## pattern = "a*"
`
}
