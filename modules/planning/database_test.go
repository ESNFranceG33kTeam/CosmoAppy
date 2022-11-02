package planning

import (
	"log"
	"testing"
)

func TestCreatePlanningsTable(t *testing.T) {
	_, err := TheDb().Exec("SHOW TABLES LIKE 'plannings';")
	if err != nil {
		log.Fatal(err)
	}

	_, err = TheDb().Exec("SHOW TABLES LIKE 'planning_attendees';")
	if err != nil {
		log.Fatal(err)
	}
}
