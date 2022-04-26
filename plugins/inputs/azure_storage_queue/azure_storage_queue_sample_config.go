//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package azure_storage_queue

func (a *AzureStorageQueue) SampleConfig() string {
	return `# Gather Azure Storage Queue metrics
[[inputs.azure_storage_queue]]
  ## Required Azure Storage Account name
  account_name = "mystorageaccount"

  ## Required Azure Storage Account access key
  account_key = "storageaccountaccesskey"

  ## Set to false to disable peeking age of oldest message (executes faster)
  # peek_oldest_message_age = true
`
}
