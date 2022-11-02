package health

import (
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
