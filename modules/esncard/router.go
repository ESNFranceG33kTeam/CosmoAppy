package esncard

import "github.com/ESNFranceG33kTeam/sAPI/router"

func InitRoutes() {
	// routes /esncards
	router.GetSecureRouter().Methods("GET").Path("/esncards").Name("Index").HandlerFunc(ESNcardsIndex)
	router.GetSecureRouter().Methods("POST").Path("/esncards").Name("Create").HandlerFunc(ESNcardsCreate)
	router.GetSecureRouter().Methods("GET").Path("/esncards/id_adherent/{id_adherent}").Name("Show").HandlerFunc(ESNcardsShowByIdAdherent)
	router.GetSecureRouter().Methods("GET").Path("/esncards/esncard/{esncard}").Name("Show").HandlerFunc(ESNcardsShowByESNcard)
	router.GetSecureRouter().Methods("DELETE").Path("/esncards/{id}").Name("DELETE").HandlerFunc(ESNcardsDelete)
}
