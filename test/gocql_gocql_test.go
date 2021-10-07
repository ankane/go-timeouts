package main

import (
	"testing"
	"time"

	"github.com/gocql/gocql"
)

func TestGocqlGocqlConnect(t *testing.T) {
	t.Parallel()

	cluster := gocql.NewCluster(ConnectHost())
	cluster.ConnectTimeout = time.Second

	start := time.Now()
	_, err := cluster.CreateSession()
	assertTimeout(t, start, err, "i/o timeout")
}

func TestGocqlGocqlRead(t *testing.T) {
	t.Parallel()

	cluster := gocql.NewCluster(ReadHostAndPort())
	cluster.ConnectTimeout = time.Second

	start := time.Now()
	_, err := cluster.CreateSession()
	assertTimeout(t, start, err, "timeout")
}
