package interfaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elaurendeau/Rainbow6Stats/domain"
)

type HttpHandler interface {
	Get(url string) (HttpContent, error)
}

type HttpRepository struct {
	HttpHandler HttpHandler
}

type SeasonRepository HttpRepository
type PlayerRepository HttpRepository
type OperatorRepository HttpRepository

const seasonsURL string = "https://api.r6stats.com/api/v1/players/%v/seasons?platform=%v"
const playerURL string = "https://api.r6stats.com/api/v1/players/%v?platform=%v"
const operatorsURL string = "https://api.r6stats.com/api/v1/players/%v/operators?platform=%v"

const statusOK string = "200 OK"
const statusCode200 int = 200

const errorMessageInvalidStatus string = "Error invalid request status"

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

func (seasonRepository *OperatorRepository) FindByProfileNameAndPlatform(profileName string, platform string) (*domain.Operators, error) {

	httpContent, err := seasonRepository.HttpHandler.Get(fmt.Sprintf(operatorsURL, profileName, platform))

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

func (seasonRepository *PlayerRepository) FindByProfileNameAndPlatform(profileName string, platform string) (*domain.Player, error) {

	httpContent, err := seasonRepository.HttpHandler.Get(fmt.Sprintf(playerURL, profileName, platform))

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
