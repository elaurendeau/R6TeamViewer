package interfaces

import (
	"bitbucket.org/elaurendeau/R6Stats/domain"
	"bitbucket.org/elaurendeau/R6Stats/usecases"
)

type ProfileInteractor interface {
	FetchSeasons(profileName string, platform string) (*domain.Profile, error)
}

type HttpContent struct {
	Status string
	StatusCode int
	Content string
}

type RequestHandler struct {
	ProfileInteractor *usecases.ProfileInteractor
}

func (RequestHandler *RequestHandler) FetchProfile(profileName string, platform string) (*domain.Profile, error) {
	return RequestHandler.ProfileInteractor.FetchProfile(profileName, platform)
}
