package models

//This interface sets the methods to be implemented clay and concrete court methods

type SetupCourtInterface interface {
	GenerateNewGame()
	AddPlayers()
	ListPlayers()
}
