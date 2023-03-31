package volunteer

import (
	"log"
	"testing"
)

func TestCreateVolunteersTable(t *testing.T) {
	_, err := TheDb().Exec("SHOW TABLES LIKE '" + db_name + "';")
	if err != nil {
		log.Fatal(err)
	}
}
