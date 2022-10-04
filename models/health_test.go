package models

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
