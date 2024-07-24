// This program demonstrates the usage of regular expressions
// to replace all substrings matching the pattern "l[a-z]+" in the constant refString
// with the string "replacement".
package main

import (
	"fmt"
	"regexp"
)

const refString = "Mary had a little lamb"

func main() {
	regex := regexp.MustCompile("l[a-z]+")
	out := regex.ReplaceAllString(refString, "replacement")
	fmt.Println(out)
}