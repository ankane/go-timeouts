package main

import (
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

func TestBradfitzGomemcacheConnect(t *testing.T) {
	t.Parallel()

	mc := memcache.New(ConnectHostAndPort())
	mc.Timeout = time.Second
	start := time.Now()
	_, err := mc.Get("key")
	assertTimeout(t, start, err, "connect timeout")
}

func TestBradfitzGomemcacheRead(t *testing.T) {
	t.Parallel()

	mc := memcache.New(ReadHostAndPort())
	mc.Timeout = time.Second
	start := time.Now()
	_, err := mc.Get("key")
	assertTimeout(t, start, err, "i/o timeout")
}
