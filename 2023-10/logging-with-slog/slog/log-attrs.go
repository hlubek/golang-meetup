package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"time"
)

func main() {
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	ctx := context.Background()
	// START OMIT
	slog.LogAttrs(ctx, slog.LevelDebug, "Debug message", slog.Int("count", 10))
	slog.LogAttrs(ctx, slog.LevelInfo, "Info message", slog.String("request_method", "POST"))
	slog.LogAttrs(ctx, slog.LevelWarn, "Warning message", slog.Duration("threshold", 10*time.Minute))
	slog.LogAttrs(ctx, slog.LevelError, "Error message", slog.Any("err", errors.New("something went wrong")))
	// END OMIT
}
