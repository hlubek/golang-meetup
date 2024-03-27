package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()
	// START OMIT
	const (
		LevelTrace  = slog.Level(-8)
		LevelNotice = slog.Level(2)
		LevelFatal  = slog.Level(12)
	)
	var LevelNames = map[slog.Leveler]string{LevelTrace: "TRACE", LevelNotice: "NOTICE", LevelFatal: "FATAL"}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: LevelTrace,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}
				a.Value = slog.StringValue(levelLabel)
			}
			return a
		},
	}))

	logger.Log(ctx, LevelNotice, "A notice message")
	// END OMIT
}
