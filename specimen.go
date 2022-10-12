package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/models"
)

func PopulateDb() {
	datedate, _ := time.Parse("2006-01-02", "1993-09-23")
	dateofbirth := datedate.Format("2006-01-02")

	// Populate database
	// Adherents
	models.NewAdherent(&models.Adherent{Firstname: "Titi", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: dateofbirth, Student: false, University: "UBFC", Homeland: "Mexique", Speakabout: "Twitter", Newsletter: false})
	// ESNcards
	rand.Seed(time.Now().UnixNano())
	rand_numb := strconv.Itoa(rand.Intn(9999999-1000000) + 1000000)
	models.NewESNcard(&models.ESNcard{Id_adherent: 1, Esncard: rand_numb})

	logger.LogInfo("Specimen", "Specimen data charged up.")
}
