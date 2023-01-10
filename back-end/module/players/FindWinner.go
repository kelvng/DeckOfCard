package players

import "DeckofCards/structs"

func FindWinner(players []structs.PlayerResponse) structs.PlayerResponse {
	winner := players[0]
	for _, player := range players {
		if player.Points > winner.Points {
			winner = player
		}
	}
	return winner
}
