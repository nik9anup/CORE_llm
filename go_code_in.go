package main

import (
	"fmt"
	"strings"
)

// add function: adds two integers
func add(a, b int) int {
	return a + b
}

// subtract function: subtracts the second integer from the first
func subtract(a, b int) int {
	return a - b
}

// multiply function: multiplies two integers
func multiply(a, b int) int {
	return a * b
}

// sumArray function: sums elements of an integer slice, calls add function
func sumArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum = add(sum, num)
	}
	return sum
}

// capitalize function: capitalizes each word in the string
func capitalize(s string) string {
	return strings.Title(s)
}

// reverse function: reverses the string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// capitalizeAndReverse function: capitalizes and then reverses the string, calls capitalize and reverse functions
func capitalizeAndReverse(s string) string {
	capitalized := capitalize(s)
	return reverse(capitalized)
}

// Person struct: contains basic person information
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// FullName method: returns the full name of the person
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Introduce method: returns an introduction string, calls FullName method
func (p Person) Introduce() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.FullName(), p.Age)
}

func POWER(num int, power int) int {
   var result int = 1
   if power != 0 {
        result = (num * POWER(num, power-1))
   }
   return result
}

func main() {
	// Arithmetic functions
	fmt.Println("Add 10 + 5 =", add(10, 5))
	fmt.Println("Subtract 10 - 5 =", subtract(10, 5))
	fmt.Println("Multiply 10 * 5 =", multiply(10, 5))

	// Sum array
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of array =", sumArray(nums))

	// String functions
	s := "hello world"
	fmt.Println("Capitalized:", capitalize(s))
	fmt.Println("Reversed:", reverse(s))
	fmt.Println("Capitalized and Reversed:", capitalizeAndReverse(s))

	// Struct and method
	p := Person{FirstName: "John", LastName: "Doe", Age: 30}
	fmt.Println("Full Name:", p.FullName())
	fmt.Println(p.Introduce())
}
