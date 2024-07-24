// This program demonstrates the usage of strings.NewReplacer to create a custom string replacer,
// replacing occurrences of "lamb" with "wolf" and "Mary" with "Jack" in the constant refString.
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"

func main() {
	replacer := strings.NewReplacer("lamb", "wolf", "Mary", "Jack")
	out := replacer.Replace(refString)
	fmt.Println(out)
}