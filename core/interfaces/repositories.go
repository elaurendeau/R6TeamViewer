package interfaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elaurendeau/R6TeamViewer/core/domain"
)

//HttpHandler is the http interface used to get or post data
type HttpHandler interface {
	//Get is called to perform an HTTP Get. Provide an URL and obtain a HttpContent and an error as the result.
	Get(url string) (HttpContent, error)
}

//HttpRepository is a structure that contains the handler for the repository requests.
type HttpRepository struct {
	HttpHandler HttpHandler
}

//SeasonRepository is just a facade for the season HttpRepository
type SeasonRepository HttpRepository

//PlayerRepository is just a facade for the player HttpRepository
type PlayerRepository HttpRepository

//OperatorRepository is just a facade for the operator HttpRepository
type OperatorRepository HttpRepository

const seasonsURL string = "https://api.r6stats.com/api/v1/players/%v/seasons?platform=%v"
const playerURL string = "https://api.r6stats.com/api/v1/players/%v?platform=%v"
const operatorsURL string = "https://api.r6stats.com/api/v1/players/%v/operators?platform=%v"

const statusOK string = "200 OK"
const statusCode200 int = 200

const errorMessageInvalidStatus string = "Error invalid request status"

//FindByProfileNameAndPlatform finds an instance of the seasons for a specific profileName and platform
func (seasonRepository *SeasonRepository) FindByProfileNameAndPlatform(profileName string, platform string) (*domain.Seasons, error) {

	httpContent, err := seasonRepository.HttpHandler.Get(fmt.Sprintf(seasonsURL, profileName, platform))

	if err != nil {
		return nil, err
	}

	if httpContent.Status != statusOK || httpContent.StatusCode != statusCode200 {
		return nil, errors.New(errorMessageInvalidStatus)
	}

	seasons := new(domain.Seasons)
	json.Unmarshal([]byte(httpContent.Content), seasons)

	return seasons, nil
}

//FindByProfileNameAndPlatform finds an instance of the operators for a specific profileName and platform
func (operatorRepository *OperatorRepository) FindByProfileNameAndPlatform(profileName string, platform string) (*domain.Operators, error) {

	httpContent, err := operatorRepository.HttpHandler.Get(fmt.Sprintf(operatorsURL, profileName, platform))

	if err != nil {
		return nil, err
	}

	if httpContent.Status != statusOK || httpContent.StatusCode != statusCode200 {
		return nil, errors.New(errorMessageInvalidStatus)
	}

	operators := new(domain.Operators)
	json.Unmarshal([]byte(httpContent.Content), operators)

	return operators, nil
}

//FindByProfileNameAndPlatform finds an instance of the player for a specific profileName and platform
func (playerRepository *PlayerRepository) FindByProfileNameAndPlatform(profileName string, platform string) (*domain.Player, error) {

	httpContent, err := playerRepository.HttpHandler.Get(fmt.Sprintf(playerURL, profileName, platform))

	if err != nil {
		return nil, err
	}

	if httpContent.Status != statusOK || httpContent.StatusCode != statusCode200 {
		return nil, errors.New(errorMessageInvalidStatus)
	}

	player := new(domain.Player)
	json.Unmarshal([]byte(httpContent.Content), player)

	return player, nil
}
