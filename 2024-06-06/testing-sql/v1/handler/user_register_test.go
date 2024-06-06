package handler_test

import (
	"context"
	"errors"
	"testing"
	"testing-sql/v1/dbtest"
	"testing-sql/v1/domain"
	"testing-sql/v1/domain/command"
	"testing-sql/v1/handler"
)

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		name               string
		cmd                command.UserRegister
		additionalFixtures []string
		expected           error
	}{
		{
			name:     "valid",
			cmd:      command.NewUserRegister("f.gopher@example.com", "Furry", "Gopher"),
			expected: nil,
		},
		{
			name:     "duplicate email",
			cmd:      command.NewUserRegister("admin@example.com", "Edmond", "Admin"),
			expected: domain.FieldError{Field: "email", Err: domain.ErrAlreadyExists},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := dbtest.CreateDatabase(t)
			dbtest.ExecFixtures(t, db, append([]string{"base"}, tc.additionalFixtures...)...)

			handler := handler.NewHandler(db)

			err := handler.UserRegister(context.Background(), tc.cmd)
			if err != nil {
				if tc.expected != nil {
					return
				}
				if errors.Is(err, tc.expected) {
					return
				}

				t.Fatalf("unexpected error: %v", err)
			}

			// TODO Implement assert callback (bonus)
		})
	}
}
