//go:build !windows
// +build !windows

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package lustre2

func (l *Lustre2) SampleConfig() string {
	return `# Read metrics from local Lustre service on OST, MDS
[[inputs.lustre2]]
  ## An array of /proc globs to search for Lustre stats
  ## If not specified, the default will work on Lustre 2.5.x
  ##
  # ost_procfiles = [
  #   "/proc/fs/lustre/obdfilter/*/stats",
  #   "/proc/fs/lustre/osd-ldiskfs/*/stats",
  #   "/proc/fs/lustre/obdfilter/*/job_stats",
  #   "/proc/fs/lustre/obdfilter/*/exports/*/stats",
  # ]
  # mds_procfiles = [
  #   "/proc/fs/lustre/mdt/*/md_stats",
  #   "/proc/fs/lustre/mdt/*/job_stats",
  #   "/proc/fs/lustre/mdt/*/exports/*/stats",
  # ]
`
}
