


// Connect to a SQLite database
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}



// Define a User model
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}



// Create a record in the database
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	user := User{Name: "John"}
	db.Create(&user)
}



// Read a record from the database
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var user User
	db.First(&user, 1) // find user with integer primary key
	fmt.Println(user)
}



// Update a record in the database
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var user User
	db.First(&user, 1) // find user with integer primary key
	user.Name = "Jane"
	db.Save(&user)
}



// Delete a record from the database
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var user User
	db.First(&user, 1) // find user with integer primary key
	db.Delete(&user)
}



// Batch insert multiple records into the database
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	users := []User{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Charlie"},
	}
	db.Create(&users)
}




// Find records with conditions from the database
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var users []User
	db.Where("name = ?", "Alice").Find(&users)
	fmt.Println(users)
}



// Find records with multiple conditions from the database
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var users []User
	db.Where("name = ? AND id = ?", "Alice", 1).Find(&users)
	fmt.Println(users)
}



// Find records with IN condition from the database
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var users []User
	db.Where("name IN ?", []string{"Alice", "Bob"}).Find(&users)
	fmt.Println(users)
}




// Find records with LIKE condition from the database
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var users []User
	db.Where("name LIKE ?", "%li%").Find(&users)
	fmt.Println(users)
}





// Find records with OR condition from the database
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var users []User
	db.Where("name = ? OR name = ?", "Alice", "Bob").Find(&users)
	fmt.Println(users)
}



// Order records by a field from the database
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var users []User
	db.Order("name desc").Find(&users)
	fmt.Println(users)
}



// Limit the number of records fetched from the database
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var users []User
	db.Limit(2).Find(&users)
	fmt.Println(users)
}



// Offset the records fetched from the database
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var users []User
	db.Offset(1).Find(&users)
	fmt.Println(users)
}



// Count the number of records in the database
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var count int64
	db.Model(&User{}).Count(&count)
	fmt.Println("Total users:", count)
}





// Execute raw SQL queries
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var users []User
	db.Raw("SELECT * FROM users WHERE name = ?", "Alice").Scan(&users)
	fmt.Println(users)
}





// Use transactions to ensure atomicity
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	tx := db.Begin()
	if err := tx.Error; err != nil {
		panic(err)
	}

	user := User{Name: "Alice"}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}
}



// Eager load related models
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID uint
	Email  string
}

type User struct {
	gorm.Model
	Name    string
	Profile Profile
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Profile{})

	var user User
	db.Preload("Profile").First(&user, 1)
	fmt.Println(user)
}




// Lazy load related models
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID uint
	Email  string
}

type User struct {
	gorm.Model
	Name    string
	Profile Profile `gorm:"foreignKey:UserID"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Profile{})

	var user User
	db.First(&user, 1)
	db.Model(&user).Association("Profile").Find(&user.Profile)
	fmt.Println(user)
}




// Define one-to-many relationship
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title  string
	UserID uint
}

type User struct {
	gorm.Model
	Name  string
	Posts []Post
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Post{})

	user := User{Name: "Alice", Posts: []Post{{Title: "First Post"}, {Title: "Second Post"}}}
	db.Create(&user)

	var users []User
	db.Preload("Posts").Find(&users)
	fmt.Println(users)
}




// Define many-to-many relationship
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:user_roles;"`
}

type User struct {
	gorm.Model
	Name  string
	Roles []Role `gorm:"many2many:user_roles;"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Role{})

	role := Role{Name: "Admin"}
	user := User{Name: "Alice", Roles: []Role{role}}
	db.Create(&user)

	var users []User
	db.Preload("Roles").Find(&users)
	fmt.Println(users)
}



// Create table with unique constraints
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Email string `gorm:"unique"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}



// Add index to a column
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"index"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}



// Add composite index to multiple columns
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"index:idx_name"`
	LastName  string `gorm:"index:idx_name"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}




// Add a unique index to a column
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"uniqueIndex"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}




// Add multiple unique indexes to a model
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"uniqueIndex:idx_name"`
	LastName  string `gorm:"uniqueIndex:idx_name"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}



// Add foreign key constraint to a column
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID uint
	Email  string
}

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Profile{})
}




// Add NOT NULL constraint to a column
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}


// Add default value to a column
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Age   int `gorm:"default:18"`
	Email string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}


// Implement soft delete for a model
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	user := User{Name: "Alice"}
	db.Create(&user)

	db.Delete(&user)
}



// Find records including soft deleted ones
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var users []User
	db.Unscoped().Find(&users)
	fmt.Println(users)
}



// Paginate results using limit and offset
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var users []User
	db.Limit(2).Offset(2).Find(&users)
	fmt.Println(users)
}



// Update specific fields of a record
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var user User
	db.First(&user, 1)
	db.Model(&user).Update("Age", 25)
}



// Update multiple fields of a record
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var user User
	db.First(&user, 1)
	db.Model(&user).Updates(User{Name: "John", Age: 30})
}



// Find the first record matching conditions
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var user User
	db.Where("name = ?", "Alice").First(&user)
	fmt.Println(user)
}



// Find the last record matching conditions
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var user User
	db.Where("name = ?", "Alice").Last(&user)
	fmt.Println(user)
}




// Find records and preload associations
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID uint
	Email  string
}

type User struct {
	gorm.Model
	Name    string
	Profile Profile
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Profile{})

	var users []User
	db.Preload("Profile").Find(&users)
	fmt.Println(users)
}



// Find records and preload nested associations
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID uint
	Email  string
}

type Post struct {
	gorm.Model
	Title  string
	UserID uint
}

type User struct {
	gorm.Model
	Name    string
	Profile Profile
	Posts   []Post
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Profile{}, &Post{})

	var users []User
	db.Preload("Profile").Preload("Posts").Find(&users)
	fmt.Println(users)
}



// Scan query results into a struct
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

type Result struct {
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var result Result
	db.Table("users").Select("name, age").Where("name = ?", "Alice").Scan(&result)
	fmt.Println(result)
}





// Scan query results into a slice of structs
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

type Result struct {
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	var results []Result
	db.Table("users").Select("name, age").Where("age > ?", 20).Scan(&results)
	fmt.Println(results)
}





// Use raw SQL to insert records
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 25)
}




// Use raw SQL to update records
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	db.Exec("UPDATE users SET age = ? WHERE name = ?", 26, "Alice")
}




// Use raw SQL to delete records
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	db.Exec("DELETE FROM users WHERE name = ?", "Alice")
}




// Use Gorm hooks for actions before/after create/update/delete
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before creating:", u.Name)
	return
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	user := User{Name: "Alice", Age: 25}
	db.Create(&user)
}




// Define a custom primary key for a model
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   string `gorm:"primaryKey"`
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}





// Define a composite primary key for a model
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	FirstName string `gorm:"primaryKey"`
	LastName  string `gorm:"primaryKey"`
	Age       int
}

func main() {
	db, err := gorm.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  db.AutoMigrate(&User{})
}



// Define a custom table name for a model
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func (User) TableName() string {
	return "my_users"
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}




// Define a custom column name for a field in a model
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"column:given_name"`
	LastName  string `gorm:"column:surname"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}




// Customize the SQL data type for a field in a model
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}






