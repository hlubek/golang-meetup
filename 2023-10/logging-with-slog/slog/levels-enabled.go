package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()

	// START OMIT
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	})
	logger := slog.New(handler)

	if logger.Enabled(ctx, slog.LevelDebug) {
		// Prepare some expensive debug data
		logger.Debug("Debug message", slog.Int("count", 10))
	}

	logger.Info("Info message", slog.String("request_method", "POST"))
	// END OMIT
}
