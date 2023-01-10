package players

import "DeckofCards/structs"

func DetermineWinnersAndLosers(players []structs.PlayerResponse) (structs.PlayerResponse, structs.PlayerResponse) {
	// Determine the player with the highest points
	winner := FindWinner(players)

	// Set the winning player and the losing player
	if players[0].Name == winner.Name {
		return winner, players[1]
	}
	return winner, players[0]
}
