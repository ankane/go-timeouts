package main

import (
	"testing"
	"time"

	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func TestGoSqlDriverMysqlConnect(t *testing.T) {
	t.Parallel()

	cfg := mysql.Config{
		Net:     "tcp",
		Addr:    ConnectHostAndPort(),
		Timeout: time.Second,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	start := time.Now()
	_, err = db.Exec("SELECT 1")
	assertTimeout(t, start, err, "i/o timeout")
}

func TestGoSqlDriverMysqlRead(t *testing.T) {
	t.Parallel()

	cfg := mysql.Config{
		Net:         "tcp",
		Addr:        ReadHostAndPort(),
		ReadTimeout: time.Second,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	start := time.Now()
	_, err = db.Exec("SELECT 1")
	// performs two retries
	assertTimeout(t, start.Add(2*time.Second), err, "bad connection")
}
