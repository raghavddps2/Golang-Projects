package main

import "fmt"

/**
Go is NOT an Object oriented programming language
*/

func main() {
	a := 11
	if a%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
	printRecursiive(10)
}

func printRecursiive(n int) {
	if n == -1 {
		return
	}
	fmt.Println(n)
	printRecursiive(n - 1)
}
