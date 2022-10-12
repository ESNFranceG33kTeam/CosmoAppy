package config

import (
	"database/sql"
	"log"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/helpers"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
)

func TestDatabaseInit(t *testing.T) {
	helpers.InitFile("../test/conf_local.yaml")
	helpers.ReadConfig()
	logger.LogInit(helpers.AppConfig.Loglevel)
	DatabaseInit(helpers.AppConfig.Userdb, helpers.AppConfig.Passdb, helpers.AppConfig.Ipdb, helpers.AppConfig.Portdb, helpers.AppConfig.Namedb, helpers.AppConfig.Extradb)

}

func TestCreateAdherentsTable(t *testing.T) {
	createAdherentsTable()

	_, err := Db().Exec("SHOW TABLES LIKE 'adherents';")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateESNcardsTable(t *testing.T) {
	createESNcardsTable()

	_, err := Db().Exec("SHOW TABLES LIKE 'esncards';")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateVolunteersTable(t *testing.T) {
	createVolunteersTable()

	_, err := Db().Exec("SHOW TABLES LIKE 'volunteers';")
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
