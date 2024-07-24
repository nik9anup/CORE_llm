// This code retrieves the current process ID, runs the "ps" command to display process information, and prints the output.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	// Get the current process ID.
	pid := os.Getpid()
	fmt.Printf("Process PID: %d \n", pid)

	// Execute the "ps" command to display process information for the current process.
	prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	out, err := prc.Output()
	if err != nil {
		panic(err)
	}

	// Print the output of the "ps" command.
	fmt.Println(string(out))

}