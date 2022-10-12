package models

import (
	"log"
	"os"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/config"
	"github.com/ESNFranceG33kTeam/sAPI/helpers"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
)

func TestMain(m *testing.M) {
	testMainSetup()

	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")
	testMainTeardown()
	os.Exit(exitVal)

}

func testMainSetup() {
	helpers.InitFile("../test/conf_local.yaml")
	helpers.ReadConfig()
	logger.LogInit(helpers.AppConfig.Loglevel)
	config.DatabaseInit(helpers.AppConfig.Userdb, helpers.AppConfig.Passdb, helpers.AppConfig.Ipdb, helpers.AppConfig.Portdb, helpers.AppConfig.Namedb, helpers.AppConfig.Extradb)

	setUpAdherent()
	setUpESNcard()
	setUpVolunteer()
}

func testMainTeardown() {
	_, err := config.Db().Exec("DROP DATABASE IF EXISTS sapi ;")
	if err != nil {
		log.Fatal(err)
	}
}
