// This code prints the path to the current executable, resolves and prints its directory, and evaluates and prints any symbolic links in the directory path.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Path to executable file
	fmt.Println(ex)

	// Resolve the directory
	// of the executable
	exPath := filepath.Dir(ex)
	fmt.Println("Executable path :" + exPath)

	// Use EvalSymlinks to get
	// the real path.
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Symlink evaluated:" + realPath)
}