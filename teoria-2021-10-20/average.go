package main

import "fmt"

func main() {
	// Write a program that retrieve a sequence of integer numbers and then calculate the average
	var counter, input, sum float64
	for {
		fmt.Print("Insert a number: ")
		fmt.Scan(&input)

		// If user insert 0 then stop the code
		if input == 0 {
			break
		}

		sum += input
		counter++
	}
	// If counter less then 0
	if counter <= 0 {
		fmt.Println("No number inserted")
		return
	}
	// If there's at least a number
	fmt.Println("the average is:", sum/counter)
}
