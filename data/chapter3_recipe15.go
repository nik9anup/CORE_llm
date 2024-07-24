// This program demonstrates how to create a new SHA-1 hash instance using the crypto package in Go.

package main

import (
	"crypto"
)

func main() {
	crypto.SHA1.New()
}
