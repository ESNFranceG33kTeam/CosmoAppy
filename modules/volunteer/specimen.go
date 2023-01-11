package volunteer

func Specimen() {
	NewVolunteer(&Volunteer{Firstname: "Ash", Lastname: "Ketchum", Email: "ash.ketchum@master.com", Actif: true, Bureau: false})
	NewVolunteer(&Volunteer{Firstname: "Gary", Lastname: "Oak", Email: "gary.oak@rival.kt", Actif: true, Bureau: false})

	NewVolunteer(&Volunteer{Firstname: "Cynthia", Lastname: "Shirona", Email: "master@league.si", Actif: true, Bureau: true})
	NewVolunteer(&Volunteer{Firstname: "Peter", Lastname: "Wataru", Email: "master@indigo.kt", Actif: true, Bureau: true})
	NewVolunteer(&Volunteer{Firstname: "Pierre", Lastname: "Takeshi", Email: "pierre@argenta.kt", Actif: true, Bureau: true})
	NewVolunteer(&Volunteer{Firstname: "Ondine", Lastname: "Kasumi", Email: "ondine@azuria.kt", Actif: true, Bureau: true})

	NewVolunteer(&Volunteer{Firstname: "Prof.", Lastname: "Oak", Email: "prof.oak@kanto.kt", Actif: false, Bureau: false})
	NewVolunteer(&Volunteer{Firstname: "Prof.", Lastname: "Orm", Email: "prof.orm@johto.jt", Actif: false, Bureau: false})

	NewVolunteer(&Volunteer{Firstname: "Flora", Lastname: "Saphir", Email: "flora@hoenn.com", Actif: true, Bureau: true})
}
