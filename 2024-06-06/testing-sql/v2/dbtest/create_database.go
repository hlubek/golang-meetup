package dbtest

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/apex/log"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	apexlogutils_pgx "github.com/networkteam/apexlogutils/pgx"
	"github.com/pressly/goose"

	"testing-sql/v2/security/helper"
)

const (
	dbPort = 5432
	// The database needs to be created (once) before running the tests.
	dbName = "dbtest-test"
)

func CreateDatabase(t *testing.T) *sql.DB {
	t.Helper()

	schemaName := generateSchemaName(t)
	connStr := configureConnection(t, schemaName)

	db := openDatabase(t, connStr)
	createSchema(t, db, schemaName)
	runMigrations(t, db)
	cleanup(t, db, schemaName, connStr)
	return db
}

func generateSchemaName(t *testing.T) string {
	randomSuffix, err := helper.GenerateRandomString(12)
	if err != nil {
		t.Fatalf("Failed to generate random string: %v", err)
	}
	return "test-" + strings.ToLower(randomSuffix)
}

func generatePostgresDSN(schemaName string) string {
	return fmt.Sprintf("host=localhost port=%d dbname=%s sslmode=disable search_path=%s", dbPort, dbName, schemaName)
}

func configureConnection(t *testing.T, schemaName string) string {
	dsn := generatePostgresDSN(schemaName)

	connConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		t.Fatalf("Failed to parse PostgreSQL connection string: %v", err)
	}
	connConfig.Logger = apexlogutils_pgx.NewLogger(log.Log)
	connConfig.LogLevel = pgx.LogLevelDebug

	return stdlib.RegisterConnConfig(connConfig)
}

func openDatabase(t *testing.T, connStr string) *sql.DB {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		t.Fatalf("Failed to open database connection: %v", err)
	}
	return db
}

func createSchema(t *testing.T, db *sql.DB, schemaName string) {
	_, err := db.Exec("CREATE SCHEMA \"" + schemaName + "\"")
	if err != nil {
		t.Fatalf("Failed to create schema: %v", err)
	}
}

func runMigrations(t *testing.T, db *sql.DB) {
	_, filename, _, _ := runtime.Caller(0)
	migrationSource := filepath.Join(filepath.Dir(filename), "../repository/migrations")

	goose.SetLogger(noopLogger{})

	err := goose.Up(db, migrationSource)
	if err != nil {
		t.Fatalf("Failed to execute migrations: %v", err)
	}
}

func cleanup(t *testing.T, db *sql.DB, schemaName string, connStr string) {
	t.Cleanup(func() {
		defer stdlib.UnregisterConnConfig(connStr)

		_, err := db.Exec("DROP SCHEMA \"" + schemaName + "\" CASCADE")
		if err != nil {
			t.Fatalf("Failed to drop schema: %v", err)
		}
		err = db.Close()
		if err != nil {
			t.Logf("Error closing test DB: %v", err)
		}
	})
}

type noopLogger struct{}

func (n noopLogger) Fatal(...interface{}) {}

func (n noopLogger) Fatalf(string, ...interface{}) {}

func (n noopLogger) Print(...interface{}) {}

func (n noopLogger) Println(...interface{}) {}

func (n noopLogger) Printf(string, ...interface{}) {}
