package money

import (
	"log"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/database"
)

func TestCreateMoneysTable(t *testing.T) {
	_, err := database.Db().Exec("SHOW TABLES LIKE 'moneys';")
	if err != nil {
		log.Fatal(err)
	}
}
