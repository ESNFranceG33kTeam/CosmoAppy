package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ESNFranceG33kTeam/sAPI/database"
	"github.com/ESNFranceG33kTeam/sAPI/helpers"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/modules/launcher"
	"github.com/ESNFranceG33kTeam/sAPI/router"
)

func startoptions() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("Abs path doesn't exist !")
	}

	flag.StringVar(&helpers.Confpathflag, "conf", filepath.Join(path, "test/conf_local.yaml"), "path for the configuration file.")
	flag.StringVar(&helpers.Swaggerpathflag, "swagger", "/swagger.yaml", "relative path for the swagger file.")

	flag.Parse()
}

func InitConf() {
	helpers.InitFile()
	helpers.ReadConfig()
	logger.GetLogger().LogInit()
}

func main() {
	startoptions()
	InitConf()
	database.DatabaseInit()
	router.InitializeRouter()

	// Loading modules
	launcher.LauncherModules()

	logger.GetLogger().LogInfo("main", "Conf and modules loaded ; app starting...")

	// Specimen data
	if helpers.TheAppConfig().Specimen {
		PopulateDb()
	}

	logger.GetLogger().LogInfo("main", "API ready.")

	logger.GetLogger().LogCritical("main", "listen error", http.ListenAndServe(":8080", router.GetRouter()))

}
