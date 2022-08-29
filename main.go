package main

import (
	"net/http"

	"github.com/ESNFranceG33kTeam/sAPI/config"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/models"
)

func main() {
	config.DatabaseInit()
	router := InitializeRouter()

	// Populate database
	models.NewAdherent(&models.Adherent{Firstname: "Titi", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", ESNcard: "grgerrbrbreht", Student: false, University: "UBFC", Homeland: "Mexique", Speakabout: "Twitter", Newsletter: false})

	logger.LogInfo("main", "API starting.")

	logger.LogCritical("main", "listen error", http.ListenAndServe(":8080", router))

}
