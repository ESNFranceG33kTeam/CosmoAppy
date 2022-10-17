package money

import (
	"log"
	"os"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/database"
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
	helpers.Confpathflag = "../../test/conf_local.yaml"
	helpers.InitFile()
	helpers.ReadConfig()
	logger.LogInit()
	database.DatabaseInit()

	CreateMoneysTable()

	setUpModel()
	setUpController()
}

func testMainTeardown() {
	_, err := database.Db().Exec("DROP DATABASE IF EXISTS " + helpers.TheAppConfig().Namedb + ";")
	if err != nil {
		log.Fatal(err)
	}
}
