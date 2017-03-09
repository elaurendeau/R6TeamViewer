package infrastructure

import (
	"errors"
	"fmt"
)

type Logger struct {
}

const levelInfo string = "INFO"

func (logger Logger) Log(level string, message string) error {

	if level == levelInfo {
		fmt.Println(message)
		return nil
	}

	return errors.New("invalid log level")
}
