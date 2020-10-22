package main

import "fmt"

/**
1. We want to extend a base type and add some extra functionality to it.
2. We need to tell go we want to create an array of strings and add a bunch of instructions specifically
	made to work with it
3. A function with a receiver is like a "method" - a function that belongs to an instance
4. 	main.go - code o create and manpulate the deck
	deck.go - all about the deck
	deck_test.go - for testing
*/
func main() {

	cards := newDeck()
	cards.printDeck(len(cards))
	hand := cards.dealMyVariant(3)
	hand.printDeck(len(hand))

	firstHand, secondHand := deal(cards, 32)

	fmt.Println("-------------------The first hand----------------------")
	firstHand.printDeck(len(firstHand))

	fmt.Println("-------------------The Second hand----------------------")
	secondHand.printDeck(len(secondHand))

	fmt.Println(cards.toString())
	cards.saveToFile("abc.txt")
}
