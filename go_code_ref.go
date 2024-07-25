package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// using the standard library sort function
	sort.Strings(lines)

	// using the standard library to remove duplicates
	lines = removeDuplicatesUsingLibrary(lines)

	for _, line := range lines {
		fmt.Println(line)
	}
}

// helper function to remove duplicates using standard library
func removeDuplicatesUsingLibrary(strs []string) []string {
	var result []string
	encountered := map[string]bool{}
	for _, str := range strs {
		if !encountered[str] {
			encountered[str] = true
			result = append(result, str)
		}
	}
	return result
}
