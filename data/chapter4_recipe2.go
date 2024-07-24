// This program demonstrates various ways to format a specific time value using the time package in Go.

package main

import (
	"fmt"
	"time"
)

func main() {
	tTime := time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)

	// Formatting with a custom layout
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/2"))

	// Formatting hours and minutes
	fmt.Printf("The time is: %s\n", tTime.Format("15:04"))

	// Using predefined RFC1123 format
	fmt.Printf("The time is: %s\n", tTime.Format(time.RFC1123))

	// Space padding for days (Go 1.9.2+)
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/_2"))

	// Zero-padding for days, months, and hours
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/01/02"))

	// Fractional seconds with leading zeros
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.00"))

	// Fractional seconds without leading zeros
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.999"))

	// AppendFormat example
	fmt.Println(string(tTime.AppendFormat([]byte("The time is up: "), "03:04PM")))
}