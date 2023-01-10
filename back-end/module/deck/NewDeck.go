package deck

import "DeckofCards/structs"

// NewDeck creates a new deck of cards with one of each suit and value
func NewDeck() {
	suits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	structs.Decks = nil
	for _, suit := range suits {
		for _, value := range values {
			card := structs.Card{suit, value}
			structs.Decks = append(structs.Decks, card)
		}
	}
}
