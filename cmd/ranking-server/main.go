// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"ranking/models"
	"ranking/playerstorage"
	"ranking/restapi"
	"ranking/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
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

	api := operations.NewRankingAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Ranking"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
