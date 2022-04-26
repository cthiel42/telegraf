//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package openntpd

func (n *Openntpd) SampleConfig() string {
	return `# Get standard NTP query metrics from OpenNTPD.
[[inputs.openntpd]]
  ## Run ntpctl binary with sudo.
  # use_sudo = false

  ## Location of the ntpctl binary.
  # binary = "/usr/sbin/ntpctl"

  ## Maximum time the ntpctl binary is allowed to run.
  # timeout = "5ms"
`
}
