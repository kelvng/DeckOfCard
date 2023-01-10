package players

import "DeckofCards/structs"

// Initialize the players slice with 4 players
func InitPlayers() {
	structs.Players = nil
	names := []string{"Alice", "Bob", "Charlie", "Dave"}
	for i := 0; i < 4; i++ {
		player := structs.PlayerResponse{
			ID:     i + 1,
			Name:   names[i],
			Points: 0,
			Coins:  5000,
		}
		structs.Players = append(structs.Players, player)
	}
}
