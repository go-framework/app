package sql

type Config struct {
	// Database driver name.
	DriverName string `json:"driverName" yaml:"driverName"`
	// The Driver-specific data source name.
	// See https://github.com/go-sql-driver/mysql#dsn-data-source-name
	DSN string `json:"dsn" yaml:"dsn"`
	// The maximum number of connections in the idle connection pool
	// If n <= 0, no idle connections are retained, default is 2.
	MaxIdleConns int `json:"maxIdleConns" yaml:"maxIdleConns"`
	// The maximum number of open connections to the database.
	// If n <= 0, then there is no limit on the number of open connections default is 1.
	MaxOpenConns int `json:"maxOpenConns" yaml:"maxOpenConns"`
}
