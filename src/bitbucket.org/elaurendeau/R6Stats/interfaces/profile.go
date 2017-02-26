package interfaces

import (
	"bitbucket.org/elaurendeau/R6Stats/domain"
	"bitbucket.org/elaurendeau/R6Stats/usecases"
)

type HttpContent struct {
	Status string
	StatusCode int
	Content string
}

type RequestHandler struct {
	UsecaseFetcher usecases.UsecaseFetcher
}

func (RequestHandler *RequestHandler) FetchProfile(profileName string, platform string) (*domain.Profile, error) {
	return RequestHandler.UsecaseFetcher.FetchProfile(profileName, platform)
}
