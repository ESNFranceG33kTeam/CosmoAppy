package money

import (
	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	// route /moneys
	router.Methods("GET").Path("/moneys").Name("Index").HandlerFunc(MoneysIndex)
	router.Methods("POST").Path("/moneys").Name("Create").HandlerFunc(MoneysCreate)
	router.Methods("GET").Path("/moneys/label/{label}").Name("Show").HandlerFunc(MoneysShowByLabel)
}
