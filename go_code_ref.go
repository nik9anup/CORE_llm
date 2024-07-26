package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, world!"
	substr := "world"

	// using the standard library function
	if strings.Contains(str, substr) {
		fmt.Println("Substring found!")
	} else {
		fmt.Println("Substring not found.")
	}
}