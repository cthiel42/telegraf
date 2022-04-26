//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package lvm

func (lvm *LVM) SampleConfig() string {
	return `# Read metrics about LVM physical volumes, volume groups, logical volumes.
[[inputs.lvm]]
  ## Use sudo to run LVM commands
  use_sudo = false
`
}
