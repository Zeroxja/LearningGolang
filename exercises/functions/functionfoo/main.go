package main

import "fmt"

var x int

// If any numbers are even return true otherwise return false
func foo(n ...int) func() bool {
	return func() bool {
		for _, i := range n {
			if i%2 == 0 {
				return true
			}
		}
		return false
	}
}

// Returns the highest nuber in the slice of int
func bar(numbers []int) {
	for _, n := range numbers {
		if n > x {
			x = n
		}
	}
	fmt.Println(x)
}

func empty() {
	fmt.Println("End of exercise.")
}

func main() {
	s := foo(1, 2, 5)
	fmt.Println(s())

	x := foo(1, 3, 5)
	fmt.Println(x())

	aSlice := []int{1, 2, 3, 4}
	bar(aSlice)
	empty()
}
