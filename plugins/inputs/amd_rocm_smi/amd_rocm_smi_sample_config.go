//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package amd_rocm_smi

func (rsmi *ROCmSMI) SampleConfig() string {
	return `# Query statistics from AMD Graphics cards using rocm-smi binary
[[inputs.amd_rocm_smi]]
  ## Optional: path to rocm-smi binary, defaults to $PATH via exec.LookPath
  # bin_path = "/opt/rocm/bin/rocm-smi"

  ## Optional: timeout for GPU polling
  # timeout = "5s"
`
}
