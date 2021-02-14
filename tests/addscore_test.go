package tests

import (
	"ranking/handlers"
	"ranking/models"
	"ranking/playerstorage"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAddPlayerScore test method relative
func TestAddPlayerScoreRelative(t *testing.T) {

	MockData()
	handlers.AddPlayerScore(1, "-20", true)
	affectedPlayer := models.Player{}
	list := playerstorage.GetInstance().PlayersList
	for _, player := range list {
		if player.ID == 1 {
			affectedPlayer = *player
			break
		}
	}
	assert.Equal(t, int(affectedPlayer.TotalScore), 100, "The score of the player is not correct")
}

func TestAddPlayerScoreSet(t *testing.T) {

	MockData()
	handlers.AddPlayerScore(1, "20", false)
	affectedPlayer := models.Player{}
	list := playerstorage.GetInstance().PlayersList
	for _, player := range list {
		if player.ID == 1 {
			affectedPlayer = *player
			break
		}
	}
	assert.Equal(t, int(affectedPlayer.TotalScore), 20, "The score of the player is not correct")
}

func TestAddPlayerWrongInput(t *testing.T) {

	MockData()
	_, err := handlers.AddPlayerScore(1, "asd20", false)

	assert.Error(t, err, "There should be an error returned")
}
