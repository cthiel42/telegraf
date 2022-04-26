//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package salesforce

func (s *Salesforce) SampleConfig() string {
	return `# Read API usage and limits for a Salesforce organisation
[[inputs.salesforce]]
  ## specify your credentials
  ##
  username = "your_username"
  password = "your_password"
  ##
  ## (optional) security token
  # security_token = "your_security_token"
  ##
  ## (optional) environment type (sandbox or production)
  ## default is: production
  ##
  # environment = "production"
  ##
  ## (optional) API version (default: "39.0")
  ##
  # version = "39.0"
`
}
