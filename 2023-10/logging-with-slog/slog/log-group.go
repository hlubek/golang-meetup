package main

import (
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
	slog.Debug(
		"Debug message",
		slog.Group(
			"request",
			slog.String("method", "POST"),
			slog.Duration("duration", 17*time.Millisecond),
		),
	)
	// END OMIT
}
