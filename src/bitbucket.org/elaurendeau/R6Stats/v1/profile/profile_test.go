package profile

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)


type MyMockedObject struct{
	mock.Mock
}

func (m MyMockedObject) getSeasons() (<-chan Seasons, <-chan error) {
	return nil, nil
}

func TestValidProcess(t *testing.T) {

	_, err := Process(Request{Name: "A", Platform: "test"})

	if err != nil  {
		t.Errorf("Processing error ", err)
	}



}

func TestInvalidProcess(t *testing.T) {

	_, err := Process(Request{Name: "A", Platform: "test"})

	if err == nil  {
		t.Errorf("Processing error ", err)
	}

}