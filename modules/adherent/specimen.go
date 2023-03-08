package adherent

import "time"

func dateofbirth(datebirth string) string {
	datedate, _ := time.Parse("2006-01-02", datebirth)
	dateofbirth := datedate.Format("2006-01-02")

	return dateofbirth
}

func Specimen() {
	NewAdherent(
		&Adherent{
			Firstname:    "Ash",
			Lastname:     "Ketchum",
			Email:        "ash.ketchum@master.com",
			Dateofbirth:  dateofbirth("1987-05-22"),
			Situation:    "student",
			University:   "Indigo",
			Homeland:     "Kanto",
			Speakabout:   "Twitter",
			Newsletter:   true,
			AdhesionDate: dateofbirth("2022-04-24"),
		},
	)
	NewAdherent(
		&Adherent{
			Firstname:    "Gary",
			Lastname:     "Oak",
			Email:        "gary.oak@rival.kt",
			Dateofbirth:  dateofbirth("1984-04-06"),
			Situation:    "internship",
			University:   "Indigo",
			Homeland:     "Kanto",
			Speakabout:   "Twitter",
			Newsletter:   true,
			AdhesionDate: dateofbirth("2022-01-24"),
		},
	)

	NewAdherent(
		&Adherent{
			Firstname:    "Cynthia",
			Lastname:     "Shirona",
			Email:        "master@league.si",
			Dateofbirth:  dateofbirth("1980-08-18"),
			Situation:    "worker",
			University:   "US",
			Homeland:     "Sinnoh",
			Speakabout:   "Twitter",
			Newsletter:   false,
			AdhesionDate: dateofbirth("2022-04-24"),
		},
	)
	NewAdherent(
		&Adherent{
			Firstname:    "Peter",
			Lastname:     "Wataru",
			Email:        "master@indigo.kt",
			Dateofbirth:  dateofbirth("1978-10-15"),
			Situation:    "worker",
			University:   "Indigo",
			Homeland:     "Kanto",
			Speakabout:   "Twitter",
			Newsletter:   false,
			AdhesionDate: dateofbirth("2022-04-24"),
		},
	)
	NewAdherent(
		&Adherent{
			Firstname:    "Pierre",
			Lastname:     "Takeshi",
			Email:        "pierre@argenta.kt",
			Dateofbirth:  dateofbirth("1981-01-16"),
			Situation:    "worker",
			University:   "Argenta",
			Homeland:     "Kanto",
			Speakabout:   "Facebook",
			Newsletter:   false,
			AdhesionDate: dateofbirth("2022-04-24"),
		},
	)
	NewAdherent(
		&Adherent{
			Firstname:    "Ondine",
			Lastname:     "Kasumi",
			Email:        "ondine@azuria.kt",
			Dateofbirth:  dateofbirth("1982-01-19"),
			Situation:    "worker",
			University:   "Azuria",
			Homeland:     "Kanto",
			Speakabout:   "Facebook",
			Newsletter:   false,
			AdhesionDate: dateofbirth("2022-04-24"),
		},
	)

	NewAdherent(
		&Adherent{
			Firstname:    "Pikachu",
			Lastname:     "Pika",
			Email:        "pikachu@pika.pkm",
			Dateofbirth:  dateofbirth("1990-01-19"),
			Situation:    "student",
			University:   "Bourg Palette",
			Homeland:     "Kanto",
			Speakabout:   "Event",
			Newsletter:   true,
			AdhesionDate: dateofbirth("2022-04-24"),
		},
	)
	NewAdherent(
		&Adherent{
			Firstname:    "Pichu",
			Lastname:     "Pika",
			Email:        "pichu@pika.pkm",
			Dateofbirth:  dateofbirth("1992-02-19"),
			Situation:    "internship",
			University:   "Bourg Geon",
			Homeland:     "Johto",
			Speakabout:   "Event",
			Newsletter:   true,
			AdhesionDate: dateofbirth("2022-04-24"),
		},
	)
	NewAdherent(
		&Adherent{
			Firstname:    "Gobou",
			Lastname:     "Gob",
			Email:        "gobou@bog.pkm",
			Dateofbirth:  dateofbirth("1991-02-12"),
			Situation:    "student",
			University:   "Bourg-en-Vol",
			Homeland:     "Hoenn",
			Speakabout:   "Event",
			Newsletter:   true,
			AdhesionDate: dateofbirth("2022-04-24"),
		},
	)

}
