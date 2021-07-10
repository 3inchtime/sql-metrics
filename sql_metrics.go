package sql_metrics

import "database/sql"

type DB struct {
	config   Config
	database string
	db       *sql.DB
}

type Sql struct {
	sql string
	sqlType   string
	sqlTag    string
}
