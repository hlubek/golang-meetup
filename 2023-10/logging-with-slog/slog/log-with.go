package main

import (
	"log/slog"
	"os"
)

type CreateUserAccountCmd struct {
	Username string
}

func main() {
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	createUserAccount(CreateUserAccountCmd{
		Username: "jane.doe",
	})
}

// START OMIT
func createUserAccount(cmd CreateUserAccountCmd) {
	logger := slog.With(
		slog.Group(
			"cmd",
			slog.String("type", "CreateUserAccount"),
			slog.String("username", cmd.Username),
		),
	)

	logger.Debug("Creating user account")

	// ... Do the magic
	accountID := "1234567890"

	logger.Info("Created user account", slog.String("accountID", accountID))
}

// END OMIT
