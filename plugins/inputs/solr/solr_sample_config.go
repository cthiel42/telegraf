//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package solr

func (s *Solr) SampleConfig() string {
	return `# Read stats from one or more Solr servers or cores
[[inputs.solr]]
  ## specify a list of one or more Solr servers
  servers = ["http://localhost:8983"]
  ##
  ## specify a list of one or more Solr cores (default - all)
  # cores = ["main"]
  ##
  ## Optional HTTP Basic Auth Credentials
  # username = "username"
  # password = "pa$$word"
`
}
