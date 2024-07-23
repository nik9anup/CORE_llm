// Example 1: Connecting to a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Define your database connection string
const (
	host     = "localhost"
	port     = 5432
	user     = "your_username"
	password = "your_password"
	dbname   = "your_database"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL database!")
}





// Example 2: Querying data from a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	rows, err := db.QueryContext(context.Background(), "SELECT id, username, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 3: Inserting data into a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	username := "newuser"
	email := "newuser@example.com"

	_, err := db.ExecContext(context.Background(), "INSERT INTO users (username, email) VALUES ($1, $2)", username, email)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New user inserted successfully!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 4: Updating data in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	newUsername := "updateduser"
	userID := 1

	_, err := db.ExecContext(context.Background(), "UPDATE users SET username = $1 WHERE id = $2", newUsername, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User updated successfully!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 5: Deleting data from a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	userID := 1

	_, err := db.ExecContext(context.Background(), "DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User deleted successfully!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 6: Executing transactions in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Example of a transaction with two SQL operations
	username := "transactionuser"
	email := "transactionuser@example.com"

	// Insert operation inside transaction
	_, err = tx.ExecContext(context.Background(), "INSERT INTO users (username, email) VALUES ($1, $2)", username, email)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Another operation inside the same transaction
	_, err = tx.ExecContext(context.Background(), "UPDATE users SET email = $1 WHERE username = $2", "newemail@example.com", username)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction executed successfully!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 7: Using prepared statements in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	stmt, err := db.PrepareContext(context.Background(), "INSERT INTO users (username, email) VALUES ($1, $2)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Example of inserting multiple users using prepared statements
	users := []struct {
		Username string
		Email    string
	}{
		{"user1", "user1@example.com"},
		{"user2", "user2@example.com"},
		{"user3", "user3@example.com"},
	}

	for _, user := range users {
		_, err := stmt.ExecContext(context.Background(), user.Username, user.Email)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Users inserted using prepared statements!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 8: Handling errors in SQL operations using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	username := "nonexistentuser"

	// Example of handling errors when querying for a single user
	var email string
	err := db.QueryRowContext(context.Background(), "SELECT email FROM users WHERE username = $1", username).Scan(&email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("User '%s' not found.\n", username)
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Email of user '%s' is '%s'.\n", username, email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 9: Querying specific columns from a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	db := connectDB()

	rows, err := db.QueryContext(context.Background(), "SELECT id, username, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 10: Using contexts in PostgreSQL queries with sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	db := connectDB()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT id, username, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 11: Querying data with WHERE clause in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	username := "example_user"

	var user User
	err := db.QueryRowContext(context.Background(), "SELECT id, username, email FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No user found with username '%s'.\n", username)
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("User found: ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 12: Using transactions with error handling in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	username := "newuser"
	email := "newuser@example.com"

	_, err = tx.ExecContext(context.Background(), "INSERT INTO users (username, email) VALUES ($1, $2)", username, email)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.ExecContext(context.Background(), "UPDATE users SET email = $1 WHERE username = $2", "updatedemail@example.com", username)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction executed successfully!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 13: Deleting data using prepared statements in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	userID := 1

	stmt, err := db.PrepareContext(context.Background(), "DELETE FROM users WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(context.Background(), userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User deleted successfully!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 14: Inserting data with error checking in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	username := "newuser"
	email := "newuser@example.com"

	result, err := db.ExecContext(context.Background(), "INSERT INTO users (username, email) VALUES ($1, $2)", username, email)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted %d rows successfully!\n", rowsAffected)
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 15: Updating data using named parameters in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	userID := 1
	newEmail := "updateduser@example.com"

	_, err := db.ExecContext(context.Background(), "UPDATE users SET email = $1 WHERE id = $2", newEmail, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User updated successfully!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 16: Querying data using LIKE operator in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	likePattern := "%example%"

	rows, err := db.QueryContext(context.Background(), "SELECT id, username, email FROM users WHERE username LIKE $1", likePattern)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users matching pattern:")
	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 17: Handling errors in transactions in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	username := "newuser"
	email := "newuser@example.com"

	_, err = tx.ExecContext(context.Background(), "INSERT INTO users (username, email) VALUES ($1, $2)", username, email)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Intentionally wrong column name to force an error
	_, err = tx.ExecContext(context.Background(), "UPDATE users SET invalid_column = $1 WHERE username = $2", "updatedemail@example.com", username)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction executed successfully!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 18: Using QueryRow with context in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	userID := 1

	var user User
	err := db.QueryRowContext(context.Background(), "SELECT id, username, email FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No user found with ID %d.\n", userID)
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("User found: ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 19: Using named returns in queries in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
}

func main() {
	db := connectDB()

	rows, err := db.QueryContext(context.Background(), "SELECT id, username, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 20: Using timeouts in queries in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT id, username, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 21: Batch inserting data into a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	// Example of batch inserting multiple users
	users := []User{
		{Username: "user1", Email: "user1@example.com"},
		{Username: "user2", Email: "user2@example.com"},
		{Username: "user3", Email: "user3@example.com"},
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.PrepareContext(context.Background(), "INSERT INTO users (username, email) VALUES ($1, $2)")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, user := range users {
		_, err := stmt.ExecContext(context.Background(), user.Username, user.Email)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Batch insert completed successfully!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 22: Handling NULL values in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username sql.NullString
	Email    sql.NullString
}

func main() {
	db := connectDB()

	rows, err := db.QueryContext(context.Background(), "SELECT id, username, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users:")
	for _, user := range users {
		if user.Username.Valid && user.Email.Valid {
			fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username.String, user.Email.String)
		} else {
			fmt.Printf("ID: %d, Username: NULL, Email: NULL\n", user.ID)
		}
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 23: Using LastInsertId in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	username := "newuser"
	email := "newuser@example.com"

	var userID int

	err := db.QueryRowContext(context.Background(), "INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id", username, email).Scan(&userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New user inserted with ID: %d\n", userID)
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 24: Querying data using IN operator in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	userIDs := []int{1, 2, 3}

	// Prepare the query with placeholders for user IDs
	query := fmt.Sprintf("SELECT id, username, email FROM users WHERE id IN (%s)", strings.Trim(strings.Repeat("?, ", len(userIDs)), ", "))
	args := make([]interface{}, len(userIDs))
	for i, id := range userIDs {
		args[i] = id
	}

	rows, err := db.QueryContext(context.Background(), query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 25: Using subqueries in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	// Example of using subquery to fetch data
	query := `
		SELECT id, username, email
		FROM users
		WHERE id IN (
			SELECT user_id FROM posts WHERE category = 'tech'
		)
	`

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users who posted in 'tech' category:")
	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 26: Using EXISTS clause in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	// Example of using EXISTS clause
	query := `
		SELECT id, username, email
		FROM users u
		WHERE EXISTS (
			SELECT 1 FROM posts p WHERE p.user_id = u.id
		)
	`

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users who have posted at least once:")
	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 27: Using GROUP BY and HAVING clauses in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type UserStats struct {
	Username string
	Count    int
}

func main() {
	db := connectDB()

	// Example of using GROUP BY and HAVING clauses
	query := `
		SELECT username, COUNT(*) AS count
		FROM posts
		GROUP BY username
		HAVING COUNT(*) >= 3
	`

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var userStats []UserStats
	for rows.Next() {
		var stats UserStats
		err := rows.Scan(&stats.Username, &stats.Count)
		if err != nil {
			log.Fatal(err)
		}
		userStats = append(userStats, stats)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users with at least 3 posts:")
	for _, stats := range userStats {
		fmt.Printf("Username: %s, Count: %d\n", stats.Username, stats.Count)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 28: Using ORDER BY and LIMIT clauses in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	// Example of using ORDER BY and LIMIT clauses
	query := `
		SELECT id, username, email
		FROM users
		ORDER BY id DESC
		LIMIT 5
	`

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Latest 5 users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 29: Using OFFSET and FETCH clauses in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	// Example of using OFFSET and FETCH clauses
	query := `
		SELECT id, username, email
		FROM users
		ORDER BY id ASC
		OFFSET 5
		FETCH FIRST 5 ROWS ONLY
	`

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users from 6 to 10:")
	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 30: Using joins in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

func main() {
	db := connectDB()

	// Example of using INNER JOIN
	query := `
		SELECT u.id, u.username, u.email, r.role_name
		FROM users u
		INNER JOIN roles r ON u.role_id = r.id
	`

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users with roles:")
	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s, Role: %s\n", user.ID, user.Username, user.Email, user.Role)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 31: Using transactions with rollback in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	// Example of transaction with rollback on error
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Inserting a user
	username := "newuser"
	email := "newuser@example.com"
	_, err = tx.ExecContext(context.Background(), "INSERT INTO users (username, email) VALUES ($1, $2)", username, email)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Simulating an error that triggers rollback
	_, err = tx.ExecContext(context.Background(), "INSERT INTO non_existing_table (column1) VALUES ($1)", 123)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction committed successfully!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 32: Using conditional updates in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	// Example of conditional update based on a condition
	username := "existinguser"
	newEmail := "updatedemail@example.com"
	result, err := db.ExecContext(context.Background(), "UPDATE users SET email = $1 WHERE username = $2 AND email != $1", newEmail, username)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %d rows\n", rowsAffected)
}





// Example 33: Using subtransactions in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	// Example of using subtransactions
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Outer transaction
	_, err = tx.ExecContext(context.Background(), "INSERT INTO users (username, email) VALUES ($1, $2)", "user1", "user1@example.com")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Inner transaction
	savepoint := "savepoint1"
	_, err = tx.ExecContext(context.Background(), fmt.Sprintf("SAVEPOINT %s", savepoint))
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Rollback inner transaction
	_, err = tx.ExecContext(context.Background(), fmt.Sprintf("ROLLBACK TO SAVEPOINT %s", savepoint))
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction with subtransaction committed successfully!")
}





// Example 34: Using Common Table Expressions (CTE) in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	ManagerID int
}

func main() {
	db := connectDB()

	// Example of using Common Table Expressions (CTE)
	query := `
		WITH RECURSIVE EmpPath AS (
			SELECT id, first_name, last_name, manager_id
			FROM employees
			WHERE id = $1
			UNION ALL
			SELECT e.id, e.first_name, e.last_name, e.manager_id
			FROM employees e
			JOIN EmpPath ep ON e.id = ep.manager_id
		)
		SELECT id, first_name, last_name, manager_id
		FROM EmpPath;
	`

	rows, err := db.QueryContext(context.Background(), query, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.ManagerID)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, emp)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Employee hierarchy:")
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s %s, ManagerID: %d\n", emp.ID, emp.FirstName, emp.LastName, emp.ManagerID)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 35: Using window functions in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Salary    int
}

func main() {
	db := connectDB()

	// Example of using window functions
	query := `
		SELECT id, first_name, last_name, salary,
		       RANK() OVER (ORDER BY salary DESC) AS salary_rank
		FROM employees;
	`

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Salary)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, emp)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Employees ranked by salary:")
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s %s, Salary: %d, Rank: %d\n", emp.ID, emp.FirstName, emp.LastName, emp.Salary, emp.SalaryRank)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 36: Handling JSON data in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	db := connectDB()

	// Example of handling JSON data
	query := `
		SELECT id, name, price
		FROM products
		WHERE details->>'category' = $1
	`

	category := "electronics"
	rows, err := db.QueryContext(context.Background(), query, category)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Products in category '%s':\n", category)
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %d\n", p.ID, p.Name, p.Price)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 37: Using stored procedures in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	db := connectDB()

	// Example of calling a stored procedure
	var userID int
	username := "existinguser"
	email := "updatedemail@example.com"

	err := db.QueryRowContext(context.Background(), "CALL update_user($1, $2, $3)", username, email, &userID).Scan(&userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User updated with ID: %d\n", userID)
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 38: Using arrays in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	db := connectDB()

	// Example of using arrays in a query
	categories := []string{"electronics", "clothing", "books"}

	// Generate the placeholder string for array elements
	var placeholders []string
	for i := range categories {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}
	placeholderStr := strings.Join(placeholders, ",")

	query := fmt.Sprintf(`
		SELECT id, name, price
		FROM products
		WHERE category = ANY(ARRAY[%s])
	`, placeholderStr)

	args := make([]interface{}, len(categories))
	for i, category := range categories {
		args[i] = category
	}

	rows, err := db.QueryContext(context.Background(), query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Products in categories:", categories)
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %d\n", p.ID, p.Name, p.Price)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 39: Using conditional aggregates in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Order struct {
	ID       int
	Customer string
	Total    float64
}

func main() {
	db := connectDB()

	// Example of using conditional aggregates
	query := `
		SELECT id, customer, SUM(CASE WHEN status = 'completed' THEN total ELSE 0 END) AS total_completed
		FROM orders
		GROUP BY id, customer
	`

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.Customer, &order.Total)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Order totals:")
	for _, o := range orders {
		fmt.Printf("ID: %d, Customer: %s, Total Completed: %.2f\n", o.ID, o.Customer, o.Total)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 40: Using recursive queries in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	ManagerID sql.NullInt64
}

func main() {
	db := connectDB()

	// Example of using recursive queries
	query := `
		WITH RECURSIVE EmpPath AS (
			SELECT id, first_name, last_name, manager_id
			FROM employees
			WHERE id = $1
			UNION ALL
			SELECT e.id, e.first_name, e.last_name, e.manager_id
			FROM employees e
			JOIN EmpPath ep ON e.id = ep.manager_id
		)
		SELECT id, first_name, last_name, manager_id
		FROM EmpPath;
	`

	employeeID := 1
	rows, err := db.QueryContext(context.Background(), query, employeeID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.ManagerID)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, emp)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Employee hierarchy:")
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s %s, ManagerID: %v\n", emp.ID, emp.FirstName, emp.LastName, emp.ManagerID)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 41: Using upsert (INSERT ... ON CONFLICT UPDATE) in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	db := connectDB()

	// Example of upsert (INSERT ... ON CONFLICT UPDATE)
	product := Product{
		ID:    1,
		Name:  "Updated Product",
		Price: 100,
	}

	query := `
		INSERT INTO products (id, name, price)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE
		SET name = excluded.name, price = excluded.price
	`

	_, err := db.ExecContext(context.Background(), query, product.ID, product.Name, product.Price)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Product upserted successfully!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 42: Using window functions with PARTITION BY in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Department string
	Salary    int
	Rank      int
}

func main() {
	db := connectDB()

	// Example of using window functions with PARTITION BY
	query := `
		SELECT id, first_name, last_name, department, salary,
		       RANK() OVER (PARTITION BY department ORDER BY salary DESC) AS rank
		FROM employees;
	`

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Department, &emp.Salary, &emp.Rank)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, emp)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Employees ranked by salary within each department:")
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s %s, Department: %s, Salary: %d, Rank: %d\n", emp.ID, emp.FirstName, emp.LastName, emp.Department, emp.Salary, emp.Rank)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 43: Using triggers in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	db := connectDB()

	// Example of using triggers
	triggerName := "employee_audit_trigger"
	query := fmt.Sprintf(`
		CREATE OR REPLACE FUNCTION %s()
		RETURNS TRIGGER AS $$
		BEGIN
			-- Insert audit record on employee update
			INSERT INTO employee_audit (employee_id, action)
			VALUES (NEW.id, TG_OP);
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

		CREATE TRIGGER %s
		AFTER UPDATE OF first_name, last_name ON employees
		FOR EACH ROW
		EXECUTE FUNCTION %s();
	`, triggerName, triggerName, triggerName)

	_, err := db.ExecContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Trigger created successfully!")
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 44: Using full-text search in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	db := connectDB()

	// Example of using full-text search
	searchQuery := "apple"

	query := `
		SELECT id, name, price
		FROM products
		WHERE to_tsvector('english', name) @@ to_tsquery('english', $1)
	`

	rows, err := db.QueryContext(context.Background(), query, searchQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Products matching search query '%s':\n", searchQuery)
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %d\n", p.ID, p.Name, p.Price)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 45: Using UUIDs in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Product struct {
	ID    uuid.UUID
	Name  string
	Price int
}

func main() {
	db := connectDB()

	// Example of using UUIDs
	query := `
		SELECT id, name, price
		FROM products
		WHERE id = $1
	`

	productID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")

	row := db.QueryRowContext(context.Background(), query, productID)
	var product Product
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Product found: ID=%s, Name=%s, Price=%d\n", product.ID.String(), product.Name, product.Price)
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 46: Using Common Table Expressions (CTE) for pagination in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	db := connectDB()

	// Example of using Common Table Expressions (CTE) for pagination
	pageSize := 10
	pageNumber := 2

	query := fmt.Sprintf(`
		WITH paginated_products AS (
			SELECT id, name, price,
			       ROW_NUMBER() OVER (ORDER BY id) AS rownum
			FROM products
		)
		SELECT id, name, price
		FROM paginated_products
		WHERE rownum > $1 AND rownum <= $2
		ORDER BY rownum;
	`, (pageNumber-1)*pageSize, pageNumber*pageSize)

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Page %d of products (page size: %d):\n", pageNumber, pageSize)
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %d\n", p.ID, p.Name, p.Price)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 47: Using PostGIS for spatial queries in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Location struct {
	ID       int
	Name     string
	Location string // Assuming PostGIS stores location as text for simplicity
}

func main() {
	db := connectDB()

	// Example of using PostGIS for spatial queries
	query := `
		SELECT id, name, location
		FROM locations
		WHERE ST_DWithin(location::geography, ST_GeogFromText('SRID=4326;POINT(10 20)'), 10000) -- Within 10km of point (10, 20)
	`

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var locations []Location
	for rows.Next() {
		var loc Location
		err := rows.Scan(&loc.ID, &loc.Name, &loc.Location)
		if err != nil {
			log.Fatal(err)
		}
		locations = append(locations, loc)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Locations within 10km of point (10, 20):")
	for _, loc := range locations {
		fmt.Printf("ID: %d, Name: %s, Location: %s\n", loc.ID, loc.Name, loc.Location)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}






// Example 48: Using EXPLAIN for query analysis in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connectDB()

	// Example of using EXPLAIN for query analysis
	query := `
		EXPLAIN SELECT id, name, price
		FROM products
		WHERE price > $1
	`

	explainQuery := "100"

	rows, err := db.QueryContext(context.Background(), query, explainQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var explanation string
	for rows.Next() {
		err := rows.Scan(&explanation)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(explanation)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 49: Using database constraints (CHECK) in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	db := connectDB()

	// Example of using database constraints (CHECK)
	product := Product{
		ID:    1,
		Name:  "Valid Product",
		Price: 100,
	}

	query := `
		INSERT INTO products (id, name, price)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var productID int
	err := db.QueryRowContext(context.Background(), query, product.ID, product.Name, product.Price).Scan(&productID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Product inserted with ID: %d\n", productID)
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}





// Example 50: Using materialized views in a PostgreSQL database using sqlc

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type ProductSummary struct {
	Category   string
	TotalCount int
	AvgPrice   float64
}

func main() {
	db := connectDB()

	// Example of using materialized views
	query := `
		SELECT category, COUNT(*) AS total_count, AVG(price) AS avg_price
		FROM products
		GROUP BY category
	`

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var summaries []ProductSummary
	for rows.Next() {
		var summary ProductSummary
		err := rows.Scan(&summary.Category, &summary.TotalCount, &summary.AvgPrice)
		if err != nil {
			log.Fatal(err)
		}
		summaries = append(summaries, summary)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Product summaries:")
	for _, summary := range summaries {
		fmt.Printf("Category: %s, Total Count: %d, Avg Price: %.2f\n", summary.Category, summary.TotalCount, summary.AvgPrice)
	}
}

func connectDB() *sql.DB {
	connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
