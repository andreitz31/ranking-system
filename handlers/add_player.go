package handlers

import (
	"errors"
	"ranking/models"
	"ranking/playerstorage"
	"ranking/restapi/operations/top"

	"strconv"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

type addPlayersScoreImpl struct{}

//NewAddPlayerScoreHandler handler
func NewAddPlayerScoreHandler() top.AddPlayerHandler {
	s := addPlayersScoreImpl{}
	return &s
}

func (impl *addPlayersScoreImpl) Handle(params top.AddPlayerParams) middleware.Responder {
	response := models.Player{}

	response, err := AddPlayerScore(params.PlayerID, params.Points, params.IsRelative)

	if err != nil {
		return middleware.Error(400, err.Error())
	}
	return top.NewAddPlayerOK().WithPayload(&response)
}

//AddPlayerScore method
func AddPlayerScore(PlayerID int64, Points string, IsRelative bool) (models.Player, error) {

	storage := playerstorage.GetInstance()
	response := models.Player{}

	for i, player := range storage.PlayersList {
		if PlayerID == player.ID {
			if IsRelative {
				if strings.HasPrefix(Points, "+") {
					scoreIntValue, err := strconv.Atoi(strings.Trim(Points, "+"))
					if err != nil {
						return response, err
					}
					storage.PlayersList[i].TotalScore = storage.PlayersList[i].TotalScore + int64(scoreIntValue)
				} else if strings.HasPrefix(Points, "-") {
					scoreIntValue, err := strconv.Atoi(strings.Trim(Points, "-"))
					if err != nil {
						return response, err
					}
					storage.PlayersList[i].TotalScore = storage.PlayersList[i].TotalScore - int64(scoreIntValue)
				} else {
					return response, errors.New("Wrong input format")
				}
			} else {

				absoluteScoreValue, err := strconv.Atoi(Points)
				if err != nil {
					return response, errors.New("Wrong input format")
				}

				storage.PlayersList[i].TotalScore = int64(absoluteScoreValue)
			}

			response = *player
			break

		}
	}

	return response, nil
}
