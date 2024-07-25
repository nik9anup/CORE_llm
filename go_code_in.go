package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

// manually implementing HTTP fetching function
func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// manually implementing word counting function
func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func main() {
	urls, err := readURLs("urls.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, url := range urls {
		content, err := fetchURL(url)
		if err != nil {
			fmt.Println("Error fetching URL:", err)
			continue
		}

		wordCount := countWords(content)
		fmt.Printf("URL: %s\nWord Count: %d\n", url, wordCount)
	}
}
