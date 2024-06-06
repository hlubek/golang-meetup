package command

import (
	"strings"

	"github.com/gofrs/uuid"

	"testing-sql/v1/domain"
)

type UserRegister struct {
	ID              uuid.UUID
	Email           string
	Firstname       string
	Lastname        string
	IsAdministrator bool
}

func NewUserRegister(email, firstname, lastname string) UserRegister {
	return UserRegister{
		ID:        uuid.Must(uuid.NewV7()),
		Email:     strings.TrimSpace(strings.ToLower(email)),
		Firstname: strings.TrimSpace(firstname),
		Lastname:  strings.TrimSpace(lastname),
	}
}

func (cmd UserRegister) Validate() error {
	if cmd.Email == "" {
		return domain.FieldError{Field: "email", Err: domain.ErrRequired}
	}

	return nil
}

func (cmd UserRegister) User() domain.User {
	role := domain.RoleUser
	if cmd.IsAdministrator {
		role = domain.RoleAdministrator
	}

	return domain.User{
		ID:        cmd.ID,
		Email:     cmd.Email,
		Firstname: cmd.Firstname,
		Lastname:  cmd.Lastname,
		Role:      role,
	}
}
