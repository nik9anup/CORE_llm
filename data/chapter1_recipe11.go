// This Go program executes the 'ls -a' command to list all files and directories in the current directory,
// captures the command's output into a buffer, and prints the output if the command executes successfully.
package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out

	err := prc.Run()
	if err != nil {
		fmt.Println(err)
	}

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:\n")
		fmt.Println(out.String())
	}
}