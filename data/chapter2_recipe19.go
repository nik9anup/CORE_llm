// This program reads a CSV file where the field delimiter is ';' instead of the default ','.
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file.
	file, err := os.Open("data_uncommon.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV reader with ';' as the delimiter.
	reader := csv.NewReader(file)
	reader.Comma = ';'

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