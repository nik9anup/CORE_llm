package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	nums, err := readInts("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Ints(nums)

	medianValue := median(nums)

	fmt.Println("Sorted:", nums)
	fmt.Println("Median:", medianValue)
}

func readInts(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var ints []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		ints = append(ints, num)
	}
	return ints, scanner.Err()
}

func sortInts(nums []int) {
	sort.Ints(nums)
}

func median(nums []int) float64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		return (float64(nums[n/2-1]) + float64(nums[n/2])) / 2
	}
	return float64(nums[n/2])
}