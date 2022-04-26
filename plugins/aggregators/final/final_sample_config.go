//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package final

func (m *Final) SampleConfig() string {
	return `# Report the final metric of a series
[[aggregators.final]]
  ## The period on which to flush & clear the aggregator.
  period = "30s"
  ## If true, the original metric will be dropped by the
  ## aggregator and will not get sent to the output plugins.
  drop_original = false

  ## The time that a series is not updated until considering it final.
  series_timeout = "5m"
`
}
