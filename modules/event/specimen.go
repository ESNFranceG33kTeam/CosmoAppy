package event

func Specimen() {
	NewEvent(&Event{Name: "Voyage a Hawai", Date: "2023-04-23", Location: "Hawai", Type: "voyage", Price: 120.42, Url: "facebook.com/voyageHawai", Actif: true})
	NewEvent(&Event{Name: "Saturday Night Fever", Date: "2023-05-23", Location: "3 rue Albert 1er, 69000 Lyon", Type: "soiree", Price: 20, Url: "facebook.com/sturdayfever", Actif: false})
	NewAttendee(&Attendee{Id_event: 1, Id_adherent: 2, Staff: true})
	NewAttendee(&Attendee{Id_event: 2, Id_adherent: 3, Staff: false})
}
