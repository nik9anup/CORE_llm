// This code demonstrates basic operations and functions related to complex numbers in Go.

package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// Complex numbers are defined with real and imaginary parts as float64.
	a := complex(2, 3)

	fmt.Printf("Real part: %f \n", real(a))
	fmt.Printf("Imaginary part: %f \n", imag(a))

	b := complex(6, 4)

	// Basic arithmetic operations on complex numbers
	c := a - b
	fmt.Printf("Difference : %v\n", c)
	c = a + b
	fmt.Printf("Sum : %v\n", c)
	c = a * b
	fmt.Printf("Product : %v\n", c)
	c = a / b
	fmt.Printf("Quotient : %v\n", c)

	// Calculating conjugate of a complex number
	conjugate := cmplx.Conj(a)
	fmt.Println("Complex number a's conjugate : ", conjugate)

	// Calculating cosine of a complex number
	cos := cmplx.Cos(b)
	fmt.Println("Cosine of b : ", cos)

}