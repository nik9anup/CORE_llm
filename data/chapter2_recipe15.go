// This program demonstrates reading and decoding a file encoded in Windows-1250 charset.
// It reads the file "win1250.txt", displays its content in its raw form,
// and then decodes it from Windows-1250 to Unicode using golang.org/x/text/encoding/charmap package.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	// Open windows-1250 encoded file.
	f, err := os.Open("win1250.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read all content in raw form.
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	content := string(b)

	fmt.Println("Without decode: " + content)

	// Decode to Unicode.
	decoder := charmap.Windows1250.NewDecoder()
	reader := decoder.Reader(strings.NewReader(content))
	b, err = ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded: " + string(b))
}