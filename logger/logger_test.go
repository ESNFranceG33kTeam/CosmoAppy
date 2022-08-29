package logger

import (
	"errors"
	"testing"
)

func TestLogInfo(t *testing.T) {
	LogInfo("test", "this is a unit test with info log.")
}

func TestLogWarning(t *testing.T) {
	err := errors.New("Warniiiing")
	LogWarning("test", "this is a unit test with warning log.", err)
}

func TestLogError(t *testing.T) {
	err := errors.New("Errooooor")
	LogError("test", "this is a unit test with errror log.", err)
}

func TestLogCritical(t *testing.T) {
	err := errors.New("Criticaaaal")
	LogCritical("test", "this is a unit test with critical log.", err, false)
}
