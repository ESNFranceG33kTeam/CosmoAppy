package planning

import (
	"log"
	"testing"
)

func TestCreatePlanningsTable(t *testing.T) {
	_, err := TheDb().Exec("SHOW TABLES LIKE '" + db_name + "';")
	if err != nil {
		log.Fatal(err)
	}

	_, err = TheDb().Exec("SHOW TABLES LIKE '" + db_name_att + "';")
	if err != nil {
		log.Fatal(err)
	}

	_, err = TheDb().Exec("SHOW TABLES LIKE '" + db_name_monthly_stat + "';")
	if err != nil {
		log.Fatal(err)
	}
}
