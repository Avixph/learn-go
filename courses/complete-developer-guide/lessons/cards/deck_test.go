package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spade, but got %v", d[0])
	}
	if d[len(d)-1] != "Ten of Diamonds" {
		t.Errorf("Expected last card to be Ten of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_test_deck")
	deck := newDeck()
	deck.saveToFile("_test_deck")
	loadedDeck := newDeckFromFile("_test_deck")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}
	os.Remove("_test_deck")
}
