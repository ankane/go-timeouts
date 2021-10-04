package main

import (
	"testing"
	"time"

	opensearch "github.com/opensearch-project/opensearch-go"
	"net"
	"net/http"
)

func TestOpensearchProjectOpensearchGoConnect(t *testing.T) {
	t.Parallel()

	cfg := opensearch.Config{
		Addresses: []string{ConnectUrl()},
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second,
			}).DialContext,
		},
	}
	client, err := opensearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	_, err = client.Search()
	assertTimeout(t, start, err, "i/o timeout")
}

func TestOpensearchProjectOpensearchGoRead(t *testing.T) {
	t.Parallel()

	cfg := opensearch.Config{
		Addresses: []string{ReadUrl()},
		Transport: &http.Transport{
			ResponseHeaderTimeout: time.Second,
		},
	}
	client, err := opensearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	_, err = client.Search()
	assertTimeout(t, start, err, "timeout awaiting response headers")
}
