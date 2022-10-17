package adherent

import "github.com/ESNFranceG33kTeam/sAPI/router"

func InitRoutes() {
	// routes /adherents
	router.GetSecureRouter().Methods("GET").Path("/adherents").Name("Index").HandlerFunc(AdherentsIndex)
	router.GetSecureRouter().Methods("POST").Path("/adherents").Name("Create").HandlerFunc(AdherentsCreate)
	router.GetSecureRouter().Methods("GET").Path("/adherents/{id}").Name("Show").HandlerFunc(AdherentsShow)
	router.GetSecureRouter().Methods("PUT").Path("/adherents/{id}").Name("Update").HandlerFunc(AdherentsUpdate)
	router.GetSecureRouter().Methods("DELETE").Path("/adherents/{id}").Name("DELETE").HandlerFunc(AdherentsDelete)
}
