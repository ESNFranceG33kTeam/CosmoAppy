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
	adherent.NewAdherent(
		&adherent.Adherent{
			Firstname:    "Test1",
			Lastname:     "Tutu",
			Email:        "toto@toto.fr",
			Dateofbirth:  "1995-04-24",
			Situation:    "worker",
			University:   "Nancy",
			Homeland:     "Pologne",
			Speakabout:   "Facebook",
			Newsletter:   false,
			AdhesionDate: "1995-04-24",
		},
	)
	adherent.NewAdherent(
		&adherent.Adherent{
			Firstname:    "Test2",
			Lastname:     "Tutu",
			Email:        "toto@toto.fr",
			Dateofbirth:  "1995-04-24",
			Situation:    "worker",
			University:   "Nancy",
			Homeland:     "Pologne",
			Speakabout:   "Facebook",
			Newsletter:   false,
			AdhesionDate: "1995-04-24",
		},
	)
	adherent.NewAdherent(
		&adherent.Adherent{
			Firstname:    "Test3",
			Lastname:     "Tutu",
			Email:        "toto@toto.fr",
			Dateofbirth:  "1995-04-24",
			Situation:    "worker",
			University:   "Nancy",
			Homeland:     "Pologne",
			Speakabout:   "Facebook",
			Newsletter:   false,
			AdhesionDate: "1995-04-24",
		},
	)
	adherent.NewAdherent(
		&adherent.Adherent{
			Firstname:    "Test4",
			Lastname:     "Tutu",
			Email:        "toto@toto.fr",
			Dateofbirth:  "1995-04-24",
			Situation:    "worker",
			University:   "Nancy",
			Homeland:     "Pologne",
			Speakabout:   "Facebook",
			Newsletter:   false,
			AdhesionDate: "1995-04-24",
		},
	)
	adherent.NewAdherent(
		&adherent.Adherent{
			Firstname:    "Test2",
			Lastname:     "Tutu",
			Email:        "toto@toto.fr",
			Dateofbirth:  "1995-04-24",
			Situation:    "worker",
			University:   "Nancy",
			Homeland:     "Pologne",
			Speakabout:   "Facebook",
			Newsletter:   false,
			AdhesionDate: "1995-04-24",
		},
	)
	adherent.NewAdherent(
		&adherent.Adherent{
			Firstname:    "Test3",
			Lastname:     "Tutu",
			Email:        "toto@toto.fr",
			Dateofbirth:  "1995-04-24",
			Situation:    "worker",
			University:   "Nancy",
			Homeland:     "Pologne",
			Speakabout:   "Facebook",
			Newsletter:   false,
			AdhesionDate: "1995-04-24",
		},
	)
	adherent.NewAdherent(
		&adherent.Adherent{
			Firstname:    "Test4",
			Lastname:     "Tutu",
			Email:        "toto@toto.fr",
			Dateofbirth:  "1995-04-24",
			Situation:    "worker",
			University:   "Nancy",
			Homeland:     "Pologne",
			Speakabout:   "Facebook",
			Newsletter:   false,
			AdhesionDate: "1995-04-24",
		},
	)
	adherent.NewAdherent(
		&adherent.Adherent{
			Firstname:    "Test2",
			Lastname:     "Tutu",
			Email:        "toto@toto.fr",
			Dateofbirth:  "1995-04-24",
			Situation:    "worker",
			University:   "Nancy",
			Homeland:     "Pologne",
			Speakabout:   "Facebook",
			Newsletter:   false,
			AdhesionDate: "1995-04-24",
		},
	)
	adherent.NewAdherent(
		&adherent.Adherent{
			Firstname:    "Test3",
			Lastname:     "Tutu",
			Email:        "toto@toto.fr",
			Dateofbirth:  "1995-04-24",
			Situation:    "worker",
			University:   "Nancy",
			Homeland:     "Pologne",
			Speakabout:   "Facebook",
			Newsletter:   false,
			AdhesionDate: "1995-04-24",
		},
	)
	adherent.NewAdherent(
		&adherent.Adherent{
			Firstname:    "Test4",
			Lastname:     "Tutu",
			Email:        "toto@toto.fr",
			Dateofbirth:  "1995-04-24",
			Situation:    "worker",
			University:   "Nancy",
			Homeland:     "Pologne",
			Speakabout:   "Facebook",
			Newsletter:   false,
			AdhesionDate: "1995-04-24",
		},
	)

	volunteer.CreateVolunteersTable()
	volunteer.NewVolunteer(
		&volunteer.Volunteer{
			Firstname:     "Ash",
			Lastname:      "Ketchum",
			Email:         "ash.ketchum@master.com",
			Discord:       "TheVeryBest",
			Phone:         "0123456789",
			University:    "Indigo",
			PostalAddress: "01 bis house, 00001 Pallet Town, Kanto",
			Actif:         true,
			Bureau:        false,
			HrStatus:      "volunteer",
			StartedDate:   "1997-04-01",
		},
	)
	volunteer.NewVolunteer(
		&volunteer.Volunteer{
			Firstname:     "Gary",
			Lastname:      "Oak",
			Email:         "gary.oak@rival.kt",
			Discord:       "TheRival",
			Phone:         "0123456789",
			University:    "Indigo",
			PostalAddress: "02 house, 00001 Pallet Town, Kanto",
			Actif:         true,
			Bureau:        false,
			HrStatus:      "volunteer",
			StartedDate:   "1997-04-01",
		},
	)
	volunteer.NewVolunteer(
		&volunteer.Volunteer{
			Firstname:     "Cynthia",
			Lastname:      "Shirona",
			Email:         "master@league.si",
			Discord:       "HistoryLover",
			Phone:         "0123456789",
			University:    "US",
			PostalAddress: "Salle du fond, 00019 Conseil 4, Sinnoh",
			Actif:         true,
			Bureau:        true,
			HrStatus:      "volunteer",
			StartedDate:   "1990-04-01",
		},
	)
	volunteer.NewVolunteer(
		&volunteer.Volunteer{
			Firstname:     "Peter",
			Lastname:      "Wataru",
			Email:         "master@indigo.kt",
			Discord:       "Dragoooooon",
			Phone:         "0123456789",
			University:    "Indigo",
			PostalAddress: "Salle du fond, 00019 League Indigo, Kanto",
			Actif:         true,
			Bureau:        true,
			HrStatus:      "volunteer",
			StartedDate:   "1991-12-01",
		},
	)
	volunteer.NewVolunteer(
		&volunteer.Volunteer{
			Firstname:     "Pierre",
			Lastname:      "Takeshi",
			Email:         "pierre@argenta.kt",
			Discord:       "TakeCare",
			Phone:         "0123456789",
			University:    "Argenta",
			PostalAddress: "Arena, 00003 Argenta, Kanto",
			Actif:         true,
			Bureau:        true,
			HrStatus:      "volunteer",
			StartedDate:   "1988-04-01",
		},
	)
	volunteer.NewVolunteer(
		&volunteer.Volunteer{
			Firstname:     "Ondine",
			Lastname:      "Kasumi",
			Email:         "ondine@azuria.kt",
			Discord:       "TheBeautiful",
			Phone:         "0123456789",
			University:    "Azuria",
			PostalAddress: "Arena, 00004 Azuria, Kanto",
			Actif:         true,
			Bureau:        true,
			HrStatus:      "volunteer",
			StartedDate:   "1989-04-01",
		},
	)
	volunteer.NewVolunteer(
		&volunteer.Volunteer{
			Firstname:     "Prof.",
			Lastname:      "Oak",
			Email:         "prof.oak@kanto.kt",
			Discord:       "Prof",
			Phone:         "0123456789",
			University:    "Indigo",
			PostalAddress: "01 house, 00001 Pallet Town, Kanto",
			Actif:         false,
			Bureau:        false,
			HrStatus:      "volunteer",
			StartedDate:   "1970-04-01",
		},
	)
	volunteer.NewVolunteer(
		&volunteer.Volunteer{
			Firstname:     "Prof.",
			Lastname:      "Orm",
			Email:         "prof.orm@johto.jt",
			Discord:       "EggMaster",
			Phone:         "0123456789",
			University:    "Indigo",
			PostalAddress: "01 bis house, 00001 Bourg Geon, Johto",
			Actif:         false,
			Bureau:        false,
			HrStatus:      "volunteer",
			StartedDate:   "1980-04-01",
		},
	)
	volunteer.NewVolunteer(
		&volunteer.Volunteer{
			Firstname:     "Flora",
			Lastname:      "Saphir",
			Email:         "flora@hoenn.com",
			Discord:       "Gobouuuu",
			Phone:         "0123456789",
			University:    "Bourg-en-Vol",
			PostalAddress: "01 bis house, 00001 Bourg-en-Vol, Hoenn",
			Actif:         true,
			Bureau:        true,
			HrStatus:      "volunteer",
			StartedDate:   "2000-04-01",
		},
	)
	volunteer.NewVolunteer(
		&volunteer.Volunteer{
			Firstname:     "President",
			Lastname:      "Shehroz",
			Email:         "president@league.ga",
			Discord:       "Pres",
			Phone:         "+49 0123456789",
			University:    "Macro Cosmos",
			PostalAddress: "Macro Cosmos site, 00008 Winscor, Galar",
			Actif:         true,
			Bureau:        false,
			HrStatus:      "volunteer",
			StartedDate:   "2023-01-01",
		},
	)

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
