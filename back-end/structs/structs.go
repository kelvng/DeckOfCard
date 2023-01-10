package structs

// Card represents a single playing card
type Card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

// Deck represents a deck of cards
type Deck []Card

var Decks Deck

var DrawnCards Deck

// Player represents a player

type PlayerResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Hand   []Card `json:"hand"`
	Points int    `json:"points"`
	Coins  int    `json:"coins"`
}

var Players []PlayerResponse
