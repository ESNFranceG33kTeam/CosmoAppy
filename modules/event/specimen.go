package event

func Specimen() {
	NewEvent(
		&Event{
			Name:         "League Indigo",
			Date:         "2023-04-23",
			Location:     "Plateau Indigo",
			NbSpotsMax:   30,
			NbSpotsTaken: 4,
			Type:         "championship",
			Price:        120,
			Url:          "https://www.pokepedia.fr/Ligue_Indigo",
			Actif:        true,
		},
	)
	NewEvent(
		&Event{
			Name:         "Masters tournament",
			Date:         "2023-04-30",
			Location:     "Winscor",
			NbSpotsMax:   8,
			NbSpotsTaken: 4,
			Type:         "championship",
			Price:        80,
			Url:          "https://www.pokepedia.fr/Ligue_Indigo",
			Actif:        true,
		},
	)
	NewEvent(
		&Event{
			Name:         "Qualif Masters",
			Date:         "2023-03-30",
			Location:     "Sinnoh",
			NbSpotsMax:   20,
			NbSpotsTaken: 19,
			Type:         "championship",
			Price:        50,
			Url:          "https://www.pokepedia.fr/Ligue_Indigo",
			Actif:        false,
		},
	)

	NewAttendee(&Attendee{Id_event: 1, Id_adherent: 1})
	NewAttendee(&Attendee{Id_event: 1, Id_adherent: 7})
	NewAttendee(&Attendee{Id_event: 1, Id_adherent: 2})
	NewAttendee(&Attendee{Id_event: 2, Id_adherent: 1})
	NewAttendee(&Attendee{Id_event: 2, Id_adherent: 7})

	NewStaff(&Staff{Id_event: 1, Id_volunteer: 3})
	NewStaff(&Staff{Id_event: 1, Id_volunteer: 4})
	NewStaff(&Staff{Id_event: 1, Id_volunteer: 5})
	NewStaff(&Staff{Id_event: 2, Id_volunteer: 3})
	NewStaff(&Staff{Id_event: 2, Id_volunteer: 4})
	NewStaff(&Staff{Id_event: 2, Id_volunteer: 5})
	NewStaff(&Staff{Id_event: 3, Id_volunteer: 3})
	NewStaff(&Staff{Id_event: 3, Id_volunteer: 4})
	NewStaff(&Staff{Id_event: 3, Id_volunteer: 5})

	AutoNewMonthlyStat("2023-04")

}
