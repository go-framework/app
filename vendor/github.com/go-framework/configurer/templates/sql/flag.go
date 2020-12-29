package sql

import (
	"github.com/spf13/pflag"
)

func AddSqlConfigFlags(flag *pflag.FlagSet, config *Config) {
	flag.StringVarP(&config.DriverName, "driver-name", "n", config.DriverName, "Database driver name")
	flag.StringVarP(&config.DSN, "dsn", "d", config.DSN, "The Driver-specific data source name")
	flag.IntVarP(&config.MaxIdleConns, "max-idle-conns", "", config.MaxIdleConns, "The maximum number of open connections to the database")
	flag.IntVarP(&config.MaxIdleConns, "max-idle-conns", "", config.MaxIdleConns, "The maximum number of connections in the idle connection pool")
}
