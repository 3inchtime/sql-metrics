package sql_metrics

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

//type Conn interface {
//	Insert(ctx context.Context, sql Sql) error
//	Query(ctx context.Context, sql Sql)
//	Update(ctx context.Context, sql Sql)
//	Delete(ctx context.Context, sql Sql)
//}

func (d *DB) Insert(ctx context.Context, sql Sql, args ...interface{}) (sql.Result, error) {

	start := time.Now()
	res, err := d.db.ExecContext(ctx, sql.sql, args...)
	duration := time.Since(start)
	fmt.Println(duration)
	//metrics.DBDurationsSeconds.WithLabelValues(
	//	name,
	//	query.table,
	//	query.sqlType,
	//).Observe(duration.Seconds())
	//

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *DB) Update(ctx context.Context, sql Sql, args []interface{}) (sql.Result, error) {
	res, err := d.db.ExecContext(ctx, sql.sql, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *DB) QueryRow(ctx context.Context, sql Sql) *sql.Row {
	rows := d.db.QueryRow(sql.sql)
	return rows
}

func (d *DB) Query(ctx context.Context, sql Sql) (*sql.Rows, error) {
	rows, err := d.db.Query(sql.sql)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *DB) Delete(ctx context.Context, sql Sql, args []interface{}) (sql.Result, error) {
	res, err := d.db.ExecContext(ctx, sql.sql, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
