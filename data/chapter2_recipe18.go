// This program reads a CSV file, ignores lines starting with '#', and ensures each record has exactly 3 fields.
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file.
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV reader.
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3 // Ensure each record has exactly 3 fields.
	reader.Comment = '#'       // Ignore lines starting with '#'.

	// Read and print each record.
	for {
		record, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(record)
	}
}