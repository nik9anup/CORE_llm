// This program splits the constant string refString using the underscore character "_"
// as a delimiter and prints each split word along with its index.
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary_had a little_lamb"

func main() {
	words := strings.Split(refString, "_")
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}