package event

import (
	"log"
	"testing"
)

func setUpModel() {
	NewEvent(&Event{Name: "Voyage a Hawai", Date: "2023-04-23", Location: "Hawai", Type: "voyage", Price: 120.42, Url: "facebook.com/voyageHawai", Actif: true})
	NewEvent(&Event{Name: "Saturday Night Fever", Date: "2023-05-23", Location: "3 rue Albert 1er, 69000 Lyon", Type: "soiree", Price: 20, Url: "facebook.com/sturdayfever", Actif: false})
}

func TestNewEvent(t *testing.T) {
	NewEvent(&Event{Name: "Barathon", Date: "2023-05-12", Location: "Hotel de Ville, 69000 Lyon", Type: "soiree", Price: 20, Url: "facebook.com/barathon", Actif: true})
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

func TestUpdateEvent(t *testing.T) {
	eve := Event{Id: 2, Name: "Barathon", Date: "2023-05-12", Location: "Hotel de Ville, 69000 Lyon", Type: "soiree", Price: 0, Url: "facebook.com/barathon", Actif: true}
	UpdateEvent(&eve)

	eve_2 := FindEventById(2)
	if eve_2.Price != 0 {
		log.Fatal("Eve_3 didn't updated !")
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
