package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	// check deck return 52 cards
	if len(deck) != 52 {
		t.Errorf("Expected number of card in deck of 52, but got %v", len(deck))
	}

	// check deck has first card Spades of Ace
	if deck[0] == "Spades of Ace" {
		t.Errorf("Expected first card of the deck is 'Spades of Ace', but return %v", deck[0])
	}

	// check deck has last card club of King
	if deck[len(deck)-1] == "Club of King" {
		t.Errorf("Expect last card of deck to be Club of King, but return %v", deck[len(deck)-1])
	}
}
