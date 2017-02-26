package usecases

import (
	"bitbucket.org/elaurendeau/R6Stats/domain"
	"fmt"
)

type Logger interface {
	Log(level string, message string) error
}

type ProfileInteractor struct {
	Logger Logger
	SeasonRepository domain.SeasonRepository
}

func (profileInteractor *ProfileInteractor) FetchProfile(profileName string, platform string) (*domain.Profile, error) {

	profileInteractor.Logger.Log("INFO", fmt.Sprintf("Fetching profile %v on %v", profileName, platform))
	profile := new(domain.Profile)

	profile.Name = profileName
	profile.Platform = platform

	seasons, err := profileInteractor.SeasonRepository.FindByProfileNameAndPlatform(profileName, platform)

	if err != nil {
		return profile, err
	}

	profile.Seasons = seasons

	return profile, nil
}

