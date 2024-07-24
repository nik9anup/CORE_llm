/*
Package main demonstrates how to perform batch insert operations into a PostgreSQL database table using prepared statements with Go's database/sql package and the pq driver.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Truncates the "post" table to clear existing data.
- Uses a prepared statement to insert multiple rows into the "post" table from a predefined slice of structs.
- Prints the number of rows successfully inserted.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
- Update connStr with your PostgreSQL database connection details.
- Ensure PostgreSQL server is running on localhost:5432 and the "post" table exists in the "example" database.


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
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import PostgreSQL driver package anonymously
)

// SQL statements for database operations
const (
	trunc = "TRUNCATE TABLE post;"                        // Truncate (empty) "post" table
	ins   = "INSERT INTO post(ID,TITLE,CONTENT) VALUES ($1,$2,$3)" // Insert statement with placeholders
)

// Struct for test data
var testTable = []struct {
	ID      int
	Title   string
	Content string
}{
	{1, "Title One", "Content of title one"},
	{2, "Title Two", "Content of title two"},
	{3, "Title Three", "Content of title three"},
}

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

	// Prepare the insert statement
	stm, err := db.Prepare(ins)
	defer stm.Close() // Ensure the prepared statement is closed
	if err != nil {
		panic(err)
	}

	inserted := int64(0)
	// Iterate over testTable and insert rows using the prepared statement
	for _, val := range testTable {
		fmt.Printf("Inserting record ID: %d\n", val.ID)
		// Execute the prepared statement with values from the struct
		r, err := stm.Exec(val.ID, val.Title, val.Content)
		if err != nil {
			fmt.Printf("Cannot insert record ID : %d\n", val.ID)
		}
		// Retrieve the number of affected rows and accumulate the total
		if affected, err := r.RowsAffected(); err == nil {
			inserted += affected
		}
	}

	// Print the total number of rows successfully inserted
	fmt.Printf("Result: Inserted %d rows.\n", inserted)
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