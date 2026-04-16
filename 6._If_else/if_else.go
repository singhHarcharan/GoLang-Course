package main

import "fmt"

func main() {
	// age := 18

	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// } else if age >= 12 {
	// 	fmt.Println("You are a teenager.")
	// } else {
	// 	fmt.Println("You are a minor.")
	// }

	// We can also declare and initialize a variable in the if statement itself. 
	// The scope of this variable will be limited to the if-else block.
	// Output of following code will be "You are a Teenager. 15"
	if age := 15; age >= 18 {
		fmt.Println("You are an adult.", age)
	} else if age >= 12 {
		fmt.Println("You are a teenager.", age)
	} else {
		fmt.Println("You are a minor.", age)
	}

	// Go does not have a ternary operator like in some other languages, but we can achieve similar functionality using an if-else statement.
}