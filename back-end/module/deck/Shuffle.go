package deck

import (
	"DeckofCards/structs"
	"math/rand"
	"time"
)

// Shuffle shuffles the deck of cards
func Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(structs.Decks), func(i, j int) {
		structs.Decks[i], structs.Decks[j] = structs.Decks[j], structs.Decks[i]
	})
}
