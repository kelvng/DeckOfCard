package deck

import "DeckofCards/structs"

func CalculatePoints(hand []structs.Card) int {
	pointValues := map[string]int{
		"Ace":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
		"Six":   6,
		"Seven": 7,
		"Eight": 8,
		"Nine":  9,
		"Ten":   10,
		"Jack":  10,
		"Queen": 10,
		"King":  10,
	}

	points := 0
	for _, card := range hand {
		points += pointValues[card.Value]
	}
	return points
}
