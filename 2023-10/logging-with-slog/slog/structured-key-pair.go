package main

import (
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

	// START OMIT
	slog.Debug("Debug message", "count", 10)
	slog.Info("Info message", "request_method", "POST")
	slog.Warn("Warning message", "threshold", 10*time.Minute)
	slog.Error("Error message", "err", errors.New("something went wrong"))
	// END OMIT
}
