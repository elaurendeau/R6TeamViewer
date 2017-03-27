package main

import (
	"fmt"
	"time"
	"github.com/elaurendeau/R6TeamViewer/core/infrastructure"
	"github.com/elaurendeau/R6TeamViewer/core/interfaces"
	"github.com/elaurendeau/R6TeamViewer/core/usecases"
)

func main() {

	start := time.Now()

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

	var profileNameList = make([]string, 0)

	profileNameList = append(profileNameList, "free11355")
	profileNameList = append(profileNameList, "Ralhem")
	profileNameList = append(profileNameList, "BodyLord")
	profileNameList = append(profileNameList, "QuodPrinceps")
	profileNameList = append(profileNameList, "W33dPanda.cvs")
	profileNameList = append(profileNameList, "Banzai.cvs")
	profileNameList = append(profileNameList, "Wiski.cvs")
	profileNameList = append(profileNameList, "Interitus.cvs")
	profileNameList = append(profileNameList, "iDimHi")
	profileNameList = append(profileNameList, "Bricks.IRON")

	queryAmount := len(profileNameList)

	profiles, err := requestHandler.FetchProfiles(profileNameList, "uplay")

	if err != nil {
		fmt.Errorf("Error occured: ", err.Error())
	} else {
		for _, profile := range profiles {
			fmt.Println(profile.Name, " ", profile)
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Total time ", elapsed)
	fmt.Println("Average profile time ", elapsed/time.Duration(queryAmount))
}
