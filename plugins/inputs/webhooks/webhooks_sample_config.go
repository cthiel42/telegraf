//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package webhooks

func (wb *Webhooks) SampleConfig() string {
	return `# A Webhooks Event collector
[[inputs.webhooks]]
  ## Address and port to host Webhook listener on
  service_address = ":1619"

  [inputs.webhooks.filestack]
    path = "/filestack"

    ## HTTP basic auth
    #username = ""
    #password = ""

  [inputs.webhooks.github]
    path = "/github"
    # secret = ""

    ## HTTP basic auth
    #username = ""
    #password = ""

  [inputs.webhooks.mandrill]
    path = "/mandrill"

    ## HTTP basic auth
    #username = ""
    #password = ""

  [inputs.webhooks.rollbar]
    path = "/rollbar"

    ## HTTP basic auth
    #username = ""
    #password = ""

  [inputs.webhooks.papertrail]
    path = "/papertrail"

    ## HTTP basic auth
    #username = ""
    #password = ""

  [inputs.webhooks.particle]
    path = "/particle"

    ## HTTP basic auth
    #username = ""
    #password = ""
`
}
