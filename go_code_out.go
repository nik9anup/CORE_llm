package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	reversedLines := make([]string, len(lines))
	for i, line := range lines {
		reversedLines[i] = strings.Reverse(line)
	}

	err = writeLines(reversedLines, "output.txt")
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
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
	_, err = writer.WriteString(strings.Join(lines, "\n") + "\n")
	if err != nil {
		return err
	}
	return writer.Flush()
}