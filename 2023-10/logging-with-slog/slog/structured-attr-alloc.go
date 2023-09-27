package main

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"testing"
)

func main() {
	handler := slog.NewTextHandler(io.Discard, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// START OMIT
	fns := []func(){
		func() {
			logger.Info("hello", "count", 3)
		},
		func() {
			logger.LogAttrs(nil, slog.LevelInfo, "hello", slog.Int("count", 3))
		},
		func() {
			logger.Info("hello", slog.Int("count", 3))
		},
		func() {
			slog.Error("oops", "err", net.ErrClosed, "status", 500)
		},
		func() {
			slog.LogAttrs(nil, slog.LevelError, "oops", slog.Any("err", net.ErrClosed), slog.Int("status", 500))
		},
	}

	for i, fn := range fns {
		allocs := testing.AllocsPerRun(1, fn)
		fmt.Printf("%d: allocs=%0.2f\n", i, allocs)
	}

	// END OMIT
}
