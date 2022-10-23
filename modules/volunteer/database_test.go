package volunteer

import (
	"log"
	"testing"
)

func TestCreateVolunteersTable(t *testing.T) {
	_, err := TheDb().Exec("SHOW TABLES LIKE 'volunteers';")
	if err != nil {
		log.Fatal(err)
	}
}
