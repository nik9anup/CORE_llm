/*
Package main demonstrates basic CRUD operations (Create, Read, Update, Delete) with PostgreSQL using Go's database/sql package and the pq driver.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Truncates the "post" table.
- Inserts predefined rows into the "post" table.
- Selects and counts the number of rows in the "post" table.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
- Update connStr with your PostgreSQL database connection details.
- Ensure PostgreSQL server is running on localhost:5432 and the "post" table exists in the "example" database.

*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import PostgreSQL driver package anonymously
)

// SQL statements for database operations
const (
	sel   = "SELECT * FROM post;"                                     // Select all rows from "post" table
	trunc = "TRUNCATE TABLE post;"                                     // Truncate (empty) "post" table
	ins   = "INSERT INTO post(ID,TITLE,CONTENT) VALUES (1,'Title 1','Content 1'), (2,'Title 2','Content 2');" // Insert rows into "post" table
)

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Truncate the "post" table
	_, err := db.Exec(trunc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table truncated.")

	// Insert rows into the "post" table
	r, err := db.Exec(ins)
	if err != nil {
		panic(err)
	}
	affected, err := r.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted rows count: %d\n", affected)

	// Query and count rows from the "post" table
	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}
	count := 0
	for rs.Next() {
		if rs.Err() != nil {
			fmt.Println(rs.Err())
			continue
		}
		count++
	}
	fmt.Printf("Total of %d rows selected.\n", count)
}

// createConnection establishes a connection to the PostgreSQL database and returns the *sql.DB object.
func createConnection() *sql.DB {
	// Connection string for PostgreSQL database
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}