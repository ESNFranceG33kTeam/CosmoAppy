package config

import (
	"database/sql"
	"log"
	"testing"
)

func TestDatabaseInit(t *testing.T) {
	DatabaseInit()
}

func TestCreateAdherentTable(t *testing.T) {
	createAdherentTable()

	_, err := Db().Exec("SHOW TABLES LIKE 'adherents';")
	if err != nil {
		log.Fatal(err)
	}
}

func TestDb(t *testing.T) {
	var got interface{} = Db()

	_, ok := got.(*sql.DB)

	if ok == false {
		log.Fatalln("got is not what i want !")
	}

}
