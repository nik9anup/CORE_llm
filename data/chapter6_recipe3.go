// Creates a file "sample.file", writes "Go is awesome!" and "Yeah! Go is great." to it using os.Create and io.Copy,
// demonstrating file writing and copying in Go, with deferred file closure for cleanup.
package main

import (
	"io"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("sample.file")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("Go is awesome!\n")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, strings.NewReader("Yeah! Go is great.\n"))
	if err != nil {
		panic(err)
	}
}