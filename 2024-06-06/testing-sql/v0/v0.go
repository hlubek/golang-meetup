/*
Package v0 is a simple test scenario using an in-memory SQLite database.

It has the following approach:
  - Tests use a dbtest.CreateDatabase helper that opens a new database connection and closes it after tests are done via t.Cleanup
  - Data is managed just by using repository methods
*/
package v0
