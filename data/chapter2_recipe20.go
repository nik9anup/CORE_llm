// This program demonstrates various string manipulation operations such as trimming whitespace,
// replacing multiple spaces with a single space, and padding strings with spaces based on alignment.
package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Example of trimming leading and trailing whitespace.
	stringToTrim := "\t\t\n   Go \tis\t Awesome \t\t"
	trimResult := strings.TrimSpace(stringToTrim)
	fmt.Println("Trimmed:", trimResult)

	// Example of replacing multiple spaces with a single space.
	stringWithSpaces := "\t\t\n   Go \tis\n Awesome \t\t"
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(stringWithSpaces, " ")
	fmt.Println("Spaces replaced:", replace)

	// Examples of padding strings with spaces based on alignment.
	needSpace := "need space"
	fmt.Println("Center padded:", pad(needSpace, 14, "CENTER"))
	fmt.Println("Left padded:", pad(needSpace, 14, "LEFT"))
}

// pad function pads the input string with spaces to achieve the desired length and alignment.
func pad(input string, padLen int, align string) string {
	inputLen := len(input)

	if inputLen >= padLen {
		return input
	}

	repeat := padLen - inputLen
	var output string
	switch align {
	case "RIGHT":
		output = fmt.Sprintf("% "+strconv.Itoa(-padLen)+"s", input)
	case "LEFT":
		output = fmt.Sprintf("% "+strconv.Itoa(padLen)+"s", input)
	case "CENTER":
		bothRepeat := float64(repeat) / float64(2)
		left := int(math.Floor(bothRepeat)) + inputLen
		right := int(math.Ceil(bothRepeat))
		output = fmt.Sprintf("% "+strconv.Itoa(left)+"s% "+strconv.Itoa(right)+"s", input, "")
	}
	return output
}