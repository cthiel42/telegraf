//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package nvidia_smi

func (smi *NvidiaSMI) SampleConfig() string {
	return `# Pulls statistics from nvidia GPUs attached to the host
[[inputs.nvidia_smi]]
  ## Optional: path to nvidia-smi binary, defaults "/usr/bin/nvidia-smi"
  ## We will first try to locate the nvidia-smi binary with the explicitly specified value (or default value), 
  ## if it is not found, we will try to locate it on PATH(exec.LookPath), if it is still not found, an error will be returned
  # bin_path = "/usr/bin/nvidia-smi"

  ## Optional: timeout for GPU polling
  # timeout = "5s"
`
}
