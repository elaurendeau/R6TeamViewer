package interfaces

import (
	"sync"
	"github.com/elaurendeau/R6TeamViewer/core/usecases"
	"github.com/elaurendeau/R6TeamViewer/core/domain"
)

type HttpContent struct {
	Status     string
	StatusCode int
	Content    string
}

type RequestHandler struct {
	UsecaseFetcher usecases.UseCaseFetcher
}

func (RequestHandler *RequestHandler) FetchProfile(profileName string, platform string) (*domain.Profile, error) {
	return RequestHandler.UsecaseFetcher.FetchProfile(profileName, platform)
}

func (RequestHandler *RequestHandler) FetchProfiles(profileNames []string, platform string) ([]*domain.Profile, error) {

	profiles := make([]*domain.Profile, 0)
	var globalErr error

	errorChannel := make(chan error)
	profileChannel := make(chan *domain.Profile)

	feederWaitGroup := sync.WaitGroup{}
	feederWaitGroup.Add(len(profileNames))

	for _, profileName := range profileNames {
		go func(currentProfile string) {

			evaluatedProfile, err := RequestHandler.FetchProfile(currentProfile, platform)

			if err != nil {
				errorChannel <- err
			}

			profileChannel <- evaluatedProfile

			feederWaitGroup.Done()
		}(profileName)
	}

	go func() {
		feederWaitGroup.Wait()
		close(errorChannel)
		close(profileChannel)
	}()

	receiverWaitGroup := sync.WaitGroup{}
	receiverWaitGroup.Add(2)

	go func() {
		for currentProfile := range profileChannel {
			profiles = append(profiles, currentProfile)
		}
		receiverWaitGroup.Done()
	}()

	go func() {
		for currentError := range errorChannel {
			globalErr = currentError
		}
		receiverWaitGroup.Done()
	}()

	receiverWaitGroup.Wait()

	return profiles, globalErr
}
