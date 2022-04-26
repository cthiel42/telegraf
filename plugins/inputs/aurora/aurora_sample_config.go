//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package aurora

func (a *Aurora) SampleConfig() string {
	return `# Gather metrics from Apache Aurora schedulers
[[inputs.aurora]]
  ## Schedulers are the base addresses of your Aurora Schedulers
  schedulers = ["http://127.0.0.1:8081"]

  ## Set of role types to collect metrics from.
  ##
  ## The scheduler roles are checked each interval by contacting the
  ## scheduler nodes; zookeeper is not contacted.
  # roles = ["leader", "follower"]

  ## Timeout is the max time for total network operations.
  # timeout = "5s"

  ## Username and password are sent using HTTP Basic Auth.
  # username = "username"
  # password = "pa$$word"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false
`
}
