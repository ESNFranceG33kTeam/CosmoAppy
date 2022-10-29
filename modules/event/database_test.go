package event

import (
	"log"
	"testing"
)

func TestCreateEventsTable(t *testing.T) {
	_, err := TheDb().Exec("SHOW TABLES LIKE 'events';")
	if err != nil {
		log.Fatal(err)
	}

	_, err = TheDb().Exec("SHOW TABLES LIKE 'event_attendees';")
	if err != nil {
		log.Fatal(err)
	}
}
