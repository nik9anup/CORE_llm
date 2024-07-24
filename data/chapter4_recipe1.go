// This program demonstrates how to retrieve and print the current date and time using the time package in Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now()
	fmt.Println(today)

}