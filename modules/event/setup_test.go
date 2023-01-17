package event

import (
	"log"
	"os"
	"testing"

	"github.com/ESNFranceG33kTeam/CosmoAppy/database"
	"github.com/ESNFranceG33kTeam/CosmoAppy/helpers"
	"github.com/ESNFranceG33kTeam/CosmoAppy/modules/adherent"
	"github.com/ESNFranceG33kTeam/CosmoAppy/modules/volunteer"
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
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test1", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "1995-04-24", Situation: "worker", University: "Nancy", Homeland: "Pologne", Speakabout: "Facebook", Newsletter: false, AdhesionDate: "1995-04-24"})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test2", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "1995-04-24", Situation: "worker", University: "Nancy", Homeland: "Pologne", Speakabout: "Facebook", Newsletter: false, AdhesionDate: "1995-04-24"})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test3", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "1995-04-24", Situation: "worker", University: "Nancy", Homeland: "Pologne", Speakabout: "Facebook", Newsletter: false, AdhesionDate: "1995-04-24"})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test4", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "1995-04-24", Situation: "worker", University: "Nancy", Homeland: "Pologne", Speakabout: "Facebook", Newsletter: false, AdhesionDate: "1995-04-24"})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test2", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "1995-04-24", Situation: "worker", University: "Nancy", Homeland: "Pologne", Speakabout: "Facebook", Newsletter: false, AdhesionDate: "1995-04-24"})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test3", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "1995-04-24", Situation: "worker", University: "Nancy", Homeland: "Pologne", Speakabout: "Facebook", Newsletter: false, AdhesionDate: "1995-04-24"})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test4", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "1995-04-24", Situation: "worker", University: "Nancy", Homeland: "Pologne", Speakabout: "Facebook", Newsletter: false, AdhesionDate: "1995-04-24"})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test2", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "1995-04-24", Situation: "worker", University: "Nancy", Homeland: "Pologne", Speakabout: "Facebook", Newsletter: false, AdhesionDate: "1995-04-24"})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test3", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "1995-04-24", Situation: "worker", University: "Nancy", Homeland: "Pologne", Speakabout: "Facebook", Newsletter: false, AdhesionDate: "1995-04-24"})
	adherent.NewAdherent(&adherent.Adherent{Firstname: "Test4", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "1995-04-24", Situation: "worker", University: "Nancy", Homeland: "Pologne", Speakabout: "Facebook", Newsletter: false, AdhesionDate: "1995-04-24"})

	volunteer.CreateVolunteersTable()
	volunteer.NewVolunteer(&volunteer.Volunteer{Firstname: "Cynthia", Lastname: "Shirona", Email: "master@league.si", Actif: true, Bureau: true})
	volunteer.NewVolunteer(&volunteer.Volunteer{Firstname: "Peter", Lastname: "Wataru", Email: "master@indigo.kt", Actif: true, Bureau: true})
	volunteer.NewVolunteer(&volunteer.Volunteer{Firstname: "Pierre", Lastname: "Takeshi", Email: "pierre@argenta.kt", Actif: true, Bureau: true})
	volunteer.NewVolunteer(&volunteer.Volunteer{Firstname: "Ondine", Lastname: "Kasumi", Email: "ondine@azuria.kt", Actif: true, Bureau: true})
	volunteer.NewVolunteer(&volunteer.Volunteer{Firstname: "Prof.", Lastname: "Oak", Email: "prof.oak@kanto.kt", Actif: false, Bureau: false})
	volunteer.NewVolunteer(&volunteer.Volunteer{Firstname: "Prof.", Lastname: "Orm", Email: "prof.orm@johto.jt", Actif: false, Bureau: false})
	volunteer.NewVolunteer(&volunteer.Volunteer{Firstname: "Flora", Lastname: "Saphir", Email: "flora@hoenn.com", Actif: true, Bureau: true})

	CreateEventsTable()

	setUpModel()
	setUpController()
}

func testMainTeardown() {
	_, err := TheDb().Exec("DROP DATABASE IF EXISTS " + helpers.TheAppConfig().Namedb + ";")
	if err != nil {
		log.Fatal(err)
	}
}
