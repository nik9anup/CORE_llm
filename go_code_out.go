package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Strings(lines)
	lines = removeDuplicates(lines)

	for _, line := range lines {
		fmt.Println(line)
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

func removeDuplicates(strs []string) []string {
	encountered := make(map[string]bool, len(strs))
	result := make([]string, 0, len(strs))
	for _, str := range strs {
		if !encountered[str] {
			encountered[str] = true
			result = append(result, str)
		}
	}
	return result
}