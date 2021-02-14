package tests

import (
	"ranking/handlers"
	"ranking/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopPlayers(t *testing.T) {

	MockData()

	list, _ := handlers.TopPlayers(2)

	listResult := models.PlayersList{}

	player2 := models.Player{ID: 5,
		Name:       "John5",
		TotalScore: 1211}
	player1 := models.Player{ID: 6,
		Name:       "John6",
		TotalScore: 1222}

	listResult = append(listResult, &player1, &player2)

	assert.Len(t, list, 2, "Lists should have the same length")
	assert.Equal(t, listResult[0].ID, list[0].ID, "Lists should be the same")
	assert.Equal(t, listResult[1].ID, list[1].ID, "Lists should be the same")
}

func TestTopPlayersError(t *testing.T) {

	MockData()

	_, err := handlers.TopPlayers(23)

	assert.Error(t, err, "There should be an error returned")

}
