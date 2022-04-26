//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package stackdriver

func (s *Stackdriver) SampleConfig() string {
	return `# Configuration for Google Cloud Stackdriver to send metrics to
[[outputs.stackdriver]]
  ## GCP Project
  project = "erudite-bloom-151019"

  ## The namespace for the metric descriptor
  namespace = "telegraf"

  ## Custom resource type
  # resource_type = "generic_node"

  ## Additional resource labels
  # [outputs.stackdriver.resource_labels]
  #   node_id = "$HOSTNAME"
  #   namespace = "myapp"
  #   location = "eu-north0"
`
}
