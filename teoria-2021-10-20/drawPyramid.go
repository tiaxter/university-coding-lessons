package main

import "fmt"

func main() {
	var base int

	fmt.Print("Insert the base length of a pyramid")
	fmt.Scan(&base)

	for i := 0; i < base; i++ {
		// Insert spaces
		for j := base - 1; j > 0; j-- {
			fmt.Print(" ")
		}

		fmt.Println()
	}
}
