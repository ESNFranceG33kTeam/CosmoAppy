package volunteer

import (
	"log"
	"testing"
)

func setUpModel() {
	NewVolunteer(
		&Volunteer{
			Firstname:     "Toto",
			Lastname:      "Bob",
			Email:         "toto.toto@bob.com",
			Discord:       "Totot",
			Phone:         "0123456789",
			University:    "UBSB",
			PostalAddress: "32 rue du Soleil, 00001 Mars",
			Actif:         true,
			Bureau:        false,
			HrStatus:      "employee",
			StartedDate:   "2016-04-14",
		},
	)
	NewVolunteer(
		&Volunteer{
			Firstname:     "Tata",
			Lastname:      "Bob",
			Email:         "toto.toto@bob.com",
			Discord:       "Totot",
			Phone:         "0123456789",
			University:    "UBSB",
			PostalAddress: "32 rue du Soleil, 00001 Mars",
			Actif:         false,
			Bureau:        false,
			HrStatus:      "scv",
			StartedDate:   "2016-04-14",
		},
	)
	NewVolunteer(
		&Volunteer{
			Firstname:     "Toto",
			Lastname:      "Bob",
			Email:         "toto.toto@bob.com",
			Discord:       "Totot",
			Phone:         "0123456789",
			University:    "UBSB",
			PostalAddress: "32 rue du Soleil, 00001 Mars",
			Actif:         true,
			Bureau:        true,
			HrStatus:      "volunteer",
			StartedDate:   "2016-04-14",
		},
	)
}

func TestNewVolunteer(t *testing.T) {
	NewVolunteer(
		&Volunteer{
			Firstname:     "Toto",
			Lastname:      "Bob",
			Email:         "toto.toto@bob.com",
			Discord:       "Totot",
			Phone:         "0123456789",
			University:    "UBSB",
			PostalAddress: "32 rue du Soleil, 00001 Mars",
			Actif:         false,
			Bureau:        true,
			HrStatus:      "volunteer",
			StartedDate:   "2016-04-14",
		},
	)
}

func TestFindVolunteerById(t *testing.T) {
	vlt := FindVolunteerById(2)

	if vlt.Firstname != "Tata" {
		log.Fatal("Volunteer is not the good one.")
	}
}

func TestAllVolunteers(t *testing.T) {
	vlts := AllVolunteers()

	//log.Println(adhs)
	if len(*vlts) == 0 {
		log.Fatal("Volunteer is empty")
	}
}

func TestUpdateVolunteer(t *testing.T) {
	vlt := Volunteer{
		Id:            3,
		Firstname:     "Titi",
		Lastname:      "titi",
		Email:         "titi@titi.com",
		Discord:       "Totot",
		Phone:         "0123456789",
		University:    "UBSB",
		PostalAddress: "32 rue du Soleil, 00001 Mars",
		Actif:         false,
		Bureau:        false,
		HrStatus:      "volunteer",
		StartedDate:   "2016-04-14",
	}
	UpdateVolunteer(&vlt)

	vlt_3 := FindVolunteerById(3)
	if vlt_3.Actif != false {
		log.Fatal("Volunteer_3 didn't updated !")
	}
}

func TestDeleteVolunteerById(t *testing.T) {
	err := DeleteVolunteerById(3)

	if err != nil {
		log.Fatal("Delete had a problem : ", err)
	}

	vlts := AllVolunteers()

	for _, vlt := range *vlts {
		if vlt.Id == 3 {
			log.Fatal("Volunteer_3 didn't be removed !")
		}
	}
}
