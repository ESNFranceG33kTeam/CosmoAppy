package volunteer

func Specimen() {
	NewVolunteer(
		&Volunteer{
			Firstname:     "Ash",
			Lastname:      "Ketchum",
			Email:         "ash.ketchum@master.com",
			Discord:       "TheVeryBest",
			Phone:         "+33 0123456789",
			University:    "Indigo",
			PostalAddress: "01 bis house, 00001 Pallet Town, Kanto",
			Actif:         true,
			Bureau:        false,
			HrStatus:      "volunteer",
			StartedDate:   "1997-04-01",
		},
	)
	NewVolunteer(
		&Volunteer{
			Firstname:     "Gary",
			Lastname:      "Oak",
			Email:         "gary.oak@rival.kt",
			Discord:       "TheRival",
			Phone:         "+33 0123456789",
			University:    "Indigo",
			PostalAddress: "02 house, 00001 Pallet Town, Kanto",
			Actif:         true,
			Bureau:        false,
			HrStatus:      "volunteer",
			StartedDate:   "1997-04-01",
		},
	)
	NewVolunteer(
		&Volunteer{
			Firstname:     "Flora",
			Lastname:      "Saphir",
			Email:         "flora@hoenn.com",
			Discord:       "Gobouuuu",
			Phone:         "+44 0123456789",
			University:    "Bourg-en-Vol",
			PostalAddress: "01 bis house, 00001 Bourg-en-Vol, Hoenn",
			Actif:         true,
			Bureau:        false,
			HrStatus:      "volunteer",
			StartedDate:   "2000-04-01",
		},
	)

	NewVolunteer(
		&Volunteer{
			Firstname:     "Cynthia",
			Lastname:      "Shirona",
			Email:         "master@league.si",
			Discord:       "HistoryLover",
			Phone:         "+49 0123456789",
			University:    "US",
			PostalAddress: "Salle du fond, 00019 Conseil 4, Sinnoh",
			Actif:         true,
			Bureau:        true,
			HrStatus:      "volunteer",
			StartedDate:   "1990-04-01",
		},
	)
	NewVolunteer(
		&Volunteer{
			Firstname:     "Peter",
			Lastname:      "Wataru",
			Email:         "master@indigo.kt",
			Discord:       "Dragoooooon",
			Phone:         "+34 0123456789",
			University:    "Indigo",
			PostalAddress: "Salle du fond, 00019 League Indigo, Kanto",
			Actif:         true,
			Bureau:        true,
			HrStatus:      "volunteer",
			StartedDate:   "1991-12-01",
		},
	)
	NewVolunteer(
		&Volunteer{
			Firstname:     "Pierre",
			Lastname:      "Takeshi",
			Email:         "pierre@argenta.kt",
			Discord:       "TakeCare",
			Phone:         "+33 0123456789",
			University:    "Argenta",
			PostalAddress: "Arena, 00003 Argenta, Kanto",
			Actif:         true,
			Bureau:        true,
			HrStatus:      "volunteer",
			StartedDate:   "1988-04-01",
		},
	)
	NewVolunteer(
		&Volunteer{
			Firstname:     "Ondine",
			Lastname:      "Kasumi",
			Email:         "ondine@azuria.kt",
			Discord:       "TheBeautiful",
			Phone:         "+33 0123456789",
			University:    "Azuria",
			PostalAddress: "Arena, 00004 Azuria, Kanto",
			Actif:         true,
			Bureau:        true,
			HrStatus:      "volunteer",
			StartedDate:   "1989-04-01",
		},
	)

	NewVolunteer(
		&Volunteer{
			Firstname:     "Prof.",
			Lastname:      "Oak",
			Email:         "prof.oak@kanto.kt",
			Discord:       "Prof",
			Phone:         "+33 0123456789",
			University:    "Indigo",
			PostalAddress: "01 house, 00001 Pallet Town, Kanto",
			Actif:         false,
			Bureau:        false,
			HrStatus:      "volunteer",
			StartedDate:   "1970-04-01",
		},
	)
	NewVolunteer(
		&Volunteer{
			Firstname:     "Prof.",
			Lastname:      "Orm",
			Email:         "prof.orm@johto.jt",
			Discord:       "EggMaster",
			Phone:         "+34 0123456789",
			University:    "Indigo",
			PostalAddress: "01 bis house, 00001 Bourg Geon, Johto",
			Actif:         false,
			Bureau:        false,
			HrStatus:      "volunteer",
			StartedDate:   "1980-04-01",
		},
	)

	NewVolunteer(
		&Volunteer{
			Firstname:     "President",
			Lastname:      "Shehroz",
			Email:         "president@league.ga",
			Discord:       "Pres",
			Phone:         "+49 0123456789",
			University:    "Macro Cosmos",
			PostalAddress: "Macro Cosmos site, 00008 Winscor, Galar",
			Actif:         true,
			Bureau:        false,
			HrStatus:      "employee",
			StartedDate:   "2023-01-01",
		},
	)
}
