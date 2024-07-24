/*
Package main demonstrates a simple HTTP server and client interaction in Go.

This package defines an HTTP server that responds with a static message and a client that connects to the server via plain TCP,
sends an HTTP GET request, and reads the response. The server is created using the `http` package, and the client uses the `net` package for TCP connection.
The server is gracefully shut down after handling the client request.
*/
package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// StringServer represents a simple HTTP handler that serves a static string message.
type StringServer string

// ServeHTTP responds to HTTP requests with the static string message defined in StringServer.
func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(string(s)))
}

// createServer initializes and returns an HTTP server listening on the specified address.
// The server uses StringServer as its handler to serve a static message.
func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("HELLO GOPHER!\n"),
	}
}

const addr = "localhost:7070"

func main() {
	// Create and start the HTTP server in a separate goroutine.
	s := createServer(addr)
	go s.ListenAndServe()

	// Establish a plain TCP connection to the server.
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send an HTTP GET request to the server.
	_, err = io.WriteString(conn, "GET / HTTP/1.1\r\nHost: localhost:7070\r\n\r\n")
	if err != nil {
		panic(err)
	}

	// Read and print the server's response.
	scanner := bufio.NewScanner(conn)
	conn.SetReadDeadline(time.Now().Add(time.Second)) // Set a read deadline for the TCP connection.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Gracefully shut down the HTTP server with a 5-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}