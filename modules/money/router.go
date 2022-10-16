package money

import "github.com/ESNFranceG33kTeam/sAPI/router"

func InitRoutes() {
	// routes /moneys
	router.GetSecureRouter().Methods("GET").Path("/moneys").Name("Index").HandlerFunc(MoneysIndex)
	router.GetSecureRouter().Methods("POST").Path("/moneys").Name("Create").HandlerFunc(MoneysCreate)
	router.GetSecureRouter().Methods("GET").Path("/moneys/label/{label}").Name("Show").HandlerFunc(MoneysShowByLabel)
}
