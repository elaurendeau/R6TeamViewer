package main

import (
	"./bitbucket.org/elaurendeau/R6Stats/parallel"
	"./bitbucket.org/elaurendeau/R6Stats/stats"
	"fmt"
)

func main() {

	users := make([]interface{}, 0)

	users = append(users, stats.Request{Name: "minthok", Platform: "uplay"})
	users = append(users, stats.Request{Name: "FGET-rafoufoun", Platform: "uplay"})
	users = append(users, stats.Request{Name: "sam-com", Platform: "uplay"})
	users = append(users, stats.Request{Name: "bearink", Platform: "uplay"})

	for data := range parallel.Process(users, stats.GetUserData) {
		fmt.Println("Hola this is a start")
		fmt.Println(data)
		fmt.Println("Hola this is an end")
	}

}
