package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"golang.org/x/exp/slices"
)

func readInts(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ints []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		ints = append(ints, num)
	}
	return ints, scanner.Err()
}

func main() {
	nums, err := readInts("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// using the standard library functions
	fmt.Println("Sum:", slices.Sum(nums))
	fmt.Println("Average:", slices.Average(nums))
	fmt.Println("Max:", slices.Max(nums))
	fmt.Println("Min:", slices.Min(nums))
}
