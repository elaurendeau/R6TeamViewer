package usecases

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"bitbucket.org/elaurendeau/R6Stats/domain"
	"reflect"
	"fmt"
	"encoding/json"
	"strings"
	"github.com/stretchr/testify/assert"
	"errors"
)


type MockedLogger struct {
	mock.Mock
}

func (mockedLogger *MockedLogger) Log(level string, message string) error {
	args := mockedLogger.Called(level, message)
	return args.Error(0)
}
type MockedSeasonRepository struct {
	mock.Mock
}

func (mockedSeasonRepository *MockedSeasonRepository) FindByProfileNameAndPlatform(profileName string, platform string) (*domain.Seasons, error) {
	args := mockedSeasonRepository.Called(profileName, platform)

	value := reflect.ValueOf(args.Get(0))
	seasons := value.Interface().(*domain.Seasons)

	return seasons, args.Error(1)
}

func TestValidFetchProfile(t *testing.T) {

	profileName := "profileName"
	platform := "platform"

	seasons := new(domain.Seasons)
	jsonSeason := "{\"seasons\":{\"4\":{\"ncsa\":{\"wins\":11,\"losses\":10,\"abandons\":1,\"season\":4,\"region\":\"ncsa\",\"ranking\":{\"rating\":2575.32307477,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.7532307477,\"stdev\":6,\"rank\":13}},\"emea\":{\"wins\":20,\"losses\":17,\"abandons\":1,\"season\":4,\"region\":\"emea\",\"ranking\":{\"rating\":2592.06724988,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.9206724988,\"stdev\":6,\"rank\":13}}},\"5\":{\"ncsa\":{\"wins\":20,\"losses\":16,\"abandons\":0,\"season\":5,\"region\":\"ncsa\",\"ranking\":{\"rating\":2909.35896403,\"next_rating\":3100,\"prev_rating\":2900,\"mean\":29.0935896403,\"stdev\":6,\"rank\":15}},\"emea\":{\"wins\":17,\"losses\":13,\"abandons\":1,\"season\":5,\"region\":\"emea\",\"ranking\":{\"rating\":2707.75585756,\"next_rating\":2900,\"prev_rating\":2700,\"mean\":27.0775585756,\"stdev\":6,\"rank\":14}}}}}"
	json.NewDecoder(strings.NewReader(jsonSeason)).Decode(seasons)

	expectedProfile := new(domain.Profile)
	expectedProfile.Name =  profileName
	expectedProfile.Platform = platform
	expectedProfile.Seasons = seasons

	profileInteractor := new(ProfileInteractor)

	mockedLogger := new(MockedLogger)
	mockedSeasonRepository := new(MockedSeasonRepository)

	profileInteractor.Logger = mockedLogger
	profileInteractor.SeasonRepository = mockedSeasonRepository

	mockedLogger.On("Log", "INFO", fmt.Sprintf("Fetching profile %v on %v", profileName, platform)).Return(nil)
	mockedSeasonRepository.On("FindByProfileNameAndPlatform",profileName, platform).Return(seasons, nil)

	actualProfile, err := profileInteractor.FetchProfile(profileName, platform)

	assert.Nil(t, err)
	assert.Equal(t, actualProfile, expectedProfile)

	mockedLogger.AssertExpectations(t)
}


func TestInvalidFetchProfile(t *testing.T) {

	profileName := "profileName"
	platform := "platform"

	seasons := new(domain.Seasons)

	profileInteractor := new(ProfileInteractor)

	mockedLogger := new(MockedLogger)
	mockedSeasonRepository := new(MockedSeasonRepository)

	profileInteractor.Logger = mockedLogger
	profileInteractor.SeasonRepository = mockedSeasonRepository

	mockedLogger.On("Log", "INFO", fmt.Sprintf("Fetching profile %v on %v", profileName, platform)).Return(nil)
	mockedSeasonRepository.On("FindByProfileNameAndPlatform",profileName, platform).Return(seasons, errors.New("controlled error"))

	_, err := profileInteractor.FetchProfile(profileName, platform)

	assert.Error(t, err, "controlled error")

	mockedLogger.AssertExpectations(t)
}