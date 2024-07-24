// Retrieves the current user's information using user.Current() and prints the home directory.

package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The user home directory: " + usr.HomeDir)
}