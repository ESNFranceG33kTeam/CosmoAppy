package main

import (
	"time"

	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/models"
)

func PopulateDb() {
	datedate, _ := time.Parse("2006-01-02", "1993-09-23")
	dateofbirth := datedate.Format("2006-01-02")

	// Populate database
	models.NewAdherent(&models.Adherent{Firstname: "Titi", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: dateofbirth, ESNcard: "grgerrbrbreht", Student: false, University: "UBFC", Homeland: "Mexique", Speakabout: "Twitter", Newsletter: false})

	logger.LogInfo("Specimen", "Specimen data charged up.")
}
