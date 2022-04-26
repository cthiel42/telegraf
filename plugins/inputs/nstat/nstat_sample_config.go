//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package nstat

func (ns *Nstat) SampleConfig() string {
	return `# Collect kernel snmp counters and network interface statistics
[[inputs.nstat]]
  ## file paths for proc files. If empty default paths will be used:
  ##    /proc/net/netstat, /proc/net/snmp, /proc/net/snmp6
  ## These can also be overridden with env variables, see README.
  proc_net_netstat = "/proc/net/netstat"
  proc_net_snmp = "/proc/net/snmp"
  proc_net_snmp6 = "/proc/net/snmp6"
  ## dump metrics with 0 values too
  dump_zeros       = true
`
}
