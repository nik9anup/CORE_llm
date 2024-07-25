package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
)

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

func min(nums []int) int {
	return int(math.Min(float64(nums[0]), float64(math.Min(float64(nums[1:])...))))
}

func sum(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func average(nums []int) float64 {
	return float64(sum(nums)) / float64(len(nums))
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