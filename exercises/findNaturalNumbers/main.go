package main

import "fmt"

// find the natural numbers below 100 that are multiples of 3 and 5 then work
// out the total sum of all those numbers added together
func main() {

	total := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	fmt.Println(total)
}
