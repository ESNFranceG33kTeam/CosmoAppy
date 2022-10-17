package volunteer

import (
	"log"
	"testing"
)

func setUpModel() {
	NewVolunteer(&Volunteer{Id_adherent: 2, Actif: true, Bureau: false})
	NewVolunteer(&Volunteer{Id_adherent: 3, Actif: false, Bureau: false})
	NewVolunteer(&Volunteer{Id_adherent: 4, Actif: true, Bureau: true})
}

func TestNewVolunteer(t *testing.T) {
	NewVolunteer(&Volunteer{Id_adherent: 1, Actif: false, Bureau: true})
}

func TestFindVolunteerByIdadherent(t *testing.T) {
	vlt := FindVolunteerByIdadherent(4)

	if vlt.Id_adherent != 4 {
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
	vlt := Volunteer{Id: 3, Id_adherent: 4, Actif: false, Bureau: false}
	UpdateVolunteer(&vlt)

	vlt_3 := FindVolunteerByIdadherent(4)
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
