package main

import "fmt"

// for -> Only construst available in Go for looping
func main() {
	// for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// Infinite loop
	// for {
	// 	fmt.Println("Infinite Loop")
	// }

	// for with initializer, condition and post statement
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	// Range
	// Prints from 0 to 2 only
	for i:= range 3 {
		fmt.Println(i)
	}
}