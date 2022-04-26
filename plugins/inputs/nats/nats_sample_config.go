//go:build !freebsd || (freebsd && cgo)
// +build !freebsd freebsd,cgo

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package nats

func (n *Nats) SampleConfig() string {
	return `# Provides metrics about the state of a NATS server
[[inputs.nats]]
  ## The address of the monitoring endpoint of the NATS server
  server = "http://localhost:8222"

  ## Maximum time to receive response
  # response_timeout = "5s"
`
}
