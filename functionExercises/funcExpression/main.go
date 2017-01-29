package main

import "fmt"

func main() {
	var number int
	stringText := "Hello, "

	greeting := makeGreeter(stringText)
	fmt.Print(greeting())

	fmt.Println("Please enter a number.")
	fmt.Scan(&number)

	evenOrOdd := takesAnInt(number)
	fmt.Println(evenOrOdd())
}

func takesAnInt(number int) func() (string, int, string, bool) {
	return func() (string, int, string, bool) {
		dividedByTwo := "When your number is divided by two, you will get"
		if number%2 == 0 {
			return dividedByTwo, number / 2, "and the number is even so we can return", true
		}
		return dividedByTwo, number / 2, "and the number isn't even so we can return", false
	}
}

func makeGreeter(stringText string) func() string {
	return func() string {
		return stringText
	}
}
