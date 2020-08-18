package main

import "fmt"

func main() {
	cards := newDeck()
	//cards.print()
	cardInHand, _ := deal(cards, 5)

	fmt.Println(cards.toString())
	cardInHand.shuffleCard()
	fmt.Println("Card in hands are:")
	cardInHand.print()
	cardInHand.saveToFile("card_in_hand")
}
