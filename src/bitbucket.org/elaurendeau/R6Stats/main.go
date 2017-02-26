package main

import (
	"bitbucket.org/elaurendeau/R6Stats/interfaces"
	"bitbucket.org/elaurendeau/R6Stats/usecases"
	"fmt"
	"bitbucket.org/elaurendeau/R6Stats/infrastructure"
	"time"
	"math/big"
	"sync"
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

	wg := sync.WaitGroup{}

	queryAmount := 10


	timeChannel := make(chan time.Duration)
	wg.Add(queryAmount)


	for i := 0; i < queryAmount ; i++ {


		go func(startTime time.Time, index int) {

			_, err := requestHandler.FetchProfile("free11355", "uplay")
			if err != nil {
				fmt.Println(err)
			}

			timeChannel <- time.Since(startTime)
			wg.Done()

		}(time.Now(), i)
	}

	var totalSingleQuery time.Duration = 0
	var longestSingleQuery time.Duration = 0

	for i := 0; i < queryAmount; i++ {

		currentDuration := <-timeChannel
		totalSingleQuery += currentDuration

		if currentDuration > longestSingleQuery {
			longestSingleQuery = currentDuration
		}
	}

	wg.Wait()

	close(timeChannel)

	elapsed := time.Since(start)
	fmt.Println("Total time ", elapsed)
	fmt.Println("Total single time ", totalSingleQuery)
	fmt.Println("Average profile time ", totalSingleQuery/time.Duration(queryAmount))
	fmt.Println("Slowest profile time ", longestSingleQuery)

}
