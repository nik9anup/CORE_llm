package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func readURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls, scanner.Err()
}

// using net/http and bufio to fetch and read URL content
func fetchAndCountWords(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordCount, nil
}

func main() {
	urls, err := readURLs("urls.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, url := range urls {
		wordCount, err := fetchAndCountWords(url)
		if err != nil {
			fmt.Println("Error fetching URL:", err)
			continue
		}

		fmt.Printf("URL: %s\nWord Count: %d\n", url, wordCount)
	}
}
