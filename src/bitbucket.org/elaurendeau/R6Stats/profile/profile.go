package profile

import (
	"fmt"
	"errors"
	"../utils"
)


var PlayformList = []string{"uplay", "xbox", "ps4"}


type Request struct {
	Name string
	Platform string
}

type Profile struct {

}

func Process(request Request) (Profile, error) {
	fmt.Println("Processing ", request.Name)

	if !utils.Contains(PlayformList, request.Platform)  {
		return Profile{}, errors.New("Invalid platform")
	}

	return Profile{}, nil
}