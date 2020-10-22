package main

import "fmt"

func main() {

	card := newCard()
	fmt.Println(card)

	num := numCards()
	fmt.Println(num)
	print1()
}

/**
	To declare new functions in go, all we need to do is to provide the string as the return type
	func funcName return_type
**/

func newCard() string {
	return "Ace of spades"
}

func numCards() int {
	return 10
}
