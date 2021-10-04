package main

import (
	"testing"
	"time"

	"net/http"
)

func TestNetHttpConnect(t *testing.T) {
	t.Parallel()

	client := http.Client{
		Timeout: time.Second,
	}
	start := time.Now()
	_, err := client.Get(ConnectUrl())
	assertTimeout(t, start, err, "Client.Timeout exceeded")
}

func TestNetHttpRead(t *testing.T) {
	t.Parallel()

	client := http.Client{
		Timeout: time.Second,
	}
	start := time.Now()
	_, err := client.Get(ReadUrl())
	assertTimeout(t, start, err, "Client.Timeout exceeded")
}
