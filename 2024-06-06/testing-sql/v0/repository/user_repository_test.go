package repository_test

import (
	"context"
	"errors"
	"testing"

	"testing-sql/v0/dbtest"
	"testing-sql/v0/domain"
	"testing-sql/v0/repository"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
)

func TestFindUserByID_Found(t *testing.T) {
	db := dbtest.CreateDatabase(t)

	ctx := context.Background()

	// Set up
	err := repository.InsertUser(ctx, db, domain.User{
		ID:        uuid.FromStringOrNil("db12b74a-a7d0-4c80-9a75-b54589a7c157"),
		Email:     "test@example.com",
		Role:      domain.RoleUser,
		Firstname: "Furry",
		Lastname:  "Gopher",
	})
	require.NoError(t, err)

	// Execute test
	user, err := repository.FindUserByID(ctx, db, uuid.FromStringOrNil("db12b74a-a7d0-4c80-9a75-b54589a7c157"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user == nil {
		t.Fatal("user must not be nil")
	}

	if *user != (domain.User{
		ID:        uuid.FromStringOrNil("db12b74a-a7d0-4c80-9a75-b54589a7c157"),
		Email:     "test@example.com",
		Role:      domain.RoleUser,
		Firstname: "Furry",
		Lastname:  "Gopher",
	}) {
		t.Errorf("unexpected user %#v", *user)
	}
}

func TestFindUserByID_NotFound(t *testing.T) {
	db := dbtest.CreateDatabase(t)

	ctx := context.Background()

	_, err := repository.FindUserByID(ctx, db, uuid.FromStringOrNil("db12b74a-a7d0-4c80-9a75-b54589a7c157"))
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, repository.ErrNotFound) {
		t.Errorf("unexpected error: %v", err)
	}

}
