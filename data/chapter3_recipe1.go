// This code demonstrates how to parse string representations of different numeric types (decimal, hexadecimal, binary, and floating-point) into their respective numeric values using the strconv package.

package main

import (
	"fmt"
	"strconv"
)

const bin = "00001"
const hex = "2f"
const intString = "12"
const floatString = "12.3"

func main() {

	// Decimals
	res, err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed integer: %d\n", res)

	// Parsing hexadecimals
	res64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed hexadecimal: %d\n", res64)

	// Parsing binary values
	resBin, err := strconv.ParseInt(bin, 2, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed binary: %d\n", resBin)

	// Parsing floating points
	resFloat, err := strconv.ParseFloat(floatString, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed float: %.5f\n", resFloat)

}