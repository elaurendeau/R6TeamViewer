package interfaces

import (
	"github.com/stretchr/testify/mock"
	"bitbucket.org/elaurendeau/R6Stats/domain"
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

type MockedProfileInteractor struct {
	mock.Mock
}

func (mockedLogger *MockedProfileInteractor) FetchProfile(profileName string, platform string) (*domain.Profile, error) {
	args := mockedLogger.Called(profileName, platform)

	value := reflect.ValueOf(args.Get(0))
	profile := value.Interface().(*domain.Profile)

	return profile, args.Error(1)
}

func TestValidFetchProfile(t *testing.T) {

	mockedProfileInteractor := new(MockedProfileInteractor)

	requestHandler := new(RequestHandler)
	requestHandler.UsecaseFetcher = mockedProfileInteractor

	profileName := "profileName"
	platform := "platform"

	expectedProfile := new(domain.Profile)
	expectedProfile.Name = profileName
	expectedProfile.Platform = platform

	mockedProfileInteractor.On("FetchProfile", profileName, platform).Return(expectedProfile, nil)

	actualProfile, err := requestHandler.FetchProfile("profileName", "platform")

	fmt.Println(actualProfile)
	fmt.Println(expectedProfile)

	assert.Nil(t, err)
	assert.Equal(t, actualProfile, expectedProfile)

	mockedProfileInteractor.AssertExpectations(t)
}
