package main

import (
	"bitbucket.org/elaurendeau/R6Stats/interfaces"
	"bitbucket.org/elaurendeau/R6Stats/usecases"
	"fmt"
	"bitbucket.org/elaurendeau/R6Stats/infrastructure"
)

func main() {

	requestHandler := interfaces.RequestHandler{}
	profileInteractor := new(usecases.ProfileInteractor)
	logger := infrastructure.Logger{}
	seasonRepository := new(interfaces.SeasonRepository)



	httpRepository := new(interfaces.HttpRepository)
	httpRepository.HttpHandler = new(infrastructure.HttpHandler)

	profileInteractor.Logger = logger
	profileInteractor.SeasonRepository = seasonRepository

	requestHandler.ProfileInteractor = profileInteractor


	profile, err := requestHandler.FetchProfile("minthok", "uplay")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(profile)

}
