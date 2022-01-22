package main

import "fmt"

func printOddNumber(n int) {
	for i := 1; i <= n; i += 2 {
		fmt.Println(i)
	}
}

func printEvenNumber(n int) {
	for i := 0; i <= n; i += 2 {
		fmt.Println(i)
	}
}

func main() {
	number := 50
	go printEvenNumber(number)
	go printOddNumber(number)
	fmt.Scanln()
}
