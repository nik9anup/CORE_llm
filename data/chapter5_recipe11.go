// Uses a pipe to capture output from executing "echo Hello Go!\nThis is example",
// then prints the output to the console using io.Copy and goroutines.
package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	pReader, pWriter := io.Pipe()

	cmd := exec.Command("echo", "Hello Go!\nThis is example")
	cmd.Stdout = pWriter

	go func() {
		defer pReader.Close()
		if _, err := io.Copy(os.Stdout, pReader); err != nil {
			log.Fatal(err)
		}
	}()

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}