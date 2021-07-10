package sql_metrics

import (
	"database/sql"
	"fmt"
	"time"
)

func New(database string, cfg Config) (*DB, error) {
	if database == "" {
		return nil, ErrNoDatabase
	}

	dsn := getDSN(cfg, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)

	return &DB{
		config:   cfg,
		database: database,
		db:       db,
	}, nil
}

func getDSN(cfg Config, database string) string {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8", cfg.Username, cfg.Password, cfg.Server, cfg.Port, database)
	return conn
}

func (d *DB) Ping() error {
	err := d.db.Ping()
	return err
}
