/*
Example 1: Connect to a SQLite Database

Description:
This example demonstrates how to connect to a SQLite database using the go-sqlite3 library.

Dependencies:
Make sure you have imported the go-sqlite3 package:
    go get github.com/mattn/go-sqlite3

SQLite Database:
Assume you have a SQLite database named 'example.db' in the current directory.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    fmt.Println("Connected to SQLite database")
}





/*
Example 2: Create Table in SQLite Database

Description:
This example demonstrates how to create a table in a SQLite database using the go-sqlite3 library.

SQLite Database:
Assume you have connected to a SQLite database named 'example.db'.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Create a new table
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            age INTEGER
        )
    `)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Table 'users' created successfully")
}





/*
Example 3: Insert Data into SQLite Table

Description:
This example demonstrates how to insert data into a SQLite table using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in Example 2).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Prepare statement for inserting data
    stmt, err := db.Prepare("INSERT INTO users(name, age) VALUES(?, ?)")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer stmt.Close()

    // Insert data into the table
    _, err = stmt.Exec("Alice", 28)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Data inserted successfully")
}





/*
Example 4: Query Data from SQLite Table

Description:
This example demonstrates how to query data from a SQLite table using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in Example 2).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Query data from the table
    rows, err := db.Query("SELECT id, name, age FROM users")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var id, age int
        var name string
        err := rows.Scan(&id, &name, &age)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
    }

    fmt.Println("Query executed successfully")
}





/*
Example 5: Update Data in SQLite Table

Description:
This example demonstrates how to update data in a SQLite table using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in Example 2).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Prepare statement for updating data
    stmt, err := db.Prepare("UPDATE users SET age=? WHERE name=?")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer stmt.Close()

    // Update data in the table
    _, err = stmt.Exec(30, "Alice")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Data updated successfully")
}





/*
Example 6: Delete Data from SQLite Table

Description:
This example demonstrates how to delete data from a SQLite table using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in Example 2).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Prepare statement for deleting data
    stmt, err := db.Prepare("DELETE FROM users WHERE name=?")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer stmt.Close()

    // Delete data from the table
    _, err = stmt.Exec("Alice")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Data deleted successfully")
}





/*
Example 7: Transactions in SQLite

Description:
This example demonstrates how to use transactions in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in Example 2).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Begin transaction
    tx, err := db.Begin()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Perform operations within the transaction
    _, err = tx.Exec("INSERT INTO users(name, age) VALUES(?, ?)", "Bob", 32)
    if err != nil {
        tx.Rollback()
        fmt.Println(err)
        return
    }

    _, err = tx.Exec("DELETE FROM users WHERE name=?", "Alice")
    if err != nil {
        tx.Rollback()
        fmt.Println(err)
        return
    }

    // Commit the transaction
    err = tx.Commit()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Transaction committed successfully")
}





/*
Example 8: Prepared Statements in SQLite

Description:
This example demonstrates how to use prepared statements in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in Example 2).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Prepare statement for querying data
    stmt, err := db.Prepare("SELECT id, name, age FROM users WHERE age > ?")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer stmt.Close()

    // Query data using prepared statement
    rows, err := stmt.Query(30)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var id, age int
        var name string
        err := rows.Scan(&id, &name, &age)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
    }

    fmt.Println("Query executed successfully")
}





/*
Example 9: Error Handling in SQLite Operations

Description:
This example demonstrates error handling best practices when working with SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in Example 2).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

    // Example of a query with error handling
    rows, err := db.Query("SELECT * FROM non_existing_table")
    if err != nil {
        fmt.Println("Error executing query:", err)
        return
    }
    defer rows.Close()

    fmt.Println("Query executed successfully")
}





/*
Example 10: Working with NULL Values in SQLite

Description:
This example demonstrates how to handle NULL values in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in Example 2).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Inserting NULL values into the table
    _, err = db.Exec("INSERT INTO users(name, age) VALUES(?, ?)", "Charlie", nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Data with NULL value inserted successfully")
}





/*
Example 11: Check if Table Exists in SQLite Database

Description:
This example demonstrates how to check if a table exists in a SQLite database using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db'.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func tableExists(db *sql.DB, tableName string) (bool, error) {
    var exists bool
    query := "SELECT name FROM sqlite_master WHERE type='table' AND name=?"
    err := db.QueryRow(query, tableName).Scan(&exists)
    if err != nil && err != sql.ErrNoRows {
        return false, err
    }
    return exists, nil
}

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Check if table 'users' exists
    exists, err := tableExists(db, "users")
    if err != nil {
        fmt.Println(err)
        return
    }

    if exists {
        fmt.Println("Table 'users' exists")
    } else {
        fmt.Println("Table 'users' does not exist")
    }
}





/*
Example 12: Drop Table in SQLite Database

Description:
This example demonstrates how to drop (delete) a table from a SQLite database using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Drop (delete) the table 'users'
    _, err = db.Exec("DROP TABLE IF EXISTS users")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Table 'users' dropped successfully")
}





/*
Example 13: Use Transactions for Bulk Insertion

Description:
This example demonstrates how to use transactions for bulk insertion into a SQLite database using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Begin transaction
    tx, err := db.Begin()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Prepare statement for bulk insertion
    stmt, err := tx.Prepare("INSERT INTO users(name, age) VALUES(?, ?)")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer stmt.Close()

    // Example data for bulk insertion
    users := []struct {
        Name string
        Age  int
    }{
        {"David", 35},
        {"Emma", 27},
        {"Frank", 40},
    }

    // Insert data using transaction and prepared statement
    for _, user := range users {
        _, err = stmt.Exec(user.Name, user.Age)
        if err != nil {
            tx.Rollback()
            fmt.Println(err)
            return
        }
    }

    // Commit the transaction
    err = tx.Commit()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Bulk insertion completed successfully")
}





/*
Example 14: Retrieve Last Insert ID in SQLite

Description:
This example demonstrates how to retrieve the last inserted row ID in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Insert a row into the table 'users'
    result, err := db.Exec("INSERT INTO users(name, age) VALUES(?, ?)", "Grace", 33)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Retrieve the last inserted ID
    lastID, err := result.LastInsertId()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Last Inserted ID: %d\n", lastID)
}





/*
Example 15: Perform Aggregate Function in SQLite

Description:
This example demonstrates how to perform an aggregate function (COUNT) in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Perform COUNT aggregate function
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Total number of users: %d\n", count)
}





/*
Example 16: Handle Foreign Key Constraints in SQLite

Description:
This example demonstrates how to enable and handle foreign key constraints in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with tables 'users' and 'orders' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database with foreign key constraints enabled
    db, err := sql.Open("sqlite3", "./example.db?_foreign_keys=1")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Create tables with foreign key constraints
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL
        )
    `)
    if err != nil {
        fmt.Println(err)
        return
    }

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS orders (
            id INTEGER PRIMARY KEY,
            user_id INTEGER,
            amount INTEGER,
            FOREIGN KEY(user_id) REFERENCES users(id)
        )
    `)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Tables 'users' and 'orders' created with foreign key constraints successfully")
}





/*
Example 17: Use Named Parameters in SQLite Queries

Description:
This example demonstrates how to use named parameters in SQLite queries using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Insert a row into the table 'users' using named parameters
    stmt, err := db.Prepare("INSERT INTO users(name, age) VALUES(:name, :age)")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer stmt.Close()

    // Example data for insertion
    user := map[string]interface{}{
        "name": "Sophia",
        "age":  30,
    }

    // Execute the named parameter query
    _, err = stmt.Exec(user)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Data inserted successfully using named parameters")
}





/*
Example 18: Handle SQLite NULL Values with Scan

Description:
This example demonstrates how to handle SQLite NULL values when scanning rows using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID   int
    Name string
    Age  sql.NullInt64
}

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Query data from the table 'users'
    rows, err := db.Query("SELECT id, name, age FROM users")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Age)
        if err != nil {
            fmt.Println(err)
            return
        }
        // Check for NULL value
        if user.Age.Valid {
            fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age.Int64)
        } else {
            fmt.Printf("ID: %d, Name: %s, Age: NULL\n", user.ID, user.Name)
        }
    }

    fmt.Println("Query executed successfully")
}





/*
Example 19: Use Context with SQLite Operations

Description:
This example demonstrates how to use context with SQLite operations using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

Note:
Using context is important for managing timeouts and cancellations in concurrent environments.

*/

package main

import (
    "context"
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Create a context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Connect to the SQLite database with context
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example query using context
    var name string
    err = db.QueryRowContext(ctx, "SELECT name FROM users WHERE id=?", 1).Scan(&name)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Name retrieved using context: %s\n", name)
}





/*
Example 20: Handle SQLite Errors Gracefully

Description:
This example demonstrates how to handle SQLite errors gracefully using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db'.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./non_existing_db.db")
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

    // Example of a query with error handling
    rows, err := db.Query("SELECT * FROM non_existing_table")
    if err != nil {
        fmt.Println("Error executing query:", err)
        return
    }
    defer rows.Close()

    fmt.Println("Query executed successfully")
}





/*
Example 21: Use SQLite Functions in Queries

Description:
This example demonstrates how to use SQLite built-in functions in queries using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example of using SQLite built-in function (UPPER) in query
    rows, err := db.Query("SELECT id, UPPER(name) FROM users")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var id int
        var name string
        err := rows.Scan(&id, &name)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("ID: %d, Name (uppercase): %s\n", id, name)
    }

    fmt.Println("Query executed successfully")
}





/*
Example 22: Use SQLite Extensions with go-sqlite3

Description:
This example demonstrates how to use SQLite extensions with the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

Note:
SQLite supports extensions that can be loaded dynamically. This example shows the basic setup for using extensions.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Enable SQLite extension (e.g., JSON1)
    _, err = db.Exec("SELECT load_extension('./json1')")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Use the extension functions in queries
    rows, err := db.Query("SELECT json_object('name', name, 'age', age) FROM users")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var jsonResult string
        err := rows.Scan(&jsonResult)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("JSON Object:", jsonResult)
    }

    fmt.Println("Query executed successfully with extension")
}





/*
Example 23: Use SQLite Virtual Tables with go-sqlite3

Description:
This example demonstrates how to use SQLite virtual tables with the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db'.

Note:
SQLite allows creating custom virtual tables using extensions or modules. This example uses FTS5 (Full-Text Search) as a virtual table.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Create a virtual table using FTS5
    _, err = db.Exec("CREATE VIRTUAL TABLE IF NOT EXISTS documents USING fts5(id, content)")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Insert data into the virtual table
    _, err = db.Exec("INSERT INTO documents(id, content) VALUES(1, 'Example document content')")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Virtual table 'documents' created and data inserted successfully")
}





/*
Example 24: Use SQLite Encryption with go-sqlite3

Description:
This example demonstrates how to use SQLite encryption with the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' that requires encryption.

Note:
SQLite does not natively support encryption. This example shows how to use the SQLite Encryption Extension (SEE) which provides encryption support.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database with encryption
    db, err := sql.Open("sqlite3", "./example.db?_cipher=SEE&key=your_encryption_key")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example query with encrypted database
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var id int
        var name string
        var age int
        err := rows.Scan(&id, &name, &age)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
    }

    fmt.Println("Query executed successfully with encryption")
}





/*
Example 25: Use SQLite Backup API

Description:
This example demonstrates how to use the SQLite Backup API for database backup and restore operations using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with some data.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the source SQLite database
    sourceDB, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer sourceDB.Close()

    // Connect to the destination SQLite database (backup file)
    destDB, err := sql.Open("sqlite3", "./backup.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer destDB.Close()

    // Initialize the backup process
    backup := sourceDB.DB().Backup("main", destDB.DB(), "main")
    if backup == nil {
        fmt.Println("Backup initialization failed")
        return
    }
    defer backup.Close()

    // Perform the backup operation
    _, err = backup.Step(-1)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Database backup completed successfully")
}





/*
Example 26: Use SQLite WAL Mode

Description:
This example demonstrates how to enable and use Write-Ahead Logging (WAL) mode in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db'.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database with WAL mode enabled
    db, err := sql.Open("sqlite3", "./example.db?_wal=journal")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example of querying data with WAL mode
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var id int
        var name string
        var age int
        err := rows.Scan(&id, &name, &age)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
    }

    fmt.Println("Query executed successfully with WAL mode")
}





/*
Example 27: Use SQLite User-defined Functions (UDF)

Description:
This example demonstrates how to define and use user-defined functions (UDFs) in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    "strings"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Define a custom UDF (concatenate strings)
    _, err = db.Exec(`CREATE FUNCTION concat_ws(separator TEXT, strings TEXT) RETURNS TEXT AS
    '
    return strings.join(strings.split(strings, separator), separator)
    '
    LANGUAGE SQL`)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Example query using the custom UDF
    rows, err := db.Query("SELECT concat_ws(',', name, age) AS concatenated FROM users")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var concatenated string
        err := rows.Scan(&concatenated)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Concatenated string:", concatenated)
    }

    fmt.Println("Query executed successfully with custom UDF")
}





/*
Example 28: Use SQLite Indexes

Description:
This example demonstrates how to create and use indexes in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Create index on the 'name' column
    _, err = db.Exec("CREATE INDEX idx_name ON users(name)")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Example query using the index
    rows, err := db.Query("SELECT * FROM users WHERE name=?", "Alice")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var id int
        var name string
        var age int
        err := rows.Scan(&id, &name, &age)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
    }

    fmt.Println("Query executed successfully using index")
}





/*
Example 29: Use SQLite Triggers

Description:
This example demonstrates how to create and use triggers in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Create a trigger that inserts a row into a log table on INSERT into 'users'
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS user_logs (
            id INTEGER PRIMARY KEY,
            user_id INTEGER,
            action TEXT,
            timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        fmt.Println(err)
        return
    }

    _, err = db.Exec(`
        CREATE TRIGGER IF NOT EXISTS user_insert_trigger
        AFTER INSERT ON users
        BEGIN
            INSERT INTO user_logs(user_id, action) VALUES(NEW.id, 'inserted');
        END
    `)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Example: Insert a row into 'users' table
    _, err = db.Exec("INSERT INTO users(name, age) VALUES('Bob', 25)")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Trigger executed successfully")
}





/*
Example 30: Use SQLite Views

Description:
This example demonstrates how to create and use views in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Create a view that selects specific columns from 'users'
    _, err = db.Exec("CREATE VIEW IF NOT EXISTS user_names AS SELECT id, name FROM users")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Example query using the view
    rows, err := db.Query("SELECT * FROM user_names")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var id int
        var name string
        err := rows.Scan(&id, &name)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("ID: %d, Name: %s\n", id, name)
    }

    fmt.Println("Query executed successfully using view")
}





/*
Example 31: Use SQLite Full-Text Search (FTS5)

Description:
This example demonstrates how to use Full-Text Search (FTS5) in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a virtual table 'documents' created using FTS5.

Note:
FTS5 is a SQLite extension that provides an efficient way to search for text in a large collection of documents.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example query using FTS5 full-text search
    rows, err := db.Query("SELECT * FROM documents WHERE content MATCH 'example'")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var id int
        var content string
        err := rows.Scan(&id, &content)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("ID: %d, Content: %s\n", id, content)
    }

    fmt.Println("Query executed successfully using FTS5")
}





/*
Example 32: Use SQLite Transactions

Description:
This example demonstrates how to use transactions in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Begin transaction
    tx, err := db.Begin()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Example: Insert multiple rows in a transaction
    stmt, err := tx.Prepare("INSERT INTO users(name, age) VALUES(?, ?)")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer stmt.Close()

    users := []struct {
        Name string
        Age  int
    }{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }

    for _, user := range users {
        _, err = stmt.Exec(user.Name, user.Age)
        if err != nil {
            fmt.Println(err)
            tx.Rollback() // Rollback the transaction on error
            return
        }
    }

    // Commit transaction
    err = tx.Commit()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Transaction executed successfully")
}





/*
Example 33: Use SQLite Foreign Keys

Description:
This example demonstrates how to use foreign keys in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with tables 'users' and 'orders' (as created in previous examples).

Note:
SQLite supports foreign key constraints that enforce referential integrity between tables.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database with foreign key support enabled
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Enable foreign key support
    _, err = db.Exec("PRAGMA foreign_keys = ON")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Create 'orders' table with foreign key constraint
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS orders (
            id INTEGER PRIMARY KEY,
            user_id INTEGER,
            amount REAL,
            FOREIGN KEY (user_id) REFERENCES users(id)
        )
    `)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Example: Insert order with a user ID that does not exist (should fail due to foreign key constraint)
    _, err = db.Exec("INSERT INTO orders(user_id, amount) VALUES(100, 50.0)")
    if err != nil {
        fmt.Println("Error inserting order:", err)
        return
    }

    fmt.Println("Foreign key constraint enforced successfully")
}





/*
Example 34: Use SQLite Check Constraints

Description:
This example demonstrates how to use check constraints in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

Note:
SQLite allows defining check constraints to enforce data integrity rules on column values.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Create 'users' table with a check constraint
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            name TEXT,
            age INTEGER CHECK (age >= 18) -- Check constraint: age must be >= 18
        )
    `)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Example: Insert a user with an age that violates the check constraint
    _, err = db.Exec("INSERT INTO users(name, age) VALUES('Alice', 17)")
    if err != nil {
        fmt.Println("Error inserting user:", err)
        return
    }

    fmt.Println("Check constraint enforced successfully")
}





/*
Example 35: Use SQLite Savepoints

Description:
This example demonstrates how to use savepoints in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

Note:
Savepoints allow creating named points within transactions to facilitate partial rollback and nested transactions.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Begin transaction
    tx, err := db.Begin()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Example: Insert a user with a savepoint
    _, err = tx.Exec("INSERT INTO users(name, age) VALUES('Alice', 30)")
    if err != nil {
        fmt.Println("Error inserting user:", err)
        tx.RollbackTo("start") // Rollback to savepoint 'start' on error
        return
    }

    // Create a savepoint 'start'
    _, err = tx.Exec("SAVEPOINT start")
    if err != nil {
        fmt.Println("Error creating savepoint:", err)
        tx.Rollback()
        return
    }

    // Example: Insert another user
    _, err = tx.Exec("INSERT INTO users(name, age) VALUES('Bob', 25)")
    if err != nil {
        fmt.Println("Error inserting user:", err)
        tx.RollbackTo("start") // Rollback to savepoint 'start' on error
        return
    }

    // Commit transaction
    err = tx.Commit()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Transaction with savepoint executed successfully")
}





/*
Example 36: Use SQLite Journal Mode

Description:
This example demonstrates how to set and use different journal modes in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db'.

Note:
SQLite supports different journal modes that affect how transactions are handled and data is written to disk.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database with a specific journal mode
    db, err := sql.Open("sqlite3", "./example.db?_journal_mode=WAL")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example: Insert data into 'users' table with WAL journal mode
    _, err = db.Exec("INSERT INTO users(name, age) VALUES('Alice', 30)")
    if err != nil {
        fmt.Println("Error inserting user:", err)
        return
    }

    fmt.Println("Data inserted successfully using WAL journal mode")
}





/*
Example 37: Use SQLite Write-Ahead Logging (WAL)

Description:
This example demonstrates how to enable and use Write-Ahead Logging (WAL) in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db'.

Note:
WAL mode in SQLite improves concurrency and performance by allowing multiple readers and one writer to access the database.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database with WAL mode enabled
    db, err := sql.Open("sqlite3", "./example.db?_journal_mode=WAL")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example: Insert data into 'users' table with WAL mode
    _, err = db.Exec("INSERT INTO users(name, age) VALUES('Bob', 25)")
    if err != nil {
        fmt.Println("Error inserting user:", err)
        return
    }

    fmt.Println("Data inserted successfully using WAL mode")
}





/*
Example 38: Use SQLite Memory Database

Description:
This example demonstrates how to use an in-memory SQLite database using the go-sqlite3 library.

SQLite Database:
No physical file is created for an in-memory database. It exists entirely in RAM.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to an in-memory SQLite database
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example: Create a table in the in-memory database
    _, err = db.Exec(`
        CREATE TABLE users (
            id INTEGER PRIMARY KEY,
            name TEXT,
            age INTEGER
        )
    `)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Example: Insert data into the in-memory database
    _, err = db.Exec("INSERT INTO users(name, age) VALUES('Alice', 30)")
    if err != nil {
        fmt.Println("Error inserting user:", err)
        return
    }

    fmt.Println("Data inserted successfully into in-memory database")
}





/*
Example 39: Use SQLite Vacuum

Description:
This example demonstrates how to use the VACUUM command in SQLite to rebuild the database file and improve performance.

SQLite Database:
Assume you have a SQLite database named 'example.db'.

Note:
The VACUUM command reclaims unused disk space and defragments the database file.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example: Perform VACUUM operation
    _, err = db.Exec("VACUUM")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Database vacuumed successfully")
}





/*
Example 40: Use SQLite Encrypt

Description:
This example demonstrates how to use encryption with SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' that you want to encrypt.

Note:
Encryption with SQLite can be achieved using external tools or libraries that provide encryption extensions.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database (encrypted)
    db, err := sql.Open("sqlite3", "./example.db?_crypt=secret")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example: Insert data into encrypted database
    _, err = db.Exec("INSERT INTO users(name, age) VALUES('Alice', 30)")
    if err != nil {
        fmt.Println("Error inserting user:", err)
        return
    }

    fmt.Println("Data inserted successfully into encrypted database")
}





/*
Example 41: Use SQLite Analyze

Description:
This example demonstrates how to use the ANALYZE command in SQLite to gather statistics for query optimization.

SQLite Database:
Assume you have a SQLite database named 'example.db'.

Note:
The ANALYZE command collects statistics about the distribution of key values in tables, which helps the query planner make better decisions.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example: Perform ANALYZE operation
    _, err = db.Exec("ANALYZE")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("ANALYZE operation completed successfully")
}





/*
Example 42: Use SQLite Online Backup

Description:
This example demonstrates how to perform an online backup of a SQLite database using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db'.

Note:
Online backup in SQLite allows you to make a copy of the database while it is still being accessed and modified.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Open the source SQLite database (to be backed up)
    sourceDB, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer sourceDB.Close()

    // Open the destination SQLite database (backup file)
    destDB, err := sql.Open("sqlite3", "./backup.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer destDB.Close()

    // Begin the online backup process
    _, err = destDB.Exec("ATTACH DATABASE './backup.db' AS backup")
    if err != nil {
        fmt.Println(err)
        return
    }
    _, err = destDB.Exec("SELECT sqlcipher_export('backup')")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Online backup completed successfully")
}





/*
Example 43: Use SQLite Export CSV

Description:
This example demonstrates how to export data from a SQLite table to a CSV file using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "encoding/csv"
    "fmt"
    "os"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Query the data to export
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Create a CSV file
    file, err := os.Create("users.csv")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Create a CSV writer
    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Iterate over the rows and write to CSV
    for rows.Next() {
        var id int
        var name string
        var age int
        err := rows.Scan(&id, &name, &age)
        if err != nil {
            fmt.Println(err)
            return
        }

        err = writer.Write([]string{fmt.Sprintf("%d", id), name, fmt.Sprintf("%d", age)})
        if err != nil {
            fmt.Println(err)
            return
        }
    }

    fmt.Println("Data exported to CSV successfully")
}





/*
Example 44: Use SQLite Import CSV

Description:
This example demonstrates how to import data from a CSV file into a SQLite table using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' where you want to import data into a table 'imported_users'.

Note:
Ensure the CSV file ('users.csv') exists and contains valid data.

*/

package main

import (
    "database/sql"
    "encoding/csv"
    "fmt"
    "os"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Create a table to import CSV data
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS imported_users (
            id INTEGER PRIMARY KEY,
            name TEXT,
            age INTEGER
        )
    `)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Open the CSV file
    file, err := os.Open("users.csv")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Create a CSV reader
    reader := csv.NewReader(file)

    // Read and insert each record into the database
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println(err)
        return
    }

    tx, err := db.Begin()
    if err != nil {
        fmt.Println(err)
        return
    }

    stmt, err := tx.Prepare("INSERT INTO imported_users(name, age) VALUES(?, ?)")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer stmt.Close()

    for _, record := range records {
        _, err = stmt.Exec(record[1], record[2]) // Assuming the CSV has no header row
        if err != nil {
            fmt.Println(err)
            tx.Rollback()
            return
        }
    }

    // Commit the transaction
    err = tx.Commit()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Data imported from CSV successfully")
}





/*
Example 45: Use SQLite JSON Extension

Description:
This example demonstrates how to use the JSON1 extension in SQLite for JSON processing using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'json_data' (as created in previous examples).

Note:
The JSON1 extension provides functions and operators for working with JSON data stored in SQLite columns.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example: Insert JSON data into the 'json_data' table
    _, err = db.Exec(`
        INSERT INTO json_data(data)
        VALUES('{"name": "Alice", "age": 30}')
    `)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Example: Query JSON data using JSON1 functions
    rows, err := db.Query("SELECT json_extract(data, '$.name') FROM json_data")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var name string
        err := rows.Scan(&name)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Name:", name)
    }

    fmt.Println("JSON data processed successfully using JSON1 extension")
}





/*
Example 46: Use SQLite Math Extension

Description:
This example demonstrates how to use the math extension in SQLite for mathematical computations using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'numbers' (as created in previous examples).

Note:
The math extension provides mathematical functions such as sin, cos, sqrt, etc., which can be used in SQL queries.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example: Insert numerical data into the 'numbers' table
    _, err = db.Exec(`
        INSERT INTO numbers(value)
        VALUES(25)
    `)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Example: Query mathematical computation using math extension functions
    rows, err := db.Query("SELECT sqrt(value) FROM numbers")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var sqrtValue float64
        err := rows.Scan(&sqrtValue)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Square root:", sqrtValue)
    }

    fmt.Println("Mathematical computation performed successfully using math extension")
}





/*
Example 47: Use SQLite Full-Text Search Extension

Description:
This example demonstrates how to use the full-text search (FTS5) extension in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a virtual table 'documents' created using FTS5.

Note:
FTS5 is a SQLite extension that provides an efficient way to search for text in a large collection of documents.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Example query using FTS5 full-text search
    rows, err := db.Query("SELECT * FROM documents WHERE content MATCH 'example'")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var id int
        var content string
        err := rows.Scan(&id, &content)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("ID: %d, Content: %s\n", id, content)
    }

    fmt.Println("Query executed successfully using FTS5")
}





/*
Example 48: Use SQLite User-Defined Functions (UDFs)

Description:
This example demonstrates how to define and use user-defined functions (UDFs) in SQLite using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db'.

Note:
UDFs allow extending SQLite with custom functions written in Go or registered external functions.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

// Example UDF: Custom function to calculate the square of a number
func squareFunction(context *sql.Context, value float64) (float64, error) {
    return value * value, nil
}

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Register the UDF (User-Defined Function)
    err = db.RegisterFunc("square", squareFunction, true)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Example: Use the square UDF in a SQL query
    rows, err := db.Query("SELECT square(5.0)")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Read the result of the query
    for rows.Next() {
        var result float64
        err := rows.Scan(&result)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Square result:", result)
    }

    fmt.Println("UDF executed successfully")
}





/*
Example 49: Use SQLite Backup API

Description:
This example demonstrates how to perform a SQLite database backup using the SQLite Backup API with the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db'.

Note:
The SQLite Backup API provides a way to create a backup of a database file, even while it is being used.

*/

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Open the source SQLite database (to be backed up)
    sourceDB, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer sourceDB.Close()

    // Open the destination SQLite database (backup file)
    destDB, err := sql.Open("sqlite3", "./backup.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer destDB.Close()

    // Initialize the SQLite Backup object
    backup := sqlite3.BackupInit(destDB, "main", sourceDB, "main")
    if backup == nil {
        fmt.Println("Backup initialization failed")
        return
    }

    // Perform the backup operation
    _, err = backup.Step(-1)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Finish the backup process
    err = backup.Finish()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("SQLite database backup completed successfully")
}





/*
Example 50: Use SQLite Export Table Schema

Description:
This example demonstrates how to export the schema of a SQLite table to a text file using the go-sqlite3 library.

SQLite Database:
Assume you have a SQLite database named 'example.db' with a table 'users' (as created in previous examples).

*/

package main

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Query to fetch table schema
    rows, err := db.Query("SELECT sql FROM sqlite_master WHERE type='table' AND name='users'")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Read the schema definition
    var schema string
    if rows.Next() {
        err := rows.Scan(&schema)
        if err != nil {
            fmt.Println(err)
            return
        }
    } else {
        fmt.Println("Table not found")
        return
    }

    // Write schema to a file
    file, err := os.Create("users_schema.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    _, err = file.WriteString(schema)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Table schema exported successfully")
}
