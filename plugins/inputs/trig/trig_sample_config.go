//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package trig

func (s *Trig) SampleConfig() string {
	return `# Inserts sine and cosine waves for demonstration purposes
[[inputs.trig]]
  ## Set the amplitude
  amplitude = 10.0
`
}
