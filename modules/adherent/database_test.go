package adherent

import (
	"log"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/database"
)

func TestCreateAdherentsTable(t *testing.T) {
	_, err := database.Db().Exec("SHOW TABLES LIKE 'adherents';")
	if err != nil {
		log.Fatal(err)
	}
}
