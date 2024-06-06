package repository_test

import (
	"context"
	"testing"
	"testing-sql/mocking/domain"
	"testing-sql/mocking/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindUserByID_Found(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	userID := "53b23538-36fd-44c9-899d-1819234770d5"

	mock.
		ExpectQuery("SELECT id, email, role, firstname, lastname FROM users WHERE id = $1").
		WithArgs(userID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "email", "role", "firstname", "lastname"}).
				AddRow(userID, "test@example.com", "Administrator", "John", "Doe"),
		)

	user, err := repository.FindUserByID(context.Background(), db, uuid.FromStringOrNil(userID))
	require.NoError(t, err)

	if assert.NotNil(t, user) {
		assert.Equal(t, domain.User{
			ID:        uuid.FromStringOrNil(userID),
			Email:     "test@example.com",
			Role:      domain.RoleAdministrator,
			Firstname: "John",
			Lastname:  "Doe",
		}, *user)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
