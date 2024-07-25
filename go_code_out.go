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

	scanner := bufio.NewScanner(file)
	var urls []string
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls, scanner.Err()
}

func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	return string(content), err
}

func countWords(content string) int {
	return len(strings.Fields(content))
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