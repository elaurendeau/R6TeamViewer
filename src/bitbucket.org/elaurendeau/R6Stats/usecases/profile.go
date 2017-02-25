package usecases

import (
	"bitbucket.org/elaurendeau/R6Stats/domain"
)

type Logger interface {
	Log(level string, message string) error
}



type Request struct {
	Name string
	Platform string
}

type WebserviceHandler interface {

}

type ProfileInteractor struct {
	Logger Logger


}

func (profileInteractor *ProfileInteractor) FetchSeasons(request Request) (domain.Seasons, error) {

	return domain.Seasons{}, nil
}

