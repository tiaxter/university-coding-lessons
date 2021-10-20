package main

import "fmt"

func main() {
	var numberCount, userInput, min int

	fmt.Print("Insert the number you're gonna insert: ")
	fmt.Scan(&numberCount)

	for i := 0; i < numberCount; i++ {
		fmt.Print("Insert a number: ")
		fmt.Scan(&userInput)

		if i == 0 {
			min = userInput
		}

		if userInput < min {
			min = userInput
		}
	}

	fmt.Println("The minimum value is:", min)
}
