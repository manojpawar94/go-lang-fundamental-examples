package main

/*******************************************************
Author     : Manoj Pawar
Program    : GoLang
Description: Operation on deck
********************************************************/

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// create new go type 'deck' slice of string
type deck []string

// creat the new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamond", "Heart", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}

// print the card from the deck
func (deck deck) print() {
	for _, card := range deck {
		fmt.Println(card)
	}
}

// distribut card from deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// convert deck to string
func (deck deck) toString() string {
	return strings.Join([]string(deck), ",")
}

// save to file
func (deck deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(deck.toString()), 0666)
}

// read the file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("Error %v while reading file", err)
	}
	cards := strings.Split(string(bs), ",")
}

// shuffle the cards from the deck
func (deck deck) shuffleCard() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for index := range deck {
		newPosition := random.Intn(len(deck) - 1)
		deck[index], deck[newPosition] = deck[newPosition], deck[index]
	}
}
