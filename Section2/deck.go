package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/**
Create a new type of deck, which is a type of strings
*/
type deck []string

/**
This is the function that we will be using to generate all the cards.
We will create two string arrays, that is simply the values and the 4 suites

We will use a for loop and iterate over the array to create the deck
*/

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen"}

	for _, suite := range cardSuits {
		for _, card := range cardValues {
			cards = append(cards, card+" of "+suite)
		}
	}
	return cards
}

/**
Basically when we deal with functions, we have two kind of ways we are kind of receiving arguments

1. When we create a custom type, we can have certain methods on them and that can be passed as a receiver,
	and we accept the receiver as cards deck
2. Another way is to simply accept everything as function parameters. The simplest way is to define the variable and then
	the type correspondingly.
3. Going forward with this, we simply receive the deal of cards using our slice way
*/

func (cards deck) dealMyVariant(number int) deck {
	hand := deck{}
	for _, card := range cards[0:number] {
		hand = append(hand, card)
	}
	return hand
}

/**

 */
func deal(d deck, handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:]
}

/**
We defne the receiver and receive the arguments sent while instantiating the function.
By convention, we can take the first letter.

//string as the return type.
//Cards card - as the receiver.
printDeck(int n) - arguments passed
*/
func (cards deck) printDeck(n int) string {
	fmt.Println(n)
	//Any variable of type deck gets acess to the print method
	for i, card := range cards {
		fmt.Println(i, card)
	}
	return "Hello world"
}

/**
Now we need to look as to hhow do we convert a normal deck slice to a byteslice
Byte slice is nothing but a string of characters.
This converts each character by the corresponding decimal

The standard library expects us to pass Byte slices, which is a computer friendly information for io
type conversion is a process that converts one type of value to another
[]byte("Hi there")

^
|
[74,...]
*/

func (d deck) toString() string {
	finalString := ""
	for _, card := range []string(d) {
		finalString += card
		finalString += ","
	}
	return finalString
}

func (d deck) toString1() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) saveToFile(fileName string) error {
	data := d.toString1()
	byteData := []byte(data)
	//0666 basically tells that the file has both read and the write permissions.
	//This generates the file abc.txt
	return ioutil.WriteFile(fileName, byteData, 0666)
}
