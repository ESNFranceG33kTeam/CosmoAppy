package health

import (
	"log"
	"testing"
)

func TestGetHealth(t *testing.T) {
	hea := GetHealth()

	if hea.Title != "IT works !" {
		log.Fatal("Not the good title")
	}
}

func TestGetStatus(t *testing.T) {
	sta := GetStatus()

	if sta.Title != "Status" {
		log.Fatal("Not the good title")
	}
}
