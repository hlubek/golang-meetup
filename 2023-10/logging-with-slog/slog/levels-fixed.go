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

	logger.Info("Info message", "location", "Isla Nublar")
	logger.Warn("Warn message", "cause", "dinosaurs!")
	logger.Log(ctx, slog.Level(12), "Super fatal")
	// END OMIT
}
