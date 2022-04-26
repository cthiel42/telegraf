//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package execd

func (e *Execd) SampleConfig() string {
	return `# Run executable as long-running processor plugin
[[processors.execd]]
  ## One program to run as daemon.
  ## NOTE: process and each argument should each be their own string
  ## eg: command = ["/path/to/your_program", "arg1", "arg2"]
  command = ["cat"]

  ## Delay before the process is restarted after an unexpected termination
  # restart_delay = "10s"
`
}
