/*
Package main demonstrates the creation, parsing, and serialization of URLs in Go.

This program constructs a URL using the `url.URL` struct, prints it, and then parses the constructed URL back
into a URL object. It also serializes the parsed URL object into a JSON format for display.
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {
	// Create and assemble a URL using the url.URL struct.
	u := &url.URL{}
	u.Scheme = "http"                     // Set the URL scheme (e.g., "http").
	u.Host = "localhost"                  // Set the host (e.g., "localhost").
	u.Path = "index.html"                 // Set the path (e.g., "index.html").
	u.RawQuery = "id=1&name=John"         // Set the raw query parameters (e.g., "id=1&name=John").
	u.User = url.UserPassword("admin", "1234") // Set the user credentials (e.g., username and password).

	// Print the assembled URL.
	fmt.Printf("Assembled URL:\n%v\n\n\n", u)

	// Parse the assembled URL string back into a URL object.
	parsedURL, err := url.Parse(u.String())
	if err != nil {
		panic(err)
	}

	// Serialize the parsed URL object to JSON.
	jsonURL, err := json.Marshal(parsedURL)
	if err != nil {
		panic(err)
	}

	// Print the serialized JSON representation of the parsed URL.
	fmt.Println("Parsed URL:")
	fmt.Println(string(jsonURL))
}