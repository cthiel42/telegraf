//go:build !windows
// +build !windows

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package postfix

func (p *Postfix) SampleConfig() string {
	return `# Measure postfix queue statistics
[[inputs.postfix]]
  ## Postfix queue directory. If not provided, telegraf will try to use
  ## 'postconf -h queue_directory' to determine it.
  # queue_directory = "/var/spool/postfix"
`
}
