/*
Package main demonstrates executing a large SQL query with context timeout using Go's database/sql package and the pq driver for PostgreSQL.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Creates a database connection.
- Uses a context with a timeout of 20 microseconds to limit the query execution time.
- Executes a SELECT query (`sel`) that performs a cross join with a large series to generate a significant number of rows.
- Cancels the query if it exceeds the context timeout.
- Prints the number of rows returned by the query.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
- Update connStr with your PostgreSQL database connection details.
- Ensure PostgreSQL server is running on localhost:5432 and the "post" table exists in the "example" database.
- Adjust the context timeout (`20*time.Microsecond`) as needed for your query's expected execution time.


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
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // Import PostgreSQL driver package anonymously
)

// SQL SELECT statement to generate a large number of rows
const sel = "SELECT * FROM post p CROSS JOIN (SELECT 1 FROM generate_series(1,1000000)) tbl"

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Create a context with a timeout of 20 microseconds
	ctx, canc := context.WithTimeout(context.Background(), 20*time.Microsecond)
	defer canc() // Ensure the cancellation function is called to cancel the query if it exceeds the timeout

	// Execute the query with context timeout
	rows, err := db.QueryContext(ctx, sel)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// Count the number of rows returned by the query
	count := 0
	for rows.Next() {
		if rows.Err() != nil {
			fmt.Println(rows.Err())
			continue
		}
		count++
	}

	// Print the number of rows returned
	fmt.Printf("%d rows returned\n", count)
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