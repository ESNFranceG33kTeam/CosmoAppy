package volunteer

import (
	"log"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/database"
)

func TestCreateVolunteersTable(t *testing.T) {
	_, err := database.Db().Exec("SHOW TABLES LIKE 'volunteers';")
	if err != nil {
		log.Fatal(err)
	}
}
