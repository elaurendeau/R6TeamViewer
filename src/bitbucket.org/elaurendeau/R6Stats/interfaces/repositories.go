package interfaces

import (
	"bitbucket.org/elaurendeau/R6Stats/domain"
	"fmt"
	"errors"
	"encoding/json"
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
const profileURL string = "https://api.r6stats.com/api/v1/players/%v?platform=%v"
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