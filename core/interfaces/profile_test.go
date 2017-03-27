package interfaces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"github.com/elaurendeau/R6TeamViewer/core/domain"
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

func TestValidFetchProfiles(t *testing.T) {

	mockedProfileInteractor := new(MockedProfileInteractor)

	requestHandler := new(RequestHandler)
	requestHandler.UsecaseFetcher = mockedProfileInteractor

	profileName := "profileName"
	var profileNameList = make([]string, 0)

	profileNameList = append(profileNameList, profileName)
	profileNameList = append(profileNameList, profileName)
	profileNameList = append(profileNameList, profileName)
	profileNameList = append(profileNameList, profileName)
	profileNameList = append(profileNameList, profileName)

	platform := "platform"

	expectedProfile := new(domain.Profile)
	expectedProfile.Name = profileName
	expectedProfile.Platform = platform

	expectedProfiles := make([]*domain.Profile, len(profileNameList))

	for i := range profileNameList {
		fmt.Println(i)
		expectedProfiles[i] = expectedProfile
	}

	mockedProfileInteractor.On("FetchProfile", profileName, platform).Return(expectedProfile, nil)

	actualProfiles, err := requestHandler.FetchProfiles(profileNameList, "platform")

	fmt.Println(actualProfiles)
	fmt.Println(expectedProfile)

	assert.Nil(t, err)
	assert.Equal(t, expectedProfiles, actualProfiles)

	mockedProfileInteractor.AssertExpectations(t)
}
