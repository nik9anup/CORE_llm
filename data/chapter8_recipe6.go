/*
Package main demonstrates querying PostgreSQL database for column information using Go's database/sql package and the pq driver.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Creates a database connection.
- Executes a SELECT query (`sel`) to fetch all columns from the "post" table.
- Retrieves and prints information about the selected columns, including their names, types, and other properties.

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
        INSERT INTO post(ID,TITLE,CONTENT) VALUES
                        (1,'Title One','Content One'),
                        (2,'Title Two','Content Two');
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import PostgreSQL driver package anonymously
)

// SQL SELECT statement to fetch all columns from the "post" table
const sel = "SELECT * FROM post p"

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Execute the SELECT query to fetch all columns
	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	// Retrieve and print the selected column names
	columns, err := rs.Columns()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected columns: %v\n", columns)

	// Retrieve and print information about each column type
	colTypes, err := rs.ColumnTypes()
	if err != nil {
		panic(err)
	}
	for _, col := range colTypes {
		fmt.Println()
		fmt.Printf("%+v\n", col)
	}
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