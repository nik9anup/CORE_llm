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

func median(nums []int) float64 {
	sort.Ints(nums)
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

	medianValue := median(nums)

	fmt.Println("Sorted:", nums)
	fmt.Println("Median:", medianValue)
}