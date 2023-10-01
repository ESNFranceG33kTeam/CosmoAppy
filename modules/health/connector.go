package health

import (
	"database/sql"

	"github.com/ESNFranceG33kTeam/CosmoAppy/database"
	"github.com/ESNFranceG33kTeam/CosmoAppy/docs"
	"github.com/ESNFranceG33kTeam/CosmoAppy/helpers"
	"github.com/ESNFranceG33kTeam/CosmoAppy/logger"
	"github.com/ESNFranceG33kTeam/CosmoAppy/router"
	"github.com/gorilla/mux"
)

func Launcher() {
	InitRoutes()
}

func TheLogger() *logger.Logger {
	return logger.GetLogger()
}

func TheRouter() *mux.Router {
	return router.GetRouter()
}

func TheSecureRouter() *mux.Router {
	return router.GetSecureRouter()
}

func TheAppConfig() *helpers.Cfg {
	return helpers.TheAppConfig()
}

func TheDb() *sql.DB {
	return database.Db()
}

func TheVersion() string {
	return docs.GetVersion()
}
