package main

import (
	"fmt"
	"tenisGameGenerator/models"
	"tenisGameGenerator/usecases"
)

func main() {

	allPlayers := usecases.GetPlayersFromFile("players.txt")

	clayPlayers, concretePlayers := usecases.SetCourts(&allPlayers)

	models.GenerateConcreteNewGame(concretePlayers)
	fmt.Println("Games for concrete court has just been generated.")
	models.GenerateClayNewGame(clayPlayers)
	fmt.Println("Games for clay court has just been generated.")
	fmt.Println("Games Scheduled:")
	fmt.Println(models.ListGameOpponentsConcrete(concretePlayers))
	fmt.Println(models.ListGameOpponentsClay(clayPlayers))

}
