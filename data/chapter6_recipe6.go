// Creates "test.file", retrieves its initial permissions using os.Create and f.Stat.
// Changes permissions to 0777 using f.Chmod and verifies the change with f.Stat, printing both sets of permissions.
package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.file")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Obtain current permissions
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File permissions %v\n", fi.Mode())

	// Change permissions
	err = f.Chmod(0777)
	if err != nil {
		panic(err)
	}
	fi, err = f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File permissions %v\n", fi.Mode())

}