// This program demonstrates functions for indenting and unindenting strings based on spaces and runes.
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// Example of indenting the text by prefixing with spaces.
	text := "Hi! Go is awesome."
	text = Indent(text, 6)
	fmt.Println("Indented:", text)

	// Example of unindenting the text by removing a specified number of leading spaces.
	text = Unindent(text, 3)
	fmt.Println("Unindented:", text)

	// Trying to unindent more than the current indent level won't affect the string.
	text = Unindent(text, 10)
	fmt.Println("Unindented beyond limit:", text)

	// Example of indenting the text by prefixing with a specified rune.
	text = IndentByRune(text, 10, '.')
	fmt.Println("Indented by rune:", text)
}

// Indent adds spaces to the beginning of the input string to achieve the desired indentation level.
func Indent(input string, indent int) string {
	padding := indent + len(input)
	return fmt.Sprintf("% "+strconv.Itoa(padding)+"s", input)
}

// Unindent removes a specified number of leading spaces from the input string.
// If the input is indented by fewer spaces than the specified indent, it removes all leading spaces.
func Unindent(input string, indent int) string {
	count := 0
	for _, val := range input {
		if unicode.IsSpace(val) {
			count++
		}
		if count == indent || !unicode.IsSpace(val) {
			break
		}
	}
	return input[count:]
}

// IndentByRune adds a specified rune at the beginning of the input string to achieve the desired indentation level.
func IndentByRune(input string, indent int, r rune) string {
	return strings.Repeat(string(r), indent) + input
}