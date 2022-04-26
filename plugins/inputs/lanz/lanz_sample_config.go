//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package lanz

func (l *Lanz) SampleConfig() string {
	return `# Read metrics off Arista LANZ, via socket
[[inputs.lanz]]
  ## URL to Arista LANZ endpoint
  servers = [
    "tcp://switch1.int.example.com:50001",
    "tcp://switch2.int.example.com:50001",
  ]
`
}
