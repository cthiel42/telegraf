//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package suricata

func (s *Suricata) SampleConfig() string {
	return `# Suricata stats and alerts plugin
[[inputs.suricata]]
  ## Data sink for Suricata stats log.
  # This is expected to be a filename of a
  # unix socket to be created for listening.
  source = "/var/run/suricata-stats.sock"

  # Delimiter for flattening field keys, e.g. subitem "alert" of "detect"
  # becomes "detect_alert" when delimiter is "_".
  delimiter = "_"

  # Detect alert logs
  alerts = false
`
}
