package report

import (
	"log"
	"testing"
)

func TestCreateReportsTable(t *testing.T) {
	_, err := TheDb().Exec("SHOW TABLES LIKE 'reports';")
	if err != nil {
		log.Fatal(err)
	}
}
