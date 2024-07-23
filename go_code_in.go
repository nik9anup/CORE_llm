package main

import (
    "fmt"
    "strconv"
)

func parseAndCalculateSum(input string) (int, bool) {

    input = trimWhitespace(input)
    if input == "" {
        return 0, false
    }

    
    var numbers []int
    var numStr string
    inNumber := false

    for _, ch := range input {
        if ch >= '0' && ch <= '9' {
            numStr += string(ch)
            inNumber = true
        } else if inNumber {
            num, err := atoi(numStr)
            if err != nil {
                return 0, false
            }
            numbers = append(numbers, num)
            numStr = ""
            inNumber = false
        }
    }

    
    sum := 0
    for _, num := range numbers {
        square := num * num
        sum += square
    }

    return sum, true
}

func atoi(str string) (int, bool) {
    num, err := strconv.Atoi(str)
    if err != nil {
        return 0, false
    }
    return num, true
}

func trimWhitespace(str string) string {
    
    start, end := 0, len(str)-1
    for start <= end && (str[start] == ' ' || str[start] == '\t' || str[start] == '\n' || str[start] == '\r') {
        start++
    }
    for end >= start && (str[end] == ' ' || str[end] == '\t' || str[end] == '\n' || str[end] == '\r') {
        end--
    }
    return str[start : end+1]
}

func main() {
    input := " 1, 2, 3, 4, 5 "
    fmt.Println("Hello world!")
    result, ok := parseAndCalculateSum(input)
    if !ok {
        fmt.Println("Error: invalid input format")
        return
    }
    fmt.Println("Sum of squares:", result)
}
