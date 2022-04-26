//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package neptune_apex

func (*NeptuneApex) SampleConfig() string {
	return `# Neptune Apex data collector
[[inputs.neptune_apex]]
  ## The Neptune Apex plugin reads the publicly available status.xml data from a local Apex.
  ## Measurements will be logged under "apex".

  ## The base URL of the local Apex(es). If you specify more than one server, they will
  ## be differentiated by the "source" tag.
  servers = [
    "http://apex.local",
  ]

  ## The response_timeout specifies how long to wait for a reply from the Apex.
  #response_timeout = "5s"

`
}
