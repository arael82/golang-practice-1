package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	exp := 52

	if len(d) != exp {
		t.Errorf("Expected deck length of %d, but got %v", exp, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card to be King of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToNewDeckAndNewDeckTestFromFile(t *testing.T) {

	fName := "_decktesting"

	os.Remove(fName)

	deck := newDeck()
	deck.saveToFile(fName)

	loadedDeck := newDeckFromFile(fName)

	exp := 52

	if len(loadedDeck) != exp {
		t.Errorf("Expected deck length of %d, but got %v", exp, len(loadedDeck))
	}

	os.Remove(fName)

}
