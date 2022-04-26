//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package mailchimp

func (m *MailChimp) SampleConfig() string {
	return `# Gathers metrics from the /3.0/reports MailChimp API
[[inputs.mailchimp]]
  ## MailChimp API key
  ## get from https://admin.mailchimp.com/account/api/
  api_key = "" # required

  ## Reports for campaigns sent more than days_old ago will not be collected.
  ## 0 means collect all and is the default value.
  days_old = 0

  ## Campaign ID to get, if empty gets all campaigns, this option overrides days_old
  # campaign_id = ""
`
}
