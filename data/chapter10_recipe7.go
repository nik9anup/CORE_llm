// This program detects errors in each line of the provided data using goroutines managed by errgroup,
// reporting any lines containing the substring "error:" as errors.
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"strings"

	"golang.org/x/sync/errgroup"
)

const data = `line one
line two with more words
error: This is erroneous line`

func main() {
	log.Printf("Application %s starting.", "Error Detection")
	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)

	// Use errgroup to manage multiple goroutines and errors
	g, _ := errgroup.WithContext(context.Background())
	for scanner.Scan() {
		row := scanner.Text()
		g.Go(func() error {
			if strings.Contains(row, "error:") {
				return fmt.Errorf("Error detected: %s", row)
			}
			return nil
		})
	}

	// Wait for all goroutines to complete and check for any errors
	if err := g.Wait(); err != nil {
		fmt.Println("Error while waiting: " + err.Error())
	}
}