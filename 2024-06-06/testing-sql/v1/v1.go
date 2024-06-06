/*
Package v1 is a next iteration of the test scenario using an in-memory SQLite database.

It adds the following things:
  - It uses a simple CQRS setup with command/handler and query/finder for reading and writing data
  - Repositories accept an interface for the database connection to allow use inside transactions
  - Data for tests is managed via SQL fixtures
  - Tests are using test-tables to test various scenarios in a single test
*/
package v1
