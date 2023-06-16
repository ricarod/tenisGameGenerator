package usecases

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tenisGameGenerator/models"
)

//Process1: list of players (12 in this example) are given via text file
//Process2: main application calls the GetPlayersFile method from the use case and loads it into a slice called "allplayers"
//Process3: a pointer to the "allplayers" slice is sent to the SetCourts method from this use case and it returns two
//			slices each containing a list of players categorized by age, so the younger players (bellow 30 years old)
//			will be set to play in the clay (arcilla) court while the older ones will be set to play in concrete court
//Process4: at that point each slices will be sent to it's corresponding method (clay and concrete court accordingly)
//Process5: both methods courtClay and courtConcrete implements the SetCourtInterface with the corresponding methods:
//			GenerateGame which will randomly assign opponents and ListOpponents
//FOR THIS USE CASE THE 'AddPlayers' METHOD HAS NOT BEEN IMPLEMENTED SINCE PLAYERS ARE FROM TEXT FILE

// Read provided TXT file with player's names and ages and return them in a slice
func GetPlayersFromFile(filename string) []models.Player {
	var playerSlice []models.Player
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var allInOneString string
	var delimitAllInOneString []string

	for fileScanner.Scan() {
		if line := fileScanner.Text(); line != "" {
			readLine := fileScanner.Text()
			splittedReadLine := strings.Split(readLine, "=")
			allInOneString += splittedReadLine[1]
		}
		delimitAllInOneString = strings.Split(allInOneString, ",")
	}
	for i, _ := range delimitAllInOneString {
		if i == len(delimitAllInOneString)-1 {
			break
		}
		if i%2 != 0 {
			continue
		}

		age, _ := strconv.Atoi(strings.TrimSpace(delimitAllInOneString[i+1]))

		player := models.Player{
			Name: delimitAllInOneString[i],
			Age:  age,
		}
		i--
		playerSlice = append(playerSlice, player)
	}
	readFile.Close()
	return playerSlice

}
