package adherent

import (
	"log"
	"testing"
)

func TestCreateAdherentsTable(t *testing.T) {
	_, err := TheDb().Exec("SHOW TABLES LIKE 'adherents';")
	if err != nil {
		log.Fatal(err)
	}
}
