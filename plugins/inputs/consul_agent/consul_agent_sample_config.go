//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package consul_agent

func (n *ConsulAgent) SampleConfig() string {
	return `# Read metrics from the Consul Agent API
[[inputs.consul_agent]]
  ## URL for the Consul agent
  # url = "http://127.0.0.1:8500"

  ## Use auth token for authorization.
  ## If both are set, an error is thrown.
  ## If both are empty, no token will be used.
  # token_file = "/path/to/auth/token"
  ## OR
  # token = "a1234567-40c7-9048-7bae-378687048181"

  ## Set timeout (default 5 seconds)
  # timeout = "5s"

  ## Optional TLS Config
  # tls_ca = /path/to/cafile
  # tls_cert = /path/to/certfile
  # tls_key = /path/to/keyfile
`
}
