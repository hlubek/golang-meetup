package dbtest

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// CreateDatabase creates an in-memory SQLite database for testing and returns a *sql.DB
func CreateDatabase(t *testing.T) *sql.DB {
	t.Helper()

	// Open a shared in-memory database (see https://github.com/mattn/go-sqlite3?tab=readme-ov-file#faq why `:memory:` is not enough)
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("Error opening in-memory database connection: %v", err)
	}
	// Make sure database is not closed until test is finished
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(-1)

	const schemaFile = "../repository/schema.sql"

	// Read SQL file
	bytes, err := os.ReadFile(schemaFile)
	if err != nil {
		t.Fatalf("Error reading SQL schema file %s: %v", schemaFile, err)
	}

	// Execute SQL queries to create schema
	_, err = db.Exec(string(bytes))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from file %s: %v", schemaFile, err)
	}

	t.Cleanup(func() {
		// Drop database after test
		err = db.Close()
		if err != nil {
			t.Fatalf("Failed to close the database: %v", err)
		}
	})

	return db
}
