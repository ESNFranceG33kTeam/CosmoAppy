package event

import (
	"database/sql"

	"github.com/ESNFranceG33kTeam/sAPI/database"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/router"
	"github.com/gorilla/mux"
)

func Launcher() {
	CreateEventsTable()
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
