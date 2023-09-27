package main

import (
	"log"
	"log/slog"
	"os"
)

// START OMIT
func main() {
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")

	log.Println("Ye olde logger")
}

// END OMIT
