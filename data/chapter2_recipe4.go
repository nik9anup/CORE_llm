// This program splits the constant string refString using a custom split function defined by splitFunc,
// which checks if each rune in the string is '*', ',', '%', or '_',
// and prints each split word along with its index.
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary*had,a%little_lamb"

func main() {
	// The splitFunc is called for each
	// rune in a string. If the rune
	// equals any of the characters '*', ',', '%', '_',
	// the refString is split.
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_", r)
	}

	words := strings.FieldsFunc(refString, splitFunc)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}