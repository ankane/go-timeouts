package main

import (
	"testing"
	"time"

	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"net/http"
)

func TestAwsAwsSdkGoV2Connect(t *testing.T) {
	t.Parallel()

	resolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{URL: ConnectUrl()}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(aws.AnonymousCredentials{}),
		config.WithEndpointResolver(resolver),
		config.WithHTTPClient(&http.Client{Timeout: time.Second}),
		config.WithRetryer(func() aws.Retryer { return aws.NopRetryer{} }),
	)
	if err != nil {
		panic(err)
	}
	client := s3.NewFromConfig(cfg)

	start := time.Now()
	_, err = client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	assertTimeout(t, start, err, "Client.Timeout exceeded")
}

func TestAwsAwsSdkGoV2Read(t *testing.T) {
	t.Parallel()

	resolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{URL: ReadUrl()}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(aws.AnonymousCredentials{}),
		config.WithEndpointResolver(resolver),
		config.WithHTTPClient(&http.Client{Timeout: time.Second}),
		config.WithRetryer(func() aws.Retryer { return aws.NopRetryer{} }),
	)
	if err != nil {
		panic(err)
	}
	client := s3.NewFromConfig(cfg)

	start := time.Now()
	_, err = client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	assertTimeout(t, start, err, "Client.Timeout exceeded")
}
