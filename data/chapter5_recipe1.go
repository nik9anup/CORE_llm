// The code  prompts for and reads the user's name and age, then prints a greeting with that information.
package main

import (
	"fmt"
)

func main() {

	var name string
	fmt.Println("What is your name?")
	fmt.Scanf("%s\n", &name)

	var age int
	fmt.Println("What is your age?")
	fmt.Scanf("%d\n", &age)

	fmt.Printf("Hello %s, your age is %d\n", name, age)

}