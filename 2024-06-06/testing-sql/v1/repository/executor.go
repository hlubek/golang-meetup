package repository

import (
	"context"
	"database/sql"
)

// Executor is the interface that wraps the basic Query, QueryRow and Exec methods.
// It allows to use *sql.DB and *sql.TX as an executor.
type Executor interface {
	QueryContext(ctx context.Context, sql string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, sql string, args ...any) *sql.Row
	ExecContext(ctx context.Context, sql string, args ...any) (sql.Result, error)
}

var _ Executor = &sql.DB{}
var _ Executor = &sql.Tx{}
