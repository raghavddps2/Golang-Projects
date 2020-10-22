package main

import "fmt"

/**
var - tells go that we are about to create a new variable
card - This is the name of the variable
string - this basically tells the go compiler and hence only a string will be ever assigned to this variable.

Go is a statically typed language which implies e need to know what kind of values a variable will have.

The var card will alwyas be a string. Go has a following types available
bool - true/false
string - "Hi", "How is it going ?"
int - 0,-1999,10000
float64 - 23.34, decimal numbers

card := "Ace of spades" -> This does all the tasks we were supposed to do.
Only in the initialization stage we use :=
*/
var number1 int

func main() {

	// var card string = "Ace of spades"
	card := "Ace of spades"
	card = "Five of diamonds"

	var number int = 10
	number1 = 10
	fmt.Println(number)
	fmt.Println(card)
	fmt.Println(number1)
}
