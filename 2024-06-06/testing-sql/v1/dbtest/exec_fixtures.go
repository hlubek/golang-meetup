package dbtest

import (
	"database/sql"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func ExecFixtures(t *testing.T, db *sql.DB, fixtureFilenames ...string) {
	t.Helper()

	for _, file := range fixtureFilenames {
		fixtureSource := fixtureSourcePath()
		data, err := os.ReadFile(filepath.Join(fixtureSource, file+".sql"))
		if err != nil {
			t.Fatalf("could not read fixture %s: %v", file, err)
		}
		_, err = db.Exec(string(data))
		if err != nil {
			t.Fatalf("could not execute fixture %q: %v", file, err)
		}
	}
}

func fixtureSourcePath() string {
	_, filename, _, _ := runtime.Caller(0)
	fixtureSource := filepath.Join(filepath.Dir(filename), "fixtures")
	return fixtureSource
}
