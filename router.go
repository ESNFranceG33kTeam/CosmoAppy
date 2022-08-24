package main

import (
	"github.com/ESNFranceG33kTeam/sAPI/controllers"
	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /adherents/ to /adherents
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/adherents").Name("Index").HandlerFunc(controllers.AdherentsIndex)
	router.Methods("POST").Path("/adherents").Name("Create").HandlerFunc(controllers.AdherentsCreate)
	router.Methods("GET").Path("/adherents/{id}").Name("Show").HandlerFunc(controllers.AdherentsShow)
	router.Methods("PUT").Path("/adherents/{id}").Name("Update").HandlerFunc(controllers.AdherentsUpdate)
	router.Methods("DELETE").Path("/adherents/{id}").Name("DELETE").HandlerFunc(controllers.AdherentsDelete)
	return router
}
