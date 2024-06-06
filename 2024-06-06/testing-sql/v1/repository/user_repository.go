package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gofrs/uuid"

	"testing-sql/v1/domain"
)

func FindUserByID(ctx context.Context, db Executor, id uuid.UUID) (*domain.User, error) {
	user := &domain.User{}
	err := db.
		QueryRowContext(
			ctx,
			"SELECT id, email, role, firstname, lastname FROM users WHERE id = $1",
			id,
		).
		Scan(&user.ID, &user.Email, &user.Role, &user.Firstname, &user.Lastname)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return user, nil
}

func CountUserByEmail(ctx context.Context, db Executor, email string) (int, error) {
	var count int
	err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func InsertUser(ctx context.Context, db Executor, user domain.User) error {
	_, err := db.
		ExecContext(
			ctx,
			"INSERT INTO users (id, email, role, firstname, lastname) VALUES ($1, $2, $3, $4, $5)",
			user.ID, user.Email, user.Role, user.Firstname, user.Lastname,
		)
	return err
}
