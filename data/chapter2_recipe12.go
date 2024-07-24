// This program demonstrates the usage of strings.Replace function
// to replace all occurrences of "lamb" with "wolf" in the constant refString,
// and replace only the first two occurrences of "lamb" with "wolf" in the constant refStringTwo.
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"
const refStringTwo = "lamb lamb lamb lamb"

func main() {
	out := strings.Replace(refString, "lamb", "wolf", -1)
	fmt.Println(out)

	out = strings.Replace(refStringTwo, "lamb", "wolf", 2)
	fmt.Println(out)
}