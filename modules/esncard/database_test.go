package esncard

import (
	"log"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/database"
)

func TestCreateESNcardsTable(t *testing.T) {
	_, err := database.Db().Exec("SHOW TABLES LIKE 'esncards';")
	if err != nil {
		log.Fatal(err)
	}
}
