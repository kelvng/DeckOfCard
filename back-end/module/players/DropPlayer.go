package players

import "DeckofCards/structs"

func DropPlayer(id int) {
	for i, player := range structs.Players {
		if player.ID == id {
			structs.Players = append(structs.Players[:i], structs.Players[i+1:]...)
			return
		}
	}
}
