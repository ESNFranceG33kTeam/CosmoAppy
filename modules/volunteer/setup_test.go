package volunteer

import (
	"log"
	"os"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/database"
	"github.com/ESNFranceG33kTeam/sAPI/helpers"
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
	helpers.Confpathflag = "../../test/conf_local.yaml"
	helpers.InitFile()
	helpers.ReadConfig()
	TheLogger().LogInit()
	database.DatabaseInit()

	adherent.CreateAdherentsTable()
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test1", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test2", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test3", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test4", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test2", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test3", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test4", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test2", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test3", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test4", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false})

	CreateVolunteersTable()

	setUpModel()
	setUpController()
}

func testMainTeardown() {
	_, err := TheDb().Exec("DROP DATABASE IF EXISTS " + helpers.TheAppConfig().Namedb + ";")
	if err != nil {
		log.Fatal(err)
	}
}
