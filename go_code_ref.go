package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// using the strings package from the standard library
func reverseString(s string) string {
	return reverseUsingLibrary(s)
}

// helper function to reverse a string using standard library functions
func reverseUsingLibrary(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var reversedLines []string
	for _, line := range lines {
		reversedLines = append(reversedLines, reverseString(line))
	}

	err = writeLines(reversedLines, "output.txt")
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
