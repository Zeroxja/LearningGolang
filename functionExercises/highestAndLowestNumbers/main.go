package main

import "fmt"

var n, x int

func main() {

	numbers := []int{
		1, 6, 5, 2, 9, 34, 2, 65, 4,
	}

	fmt.Print("The highest number is: ")
	fmt.Println(findTheHighestNumber(numbers))
	fmt.Print("The lowest number is: ")
	fmt.Println(findTheLowestNumber(numbers))
}

func findTheHighestNumber(numbers []int) int {
	for _, n := range numbers {
		if n > x {
			x = n
		}
	}
	return x
}

func findTheLowestNumber(numbers []int) int {
	for _, n := range numbers {
		if n < x {
			x = n
		}
	}
	return x
}
