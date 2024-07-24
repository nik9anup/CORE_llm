/*
Package main demonstrates transaction management in PostgreSQL using Go's database/sql package and the pq driver. It showcases how to perform operations inside transactions, including querying and modifying data, rolling back transactions, and using contexts for transaction control.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Starts a transaction and inserts a new row into the "post" table.
- Queries the newly inserted row both outside and within the transaction.
- Rolls back the transaction to discard changes.
- Demonstrates a transaction with context, showing how to use contexts to control transaction lifespan and ensure proper cleanup.

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
                        (1,'Title One','Content One'),
                        (2,NULL,'Content Two');
*/
package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// selOne is the SQL statement to select a specific row from the "post" table by ID.
const selOne = "SELECT id,title,content FROM post WHERE ID = $1;"

// insert is the SQL statement to insert a new row into the "post" table.
const insert = "INSERT INTO post(ID,TITLE,CONTENT) VALUES (4,'Transaction Title','Transaction Content');"

// Post represents the structure of a post entity.
type Post struct {
	ID      int
	Title   string
	Content string
}

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// Insert a new row into the "post" table within the transaction
	_, err = tx.Exec(insert)
	if err != nil {
		panic(err)
	}

	// Query the newly inserted row from a separate session (outside transaction)
	p := Post{}
	if err := db.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("Error querying outside transaction:", err)
	}
	fmt.Println("Query outside transaction:", p)

	// Query the newly inserted row from within the transaction
	if err := tx.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("Error querying within transaction:", err)
	}
	fmt.Println("Query within transaction:", p)

	// Rollback the transaction to discard changes
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction rolled back successfully.")

	// Demonstrate transaction with context
	fmt.Println("\nTransaction with context")
	ctx, cancel := context.WithCancel(context.Background())
	tx, err = db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		panic(err)
	}

	// Insert a new row into the "post" table within the transaction
	_, err = tx.Exec(insert)
	if err != nil {
		panic(err)
	}

	// Cancel the context to simulate premature transaction termination
	cancel()

	// Commit the transaction (which should fail due to canceled context)
	err = tx.Commit()
	if err != nil {
		fmt.Println("Error committing transaction with canceled context:", err)
	}
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