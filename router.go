package main

import (
	"net/http"

	"github.com/ESNFranceG33kTeam/sAPI/controllers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /adherents/ to /adherents
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// documentation for developers
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/docs", sh)

	// documentation for share
	opts1 := middleware.RedocOpts{SpecURL: "/swagger.yaml", Path: "docs1"}
	sh1 := middleware.Redoc(opts1, nil)
	router.Handle("/docs1", sh1)

	router.Methods("GET").Path("/adherents").Name("Index").HandlerFunc(controllers.AdherentsIndex)
	router.Methods("POST").Path("/adherents").Name("Create").HandlerFunc(controllers.AdherentsCreate)
	router.Methods("GET").Path("/adherents/{id}").Name("Show").HandlerFunc(controllers.AdherentsShow)
	router.Methods("PUT").Path("/adherents/{id}").Name("Update").HandlerFunc(controllers.AdherentsUpdate)
	router.Methods("DELETE").Path("/adherents/{id}").Name("DELETE").HandlerFunc(controllers.AdherentsDelete)
	return router
}
