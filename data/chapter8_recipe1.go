/*
Package main demonstrates how to connect to a PostgreSQL database using the "database/sql" package and the pq driver in Go.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It verifies the connection by pinging the database and prints "Ping OK"
if the connection is successful.
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import the PostgreSQL driver package anonymously
)

func main() {
	// Connection string for PostgreSQL database
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Print message indicating successful ping
	fmt.Println("Ping OK")
}