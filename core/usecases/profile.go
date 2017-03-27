package usecases

import (
	"fmt"
	"sync"
	"github.com/elaurendeau/R6TeamViewer/core/domain"
)

//Logger interface for a simple logger
type Logger interface {
	//Log a message at a specific level
	Log(level string, message string) error
}

type UseCaseFetcher interface {
	FetchProfile(profileName string, platform string) (*domain.Profile, error)
}

type ProfileInteractor struct {
	Logger             Logger
	SeasonRepository   domain.SeasonRepository
	PlayerRepository   domain.PlayerRepository
	OperatorRepository domain.OperatorRepository
}

func (profileInteractor *ProfileInteractor) FetchProfile(profileName string, platform string) (*domain.Profile, error) {

	profileInteractor.Logger.Log("INFO", fmt.Sprintf("Fetching profile %v on %v", profileName, platform))

	wg := sync.WaitGroup{}
	wg.Add(3)
	profile := new(domain.Profile)

	profile.Name = profileName
	profile.Platform = platform

	var seasonsError error

	go func() {
		seasons, err := profileInteractor.SeasonRepository.FindByProfileNameAndPlatform(profileName, platform)

		if err != nil {
			seasonsError = err
		}

		profile.Seasons = seasons
		wg.Done()
	}()

	var playerError error

	go func() {
		player, err := profileInteractor.PlayerRepository.FindByProfileNameAndPlatform(profileName, platform)

		if err != nil {
			playerError = err
		}

		profile.Player = player
		wg.Done()
	}()

	var operatorsError error

	go func() {
		operators, err := profileInteractor.OperatorRepository.FindByProfileNameAndPlatform(profileName, platform)

		if err != nil {
			operatorsError = err
		}

		profile.Operators = operators
		wg.Done()
	}()

	wg.Wait()

	var errorList []error

	errorList = append(errorList, seasonsError)
	errorList = append(errorList, operatorsError)
	errorList = append(errorList, playerError)

	for _, currentError := range errorList {
		if currentError != nil {
			return profile, currentError
		}
	}

	return profile, nil
}
