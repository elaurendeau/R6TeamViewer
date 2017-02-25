package interfaces

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"bitbucket.org/elaurendeau/R6Stats/domain"
	"bitbucket.org/elaurendeau/R6Stats/usecases"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"strings"
	"reflect"
	"errors"
)

type MockedHttpHandler struct{
	mock.Mock
}

func (mockedHttpHandler *MockedHttpHandler) FetchSeasons(request usecases.Request) (*domain.Seasons, error) {
	args := mockedHttpHandler.Called(request)

	//TODO find a solution without reflection
	value := reflect.ValueOf(args.Get(0))
	typedRequest := value.Interface().(*domain.Seasons)

	return typedRequest, args.Error(1)
}

func TestValidBuildProfile(t *testing.T) {

	request := usecases.Request{Name: "name", Platform: "platform"}
	expectedSeasons := new(domain.Seasons)

	jsonSeason := "{\"seasons\":{\"4\":{\"ncsa\":{\"wins\":11,\"losses\":10,\"abandons\":1,\"season\":4,\"region\":\"ncsa\",\"ranking\":{\"rating\":2575.32307477,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.7532307477,\"stdev\":6,\"rank\":13}},\"emea\":{\"wins\":20,\"losses\":17,\"abandons\":1,\"season\":4,\"region\":\"emea\",\"ranking\":{\"rating\":2592.06724988,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.9206724988,\"stdev\":6,\"rank\":13}}},\"5\":{\"ncsa\":{\"wins\":20,\"losses\":16,\"abandons\":0,\"season\":5,\"region\":\"ncsa\",\"ranking\":{\"rating\":2909.35896403,\"next_rating\":3100,\"prev_rating\":2900,\"mean\":29.0935896403,\"stdev\":6,\"rank\":15}},\"emea\":{\"wins\":17,\"losses\":13,\"abandons\":1,\"season\":5,\"region\":\"emea\",\"ranking\":{\"rating\":2707.75585756,\"next_rating\":2900,\"prev_rating\":2700,\"mean\":27.0775585756,\"stdev\":6,\"rank\":14}}}}}"

	json.NewDecoder(strings.NewReader(jsonSeason)).Decode(expectedSeasons)

	MockedHttpHandler := new(MockedHttpHandler)
	MockedHttpHandler.On("FetchSeasons", request).Return(expectedSeasons, nil)

	httpHandler := new(HttpHandler)

	httpHandler.ProfileInteractor = MockedHttpHandler

	actualSeasons, err := httpHandler.ProfileInteractor.FetchSeasons(request)

	assert.Nil(t, err)
	assert.Equal(t, actualSeasons, expectedSeasons)

	MockedHttpHandler.AssertExpectations(t)

}


func TestInvalidBuildProfileBecauseOfSeasons(t *testing.T) {

	request := usecases.Request{Name: "name", Platform: "platform"}
	expectedSeasons := new(domain.Seasons)

	MockedHttpHandler := new(MockedHttpHandler)
	MockedHttpHandler.On("FetchSeasons", request).Return(expectedSeasons, errors.New("Mocked exception"))

	httpHandler := new(HttpHandler)

	httpHandler.ProfileInteractor = MockedHttpHandler

	_, err := httpHandler.ProfileInteractor.FetchSeasons(request)

	assert.Error(t, err)

	MockedHttpHandler.AssertExpectations(t)

}