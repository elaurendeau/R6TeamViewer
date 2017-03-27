package infrastructure

import (
	"errors"
	"fmt"
)

//Logger structure used to log
type Logger struct {
}

const levelInfo string = "INFO"

//Log is a function that prints a log level ex: INFO, DEBUG, ETC and a message
func (logger Logger) Log(level string, message string) error {

	if level == levelInfo {
		fmt.Println(message)
		return nil
	}

	return errors.New("invalid log level")
}
