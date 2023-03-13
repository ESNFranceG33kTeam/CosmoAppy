package money

import (
	"log"
	"testing"
)

func TestCreateMoneysTable(t *testing.T) {
	_, err := TheDb().Exec("SHOW TABLES LIKE 'moneys';")
	if err != nil {
		log.Fatal(err)
	}

	_, err = TheDb().Exec("SHOW TABLES LIKE 'money_stats_monthly';")
	if err != nil {
		log.Fatal(err)
	}
}
