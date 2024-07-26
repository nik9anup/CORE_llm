package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// sort the dates using sort.Slice
	sort.Slice(parsedDates, func(i, j int) bool {
		return parsedDates[i].Before(parsedDates[j])
	})

	fmt.Println("Sorted Dates:")
	for _, date := range parsedDates {
		fmt.Println(date.Format("2006-01-02"))
	}
}