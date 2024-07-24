/*
Package main demonstrates calling stored procedures/functions in both PostgreSQL and MySQL databases using Go's database/sql package and respective drivers.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Calls a PostgreSQL function named format_name with three parameters and retrieves the result.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go
- github.com/go-sql-driver/mysql: MySQL driver for Go

Usage:
- Ensure PostgreSQL server is running on localhost:5432 and the "example" database contains the format_name function.

init.sql:
CREATE OR REPLACE FUNCTION format_name
        (firstname Text,lastname Text,age INT) RETURNS 
        VARCHAR AS $$
        BEGIN
          RETURN trim(firstname) ||' '||trim(lastname) ||' ('||age||')';
        END;
        $$ LANGUAGE plpgsql;
*/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	_ "github.com/lib/pq"              // PostgreSQL driver
)

// call is the SQL statement to call the format_name function in PostgreSQL.
const call = "select * from format_name($1,$2,$3)"

// callMySQL is the SQL statement to call the simpleproc stored procedure in MySQL.
const callMySQL = "CALL simpleproc(?)"

// Result represents the structure of the result from calling the stored procedure/function.
type Result struct {
	Name     string
	Category int
}

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close()

	// Initialize a Result struct to hold the returned values
	r := Result{}

	// Call the PostgreSQL function format_name with parameters and scan the result into Result struct
	if err := db.QueryRow(call, "John", "Doe", 32).Scan(&r.Name); err != nil {
		panic(err)
	}
	fmt.Printf("Result is: %+v\n", r)
}

// createConnection establishes a connection to the PostgreSQL database and returns the connection object.
func createConnection() *sql.DB {
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")
	return db
}