package event

func Specimen() {
	NewEvent(&Event{Name: "League Indigo", Date: "2023-04-23", Location: "Plateau Indigo", NbSpotsMax: 30, NbSpotsTaken: 0, Type: "championship", Price: 120, Url: "https://www.pokepedia.fr/Ligue_Indigo", Actif: true})
	NewEvent(&Event{Name: "Masters tournament", Date: "2023-04-30", Location: "Winscor", NbSpotsMax: 8, NbSpotsTaken: 0, Type: "championship", Price: 80, Url: "https://www.pokepedia.fr/Ligue_Indigo", Actif: true})

	NewAttendee(&Attendee{Id_event: 1, Id_adherent: 1, Staff: false})
	NewAttendee(&Attendee{Id_event: 1, Id_adherent: 4, Staff: true})
	NewAttendee(&Attendee{Id_event: 1, Id_adherent: 7, Staff: false})
	NewAttendee(&Attendee{Id_event: 1, Id_adherent: 2, Staff: false})
	NewAttendee(&Attendee{Id_event: 2, Id_adherent: 1, Staff: false})
	NewAttendee(&Attendee{Id_event: 2, Id_adherent: 7, Staff: false})
	NewAttendee(&Attendee{Id_event: 2, Id_adherent: 3, Staff: true})
	NewAttendee(&Attendee{Id_event: 2, Id_adherent: 4, Staff: true})

}
