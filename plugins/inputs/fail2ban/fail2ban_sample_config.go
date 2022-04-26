//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package fail2ban

func (f *Fail2ban) SampleConfig() string {
	return `# Read metrics from fail2ban.
[[inputs.fail2ban]]
  ## Use sudo to run fail2ban-client
  use_sudo = false
`
}
