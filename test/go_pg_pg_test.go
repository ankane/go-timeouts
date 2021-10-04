package main

import (
	"testing"
	"time"

	"github.com/go-pg/pg/v10"
)

func TestGoPgPgConnect(t *testing.T) {
	t.Parallel()

	db := pg.Connect(&pg.Options{
		Addr:        ConnectHostAndPort(),
		DialTimeout: time.Second,
	})
	defer db.Close()

	start := time.Now()
	_, err := db.Exec("SELECT 1")
	assertTimeout(t, start, err, "i/o timeout")
}

func TestGoPgPgRead(t *testing.T) {
	t.Parallel()

	db := pg.Connect(&pg.Options{
		Addr:        ReadHostAndPort(),
		ReadTimeout: time.Second,
	})
	defer db.Close()

	start := time.Now()
	_, err := db.Exec("SELECT 1")
	assertTimeout(t, start, err, "i/o timeout")
}
