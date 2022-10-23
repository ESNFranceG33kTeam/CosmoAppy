package logger

import (
	"errors"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setUp()
	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")
	//tearDown()
	os.Exit(exitVal)
}

func setUp() {
	GetLogger().LogInit()
}

func TestLogInfo(t *testing.T) {
	GetLogger().LogInfo("test", "this is a unit test with info log.")
}

func TestLogWarning(t *testing.T) {
	err := errors.New("Warniiiing")
	GetLogger().LogWarning("test", "this is a unit test with warning log.", err)
}

func TestLogError(t *testing.T) {
	err := errors.New("Errooooor")
	GetLogger().LogError("test", "this is a unit test with errror log.", err)
}

func TestLogCritical(t *testing.T) {
	err := errors.New("Criticaaaal")
	GetLogger().LogCritical("test", "this is a unit test with critical log.", err, false)
}

func TestTheLogger(t *testing.T) {
	var got interface{} = GetLogger()

	_, ok := got.(*Logger)

	if ok == false {
		log.Fatalln("got is not what i want !")
	}

}
