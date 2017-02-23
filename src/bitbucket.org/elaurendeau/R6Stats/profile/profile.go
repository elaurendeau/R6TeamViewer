package profile

import (
	"fmt"
	"errors"
	"bitbucket.org/elaurendeau/R6Stats/utils"
	"time"
)


var PlayformList = []string{"uplay", "xbox", "ps4"}

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

	if !utils.Contains(PlayformList, request.Platform)  {
		return UserProfile{}, errors.New("Invalid platform")
	}

	seasonsOut, err := getSeasons(request)

	if err != nil {
		return UserProfile{}, err
	}

	userProfile.Seasons = <-seasonsOut

	fmt.Println(userProfile.Seasons.Seasons.Num4.Ncsa.Wins)

	fmt.Println("Hi ", userProfile)

	return userProfile, nil
}
