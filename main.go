package main

import (
	"flag"
	"net/http"

	"github.com/ESNFranceG33kTeam/sAPI/database"
	"github.com/ESNFranceG33kTeam/sAPI/helpers"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/modules/adherent"
	"github.com/ESNFranceG33kTeam/sAPI/modules/esncard"
	"github.com/ESNFranceG33kTeam/sAPI/modules/health"
	"github.com/ESNFranceG33kTeam/sAPI/modules/money"
	"github.com/ESNFranceG33kTeam/sAPI/modules/volunteer"
	"github.com/ESNFranceG33kTeam/sAPI/router"
)

func startoptions() {
	flag.StringVar(&helpers.Confpathflag, "conf", "test/conf_local.yaml", "path for the configuration file.")
	flag.StringVar(&helpers.Swaggerpathflag, "swagger", "/test/swagger.yaml", "path for the swagger file.")
	flag.Parse()
}

func InitConf() {
	helpers.InitFile()
	helpers.ReadConfig()
	logger.LogInit()
}

func InitDb() {
	database.DatabaseInit()
	adherent.CreateAdherentsTable()
	money.CreateMoneysTable()
}

func InitRouter() {
	router.InitializeRouter()
	adherent.InitRoutes()
	volunteer.InitRoutes()
	esncard.InitRoutes()
	money.InitRoutes()

	// always last
	health.InitRoutes()
}

func main() {
	startoptions()
	InitConf()
	InitDb()
	logger.LogInfo("main", "Conf loaded ; app starting...")

	InitRouter()

	// Specimen data
	if helpers.AppConfig.Specimen {
		PopulateDb()
	}

	logger.LogInfo("main", "API ready.")

	logger.LogCritical("main", "listen error", http.ListenAndServe(":8080", router.GetRouter()))

}
