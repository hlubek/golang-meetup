package domain

import (
	"errors"
	"fmt"
)

type FieldError struct {
	Field string
	Err   error
}

var _ error = FieldError{}

func (e FieldError) Error() string {
	return fmt.Sprintf("field %s: %v", e.Field, e.Err)
}

func (e FieldError) Unwrap() error { return e.Err }

var ErrRequired = errors.New("required")
var ErrAlreadyExists = errors.New("already exists")
