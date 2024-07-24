/*
Package main demonstrates creating an HTTP server in Go and sending a POST request with form data and headers.

This program sets up a simple HTTP server that handles incoming requests, prints received form data and headers,
and responds with a static message. It also creates a POST request to this server with specific form data and headers,
and then prints the server's response.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// StringServer represents a simple HTTP handler that serves a static string message.
type StringServer string

// ServeHTTP responds to HTTP requests with the static string message defined in StringServer.
// It also parses and prints any form data and headers received in the request.
func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Printf("Received form data: %v\n", req.Form)
	fmt.Printf("Received header: %v\n", req.Header)
	rw.Write([]byte(string(s)))
}

// createServer initializes and returns an HTTP server listening on the specified address.
// The server uses StringServer as its handler to serve a static message.
func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hello world"),
	}
}

const addr = "localhost:7070"

func main() {
	// Create and start the HTTP server in a separate goroutine.
	s := createServer(addr)
	go s.ListenAndServe()

	// Prepare form data for the POST request.
	form := url.Values{}
	form.Set("id", "5")
	form.Set("name", "Wolfgang")

	// Create a new POST request with the form data and the appropriate content type.
	req, err := http.NewRequest(http.MethodPost,
		"http://localhost:7070",
		strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the POST request and handle the response.
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Read and print the response from the server.
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response from server:" + string(data))
}