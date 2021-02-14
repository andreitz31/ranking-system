package handlers

import (
	"errors"
	"ranking/models"
	"ranking/playerstorage"
	"ranking/restapi/operations/top"
	"ranking/utils"

	"github.com/go-openapi/runtime/middleware"
)

type topPlayersListImpl struct{}

//NewTopPlayersListHandler handler
func NewTopPlayersListHandler() top.TopListHandler {
	return &topPlayersListImpl{}
}

func (impl *topPlayersListImpl) Handle(params top.TopListParams) middleware.Responder {
	response := models.PlayersList{}
	response, err := TopPlayers(int(params.Number))
	if err != nil {
		return middleware.Error(400, err)
	}
	return top.NewTopListOK().WithPayload(response)
}

// TopPlayers returns the top "number" players
func TopPlayers(number int) (models.PlayersList, error) {

	top := models.PlayersList{}
	allPlayers := playerstorage.GetInstance().PlayersList
	playerMap := make(map[int]int, len(allPlayers))
	if number > len(allPlayers) {
		return top, errors.New("There are not so many players in total")
	}
	for _, player := range allPlayers {
		playerMap[int(player.ID)] = int(player.TotalScore)
	}

	SortedKeyValue := utils.SortMap(playerMap)

	for i := 0; i < number; i++ {
		for _, player := range allPlayers {
			if player.ID == int64(SortedKeyValue[i].Key) {
				top = append(top, player)
				break
			}
		}
	}

	return top, nil

}
