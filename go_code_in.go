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

// manually implementing sort function
func sortInts(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// manually implementing median finding function
func median(nums []int) float64 {
	n := len(nums)
	if n%2 == 0 {
		return float64(nums[n/2-1]+nums[n/2]) / 2
	}
	return float64(nums[n/2])
}

func main() {
	nums, err := readInts("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sortInts(nums)

	medianValue := median(nums)

	fmt.Println("Sorted:", nums)
	fmt.Println("Median:", medianValue)
}
