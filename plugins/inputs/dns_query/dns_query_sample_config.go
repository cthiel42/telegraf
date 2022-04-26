//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package dns_query

func (d *DNSQuery) SampleConfig() string {
	return `# Query given DNS server and gives statistics
[[inputs.dns_query]]
  ## servers to query
  servers = ["8.8.8.8"]

  ## Network is the network protocol name.
  # network = "udp"

  ## Domains or subdomains to query.
  # domains = ["."]

  ## Query record type.
  ## Possible values: A, AAAA, CNAME, MX, NS, PTR, TXT, SOA, SPF, SRV.
  # record_type = "A"

  ## Dns server port.
  # port = 53

  ## Query timeout in seconds.
  # timeout = 2
`
}
