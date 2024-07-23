// Example 1: Creating a new constraint
//
// This example demonstrates how to create a new constraint using the constraints library.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    fmt.Println(c)
}





// Example 2: Adding a validation rule to a constraint
//
// This example demonstrates how to add a validation rule to a constraint using the constraints library.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 10 {
            return fmt.Errorf("value must be greater than or equal to 10")
        }
        return nil
    })
    fmt.Println(c.Validate(5))
}





// Example 3: Validating a value against a constraint
//
// This example demonstrates how to validate a value against a constraint using the constraints library.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 10 {
            return fmt.Errorf("value must be greater than or equal to 10")
        }
        return nil
    })
    fmt.Println(c.Validate(15))
}





// Example 4: Adding multiple validation rules
//
// This example demonstrates how to add multiple validation rules to a constraint.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 10 {
            return fmt.Errorf("value must be greater than or equal to 10")
        }
        return nil
    })
    c.AddRule(func(value interface{}) error {
        if value.(int) > 20 {
            return fmt.Errorf("value must be less than or equal to 20")
        }
        return nil
    })
    fmt.Println(c.Validate(25))
}





// Example 5: Custom error message
//
// This example demonstrates how to set a custom error message for a validation rule.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 10 {
            return fmt.Errorf("custom error: value must be >= 10")
        }
        return nil
    })
    fmt.Println(c.Validate(5))
}





// Example 6: Chaining constraints
//
// This example demonstrates how to chain multiple constraints together.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c1 := constraints.NewConstraint()
    c1.AddRule(func(value interface{}) error {
        if value.(int) < 10 {
            return fmt.Errorf("value must be greater than or equal to 10")
        }
        return nil
    })

    c2 := constraints.NewConstraint()
    c2.AddRule(func(value interface{}) error {
        if value.(int) > 20 {
            return fmt.Errorf("value must be less than or equal to 20")
        }
        return nil
    })

    c3 := constraints.NewCompositeConstraint(c1, c2)
    fmt.Println(c3.Validate(15))
}





// Example 7: Using string validation
//
// This example demonstrates how to use string validation in constraints.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if len(value.(string)) < 5 {
            return fmt.Errorf("string length must be at least 5 characters")
        }
        return nil
    })
    fmt.Println(c.Validate("abc"))
}





// Example 8: Numeric range validation
//
// This example demonstrates how to validate a numeric range.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 10 || value.(int) > 20 {
            return fmt.Errorf("value must be between 10 and 20")
        }
        return nil
    })
    fmt.Println(c.Validate(15))
}





// Example 9: Regular expression validation
//
// This example demonstrates how to use regular expression validation.
//
// Code:
package main

import (
    "fmt"
    "regexp"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        matched, _ := regexp.MatchString(`^[a-z]+$`, value.(string))
        if !matched {
            return fmt.Errorf("value must match regex ^[a-z]+$")
        }
        return nil
    })
    fmt.Println(c.Validate("abc123"))
}





// Example 10: Validating arrays
//
// This example demonstrates how to validate elements within an array.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        for _, v := range value.([]int) {
            if v < 0 {
                return fmt.Errorf("all elements must be non-negative")
            }
        }
        return nil
    })
    fmt.Println(c.Validate([]int{1, 2, -3}))
}





// Example 11: Validating maps
//
// This example demonstrates how to validate values within a map.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        for _, v := range value.(map[string]int) {
            if v < 0 {
                return fmt.Errorf("all map values must be non-negative")
            }
        }
        return nil
    })
    fmt.Println(c.Validate(map[string]int{"a": 1, "b": -2}))
}





// Example 12: Nested constraints
//
// This example demonstrates how to use nested constraints.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c1 := constraints.NewConstraint()
    c1.AddRule(func(value interface{}) error {
        if value.(int) < 0 {
            return fmt.Errorf("value must be non-negative")
        }
        return nil
    })

    c2 := constraints.NewConstraint()
    c2.AddRule(func(value interface{}) error {
        for _, v := range value.([]int) {
            if err := c1.Validate(v); err != nil {
                return fmt.Errorf("nested validation failed: %v", err)
            }
        }
        return nil
    })

    fmt.Println(c2.Validate([]int{1, -2, 3}))
}





// Example 13: Constraint inheritance
//
// This example demonstrates how to use constraint inheritance.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type CustomConstraint struct {
    *constraints.Constraint
}

func NewCustomConstraint() *CustomConstraint {
    return &CustomConstraint{constraints.NewConstraint()}
}

func main() {
    c := NewCustomConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 10 {
            return fmt.Errorf("value must be >= 10")
        }
        return nil
    })
    fmt.Println(c.Validate(5))
}






// Example 14: Custom constraint types
//
// This example demonstrates how to create custom constraint types.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type PositiveConstraint struct {
    *constraints.Constraint
}

func NewPositiveConstraint() *PositiveConstraint {
    pc := &PositiveConstraint{constraints.NewConstraint()}
    pc.AddRule(func(value interface{}) error {
        if value.(int) <= 0 {
            return fmt.Errorf("value must be positive")
        }
        return nil
    })
    return pc
}

func main() {
    c := NewPositiveConstraint()
    fmt.Println(c.Validate(-1))
}





// Example 15: Combining constraints
//
// This example demonstrates how to combine multiple constraints.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    isNonNegative := constraints.NewConstraint()
    isNonNegative.AddRule(func(value interface{}) error {
        if value.(int) < 0 {
            return fmt.Errorf("value must be non-negative")
        }
        return nil
    })

    isEven := constraints.NewConstraint()
    isEven.AddRule(func(value interface{}) error {
        if value.(int)%2 != 0 {
            return fmt.Errorf("value must be even")
        }
        return nil
    })

    combined := constraints.NewCompositeConstraint(isNonNegative, isEven)
    fmt.Println(combined.Validate(4))
}





// Example 16: Constraints with custom structs
//
// This example demonstrates how to use constraints with custom structs.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type User struct {
    Age int
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        user := value.(User)
        if user.Age < 18 {
            return fmt.Errorf("user must be at least 18 years old")
        }
        return nil
    })
    fmt.Println(c.Validate(User{Age: 15}))
}





// Example 17: Validating nested structs
//
// This example demonstrates how to validate nested structs.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type Address struct {
    City string
}

type User struct {
    Name    string
    Address Address
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        user := value.(User)
        if user.Address.City == "" {
            return fmt.Errorf("city must not be empty")
        }
        return nil
    })
    fmt.Println(c.Validate(User{Name: "John", Address: Address{City: ""}}))
}





// Example 18: Conditional validation
//
// This example demonstrates how to use conditional validation.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type User struct {
    Age   int
    Email string
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        user := value.(User)
        if user.Age >= 18 && user.Email == "" {
            return fmt.Errorf("email is required for users 18 or older")
        }
        return nil
    })
    fmt.Println(c.Validate(User{Age: 20}))
}





// Example 19: Using interface types
//
// This example demonstrates how to use interface types with constraints.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type Validatable interface {
    Validate() error
}

type Product struct {
    Price int
}

func (p Product) Validate() error {
    if p.Price <= 0 {
        return fmt.Errorf("price must be positive")
    }
    return nil
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        return value.(Validatable).Validate()
    })
    fmt.Println(c.Validate(Product{Price: -1}))
}





// Example 20: Validating with custom functions
//
// This example demonstrates how to validate using custom functions.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func isValidEmail(email string) bool {
    return len(email) > 5
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if !isValidEmail(value.(string)) {
            return fmt.Errorf("invalid email address")
        }
        return nil
    })
    fmt.Println(c.Validate("a@b.com"))
}





// Example 21: Handling validation errors
//
// This example demonstrates how to handle validation errors.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 0 {
            return fmt.Errorf("value must be non-negative")
        }
        return nil
    })
    err := c.Validate(-5)
    if err != nil {
        fmt.Printf("Validation failed: %v\n", err)
    }
}






// Example 22: Asynchronous validation
//
// This example demonstrates how to perform asynchronous validation.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
    "sync"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 0 {
            return fmt.Errorf("value must be non-negative")
        }
        return nil
    })

    var wg sync.WaitGroup
    values := []int{-1, 0, 1}
    for _, v := range values {
        wg.Add(1)
        go func(val int) {
            defer wg.Done()
            fmt.Println(c.Validate(val))
        }(v)
    }
    wg.Wait()
}





// Example 23: Combining synchronous and asynchronous validation
//
// This example demonstrates how to combine synchronous and asynchronous validation.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
    "sync"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 0 {
            return fmt.Errorf("value must be non-negative")
        }
        return nil
    })

    var wg sync.WaitGroup
    values := []int{-1, 0, 1}
    results := make(chan error, len(values))
    for _, v := range values {
        wg.Add(1)
        go func(val int) {
            defer wg.Done()
            results <- c.Validate(val)
        }(v)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    for res := range results {
        fmt.Println(res)
    }
}





// Example 23: Combining synchronous and asynchronous validation
//
// This example demonstrates how to combine synchronous and asynchronous validation.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
    "sync"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 0 {
            return fmt.Errorf("value must be non-negative")
        }
        return nil
    })

    var wg sync.WaitGroup
    values := []int{-1, 0, 1}
    results := make(chan error, len(values))
    for _, v := range values {
        wg.Add(1)
        go func(val int) {
            defer wg.Done()
            results <- c.Validate(val)
        }(v)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    for res := range results {
        fmt.Println(res)
    }
}





// Example 24: Dynamic constraints
//
// This example demonstrates how to use dynamic constraints based on input.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    minAge := 18
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < minAge {
            return fmt.Errorf("age must be at least %d", minAge)
        }
        return nil
    })
    fmt.Println(c.Validate(17))
}





// Example 25: Validating custom data types
//
// This example demonstrates how to validate custom data types.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type Product struct {
    Name  string
    Price float64
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        product := value.(Product)
        if product.Price < 0 {
            return fmt.Errorf("price must be positive")
        }
        return nil
    })
    fmt.Println(c.Validate(Product{Name: "Widget", Price: -9.99}))
}





// Example 26: Using constraints with interfaces
//
// This example demonstrates how to use constraints with interfaces.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(Shape).Area() <= 0 {
            return fmt.Errorf("area must be positive")
        }
        return nil
    })
    fmt.Println(c.Validate(Circle{Radius: -5}))
}





// Example 27: Using dependency injection
//
// This example demonstrates how to use dependency injection with constraints.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type Validator struct {
    Constraint *constraints.Constraint
}

func NewValidator(c *constraints.Constraint) *Validator {
    return &Validator{Constraint: c}
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) <= 0 {
            return fmt.Errorf("value must be positive")
        }
        return nil
    })
    validator := NewValidator(c)
    fmt.Println(validator.Constraint.Validate(-1))
}





// Example 28: Reusing constraints
//
// This example demonstrates how to reuse constraints across different contexts.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    nonNegative := constraints.NewConstraint()
    nonNegative.AddRule(func(value interface{}) error {
        if value.(int) < 0 {
            return fmt.Errorf("value must be non-negative")
        }
        return nil
    })

    fmt.Println(nonNegative.Validate(-1))
    fmt.Println(nonNegative.Validate(10))
}





// Example 29: Using constraints with HTTP handlers
//
// This example demonstrates how to use constraints with HTTP handlers.
//
// Code:
package main

import (
    "fmt"
    "net/http"
    "github.com/go-playground/constraints"
)

func validateUser(c *constraints.Constraint, age int) error {
    return c.Validate(age)
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 18 {
            return fmt.Errorf("user must be at least 18 years old")
        }
        return nil
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        age := 17
        err := validateUser(c, age)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
        } else {
            fmt.Fprintf(w, "Validation passed")
        }
    })
    http.ListenAndServe(":8080", nil)
}





// Example 30: Handling multiple validation errors
//
// This example demonstrates how to handle multiple validation errors.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 0 {
            return fmt.Errorf("value must be non-negative")
        }
        return nil
    })
    c.AddRule(func(value interface{}) error {
        if value.(int)%2 != 0 {
            return fmt.Errorf("value must be even")
        }
        return nil
    })
    fmt.Println(c.Validate(-3))
}





// Example 31: Logging validation process
//
// This example demonstrates how to log the validation process.
//
// Code:
package main

import (
    "fmt"
    "log"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        log.Printf("Validating value: %v", value)
        if value.(int) <= 0 {
            return fmt.Errorf("value must be positive")
        }
        return nil
    })
    fmt.Println(c.Validate(-1))
}





// Example 32: Extending constraints library
//
// This example demonstrates how to extend the constraints library with custom logic.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type CustomValidator struct {
    *constraints.Constraint
}

func NewCustomValidator() *CustomValidator {
    return &CustomValidator{constraints.NewConstraint()}
}

func (cv *CustomValidator) AddCustomRule(f func(interface{}) error) {
    cv.AddRule(f)
}

func main() {
    cv := NewCustomValidator()
    cv.AddCustomRule(func(value interface{}) error {
        if value.(int) != 42 {
            return fmt.Errorf("value must be 42")
        }
        return nil
    })
    fmt.Println(cv.Validate(100))
}





// Example 33: Constraints with database records
//
// This example demonstrates how to use constraints to validate database records.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type User struct {
    Name string
    Age  int
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        user := value.(User)
        if user.Age < 0 {
            return fmt.Errorf("age must be non-negative")
        }
        return nil
    })

    users := []User{{Name: "Alice", Age: 30}, {Name: "Bob", Age: -5}}
    for _, user := range users {
        err := c.Validate(user)
        if err != nil {
            fmt.Printf("Validation failed for user %s: %v\n", user.Name, err)
        }
    }
}





// Example 34: Dynamic rule based on context
//
// This example demonstrates how to use a dynamic rule based on context.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    maxAge := 65
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) > maxAge {
            return fmt.Errorf("age must be less than or equal to %d", maxAge)
        }
        return nil
    })
    fmt.Println(c.Validate(70))
}





// Example 35: Using constraints in a middleware
//
// This example demonstrates how to use constraints in middleware.
//
// Code:
package main

import (
    "fmt"
    "net/http"
    "github.com/go-playground/constraints"
)

func validateMiddleware(c *constraints.Constraint, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        age := 17
        if err := c.Validate(age); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 18 {
            return fmt.Errorf("user must be at least 18 years old")
        }
        return nil
    })

    http.Handle("/", validateMiddleware(c, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, world!")
    })))
    http.ListenAndServe(":8080", nil)
}





// Example 36: Implementing validation for complex types
//
// This example demonstrates how to implement validation for complex types.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type Address struct {
    City string
}

type User struct {
    Name    string
    Address Address
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        user := value.(User)
        if user.Name == "" {
            return fmt.Errorf("name must not be empty")
        }
        return nil
    })
    fmt.Println(c.Validate(User{Name: "", Address: Address{City: "New York"}}))
}





// Example 37: Validation with optional fields
//
// This example demonstrates how to perform validation with optional fields.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type User struct {
    Name  string
    Email string
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        user := value.(User)
        if user.Name == "" {
            return fmt.Errorf("name must not be empty")
        }
        return nil
    })
    fmt.Println(c.Validate(User{Email: "test@example.com"}))
}





// Example 38: Validation with regular expressions
//
// This example demonstrates how to validate fields using regular expressions.
//
// Code:
package main

import (
    "fmt"
    "regexp"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        email := value.(string)
        if matched, _ := regexp.MatchString(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`, email); !matched {
            return fmt.Errorf("invalid email format")
        }
        return nil
    })
    fmt.Println(c.Validate("invalid-email"))
}





// Example 39: Validating with custom error messages
//
// This example demonstrates how to use custom error messages for validation.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) <= 0 {
            return fmt.Errorf("value must be greater than zero")
        }
        return nil
    })
    fmt.Println(c.Validate(-1))
}





// Example 40: Validating custom collections
//
// This example demonstrates how to validate custom collections.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type Collection []int

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        coll := value.(Collection)
        for _, v := range coll {
            if v <= 0 {
                return fmt.Errorf("collection must contain only positive numbers")
            }
        }
        return nil
    })
    fmt.Println(c.Validate(Collection{-1, 2, 3}))
}





// Example 41: Validation with custom tags
//
// This example demonstrates how to perform validation with custom tags.
//
// Code:
package main

import (
    "fmt"
    "reflect"
    "github.com/go-playground/constraints"
)

func main() {
    type User struct {
        Name string `validate:"required"`
    }

    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        v := reflect.ValueOf(value)
        t := v.Type()
        for i := 0; i < t.NumField(); i++ {
            field := t.Field(i)
            tag := field.Tag.Get("validate")
            if tag == "required" && v.Field(i).String() == "" {
                return fmt.Errorf("%s is required", field.Name)
            }
        }
        return nil
    })
    fmt.Println(c.Validate(User{}))
}





// Example 42: Validation with custom struct tags
//
// This example demonstrates how to validate struct fields with custom tags.
//
// Code:
package main

import (
    "fmt"
    "reflect"
    "github.com/go-playground/constraints"
)

type User struct {
    Name string `validate:"required"`
    Age  int    `validate:"min:18"`
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        v := reflect.ValueOf(value)
        t := v.Type()
        for i := 0; i < t.NumField(); i++ {
            field := t.Field(i)
            tag := field.Tag.Get("validate")
            if tag == "required" && v.Field(i).String() == "" {
                return fmt.Errorf("%s is required", field.Name)
            }
            if tag == "min:18" && v.Field(i).Int() < 18 {
                return fmt.Errorf("%s must be at least 18", field.Name)
            }
        }
        return nil
    })
    fmt.Println(c.Validate(User{Name: "John", Age: 15}))
}





// Example 43: Chaining validation rules
//
// This example demonstrates how to chain multiple validation rules.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) < 0 {
            return fmt.Errorf("value must be non-negative")
        }
        return nil
    })
    c.AddRule(func(value interface{}) error {
        if value.(int)%2 != 0 {
            return fmt.Errorf("value must be even")
        }
        return nil
    })
    fmt.Println(c.Validate(-3))
}





// Example 44: Validating map entries
//
// This example demonstrates how to validate map entries.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type UserMap map[string]int

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        userMap := value.(UserMap)
        for k, v := range userMap {
            if v < 0 {
                return fmt.Errorf("%s has an invalid age: %d", k, v)
            }
        }
        return nil
    })
    fmt.Println(c.Validate(UserMap{"Alice": 25, "Bob": -5}))
}





// Example 45: Using multiple constraints
//
// This example demonstrates how to use multiple constraints together.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    nonNegative := constraints.NewConstraint()
    nonNegative.AddRule(func(value interface{}) error {
        if value.(int) < 0 {
            return fmt.Errorf("value must be non-negative")
        }
        return nil
    })

    even := constraints.NewConstraint()
    even.AddRule(func(value interface{}) error {
        if value.(int)%2 != 0 {
            return fmt.Errorf("value must be even")
        }
        return nil
    })

    fmt.Println(nonNegative.Validate(-1))
    fmt.Println(even.Validate(3))
}





// Example 46: Grouping constraints
//
// This example demonstrates how to group constraints.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddGroup(func(g *constraints.Group) {
        g.AddRule(func(value interface{}) error {
            if value.(int) < 0 {
                return fmt.Errorf("value must be non-negative")
            }
            return nil
        })
        g.AddRule(func(value interface{}) error {
            if value.(int)%2 != 0 {
                return fmt.Errorf("value must be even")
            }
            return nil
        })
    })

    fmt.Println(c.Validate(-1))
    fmt.Println(c.Validate(3))
}





// Example 47: Using constraints with structs
//
// This example demonstrates how to use constraints with structs.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type User struct {
    Name string
    Age  int
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        user := value.(User)
        if user.Age < 0 {
            return fmt.Errorf("age must be non-negative")
        }
        return nil
    })

    fmt.Println(c.Validate(User{Name: "Alice", Age: -5}))
}





// Example 48: Customizing validation errors
//
// This example demonstrates how to customize validation errors.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) <= 0 {
            return fmt.Errorf("value must be greater than zero")
        }
        return nil
    })

    result := c.Validate(-1)
    if result != nil {
        fmt.Println("Validation failed with error:", result.Error())
    } else {
        fmt.Println("Validation passed")
    }
}





// Example 49: Validating nested structs
//
// This example demonstrates how to validate nested structs.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

type Address struct {
    City string
}

type User struct {
    Name    string
    Address Address
}

func main() {
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        user := value.(User)
        if user.Name == "" {
            return fmt.Errorf("name must not be empty")
        }
        if user.Address.City == "" {
            return fmt.Errorf("address.city must not be empty")
        }
        return nil
    })
    fmt.Println(c.Validate(User{Name: "John", Address: Address{City: ""}}))
}





// Example 50: Validation with custom comparisons
//
// This example demonstrates how to perform validation with custom comparisons.
//
// Code:
package main

import (
    "fmt"
    "github.com/go-playground/constraints"
)

func main() {
    targetValue := 42
    c := constraints.NewConstraint()
    c.AddRule(func(value interface{}) error {
        if value.(int) != targetValue {
            return fmt.Errorf("value must be %d", targetValue)
        }
        return nil
    })
    fmt.Println(c.Validate(50))
}





