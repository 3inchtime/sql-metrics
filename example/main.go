package main

import "sql-metrics"

import (
	"context"
	"testing"
)

var cfg = &sql-metrics.Config{
	Username: "root",
	Password: "123456",
	Server: "192.168.1.103",
	Port: "3306",
	MaxLifetime: 30,
	MaxIdleConn: 10,
	MaxOpenConn: 100,
}

var testDB, _ = New("user", *cfg)

func TestInsert(t *testing.T) {
	sql := &Sql{
		sql: "INSERT INTO ts_user (user_id, create_time, update_time) VALUES (?,?,?)",
		sqlType: "insert",
		sqlTag: "insert user",
	}
	testDB.Insert(context.Background(), *sql, "test1", 2312312312312, 21312312312)
}