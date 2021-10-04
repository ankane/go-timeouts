package main

import (
	"testing"
	"time"

	"context"
	"os/exec"
)

func TestOsExecTimeout(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	start := time.Now()
	err := exec.CommandContext(ctx, "sleep", "10").Run()
	assertTimeout(t, start, err, "signal: killed")
}
