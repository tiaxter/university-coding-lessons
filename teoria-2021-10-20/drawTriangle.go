package main

import "fmt"

func main() {
	var base int

	fmt.Print("Insert the base: ")
	fmt.Scan(&base)

	for i := 0; i < base; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
