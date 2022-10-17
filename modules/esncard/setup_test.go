package esncard

import (
	"log"
	"os"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/database"
	"github.com/ESNFranceG33kTeam/sAPI/helpers"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/modules/adherent"
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
	helpers.InitFile()
	helpers.ReadConfig()
	logger.LogInit()
	database.DatabaseInit()

	adherent.CreateAdherentsTable()
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test1", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test2", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test3", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test4", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})

	setUpModel()
	setUpController()
}

func testMainTeardown() {
	_, err := database.Db().Exec("DROP DATABASE IF EXISTS " + helpers.TheAppConfig().Namedb + ";")
	if err != nil {
		log.Fatal(err)
	}
}
