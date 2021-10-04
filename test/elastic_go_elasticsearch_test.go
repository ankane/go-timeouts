package main

import (
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"net"
	"net/http"
)

func TestElasticGoElastisearchConnect(t *testing.T) {
	t.Parallel()

	cfg := elasticsearch.Config{
		Addresses: []string{ConnectUrl()},
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second,
			}).DialContext,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	_, err = es.Search()
	assertTimeout(t, start, err, "i/o timeout")
}

func TestElasticGoElastisearchRead(t *testing.T) {
	t.Parallel()

	cfg := elasticsearch.Config{
		Addresses: []string{ReadUrl()},
		Transport: &http.Transport{
			ResponseHeaderTimeout: time.Second,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	_, err = es.Search()
	assertTimeout(t, start, err, "timeout awaiting response headers")
}
