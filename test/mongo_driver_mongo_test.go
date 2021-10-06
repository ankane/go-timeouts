package main

import (
	"testing"
	"time"

	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoDriverMongoConnect(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+ConnectHostAndPort()))

	start := time.Now()
	err = client.Ping(ctx, nil)
	assertTimeout(t, start, err, "server selection error: context deadline exceeded")
}

func TestMongoDriverMongoRead(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+ReadHostAndPort()))

	start := time.Now()
	err = client.Ping(ctx, nil)
	assertTimeout(t, start, err, "server selection error: context deadline exceeded")
}
