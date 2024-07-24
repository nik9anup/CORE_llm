/*
Package main demonstrates querying a PostgreSQL database using Go's database/sql package and the pq driver. It retrieves a specific row from the "post" table based on the ID and demonstrates two methods to parse the result set into a map.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Queries a single row from the "post" table based on ID = 1.
- Parses the result set using both RawBytes and standard interface{} methods.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
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
                        (1,NULL,'Content One'),
                        (2,'Title Two','Content Two');
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// selOne is the SQL statement to select a specific row from the "post" table by ID.
const selOne = "SELECT id,title,content FROM post WHERE ID = $1;"

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Query the specific row from the "post" table based on ID = 1
	rows, err := db.Query(selOne, 1)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Get column names from the result set
	cols, _ := rows.Columns()

	// Iterate over the result set
	for rows.Next() {
		// Parse the row into a map using RawBytes method
		m := parseWithRawBytes(rows, cols)
		fmt.Println("Parsed with RawBytes:", m)

		// Parse the row into a map using standard interface{} method
		m = parseToMap(rows, cols)
		fmt.Println("Parsed with interface{}:", m)
	}
}

// parseWithRawBytes parses a row from sql.Rows into a map[string]interface{} using RawBytes method.
func parseWithRawBytes(rows *sql.Rows, cols []string) map[string]interface{} {
	vals := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(vals))
	for i := range vals {
		scanArgs[i] = &vals[i]
	}
	if err := rows.Scan(scanArgs...); err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	for i, col := range vals {
		if col == nil {
			m[cols[i]] = nil
		} else {
			m[cols[i]] = string(col)
		}
	}
	return m
}

// parseToMap parses a row from sql.Rows into a map[string]interface{} using standard interface{} method.
func parseToMap(rows *sql.Rows, cols []string) map[string]interface{} {
	values := make([]interface{}, len(cols))
	pointers := make([]interface{}, len(cols))
	for i := range values {
		pointers[i] = &values[i]
	}

	if err := rows.Scan(pointers...); err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	for i, colName := range cols {
		if values[i] == nil {
			m[colName] = nil
		} else {
			m[colName] = values[i]
		}
	}
	return m
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
