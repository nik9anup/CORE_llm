package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

	// Using the standard library function to calculate the sum
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// Calculating the average
	average := float64(sum) / float64(len(nums))

	// Using the standard library sort function to find max and min
	sort.Ints(nums)
	min := nums[0]

	fmt.Println("Sum:", sum)
	fmt.Println("Average:", average)
	fmt.Println("Min:", min)
}