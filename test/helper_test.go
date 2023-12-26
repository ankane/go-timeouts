package main

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func assertTimeout(t *testing.T, start time.Time, err error, message string) {
	assert.Greater(t, time.Since(start).Seconds(), 0.99)
	assert.Less(t, time.Since(start).Seconds(), 1.25)
	assert.Contains(t, fmt.Sprint(err), message)
}

func ConnectHost() string {
	return "10.255.255.1"
}

func ConnectHostAndPort() string {
	return ConnectHost() + ":80"
}

func ConnectUrl() string {
	return "http://" + ConnectHost()
}

func ReadHost() string {
	return "127.0.0.1"
}

func ReadPort() int {
	return 4567
}

func ReadHostAndPort() string {
	return fmt.Sprintf("%s:%d", ReadHost(), ReadPort())
}

func ReadUrl() string {
	return "http://" + ReadHostAndPort()
}

func TestMain(m *testing.M) {
	l, err := net.Listen("tcp", ReadHostAndPort())
	if err != nil {
		panic(err)
	}
	defer l.Close()

	os.Exit(m.Run())
}
