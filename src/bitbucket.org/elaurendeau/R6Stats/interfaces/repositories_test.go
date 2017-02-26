package interfaces

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"reflect"
	"github.com/stretchr/testify/assert"
	"fmt"
	"bitbucket.org/elaurendeau/R6Stats/domain"
	"encoding/json"
	"strings"
	"errors"
)

type MockedRepositoryHttpHandler struct {
	mock.Mock
}

func (mockedHttpHandler *MockedRepositoryHttpHandler) Get(url string) (HttpContent, error) {
	args := mockedHttpHandler.Called(url)

	value := reflect.ValueOf(args.Get(0))
	httpContent := value.Interface().(HttpContent)

	return httpContent, args.Error(1)
}

func TestSeasonValidFindByProfileNameAndPlatform(t *testing.T) {

	accountName := "accountName"
	platform := "platform"
	unformatedUrl := "https://api.r6stats.com/api/v1/players/%v/seasons?platform=%v"

	url := fmt.Sprintf(unformatedUrl, accountName, platform)

	mockedHttpContent := HttpContent{StatusCode: 200, Status: "200 OK", Content: "{\"seasons\":{\"4\":{\"ncsa\":{\"wins\":11,\"losses\":10,\"abandons\":1,\"season\":4,\"region\":\"ncsa\",\"ranking\":{\"rating\":2575.32307477,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.7532307477,\"stdev\":6,\"rank\":13}},\"emea\":{\"wins\":20,\"losses\":17,\"abandons\":1,\"season\":4,\"region\":\"emea\",\"ranking\":{\"rating\":2592.06724988,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.9206724988,\"stdev\":6,\"rank\":13}}},\"5\":{\"ncsa\":{\"wins\":20,\"losses\":16,\"abandons\":0,\"season\":5,\"region\":\"ncsa\",\"ranking\":{\"rating\":2909.35896403,\"next_rating\":3100,\"prev_rating\":2900,\"mean\":29.0935896403,\"stdev\":6,\"rank\":15}},\"emea\":{\"wins\":17,\"losses\":13,\"abandons\":1,\"season\":5,\"region\":\"emea\",\"ranking\":{\"rating\":2707.75585756,\"next_rating\":2900,\"prev_rating\":2700,\"mean\":27.0775585756,\"stdev\":6,\"rank\":14}}}}}"}

	seasonRepository := new(SeasonRepository)
	mockedRepositoryHttpHandler := new(MockedRepositoryHttpHandler)

	seasonRepository.HttpHandler = mockedRepositoryHttpHandler

	mockedRepositoryHttpHandler.On("Get", url).Return(mockedHttpContent, nil)

	expectedSeasons := new(domain.Seasons)
	jsonSeason := "{\"seasons\":{\"4\":{\"ncsa\":{\"wins\":11,\"losses\":10,\"abandons\":1,\"season\":4,\"region\":\"ncsa\",\"ranking\":{\"rating\":2575.32307477,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.7532307477,\"stdev\":6,\"rank\":13}},\"emea\":{\"wins\":20,\"losses\":17,\"abandons\":1,\"season\":4,\"region\":\"emea\",\"ranking\":{\"rating\":2592.06724988,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.9206724988,\"stdev\":6,\"rank\":13}}},\"5\":{\"ncsa\":{\"wins\":20,\"losses\":16,\"abandons\":0,\"season\":5,\"region\":\"ncsa\",\"ranking\":{\"rating\":2909.35896403,\"next_rating\":3100,\"prev_rating\":2900,\"mean\":29.0935896403,\"stdev\":6,\"rank\":15}},\"emea\":{\"wins\":17,\"losses\":13,\"abandons\":1,\"season\":5,\"region\":\"emea\",\"ranking\":{\"rating\":2707.75585756,\"next_rating\":2900,\"prev_rating\":2700,\"mean\":27.0775585756,\"stdev\":6,\"rank\":14}}}}}"
	json.NewDecoder(strings.NewReader(jsonSeason)).Decode(expectedSeasons)

	seasonRepository.HttpHandler.Get(url)

	actualSeasons, err := seasonRepository.FindByProfileNameAndPlatform(accountName, platform)

	assert.Nil(t, err)

	assert.Equal(t, actualSeasons, expectedSeasons)

	mockedRepositoryHttpHandler.AssertExpectations(t)

}

func TestSeasonInvalidFindByProfileNameAndPlatform(t *testing.T) {

	accountName := "accountName"
	platform := "platform"
	unformatedUrl := "https://api.r6stats.com/api/v1/players/%v/seasons?platform=%v"

	url := fmt.Sprintf(unformatedUrl, accountName, platform)

	mockedHttpContent := HttpContent{StatusCode: 200, Status: "200 OK", Content: "{\"seasons\":{\"4\":{\"ncsa\":{\"wins\":11,\"losses\":10,\"abandons\":1,\"season\":4,\"region\":\"ncsa\",\"ranking\":{\"rating\":2575.32307477,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.7532307477,\"stdev\":6,\"rank\":13}},\"emea\":{\"wins\":20,\"losses\":17,\"abandons\":1,\"season\":4,\"region\":\"emea\",\"ranking\":{\"rating\":2592.06724988,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.9206724988,\"stdev\":6,\"rank\":13}}},\"5\":{\"ncsa\":{\"wins\":20,\"losses\":16,\"abandons\":0,\"season\":5,\"region\":\"ncsa\",\"ranking\":{\"rating\":2909.35896403,\"next_rating\":3100,\"prev_rating\":2900,\"mean\":29.0935896403,\"stdev\":6,\"rank\":15}},\"emea\":{\"wins\":17,\"losses\":13,\"abandons\":1,\"season\":5,\"region\":\"emea\",\"ranking\":{\"rating\":2707.75585756,\"next_rating\":2900,\"prev_rating\":2700,\"mean\":27.0775585756,\"stdev\":6,\"rank\":14}}}}}"}

	seasonRepository := new(SeasonRepository)
	mockedRepositoryHttpHandler := new(MockedRepositoryHttpHandler)

	seasonRepository.HttpHandler = mockedRepositoryHttpHandler

	mockedRepositoryHttpHandler.On("Get", url).Return(mockedHttpContent, errors.New("Mocked error"))

	seasonRepository.HttpHandler.Get(url)

	_, err := seasonRepository.FindByProfileNameAndPlatform(accountName, platform)

	assert.Error(t, err)

	mockedRepositoryHttpHandler.AssertExpectations(t)

}

func TestSeasonValidButWithInvalidHttpRequestFindByProfileNameAndPlatform(t *testing.T) {

	accountName := "accountName"
	platform := "platform"
	unformatedUrl := "https://api.r6stats.com/api/v1/players/%v/seasons?platform=%v"

	url := fmt.Sprintf(unformatedUrl, accountName, platform)

	mockedHttpContent := HttpContent{StatusCode: 400, Status: "400 BAD REQUEST", Content: "{\"seasons\":{\"4\":{\"ncsa\":{\"wins\":11,\"losses\":10,\"abandons\":1,\"season\":4,\"region\":\"ncsa\",\"ranking\":{\"rating\":2575.32307477,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.7532307477,\"stdev\":6,\"rank\":13}},\"emea\":{\"wins\":20,\"losses\":17,\"abandons\":1,\"season\":4,\"region\":\"emea\",\"ranking\":{\"rating\":2592.06724988,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.9206724988,\"stdev\":6,\"rank\":13}}},\"5\":{\"ncsa\":{\"wins\":20,\"losses\":16,\"abandons\":0,\"season\":5,\"region\":\"ncsa\",\"ranking\":{\"rating\":2909.35896403,\"next_rating\":3100,\"prev_rating\":2900,\"mean\":29.0935896403,\"stdev\":6,\"rank\":15}},\"emea\":{\"wins\":17,\"losses\":13,\"abandons\":1,\"season\":5,\"region\":\"emea\",\"ranking\":{\"rating\":2707.75585756,\"next_rating\":2900,\"prev_rating\":2700,\"mean\":27.0775585756,\"stdev\":6,\"rank\":14}}}}}"}

	seasonRepository := new(SeasonRepository)
	mockedRepositoryHttpHandler := new(MockedRepositoryHttpHandler)

	seasonRepository.HttpHandler = mockedRepositoryHttpHandler

	mockedRepositoryHttpHandler.On("Get", url).Return(mockedHttpContent, nil)

	seasonRepository.HttpHandler.Get(url)

	_, err := seasonRepository.FindByProfileNameAndPlatform(accountName, platform)

	assert.Error(t, err)

	mockedRepositoryHttpHandler.AssertExpectations(t)

}