package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func readFromFile(filename string) deck {

	byteContent, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There is an error in the reading of the file")
		fmt.Println(err)
		os.Exit(1)
		return deck{}
	} else {
		cardsString := string(byteContent)
		cards := strings.Split(cardsString, "\n")
		res := deck(cards)
		return res
	}
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	//We will get the rand object and then use rand.r() later to get the random number
	random_gen := rand.New(source)

	for i, card := range d {
		//Second time I ran this, I gor the same answer because they are made on some seed value
		//And this seed value is fixed. So, we need a different way and change the seed value

		//Now, i will get the random values
		random_num := random_gen.Intn(len(d) - 1)
		d[i] = d[random_num]
		d[random_num] = card
	}
}
