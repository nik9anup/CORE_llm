/*
Package main demonstrates basic operations on HTTP headers using the `http.Header` type in Go.

This program performs various operations on HTTP headers, including setting, adding, retrieving, replacing,
and deleting header values. It prints the header state after each operation to illustrate the changes.
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new HTTP header.
	header := http.Header{}

	// Set the header "Auth-X" to a single value "abcdef1234".
	header.Set("Auth-X", "abcdef1234")

	// Add another value "defghijkl" to the header "Auth-X".
	header.Add("Auth-X", "defghijkl")

	// Print the current state of the header.
	fmt.Println(header)

	// Retrieve all values associated with the header "Auth-X".
	resSlice := header["Auth-X"]
	fmt.Println(resSlice)

	// Get the first value associated with the header "Auth-X".
	resFirst := header.Get("Auth-X")
	fmt.Println(resFirst)

	// Replace all existing values of the header "Auth-X" with a new value "newvalue".
	header.Set("Auth-X", "newvalue")
	fmt.Println(header)

	// Remove the header "Auth-X".
	header.Del("Auth-X")
	fmt.Println(header)
}