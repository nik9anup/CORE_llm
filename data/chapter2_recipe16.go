// This program demonstrates writing Unicode text "Gdańsk" to a file "out.txt"
// encoded in Windows-1250 charset using golang.org/x/text/encoding/charmap package.
package main

import (
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	f, err := os.OpenFile("out.txt", os.O_CREATE|os.O_RDWR, os.ModePerm|os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Encode Unicode to Windows-1250.
	encoder := charmap.Windows1250.NewEncoder()
	writer := encoder.Writer(f)
	io.WriteString(writer, "Gdańsk")
}