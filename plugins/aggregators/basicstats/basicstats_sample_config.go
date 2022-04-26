//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package basicstats

func (*BasicStats) SampleConfig() string {
	return `# Keep the aggregate basicstats of each metric passing through.
[[aggregators.basicstats]]
  ## The period on which to flush & clear the aggregator.
  period = "30s"

  ## If true, the original metric will be dropped by the
  ## aggregator and will not get sent to the output plugins.
  drop_original = false

  ## Configures which basic stats to push as fields
  # stats = ["count","diff","rate","min","max","mean","non_negative_diff","non_negative_rate","stdev","s2","sum","interval"]
`
}
