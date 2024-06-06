package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"testing-sql/v1/domain"
	"testing-sql/v1/domain/command"
	"testing-sql/v1/repository"
)

func (h *Handler) UserRegister(ctx context.Context, cmd command.UserRegister) error {
	slog.Debug("Handling command UserRegister")

	err := cmd.Validate()
	if err != nil {
		return err
	}

	user := cmd.User()

	err = h.transactional(func(tx *sql.Tx) error {
		count, err := repository.CountUserByEmail(ctx, tx, cmd.Email)
		if err != nil {
			return fmt.Errorf("counting users: %w", err)
		}
		if count > 0 {
			return domain.FieldError{Field: "email", Err: domain.ErrAlreadyExists}
		}

		err = repository.InsertUser(ctx, tx, user)
		if err != nil {
			return fmt.Errorf("inserting user: %w", err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	slog.Info("Registered user", "email", cmd.Email, "userID", cmd.ID)

	return nil
}
