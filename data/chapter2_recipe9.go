// This program efficiently concatenates multiple strings from the slice `strings`
// into a single string using a bytes.Buffer to accumulate the result,
// demonstrating a more performant approach compared to direct string concatenation.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	strings := []string{"This ", "is ", "even ", "more ", "performant "}
	buffer := bytes.Buffer{}
	for _, val := range strings {
		buffer.WriteString(val)
	}

	fmt.Println(buffer.String())
}