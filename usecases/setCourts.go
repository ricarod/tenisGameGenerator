package usecases

import (
	"tenisGameGenerator/models"
)

// Sets courts based on players age, players
// Receive a pointer to the players slice and returns two slices
func SetCourts(players *[]models.Player) ([]models.Player, []models.Player) {

	var clayPlayers []models.Player
	var concretePlayers []models.Player

	for _, player := range *players {
		if player.Age > 30 {
			concretePlayers = append(concretePlayers, player)
		} else {
			clayPlayers = append(clayPlayers, player)
		}
	}

	return clayPlayers, concretePlayers
}
