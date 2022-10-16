package main

import (
	"flag"
	"net/http"

	"github.com/ESNFranceG33kTeam/sAPI/config"
	"github.com/ESNFranceG33kTeam/sAPI/helpers"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/modules/money"
	"github.com/ESNFranceG33kTeam/sAPI/router"
)

func startoptions() {
	flag.StringVar(&helpers.Confpathflag, "conf", "test/conf_local.yaml", "path for the configuration file.")
	flag.StringVar(&helpers.Swaggerpathflag, "swagger", "/test/swagger.yaml", "path for the swagger file.")
	flag.Parse()
}

func InitConf() {
	helpers.InitFile(helpers.Confpathflag)
	helpers.ReadConfig()
	logger.LogInit(helpers.AppConfig.Loglevel)
	config.DatabaseInit(helpers.AppConfig.Userdb, helpers.AppConfig.Passdb, helpers.AppConfig.Ipdb, helpers.AppConfig.Portdb, helpers.AppConfig.Namedb, helpers.AppConfig.Extradb)
}

func CreateDbTables() {
	money.CreateMoneysTable()
}

func CreateRoutes() {
	money.InitRoutes()
}

func main() {
	startoptions()
	InitConf()
	CreateDbTables()
	logger.LogInfo("main", "Conf loaded ; app starting...")

	router.InitializeRouter()
	CreateRoutes()

	// Specimen data
	if helpers.AppConfig.Specimen {
		PopulateDb()
	}

	logger.LogInfo("main", "API ready.")

	logger.LogCritical("main", "listen error", http.ListenAndServe(":8080", router.GetRouter()))

}
