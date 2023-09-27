package main

import (
	"errors"
	"log/slog"
	"os"
	"testing"
	"time"
)

func main() {
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// START OMIT
	slog.Debug("Debug message", slog.Int("count", 10))
	slog.Info("Info message", slog.String("request_method", "POST"))
	slog.Warn("Warning message", slog.Duration("threshold", 10*time.Minute))
	slog.Error("Error message", slog.Any("err", errors.New("something went wrong")))
	// END OMIT

	testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			slog.Debug("Debug message", slog.Int("count", 10))
			slog.Info("Info message", slog.String("request_method", "POST"))
			slog.Warn("Warning message", slog.Duration("threshold", 10*time.Minute))
			slog.Error("Error message", slog.Any("err", errors.New("something went wrong")))
		}
	})
}
