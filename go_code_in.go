package main

import (
	"bufio"
	"fmt"
	"os"
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

// manually implement a sorting function
func sortStrings(strs []string) {
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if strs[i] > strs[j] {
				strs[i], strs[j] = strs[j], strs[i]
			}
		}
	}
}

// manually remove duplicates
func removeDuplicates(strs []string) []string {
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

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sortStrings(lines)
	lines = removeDuplicates(lines)

	for _, line := range lines {
		fmt.Println(line)
	}
}
