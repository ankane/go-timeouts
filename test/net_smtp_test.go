package main

import (
	"testing"
	"time"

	"net"
	"net/smtp"
)

func TestNetSmtpConnect(t *testing.T) {
	t.Parallel()

	start := time.Now()
	_, err := net.DialTimeout("tcp", ConnectHostAndPort(), time.Second)
	assertTimeout(t, start, err, "i/o timeout")
}

func TestNetSmtpRead(t *testing.T) {
	t.Parallel()

	conn, err := net.Dial("tcp", ReadHostAndPort())
	if err != nil {
		panic(err)
	}
	err = conn.SetDeadline(time.Now().Add(time.Second))
	if err != nil {
		panic(err)
	}
	start := time.Now()
	_, err = smtp.NewClient(conn, "example.com")
	assertTimeout(t, start, err, "i/o timeout")
}
