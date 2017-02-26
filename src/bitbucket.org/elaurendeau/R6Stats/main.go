package main

import (
	"bitbucket.org/elaurendeau/R6Stats/interfaces"
	"bitbucket.org/elaurendeau/R6Stats/usecases"
	"fmt"
	"bitbucket.org/elaurendeau/R6Stats/infrastructure"
	"time"
	"math/big"
)

func main() {

	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	requestHandler := interfaces.RequestHandler{}
	profileInteractor := new(usecases.ProfileInteractor)
	logger := infrastructure.Logger{}

	httpHandler := new(infrastructure.HttpHandler)

	seasonRepository := new(interfaces.SeasonRepository)
	playerRepository := new(interfaces.PlayerRepository)
	operatorRepository := new(interfaces.OperatorRepository)

	seasonRepository.HttpHandler = httpHandler
	playerRepository.HttpHandler = httpHandler
	operatorRepository.HttpHandler = httpHandler

	profileInteractor.Logger = logger
	profileInteractor.SeasonRepository = seasonRepository
	profileInteractor.PlayerRepository = playerRepository
	profileInteractor.OperatorRepository = operatorRepository

	requestHandler.UsecaseFetcher = profileInteractor


	profile, err := requestHandler.FetchProfile("minthok", "uplay")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(profile)

	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)

}
