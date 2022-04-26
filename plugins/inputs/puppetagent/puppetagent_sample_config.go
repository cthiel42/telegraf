//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package puppetagent

func (pa *PuppetAgent) SampleConfig() string {
	return `# Reads last_run_summary.yaml file and converts to measurements
[[inputs.puppetagent]]
  ## Location of puppet last run summary file
  location = "/var/lib/puppet/state/last_run_summary.yaml"
`
}
