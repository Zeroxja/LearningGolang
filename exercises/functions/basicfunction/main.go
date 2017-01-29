package main

import "fmt"

func main() {
	var number int
	fmt.Println("Please enter a number.")
	fmt.Scanln(&number)
	fmt.Println(takesAnInt(number))
}

// Takes two numbers and adds them together
func takesAnInt(number int) (int, bool) {

	if number%2 == 0 {
		return number / 2, true
	}
	return number / 2, false
}
