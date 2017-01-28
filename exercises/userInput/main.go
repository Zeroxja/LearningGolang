package main

import "fmt"

// Ask the user to enter their name and then print
func main() {

	var username string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&username)
	fmt.Println("Hello,", username)
}
