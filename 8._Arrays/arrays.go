package main

import "fmt"

func main() {
	var nums [4] int
	nums[0] = 10
	nums[1] = 20

	// Array Length
	fmt.Println(len(nums))
	fmt.Println(nums)
	fmt.Println(nums[0])
}