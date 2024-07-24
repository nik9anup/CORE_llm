// This code logs a formatted message indicating that the application "Example" is starting and specifies the version of Go used to build the binary.
package main

import (
	"log"
	"runtime"
)

const info = `
Application %s starting.
The binary was build by GO: %s`

func main() {
	log.Printf(info, "Example", runtime.Version())
}
