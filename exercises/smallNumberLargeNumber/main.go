package main

import "fmt"

// small number larger number
// print remander of the large number when divided by the small number
func main() {

	var number1 int
	var number2 int
	var remainder int

	fmt.Println("Please enter your first number:")
	fmt.Scan(&number1)
	fmt.Println("Please enter another number:")
	fmt.Scan(&number2)

	if number1 > number2 {
		remainder = number1 / number2
	} else {
		remainder = number2 / number1
	}
	fmt.Println("larger number / small number = ", remainder)
}
