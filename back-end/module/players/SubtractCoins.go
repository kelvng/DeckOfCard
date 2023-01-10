package players

import (
	"DeckofCards/structs"
	"fmt"
)

func SubtractCoins(winner structs.PlayerResponse, players []structs.PlayerResponse) {
	for i := 0; i < len(players); i++ {
		if players[i].ID != winner.ID {
			players[i].Coins -= 900
			if players[i].Coins < 900 {
				//players[i].Coins = 0
				DropPlayer(players[i].ID)
				fmt.Printf("Player %d (%s) has run out of coins and can no longer play.\n", players[i].ID, players[i].Name)
			}
		}
	}

}
