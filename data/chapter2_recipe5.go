// This program splits the constant string refString into words using whitespace characters as delimiters
// and prints each word along with its index.
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had	a little lamb"

func main() {
	words := strings.Fields(refString)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}