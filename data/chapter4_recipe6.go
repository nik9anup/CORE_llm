// This program demonstrates manipulating dates and times using the time package in Go,
// including adding and subtracting durations and using a more convenient API for adding years, months, and days.

package main

import (
	"fmt"
	"time"
)

func main() {

	l, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}
	t := time.Date(2017, 11, 30, 11, 10, 20, 0, l)
	fmt.Printf("Default date is: %v\n", t)

	// Add 3 days
	r1 := t.Add(72 * time.Hour)
	fmt.Printf("Default date +3 days is: %v\n", r1)

	// Subtract 3 days
	r1 = t.Add(-72 * time.Hour)
	fmt.Printf("Default date -3 days is: %v\n", r1)

	// Using AddDate to add years, months, and days
	r1 = t.AddDate(1, 3, 2)
	fmt.Printf("Default date +1 year +3 months +2 days is: %v\n", r1)
}