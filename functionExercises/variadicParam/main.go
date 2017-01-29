package main

import "fmt"

var n int

func main() {
	firstNumber := 76
	secondNumber := 24
	thirdNumber := 2
	fourthNumber := 17

	fmt.Print("The highest number is: ")
	fmt.Println(highestNumber(firstNumber, secondNumber, thirdNumber, fourthNumber))
	fmt.Print("The lowest number is: ")
	fmt.Println(lowestNumber(firstNumber, secondNumber, thirdNumber, fourthNumber))
}

func highestNumber(s ...int) int {
	for _, x := range s {
		if x > n {
			n = x
		}
	}
	return n
}

func lowestNumber(s ...int) int {
	for _, x := range s {
		if x < n {
			n = x
		}
	}
	return n
}
