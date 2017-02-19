package stats

import (
	"../parallel"
	"fmt"
)

const seasonsURL string = "https://api.r6stats.com/api/v1/players/%v/seasons?platform=%v"
const profileURL string = "https://api.r6stats.com/api/v1/players/%v?platform=%v"
const operatorsURL string = "https://api.r6stats.com/api/v1/players/%v/operators?platform=%v"

type RequestStatus int

const (
	New RequestStatus = 1 << iota
	HttpPending
	HttpProcessed
	Completed
	Error
)

//UserRequest is a structure that represent a request to get complete user information
type Request struct {
	Name string
	Platform string
	Status RequestStatus

	Seasons HttpRequest
	Profile HttpRequest
	Operators HttpRequest
}

type HttpRequest struct {
	URL string
	Status RequestStatus
	Data string
}


func GetUserData(request Request) struct{} {

	dataSlice := make([]HttpRequest, 0)


	dataSlice = append(dataSlice, HttpRequest{Status: HttpPending, URL: fmt.Sprintf(seasonsURL, request.Name, request.Platform)})
	dataSlice = append(dataSlice, HttpRequest{Status: HttpPending, URL: fmt.Sprintf(profileURL, request.Name, request.Platform)})
	dataSlice = append(dataSlice, HttpRequest{Status: HttpPending, URL: fmt.Sprintf(operatorsURL, request.Name, request.Platform)})

	var result Request
	for data := range parallel.Process(dataSlice, getSingleDataNode) {
		fmt.Println(data)
	}

	return result
}

func getSingleDataNode(httpRequest HttpRequest) HttpRequest {
	fmt.Println("HTTP GET for " + httpRequest.URL)
	httpRequest.Data = "Some data"
	httpRequest.Status = HttpProcessed
	return httpRequest
}