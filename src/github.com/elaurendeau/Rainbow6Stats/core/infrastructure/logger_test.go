package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidLoggerCall(t *testing.T) {
	logger := Logger{}

	err := logger.Log("INFO", "Test message")

	assert.Nil(t, err)
}
func TestInvalidValidLoggerLevel(t *testing.T) {
	logger := Logger{}

	err := logger.Log("TEST", "Test message")

	assert.Error(t, err, "invalid log level")
}
