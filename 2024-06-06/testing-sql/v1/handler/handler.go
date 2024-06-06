package handler

import (
	"database/sql"
	"fmt"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) transactional(f func(*sql.Tx) error) error {
	tx, err := h.db.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	err = f(tx)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("rolling back transaction: %v, caused by error: %w", rollbackErr, err)
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
