package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func readWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, strings.Split(line, " ")...)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func countFrequencies(words []string) map[string]int {
	frequencies := make(map[string]int)
	for _, word := range words {
		frequencies[word]++
	}
	return frequencies
}

func sortWordsByFrequency(frequencies map[string]int) []string {
	var sortedWords []string
	for word, freq := range frequencies {
		sortedWords = append(sortedWords, fmt.Sprintf("%s:%d", word, freq))
	}
	sort.Strings(sortedWords)
	return sortedWords
}

func main() {
	words, err := readWords("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	frequencies := countFrequencies(words)
	sortedWords := sortWordsByFrequency(frequencies)

	fmt.Println("Word Frequencies:")
	for _, wordFreq := range sortedWords {
		parts := strings.Split(wordFreq, ":")
		fmt.Printf("%s: %s\n", parts[0], parts[1])
	}
}