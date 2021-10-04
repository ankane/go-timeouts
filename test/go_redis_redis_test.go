package main

import (
	"testing"
	"time"

	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func TestGoRedisRedisConnect(t *testing.T) {
	t.Parallel()

	rdb := redis.NewClient(&redis.Options{
		Addr:        ConnectHostAndPort(),
		Password:    "",
		DB:          0,
		DialTimeout: time.Second,
		MaxRetries:  -1,
	})
	start := time.Now()
	_, err := rdb.Get(ctx, "key").Result()
	assertTimeout(t, start, err, "i/o timeout")
}

func TestGoRedisRedisRead(t *testing.T) {
	t.Parallel()

	rdb := redis.NewClient(&redis.Options{
		Addr:        ReadHostAndPort(),
		Password:    "",
		DB:          0,
		ReadTimeout: time.Second,
		MaxRetries:  -1,
	})
	start := time.Now()
	_, err := rdb.Get(ctx, "key").Result()
	assertTimeout(t, start, err, "i/o timeout")
}
