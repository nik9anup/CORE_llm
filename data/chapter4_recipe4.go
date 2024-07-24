// This program demonstrates how to work with epoch time (Unix time) using the time package in Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Set the epoch time from int64
	t := time.Unix(0, 0)
	fmt.Println(t)

	// Get the epoch time from a Time instance
	epoch := t.Unix()
	fmt.Println(epoch)

	// Current epoch time in seconds
	epochNow := time.Now().Unix()
	fmt.Printf("Epoch time in seconds: %d\n", epochNow)

	// Current epoch time in nanoseconds
	epochNano := time.Now().UnixNano()
	fmt.Printf("Epoch time in nano-seconds: %d\n", epochNano)

}