package main

import "fmt"

/**
Array - Fixed length list of things for which length cannot be changed
Slice - This is basically ArrayList equivalent of Java in which array size can grow or shrink

We can have elements of only ssame type in a slice or an array

*/
func main() {
	cards := []string{"Ace of spades", "4 of diamonds", "Jack of clubs", newCard()}
	fmt.Println(cards[0])
	cards = append(cards, "Six of hearts")
	fmt.Println(cards[4])

	/**
	Iterating over the slice
	i - tells the index
	card - tells the element at the current index
	range cards - Basically for iterating over every single record inside a slice
	:= we are redeclaring or re initilaizing
	*/
	for i, card := range cards {
		//run this one for each time in the slice
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Joker"
}
