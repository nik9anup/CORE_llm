/*
Package main demonstrates how to use PostgreSQL database connection and context handling with "database/sql" package and the pq driver in Go.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It verifies the connection by performing pings with and without a context.
It also demonstrates creating a connection using db.Conn() and verifying its ping using a context.

Note: The use of time.Nanosecond for context timeout is not practical and is only used here to demonstrate context handling.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
- Update connStr with your PostgreSQL database connection details.
- Ensure PostgreSQL server is running on localhost:5432.

init.sql:
        DROP TABLE IF EXISTS post;
        CREATE TABLE post (
          ID serial,
          TITLE varchar(40),
          CONTENT varchar(255),
          CONSTRAINT pk_post PRIMARY KEY(ID)
        );
        SELECT * FROM post;
*/
package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // Import PostgreSQL driver package anonymously
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
	fmt.Println("Ping OK.")

	// Create a context with a timeout of 1 nanosecond (not practical, for demonstration purposes only)
	ctx, _ := context.WithTimeout(context.Background(), time.Nanosecond)

	// Ping the database with the context
	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	// Create a connection using db.Conn()
	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Close() // Ensure the connection is closed when function exits

	// Ping the connection with a context
	err = conn.PingContext(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Ping OK.")
}