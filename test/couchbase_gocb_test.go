package main

import (
	"testing"
	"time"

	"github.com/couchbase/gocb/v2"
)

// TODO figure out what ConnectTimeout affects

func TestCouchbaseGocbKV(t *testing.T) {
	t.Parallel()

	cluster, err := gocb.Connect("couchbase://"+ReadHostAndPort(), gocb.ClusterOptions{
		TimeoutsConfig: gocb.TimeoutsConfig{
			KVTimeout: time.Second,
		},
	})
	collection := cluster.Bucket("bucket").Collection("collection")

	start := time.Now()
	_, err = collection.Get("key", &gocb.GetOptions{})
	assertTimeout(t, start, err, "ambiguous timeout")
}

func TestCouchbaseGocbQuery(t *testing.T) {
	t.Parallel()

	cluster, err := gocb.Connect("couchbase://"+ReadHostAndPort(), gocb.ClusterOptions{
		TimeoutsConfig: gocb.TimeoutsConfig{
			QueryTimeout: time.Second,
		},
	})

	start := time.Now()
	_, err = cluster.Query("SELECT 1", &gocb.QueryOptions{})
	assertTimeout(t, start, err, "ambiguous timeout")
}

func TestCouchbaseGocbManagement(t *testing.T) {
	t.Parallel()

	cluster, err := gocb.Connect("couchbase://"+ReadHostAndPort(), gocb.ClusterOptions{
		TimeoutsConfig: gocb.TimeoutsConfig{
			ManagementTimeout: time.Second,
		},
	})

	start := time.Now()
	_, err = cluster.Buckets().GetBucket("bucket", &gocb.GetBucketOptions{})
	assertTimeout(t, start, err, "ambiguous timeout")
}
