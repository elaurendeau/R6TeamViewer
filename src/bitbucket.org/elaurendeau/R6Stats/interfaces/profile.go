package interfaces

import (
	"bitbucket.org/elaurendeau/R6Stats/usecases"
	"bitbucket.org/elaurendeau/R6Stats/domain"
	"fmt"
)

type ProfileInteractor interface {
	FetchSeasons(request usecases.Request) (*domain.Seasons, error)
}



type HttpHandler struct {
	ProfileInteractor ProfileInteractor

}

type HttpContent struct {
	Status string
	StatusCode int
	Content string
}

func (handler HttpHandler) buildProfile(name string, platform string) (domain.Profile, error) {
	request := usecases.Request{Name: name, Platform: platform}

	profile := domain.Profile{Name: name, Platform: platform}
	seasons, err := handler.ProfileInteractor.FetchSeasons(request)

	if err != nil {
		return domain.Profile{}, err
	}

	profile.Seasons = seasons

	fmt.Println(profile)
	return profile, nil

}

func (handler HttpHandler) getSeasons(url string) (*domain.Seasons, error) {
	return new(domain.Seasons), nil
}
