// This code reads lines of input from the user and echoes each line back to the console prefixed with "Echo:".
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// The Scanner is able to
	// scan input by lines
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Echo: %s\n", txt)
	}

}