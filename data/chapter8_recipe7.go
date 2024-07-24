/*
Package main demonstrates querying PostgreSQL database for multiple result sets and single row using Go's database/sql package and the pq driver.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Creates a database connection.
- Executes multiple SELECT queries (`sel` and `selOne`) to fetch multiple result sets and a single row.
- Retrieves and prints the selected posts and number.
- Handles errors gracefully using panic and printing error messages.

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
                       (2,NULL,'Content Two');
*/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import PostgreSQL driver package anonymously
)

// SQL SELECT statement to fetch title and content from the "post" table and a constant number
const sel = `SELECT title,content FROM post;
			SELECT 1234 NUM;`

// SQL SELECT statement to fetch title and content from the "post" table based on ID parameter
const selOne = "SELECT title,content FROM post WHERE ID = $1;"

// Post struct represents a row from the "post" table
type Post struct {
	Name sql.NullString
	Text sql.NullString
}

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Execute the first SELECT query (`sel`) to fetch multiple result sets
	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	// Retrieve and store posts from the first result set
	posts := []Post{}
	for rs.Next() {
		p := Post{}
		if err := rs.Scan(&p.Name, &p.Text); err != nil {
			panic(err)
		}
		posts = append(posts, p)
	}

	// Move to the next result set to retrieve the number
	var num int
	if rs.NextResultSet() {
		for rs.Next() {
			if err := rs.Scan(&num); err != nil {
				panic(err)
			}
		}
	}

	// Print retrieved posts and number
	fmt.Printf("Retrieved posts: %+v\n", posts)
	fmt.Printf("Retrieved number: %d\n", num)

	// Execute the second SELECT query (`selOne`) to fetch a single row based on ID parameter
	row := db.QueryRow(selOne, 100)
	or := Post{}
	if err := row.Scan(&or.Name, &or.Text); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	// Print retrieved single post
	fmt.Printf("Retrieved one post: %+v\n", or)
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