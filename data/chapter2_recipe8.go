// This program demonstrates efficient string concatenation by copying each string
// from the slice `strings` into a byte slice `bs` using the `copy` function,
// and then converting the byte slice back to a string for output.
package main

import (
	"fmt"
)

func main() {
	strings := []string{"This ", "is ", "even ", "more ", "performant "}

	bs := make([]byte, 100)
	bl := 0

	for _, val := range strings {
		bl += copy(bs[bl:], []byte(val))
	}

	fmt.Println(string(bs[:bl]))
}