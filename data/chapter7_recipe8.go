/*
Package main demonstrates how to handle HTTP redirects and track redirection counts in Go.

This program sets up an HTTP server that responds with temporary redirects (/redirect1, /redirect2, etc.) based on the number of redirections.
It also creates an HTTP client that handles redirects and limits the maximum number of redirects to 2. It prints details of each redirect
and stops if the maximum redirect count is exceeded.
*/
package main

import (
	"fmt"
	"net/http"
)

// addr represents the address and port on which the HTTP server listens.
const addr = "localhost:7070"

// RedirecServer is a struct that implements the http.Handler interface.
// It handles incoming HTTP requests and performs redirects based on a redirect count.
type RedirecServer struct {
	redirectCount int
}

// ServeHTTP handles incoming HTTP requests and performs a temporary redirect to the next redirect path.
// It increments the redirect count and sets a custom header "Known-redirects" with the current count.
func (s *RedirecServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.redirectCount++
	fmt.Println("Received header: " + req.Header.Get("Known-redirects"))
	// Perform a temporary redirect to the next redirection path (/redirect1, /redirect2, etc.).
	http.Redirect(rw, req, fmt.Sprintf("/redirect%d", s.redirectCount), http.StatusTemporaryRedirect)
}

func main() {
	// Create an HTTP server with a RedirecServer handler.
	s := http.Server{
		Addr:    addr,
		Handler: &RedirecServer{0}, // Start with redirectCount set to 0.
	}
	go s.ListenAndServe() // Start the HTTP server in a separate goroutine.

	// Create an HTTP client.
	client := http.Client{}
	redirectCount := 0

	// Configure the client to handle redirects and limit the maximum number of redirects to 2.
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		fmt.Println("Redirected")
		// Limit the maximum number of redirects to 2.
		if redirectCount > 2 {
			return fmt.Errorf("Too many redirects")
		}
		// Set a custom header "Known-redirects" with the current redirect count.
		req.Header.Set("Known-redirects", fmt.Sprintf("%d", redirectCount))
		redirectCount++
		// Print details of each previous request in the redirection chain.
		for _, prReq := range via {
			fmt.Printf("Previous request: %v\n", prReq.URL)
		}
		return nil
	}

	// Perform a GET request to the HTTP server.
	_, err := client.Get("http://" + addr)
	if err != nil {
		panic(err)
	}
}