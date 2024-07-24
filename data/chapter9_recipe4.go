// This code starts an HTTP server on port 8080 that responds with "Hello world" to incoming requests.
package main

import (
	"fmt"
	"net/http"
)

type SimpleHTTP struct{}

func (s SimpleHTTP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello world")
}

func main() {
	fmt.Println("Starting HTTP server on port 8080")
	s := &http.Server{Addr: ":8080", Handler: SimpleHTTP{}}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}