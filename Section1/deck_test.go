package main

import (
	"fmt"
	"os"
	"testing"
)
/**
	t - is basically a test handler.
 */
func TestNewDeck(t *testing.T) {
	deck := newDeck()
	fmt.Println(len(deck))
	if len(deck) != 52 {
		t.Errorf("Expected deck leght of 52 but got a length of %d ", len(deck))
	} else {
		if deck[0] != "Ace of Spades" {
			t.Errorf("Expected first card value is Ace of Spades but we got %v",deck[0])
		}

		if deck[len(deck)-1] != "Queen of Hearts" {
			t.Errorf("Expected last card of four of clubs bot got %v ",deck[len(deck)-1])
		}
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	//we create a new deck
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards and got %d",len(loadedDeck))
	}
	os.Remove("_decktesting")
}