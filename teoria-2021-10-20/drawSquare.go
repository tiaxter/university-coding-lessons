package main

import "fmt"

func main() {
	var length int

	fmt.Print("Insert the square side length: ")
	fmt.Scan(&length)

	for i := 0; i < length; i++ {
		for i := 0; i < length; i++ {
			fmt.Print("$")
		}
		fmt.Println()
	}
}
