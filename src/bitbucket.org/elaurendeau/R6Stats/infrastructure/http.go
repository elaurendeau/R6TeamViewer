package infrastructure

import (
	"time"
	"io/ioutil"
	"bitbucket.org/elaurendeau/R6Stats/interfaces"
	"net/http"
)

type HttpHandler struct {
	HttpTimeout time.Duration
}

const defaultHttpTimeout time.Duration = 30

func (httpHandler *HttpHandler) Get(url string) (interfaces.HttpContent, error)  {

	if httpHandler.HttpTimeout <= 0 {
		httpHandler.HttpTimeout = defaultHttpTimeout
	}

	timeout := time.Duration(httpHandler.HttpTimeout * time.Second)

	client := http.Client{
		Timeout: timeout,
	}

	response,err := client.Get(url)

	if err != nil {
		return interfaces.HttpContent{}, err
	}

	defer response.Body.Close()

	actualContent, err :=  ioutil.ReadAll(response.Body)
	httpContent := interfaces.HttpContent{Status: response.Status, StatusCode: response.StatusCode, Content: string(actualContent)}

	return httpContent,err
}


//TODO POST