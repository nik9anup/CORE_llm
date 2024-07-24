// This code executes the "ls -a" command to list all files, captures the output, and prints it if the command succeeds.
package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	// Create a command to execute "ls -a"
	prc := exec.Command("ls", "-a")

	// Create a buffer to capture the command's output
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out

	// Start the command
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	// Wait for the command to complete
	prc.Wait()

	// Check if the command executed successfully
	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:\n")
		fmt.Println(out.String())
	}
}