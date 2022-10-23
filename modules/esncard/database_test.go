package esncard

import (
	"log"
	"testing"
)

func TestCreateESNcardsTable(t *testing.T) {
	_, err := TheDb().Exec("SHOW TABLES LIKE 'esncards';")
	if err != nil {
		log.Fatal(err)
	}
}
