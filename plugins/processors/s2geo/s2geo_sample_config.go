//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package s2geo

func (g *Geo) SampleConfig() string {
	return `# Add the S2 Cell ID as a tag based on latitude and longitude fields
[[processors.s2geo]]
  ## The name of the lat and lon fields containing WGS-84 latitude and
  ## longitude in decimal degrees.
  # lat_field = "lat"
  # lon_field = "lon"

  ## New tag to create
  # tag_key = "s2_cell_id"

  ## Cell level (see https://s2geometry.io/resources/s2cell_statistics.html)
  # cell_level = 9
`
}
