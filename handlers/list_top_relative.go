package handlers

import (
	"ranking/models"
	"ranking/restapi/operations/top"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
)

type topPlayersRelativeListImpl struct{}

//NewTopPlayersRelativeListHandler handler
func NewTopPlayersRelativeListHandler() top.TopListRelativeHandler {
	s := &topPlayersRelativeListImpl{}
	return s
}

func (impl *topPlayersRelativeListImpl) Handle(params top.TopListRelativeParams) middleware.Responder {
	response := models.PlayersList{}

	if params.Position-params.Players < 0 {
		return middleware.Error(400, "We cannot display the player on the position: "+strconv.Itoa(int(params.Position-params.Players)))
	}

	partialResponse, err := TopPlayers(int(params.Position) + int(params.Players))

	if err != nil {
		return middleware.Error(400, err)
	}

	for i := len(partialResponse) - 1; i >= len(partialResponse)-2*int(params.Players)-1; i-- {

		response = append(response, partialResponse[i])
	}
	return top.NewTopListOK().WithPayload(response)
}
