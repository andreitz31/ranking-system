package tests

import (
	"ranking/models"
	"ranking/playerstorage"
)

// MockData adds fake data to the Storage
func MockData() {
	player1 := models.Player{ID: 1,
		Name:       "John1",
		TotalScore: 120}

	player2 := models.Player{ID: 2,
		Name:       "John2",
		TotalScore: 0}
	player3 := models.Player{ID: 3,
		Name:       "John3",
		TotalScore: 12}
	player4 := models.Player{ID: 4,
		Name:       "John4",
		TotalScore: 121}
	player5 := models.Player{ID: 5,
		Name:       "John5",
		TotalScore: 1211}
	player6 := models.Player{ID: 6,
		Name:       "John6",
		TotalScore: 1222}
	player7 := models.Player{ID: 7,
		Name:       "John7",
		TotalScore: 1}
	playerstorage.GetInstance().PlayersList = append(playerstorage.GetInstance().PlayersList, &player1, &player2, &player3, &player4, &player5, &player6, &player7)
}
