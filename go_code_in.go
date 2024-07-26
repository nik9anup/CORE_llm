package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, strings.Fields(scanner.Text())...)
	}
	return words, scanner.Err()
}

// manually implementing frequency counting function
func countFrequencies(words []string) map[string]int {
	frequencies := make(map[string]int)
	for _, word := range words {
		frequencies[word]++
	}
	return frequencies
}

// manually implementing sorting function
func sortWordsByFrequency(frequencies map[string]int) []string {
	type wordFrequency struct {
		word  string
		count int
	}

	var wordFrequencies []wordFrequency
	for word, count := range frequencies {
		wordFrequencies = append(wordFrequencies, wordFrequency{word, count})
	}

	// sort by frequency in descending order
	for i := 0; i < len(wordFrequencies); i++ {
		for j := i + 1; j < len(wordFrequencies); j++ {
			if wordFrequencies[i].count < wordFrequencies[j].count {
				wordFrequencies[i], wordFrequencies[j] = wordFrequencies[j], wordFrequencies[i]
			}
		}
	}

	var sortedWords []string
	for _, wf := range wordFrequencies {
		sortedWords = append(sortedWords, wf.word)
	}
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
	for _, word := range sortedWords {
		fmt.Printf("%s: %d\n", word, frequencies[word])
	}
}