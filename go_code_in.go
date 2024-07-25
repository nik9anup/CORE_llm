package main

import "fmt"

// manually implementing contains function
func contains(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		j := 0
		for ; j < len(substr); j++ {
			if str[i+j] != substr[j] {
				break
			}
		}
		if j == len(substr) {
			return true
		}
	}
	return false
}

func main() {
	str := "Hello, world!"
	substr := "world"

	if contains(str, substr) {
		fmt.Println("Substring found!")
	} else {
		fmt.Println("Substring not found.")
	}
}
