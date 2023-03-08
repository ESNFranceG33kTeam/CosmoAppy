package event

import (
	"log"
	"testing"
)

func setUpModel() {
	NewEvent(
		&Event{
			Name:         "Voyage a Hawai",
			Date:         "2023-04-23",
			Location:     "Hawai",
			NbSpotsMax:   30,
			NbSpotsTaken: 0,
			Type:         "voyage",
			Price:        120.42,
			Url:          "facebook.com/voyageHawai",
			Actif:        true,
		},
	)
	NewEvent(
		&Event{
			Name:         "Saturday Night Fever",
			Date:         "2023-05-23",
			Location:     "3 rue Albert 1er, 69000 Lyon",
			NbSpotsMax:   30,
			NbSpotsTaken: 0,
			Type:         "soiree",
			Price:        20,
			Url:          "facebook.com/sturdayfever",
			Actif:        false,
		},
	)
	NewAttendee(&Attendee{Id_event: 1, Id_adherent: 2})
	NewAttendee(&Attendee{Id_event: 1, Id_adherent: 3})
	NewStaff(&Staff{Id_event: 1, Id_volunteer: 2})
	NewStaff(&Staff{Id_event: 1, Id_volunteer: 3})
}

func TestNewEvent(t *testing.T) {
	NewEvent(
		&Event{
			Name:         "Barathon",
			Date:         "2023-05-12",
			Location:     "Hotel de Ville, 69000 Lyon",
			NbSpotsMax:   30,
			NbSpotsTaken: 0,
			Type:         "soiree",
			Price:        20,
			Url:          "facebook.com/barathon",
			Actif:        true,
		},
	)
}

func TestFindEventById(t *testing.T) {
	eve := FindEventById(2)

	if eve.Id != 2 {
		log.Fatal("Event is not him")
	}
}

func TestAllEvents(t *testing.T) {
	eves := AllEvents()

	//log.Println(adhs)
	if len(*eves) == 0 {
		log.Fatal("Event is empty")
	}
}

func TestUpdateSpotsAvaiEvent(t *testing.T) {
	eve := FindEventById(2)
	eve.NbSpotsTaken += 1
	UpdateSpotsTakenEvent(eve)

	eve_2 := FindEventById(2)
	if eve_2.NbSpotsTaken != 1 {
		log.Fatal("Eve_2 Spots available didn't updated !")
	}
}

func TestUpdateEvent(t *testing.T) {
	eve := Event{
		Id:           2,
		Name:         "Barathon",
		Date:         "2023-05-12",
		Location:     "Hotel de Ville, 69000 Lyon",
		NbSpotsMax:   30,
		NbSpotsTaken: 0,
		Type:         "soiree",
		Price:        0,
		Url:          "facebook.com/barathon",
		Actif:        true,
	}
	UpdateEvent(&eve)

	eve_2 := FindEventById(2)
	if eve_2.Price != 0 {
		log.Fatal("Eve_2 didn't updated !")
	}
}

func TestDeleteEventById(t *testing.T) {
	err := DeleteEventById(3)

	if err != nil {
		log.Fatal("Delete had a problem : ", err)
	}

	eves := AllEvents()

	for _, eve := range *eves {
		if eve.Id == 3 {
			log.Fatal("Eve_3 didn't be removed !")
		}
	}
}

func TestNewAttendee(t *testing.T) {
	NewAttendee(&Attendee{Id_event: 2, Id_adherent: 3})
}

func TestAllAttendees(t *testing.T) {
	atts := AllAttendees()

	//log.Println(adhs)
	if len(*atts) == 0 {
		log.Fatal("Attendee is empty")
	}
}

func TestFindAttendeeById(t *testing.T) {
	att := FindAttendeeById(2)

	if att.Id != 2 {
		log.Fatal("Attendee is not him")
	}
}

func TestFindAttendeeByEventId(t *testing.T) {
	atts := FindAttendeeByEventId(1)

	if len(*atts) == 0 {
		log.Fatal("Attendee is empty")
	}
}

func TestFindAttendeeByAdherentId(t *testing.T) {
	atts := FindAttendeeByAdherentId(3)

	if len(*atts) == 0 {
		log.Fatal("Attendee is empty")
	}
}

func TestUpdateAttendee(t *testing.T) {
	att := Attendee{Id: 1, Id_event: 1, Id_adherent: 2}
	UpdateAttendee(&att)
}

func TestDeleteAttendeeById(t *testing.T) {
	err := DeleteAttendeeById(3)

	if err != nil {
		log.Fatal("Delete had a problem : ", err)
	}

	atts := AllAttendees()

	for _, att := range *atts {
		if att.Id == 3 {
			log.Fatal("Att_3 didn't be removed !")
		}
	}
}

func TestNewStaff(t *testing.T) {
	NewStaff(&Staff{Id_event: 2, Id_volunteer: 3})
}

func TestAllStaffs(t *testing.T) {
	stas := AllStaffs()

	//log.Println(adhs)
	if len(*stas) == 0 {
		log.Fatal("Staff is empty")
	}
}

func TestFindStaffById(t *testing.T) {
	sta := FindStaffById(2)

	if sta.Id != 2 {
		log.Fatal("Staff is not him")
	}
}

func TestFindStaffByEventId(t *testing.T) {
	stas := FindStaffByEventId(1)

	if len(*stas) == 0 {
		log.Fatal("Staff is empty")
	}
}

func TestFindStaffByVolunterreId(t *testing.T) {
	stas := FindStaffByVolunteerId(3)

	if len(*stas) == 0 {
		log.Fatal("Staff is empty")
	}
}

func TestUpdateStaff(t *testing.T) {
	sta := Staff{Id: 1, Id_event: 1, Id_volunteer: 2}
	UpdateStaff(&sta)
}

func TestDeleteStaffById(t *testing.T) {
	err := DeleteStaffById(3)

	if err != nil {
		log.Fatal("Delete had a problem : ", err)
	}

	stas := AllStaffs()

	for _, sta := range *stas {
		if sta.Id == 3 {
			log.Fatal("Sta_3 didn't be removed !")
		}
	}
}
