/*
Package v2 is showing parts of an approach using PostgreSQL with generated schemas for isolated tests.

The basic setup works like in v1 or v0, but CreateDatabase will not create a new database, but generate a random schema per test and set it as the
`search_path` for the connection.
*/
package v2
