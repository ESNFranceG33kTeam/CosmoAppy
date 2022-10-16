package main

import (
	"flag"
	"net/http"

	"github.com/ESNFranceG33kTeam/sAPI/config"
	"github.com/ESNFranceG33kTeam/sAPI/helpers"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/modules/money"
)

// flags
var confpathflag string
var swaggerpathflag string

func startoptions() {
	flag.StringVar(&confpathflag, "conf", "test/conf_local.yaml", "path for the configuration file.")
	flag.StringVar(&swaggerpathflag, "swagger", "/test/swagger.yaml", "path for the swagger file.")
	flag.Parse()
}

func InitConf() {
	helpers.InitFile(confpathflag)
	helpers.ReadConfig()
	logger.LogInit(helpers.AppConfig.Loglevel)
	config.DatabaseInit(helpers.AppConfig.Userdb, helpers.AppConfig.Passdb, helpers.AppConfig.Ipdb, helpers.AppConfig.Portdb, helpers.AppConfig.Namedb, helpers.AppConfig.Extradb)
}

func CreateDbTables() {
	money.CreateMoneysTable()
}

func main() {
	startoptions()
	InitConf()
	CreateDbTables()
	logger.LogInfo("main", "Conf loaded ; app starting...")

	router := InitializeRouter(helpers.AppConfig.Usersapi, helpers.AppConfig.Tokensapi)

	// Specimen data
	if helpers.AppConfig.Specimen {
		PopulateDb()
	}

	logger.LogInfo("main", "API ready.")

	logger.LogCritical("main", "listen error", http.ListenAndServe(":8080", router))

}
