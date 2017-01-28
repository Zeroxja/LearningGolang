package main

import "fmt"

// print the numbers between 0 and 100
// If there is a multiple of 3 then print Fizz instead
// Multiples of 5 print Buzz
// If there are multiples of both 3 and 5 then print FizzBuzz
func main() {

	for i := 0; i <= 100; i++ {
		if i == 0 {
			fmt.Println(i)
		} else if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
