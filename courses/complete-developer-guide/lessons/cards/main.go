package main

// type deck []string

// func newCard() string {
// 	return "Five of Diamonds"
// }

func main() {
	cards := newDeck()
	cards.shuffle()
	// cards.saveToFile("card_deck")
	// hand, remainingCards := deal(cards, 5)
	hand, _ := deal(cards, 5)
	hand.printDeck()
	// remainingCards.printDeck()
	// hand.saveToFile("card_hand")
	// oldCards := newDeckFromFile("card_deck")
	// oldCards.printDeck()
}
