 package main

import "fmt"

func main() {
	cards := readFromFile("abc.txt")
	cards.printDeck(len(cards))
	cards.shuffle()
	fmt.Println("Shuffling")
	cards.printDeck(len(cards))
}
