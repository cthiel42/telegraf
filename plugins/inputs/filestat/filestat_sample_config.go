//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package filestat

func (*FileStat) SampleConfig() string {
	return `# Read stats about given file(s)
[[inputs.filestat]]
  ## Files to gather stats about.
  ## These accept standard unix glob matching rules, but with the addition of
  ## ** as a "super asterisk". See https://github.com/gobwas/glob.
  files = ["/etc/telegraf/telegraf.conf", "/var/log/**.log"]

  ## If true, read the entire file and calculate an md5 checksum.
  md5 = false
`
}
