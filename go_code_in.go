package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func readDates(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var dates []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dates = append(dates, scanner.Text())
	}
	return dates, scanner.Err()
}

// manually implementing date parsing function
func parseDates(dates []string) ([]time.Time, error) {
	var parsedDates []time.Time
	for _, date := range dates {
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			return nil, err
		}
		parsedDates = append(parsedDates, parsedDate)
	}
	return parsedDates, nil
}

// manually implementing sorting function
func sortDates(dates []time.Time) {
	for i := 0; i < len(dates); i++ {
		for j := i + 1; j < len(dates); j++ {
			if dates[i].After(dates[j]) {
				dates[i], dates[j] = dates[j], dates[i]
			}
		}
	}
}

func main() {
	dates, err := readDates("dates.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	parsedDates, err := parseDates(dates)
	if err != nil {
		fmt.Println("Error parsing dates:", err)
		return
	}

	sortDates(parsedDates)

	fmt.Println("Sorted Dates:")
	for _, date := range parsedDates {
		fmt.Println(date.Format("2006-01-02"))
	}
}