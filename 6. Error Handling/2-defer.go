package main

import "fmt"

// Functions
func add(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Sum : ", res)
	return res
}

func subtract(n1, n2 int) int {
	diff := n1 - n2
	fmt.Println("Difference : ", diff)
	return diff
}

// Main function
func main() {
	//defer statement executes in the last
	defer fmt.Println("End")
	add(56, 34)
	subtract(56, 34)
}
