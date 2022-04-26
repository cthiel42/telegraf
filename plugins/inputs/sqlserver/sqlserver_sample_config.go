//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package sqlserver

func (s *SQLServer) SampleConfig() string {
	return `# Read metrics from Microsoft SQL Server
[[inputs.sqlserver]]
  ## Specify instances to monitor with a list of connection strings.
  ## All connection parameters are optional.
  ## By default, the host is localhost, listening on default port, TCP 1433.
  ##   for Windows, the user is the currently running AD user (SSO).
  ##   See https://github.com/denisenkom/go-mssqldb for detailed connection
  ##   parameters, in particular, tls connections can be created like so:
  ##   "encrypt=true;certificate=<cert>;hostNameInCertificate=<SqlServer host fqdn>"
  servers = [
    "Server=192.168.1.10;Port=1433;User Id=<user>;Password=<pw>;app name=telegraf;log=1;",
  ]

  ## Authentication method
  ## valid methods: "connection_string", "AAD"
  # auth_method = "connection_string"

  ## "database_type" enables a specific set of queries depending on the database type. If specified, it replaces azuredb = true/false and query_version = 2
  ## In the config file, the sql server plugin section should be repeated each with a set of servers for a specific database_type.
  ## Possible values for database_type are - "SQLServer" or "AzureSQLDB" or "AzureSQLManagedInstance" or "AzureSQLPool"

  database_type = "SQLServer"

  ## A list of queries to include. If not specified, all the below listed queries are used.
  include_query = []

  ## A list of queries to explicitly ignore.
  exclude_query = ["SQLServerAvailabilityReplicaStates", "SQLServerDatabaseReplicaStates"]

  ## Queries enabled by default for database_type = "SQLServer" are -
  ## SQLServerPerformanceCounters, SQLServerWaitStatsCategorized, SQLServerDatabaseIO, SQLServerProperties, SQLServerMemoryClerks,
  ## SQLServerSchedulers, SQLServerRequests, SQLServerVolumeSpace, SQLServerCpu, SQLServerAvailabilityReplicaStates, SQLServerDatabaseReplicaStates,
  ## SQLServerRecentBackups

  ## Queries enabled by default for database_type = "AzureSQLDB" are -
  ## AzureSQLDBResourceStats, AzureSQLDBResourceGovernance, AzureSQLDBWaitStats, AzureSQLDBDatabaseIO, AzureSQLDBServerProperties,
  ## AzureSQLDBOsWaitstats, AzureSQLDBMemoryClerks, AzureSQLDBPerformanceCounters, AzureSQLDBRequests, AzureSQLDBSchedulers

  ## Queries enabled by default for database_type = "AzureSQLManagedInstance" are -
  ## AzureSQLMIResourceStats, AzureSQLMIResourceGovernance, AzureSQLMIDatabaseIO, AzureSQLMIServerProperties, AzureSQLMIOsWaitstats,
  ## AzureSQLMIMemoryClerks, AzureSQLMIPerformanceCounters, AzureSQLMIRequests, AzureSQLMISchedulers

  ## Queries enabled by default for database_type = "AzureSQLPool" are -
  ## AzureSQLPoolResourceStats, AzureSQLPoolResourceGovernance, AzureSQLPoolDatabaseIO, AzureSQLPoolWaitStats,
  ## AzureSQLPoolMemoryClerks, AzureSQLPoolPerformanceCounters, AzureSQLPoolSchedulers

  ## Following are old config settings
  ## You may use them only if you are using the earlier flavor of queries, however it is recommended to use
  ## the new mechanism of identifying the database_type there by use it's corresponding queries

  ## Optional parameter, setting this to 2 will use a new version
  ## of the collection queries that break compatibility with the original
  ## dashboards.
  ## Version 2 - is compatible from SQL Server 2012 and later versions and also for SQL Azure DB
  # query_version = 2

  ## If you are using AzureDB, setting this to true will gather resource utilization metrics
  # azuredb = false

  ## Toggling this to true will emit an additional metric called "sqlserver_telegraf_health".
  ## This metric tracks the count of attempted queries and successful queries for each SQL instance specified in "servers".
  ## The purpose of this metric is to assist with identifying and diagnosing any connectivity or query issues.
  ## This setting/metric is optional and is disabled by default.
  # health_metric = false

  ## Possible queries accross different versions of the collectors
  ## Queries enabled by default for specific Database Type

  ## database_type =  AzureSQLDB  by default collects the following queries
  ## - AzureSQLDBWaitStats
  ## - AzureSQLDBResourceStats
  ## - AzureSQLDBResourceGovernance
  ## - AzureSQLDBDatabaseIO
  ## - AzureSQLDBServerProperties
  ## - AzureSQLDBOsWaitstats
  ## - AzureSQLDBMemoryClerks
  ## - AzureSQLDBPerformanceCounters
  ## - AzureSQLDBRequests
  ## - AzureSQLDBSchedulers

  ## database_type =  AzureSQLManagedInstance by default collects the following queries
  ## - AzureSQLMIResourceStats
  ## - AzureSQLMIResourceGovernance
  ## - AzureSQLMIDatabaseIO
  ## - AzureSQLMIServerProperties
  ## - AzureSQLMIOsWaitstats
  ## - AzureSQLMIMemoryClerks
  ## - AzureSQLMIPerformanceCounters
  ## - AzureSQLMIRequests
  ## - AzureSQLMISchedulers

  ## database_type =  AzureSQLPool by default collects the following queries
  ## - AzureSQLPoolResourceStats
  ## - AzureSQLPoolResourceGovernance
  ## - AzureSQLPoolDatabaseIO
  ## - AzureSQLPoolOsWaitStats,
  ## - AzureSQLPoolMemoryClerks
  ## - AzureSQLPoolPerformanceCounters
  ## - AzureSQLPoolSchedulers

  ## database_type =  SQLServer by default collects the following queries
  ## - SQLServerPerformanceCounters
  ## - SQLServerWaitStatsCategorized
  ## - SQLServerDatabaseIO
  ## - SQLServerProperties
  ## - SQLServerMemoryClerks
  ## - SQLServerSchedulers
  ## - SQLServerRequests
  ## - SQLServerVolumeSpace
  ## - SQLServerCpu
  ## - SQLServerRecentBackups
  ## and following as optional (if mentioned in the include_query list)
  ## - SQLServerAvailabilityReplicaStates
  ## - SQLServerDatabaseReplicaStates

  ## Version 2 by default collects the following queries
  ## Version 2 is being deprecated, please consider using database_type.
  ## - PerformanceCounters
  ## - WaitStatsCategorized
  ## - DatabaseIO
  ## - ServerProperties
  ## - MemoryClerk
  ## - Schedulers
  ## - SqlRequests
  ## - VolumeSpace
  ## - Cpu

  ## Version 1 by default collects the following queries
  ## Version 1 is deprecated, please consider using database_type.
  ## - PerformanceCounters
  ## - WaitStatsCategorized
  ## - CPUHistory
  ## - DatabaseIO
  ## - DatabaseSize
  ## - DatabaseStats
  ## - DatabaseProperties
  ## - MemoryClerk
  ## - VolumeSpace
  ## - PerformanceMetrics
`
}
