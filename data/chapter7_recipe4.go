/*
Package main demonstrates how to create a simple HTTP server in Go and interact with it using POST requests.

This package defines an HTTP server that responds with a static message and logs received form data. It also provides
two examples of sending POST requests to the server using the `http` package: one with `http.Post` and one with `http.NewRequest`.
The responses from the server are printed to the console.
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
// It also parses and prints any form data received in the request.
func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Printf("Received form data: %v\n", req.Form)
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

	// Send a POST request to the server using http.Post.
	simplePost()

	// Send a POST request to the server using a custom http.Request.
	useRequest()
}

// simplePost sends a POST request with form data to the HTTP server using http.Post.
// It prints the response from the server to the console.
func simplePost() {
	res, err := http.Post("http://localhost:7070",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=Radek&surname=Sohlich"))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response from server:" + string(data))
}

// useRequest sends a POST request with form data to the HTTP server using a custom http.Request.
// It prints the response from the server to the console.
func useRequest() {
	hc := http.Client{}
	form := url.Values{}
	form.Add("name", "Radek")
	form.Add("surname", "Sohlich")

	req, err := http.NewRequest("POST",
		"http://localhost:7070",
		strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := hc.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response from server:" + string(data))
}