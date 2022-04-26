//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package application_insights

func (a *ApplicationInsights) SampleConfig() string {
	return `# Send metrics to Azure Application Insights
[[outputs.application_insights]]
  ## Instrumentation key of the Application Insights resource.
  instrumentation_key = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"

  ## Regions that require endpoint modification https://docs.microsoft.com/en-us/azure/azure-monitor/app/custom-endpoints
  # endpoint_url = "https://dc.services.visualstudio.com/v2/track"

  ## Timeout for closing (default: 5s).
  # timeout = "5s"

  ## Enable additional diagnostic logging.
  # enable_diagnostic_logging = false

  ## Context Tag Sources add Application Insights context tags to a tag value.
  ##
  ## For list of allowed context tag keys see:
  ## https://github.com/microsoft/ApplicationInsights-Go/blob/master/appinsights/contracts/contexttagkeys.go
  # [outputs.application_insights.context_tag_sources]
  #   "ai.cloud.role" = "kubernetes_container_name"
  #   "ai.cloud.roleInstance" = "kubernetes_pod_name"
`
}
