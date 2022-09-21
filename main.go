package main

import (
	"net/http"

	"github.com/ESNFranceG33kTeam/sAPI/config"
	"github.com/ESNFranceG33kTeam/sAPI/helpers"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/models"
)

func InitConf() {
	helpers.InitFile("conf.yml")
	helpers.ReadConfig()
	logger.LogInit(helpers.AppConfig.Loglevel)
	config.DatabaseInit(helpers.AppConfig.Userdb, helpers.AppConfig.Passdb, helpers.AppConfig.Ipdb, helpers.AppConfig.Portdb, helpers.AppConfig.Namedb, helpers.AppConfig.Extradb)
}

func main() {
	InitConf()
	logger.LogInfo("main", "Conf loaded ; app starting.")
	router := InitializeRouter(helpers.AppConfig.Usersapi, helpers.AppConfig.Tokensapi)

	// Populate database
	models.NewAdherent(&models.Adherent{Firstname: "Titi", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", ESNcard: "grgerrbrbreht", Student: false, University: "UBFC", Homeland: "Mexique", Speakabout: "Twitter", Newsletter: false})

	logger.LogInfo("main", "API ready.")

	logger.LogCritical("main", "listen error", http.ListenAndServe(":8080", router))

}
