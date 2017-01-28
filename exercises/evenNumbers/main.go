package main

import "fmt"

// Prints even numbers between 1 and 100
func main() {

	even := 2
	for i := 1; i <= 100; i++ {
		if i%even == 0 {
			fmt.Println(i)
		}
	}
}
