package models

import (
	"fmt"
	"math/rand"
)

// Generate new game for clay court.
// This method receives a player slice and returns it's content in random order
func GenerateClayNewGame(players []Player) []Player {
	for i := range players {
		j := rand.Intn(i + 1)
		players[i], players[j] = players[j], players[i]
	}
	return players
}

// This method list game opponents randomly generated
// This method receives a list of players and return an output string
func ListGameOpponentsClay(player []Player) string {

	result :=
		`======================================
GAMES FOR CONCRETE COURT TODAY
=======================================`
	result += "\n"

	for i, p := range player {
		if i%2 == 0 {
			//			result += "\n" + p.Name + " VS "
			result += fmt.Sprintf("%-12v", p.Name) + " VS "
		} else {
			result += p.Name + "\n"
		}
	}
	return result
}
