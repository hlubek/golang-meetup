package domain

import "github.com/gofrs/uuid"

type User struct {
	ID        uuid.UUID
	Email     string
	Role      Role
	Firstname string
	Lastname  string
}
