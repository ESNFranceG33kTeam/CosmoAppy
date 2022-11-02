package volunteer

import (
	"database/sql"

	"github.com/ESNFranceG33kTeam/CosmoAppy/database"
	"github.com/ESNFranceG33kTeam/CosmoAppy/logger"
	"github.com/ESNFranceG33kTeam/CosmoAppy/router"
	"github.com/gorilla/mux"
)

func Launcher() {
	CreateVolunteersTable()
	InitRoutes()
}

func TheLogger() *logger.Logger {
	return logger.GetLogger()
}

func TheDb() *sql.DB {
	return database.Db()
}

func TheRouter() *mux.Router {
	return router.GetRouter()
}

func TheSecureRouter() *mux.Router {
	return router.GetSecureRouter()
}
