package adherent

import "time"

func Specimen() {
	datedate, _ := time.Parse("2006-01-02", "1993-09-23")
	dateofbirth := datedate.Format("2006-01-02")

	NewAdherent(&Adherent{Firstname: "Titi", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: dateofbirth, Student: false, University: "UBFC", Homeland: "Mexique", Speakabout: "Twitter", Newsletter: false})
}
