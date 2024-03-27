package main

import (
	"log/slog"
	"os"
)

func main() {
	// START OMIT
	level := new(slog.LevelVar)
	level.Set(slog.LevelWarn)

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})
	logger := slog.New(handler)

	logger.Info("Info message", "location", "Isla Nublar")
	logger.Warn("Warn message", "cause", "dinosaurs!")

	// This could be done in a different goroutine (e.g. HTTP handler)
	level.Set(slog.LevelDebug)

	logger.Debug("Debug message")
	// END OMIT
}
