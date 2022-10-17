package health

import "github.com/ESNFranceG33kTeam/sAPI/router"

func InitRoutes() {
	// route /health
	router.GetRouter().Methods("GET").Path("/health").Name("Index").HandlerFunc(HealthsCheck)
}
