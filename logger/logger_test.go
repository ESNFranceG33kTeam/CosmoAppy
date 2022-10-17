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
	LogInit()
}

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

func TestTheLogger(t *testing.T) {
	var got interface{} = TheLogger()

	_, ok := got.(*Logger)

	if ok == false {
		log.Fatalln("got is not what i want !")
	}

}
