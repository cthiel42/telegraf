//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package opensmtpd

func (s *Opensmtpd) SampleConfig() string {
	return `# A plugin to collect stats from Opensmtpd - a validating, recursive, and caching DNS resolver
 [[inputs.opensmtpd]]
   ## If running as a restricted user you can prepend sudo for additional access:
   #use_sudo = false

   ## The default location of the smtpctl binary can be overridden with:
   binary = "/usr/sbin/smtpctl"

   # The default timeout of 1s can be overridden with:
   #timeout = "1s"
`
}
