//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package bigquery

func (s *BigQuery) SampleConfig() string {
	return `# Configuration for Google Cloud BigQuery to send entries
[[outputs.bigquery]]
  ## Credentials File
  credentials_file = "/path/to/service/account/key.json"

  ## Google Cloud Platform Project
  project = "my-gcp-project"

  ## The namespace for the metric descriptor
  dataset = "telegraf"

  ## Timeout for BigQuery operations.
  # timeout = "5s"

  ## Character to replace hyphens on Metric name
  # replace_hyphen_to = "_"
`
}
