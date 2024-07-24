// Read "temp/file.txt" using bufio.Scanner line by line into bytes.Buffer,
// and print accumulated content. Then read entire file using ioutil.ReadFile,
// convert to string, and print directly, achieving the same file reading goal.
/*file content: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris id pretium eros. Aliquam imperdiet mi ut elit faucibus porta.
Donec facilisis nunc at risus dapibus elementum.
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("### Read as reader ###")
	f, err := os.Open("temp/file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the
	// file with reader
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())

	fmt.Println("### ReadFile ###")
	// for smaller files
	fContent, err := ioutil.ReadFile("temp/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))

}