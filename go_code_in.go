package main

import (
	"bufio"
	"fmt"
	"os"
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

// manually implementing sum function
func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// manually implementing average function
func average(nums []int) float64 {
	total := sum(nums)
	return float64(total) / float64(len(nums))
}


// manually implementing min function
func min(nums []int) int {
	minVal := nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

func main() {
	nums, err := readInts("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Sum:", sum(nums))
	fmt.Println("Average:", average(nums))
	fmt.Println("Min:", min(nums))
}