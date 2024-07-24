// This program demonstrates various string manipulations such as case conversion,
// matching case-insensitive strings, and converting snake_case to camelCase.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	email     = "ExamPle@domain.com"
	name      = "isaac newton"
	upc       = "upc"
	i         = "i"
	snakeCase = "first_name"
)

func main() {
	// Compare email case insensitively.
	input := "Example@domain.com"
	input = strings.ToLower(input)
	emailToCompare := strings.ToLower(email)
	matches := input == emailToCompare
	fmt.Printf("Email matches: %t\n", matches)

	// Convert to upper case.
	upcCode := strings.ToUpper(upc)
	fmt.Println("UPPER case: " + upcCode)

	// Convert to upper case and title case.
	str := "ǳ"
	fmt.Printf("%s in upper: %s and title: %s \n",
		str,
		strings.ToUpper(str),
		strings.ToTitle(str))

	// Compare ToTitle and ToTitleSpecial functions.
	title := strings.ToTitle(i)
	titleTurk := strings.ToTitleSpecial(unicode.TurkishCase, i)
	if title != titleTurk {
		fmt.Printf("ToTitle is different: %#U vs. %#U \n",
			title[0],
			[]rune(titleTurk)[0])
	}

	// Correct the case of a name.
	correctNameCase := strings.Title(name)
	fmt.Println("Corrected name: " + correctNameCase)

	// Convert snake_case to camelCase.
	firstNameCamel := toCamelCase(snakeCase)
	fmt.Println("Camel case: " + firstNameCamel)
}

func toCamelCase(input string) string {
	titleSpace := strings.Title(strings.Replace(input, "_", " ", -1))
	camel := strings.ReplaceAll(titleSpace, " ", "")
	return strings.ToLower(camel[:1]) + camel[1:]
}