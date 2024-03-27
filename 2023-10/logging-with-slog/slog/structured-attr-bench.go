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

	fns := []func(n int){
		// START OMIT
		func(n int) {
			logger.Info("hello", "count", 3)
		},
		func(n int) {
			logger.LogAttrs(nil, slog.LevelInfo, "hello", slog.Int("count", 3))
		},
		func(n int) {
			logger.Info("hello", slog.Int("count", 3))
		},
		func(n int) {
			slog.Error("oops", "err", net.ErrClosed, "status", 500)
		},
		func(n int) {
			slog.LogAttrs(nil, slog.LevelError, "oops", slog.Any("err", net.ErrClosed), slog.Int("status", 500))
		},
		// END OMIT
	}

	for i, fn := range fns {
		result := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(i)
			}
		})
		fmt.Printf("%d: alloc/ops=%d ns/op=%d\n", i, result.AllocsPerOp(), result.NsPerOp())
	}
}
