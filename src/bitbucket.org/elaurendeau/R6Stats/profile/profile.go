package profile

import (
	"fmt"
	"errors"
	"bitbucket.org/elaurendeau/R6Stats/utils"
	"time"
	"sync"
	"net/http"
)


var platformList = []string{"uplay", "xbox", "ps4"}

const seasonsURL string = "https://api.r6stats.com/api/v1/players/%v/seasons?platform=%v"
const profileURL string = "https://api.r6stats.com/api/v1/players/%v?platform=%v"
const operatorsURL string = "https://api.r6stats.com/api/v1/players/%v/operators?platform=%v"

const httpTimeout = time.Duration(5 * time.Second)

type Request struct {
	Name string
	Platform string
}

type UserProfile struct {
	Name string
	Platform string
	Seasons Seasons
}

func Process(request Request) (UserProfile, error) {
	fmt.Println("Processing ", request.Name)
	userProfile := UserProfile{Name: request.Name, Platform: request.Platform}

	var wg sync.WaitGroup
	var globalError error
	if !utils.Contains(platformList, request.Platform)  {
		return UserProfile{}, errors.New("Invalid platform")
	}

	seasonsOut, errorChannel := request.getSeasons()

	wg.Add(1)
	go func() {
		for errorChannel := range errorChannel {
			globalError = errorChannel
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		userProfile.Seasons = <-seasonsOut
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Hi ", userProfile)

	return userProfile, globalError
}

type HttpFetcher interface {
	get(url string) (resp *http.Response, err error)
}
