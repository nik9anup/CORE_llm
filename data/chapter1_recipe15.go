// This program demonstrates interprocess communication by executing "go run sample.go" command,
// sending multiple lines of input to the child process, and printing its responses until killed after 2 seconds.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func main() {
	cmd := []string{"go", "run", "sample.go"}

	// The command line tool "go run sample.go" is executed for 2 seconds
	proc := exec.Command(cmd[0], cmd[1], cmd[2])

	// The process input is obtained in the form of io.WriteCloser using os.Pipe
	stdin, _ := proc.StdinPipe()
	defer stdin.Close()

	// For capturing and printing output from the executed process
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	go func() {
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			fmt.Println("Program says:" + s.Text())
		}
	}()

	// Start the process
	proc.Start()

	// Writing input lines to the child process standard input
	fmt.Println("Writing input")
	io.WriteString(stdin, "Hello\n")
	io.WriteString(stdin, "Golang\n")
	io.WriteString(stdin, "is awesome\n")

	// Allow 2 seconds for the process to run
	time.Sleep(time.Second * 2)

	// Kill the process
	proc.Process.Kill()
}
