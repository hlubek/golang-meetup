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

	fns := []func(b *testing.B){
		// START OMIT
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				logger.Info("hello", "count")
			}
		},
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				logger.LogAttrs(nil, slog.LevelInfo, "hello", slog.Int("count", 3))
			}
		},
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				logger.Info("hello", slog.Int("count", 3))
			}
		},
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				slog.Error("oops", "err", net.ErrClosed, "status", 500)
			}
		},
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				slog.LogAttrs(nil, slog.LevelError, "oops", slog.Any("err", net.ErrClosed), slog.Int("status", 500))
			}
		},
		// END OMIT
	}

	for i, fn := range fns {
		result := testing.Benchmark(fn)
		fmt.Printf("%d: alloc/ops=%d ns/op=%d\n", i, result.AllocsPerOp(), result.NsPerOp())
	}
}
