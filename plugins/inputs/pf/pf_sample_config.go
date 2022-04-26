//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package pf

func (pf *PF) SampleConfig() string {
	return `# Gather counters from PF
[[inputs.pf]]
  ## PF require root access on most systems.
  ## Setting 'use_sudo' to true will make use of sudo to run pfctl.
  ## Users must configure sudo to allow telegraf user to run pfctl with no password.
  ## pfctl can be restricted to only list command "pfctl -s info".
  use_sudo = false
`
}
