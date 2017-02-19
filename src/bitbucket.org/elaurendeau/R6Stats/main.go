package main

import (
	"./stats"
	"./parallel"
	"fmt"
)

func main() {

	playerNames := make([]stats.UserRequest, 0)

	playerNames = append(playerNames, stats.UserRequest{Name: "minthok",Platform: "uplay"})
	playerNames = append(playerNames, stats.UserRequest{Name: "bearink",Platform: "uplay"})
	playerNames = append(playerNames, stats.UserRequest{Name: "sam-com",Platform: "uplay"})
	playerNames = append(playerNames, stats.UserRequest{Name: "FGET-rafoufoun",Platform: "uplay"})

	for data := range parallel.Process(playerNames, stats.GetUserData) {
		fmt.Println(data)
	}



}






