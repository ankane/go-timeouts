package main

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
)

func TestAwsAwsSdkGoConnect(t *testing.T) {
	t.Parallel()

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.AnonymousCredentials,
		Endpoint:    aws.String(ConnectUrl()),
		Region:      aws.String("us-east-1"),
		HTTPClient: &http.Client{
			Timeout: time.Second,
		},
		MaxRetries: aws.Int(0),
	}))
	client := s3.New(sess)

	start := time.Now()
	_, err := client.ListBuckets(&s3.ListBucketsInput{})
	assertTimeout(t, start, err, "Client.Timeout exceeded")
}

func TestAwsAwsSdkGoRead(t *testing.T) {
	t.Parallel()

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.AnonymousCredentials,
		Endpoint:    aws.String(ReadUrl()),
		Region:      aws.String("us-east-1"),
		HTTPClient: &http.Client{
			Timeout: time.Second,
		},
		MaxRetries: aws.Int(0),
	}))
	client := s3.New(sess)

	start := time.Now()
	_, err := client.ListBuckets(&s3.ListBucketsInput{})
	assertTimeout(t, start, err, "Client.Timeout exceeded")
}
