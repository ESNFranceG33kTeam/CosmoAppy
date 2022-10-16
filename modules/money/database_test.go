package money

import (
	"log"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/config"
)

func TestCreateMoneysTable(t *testing.T) {
	CreateMoneysTable()

	_, err := config.Db().Exec("SHOW TABLES LIKE 'moneys';")
	if err != nil {
		log.Fatal(err)
	}
}
