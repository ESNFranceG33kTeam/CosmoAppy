package volunteer

import "github.com/ESNFranceG33kTeam/sAPI/router"

func InitRoutes() {
	// routes /volunteers
	router.GetSecureRouter().Methods("GET").Path("/volunteers").Name("Index").HandlerFunc(VolunteersIndex)
	router.GetSecureRouter().Methods("POST").Path("/volunteers").Name("Create").HandlerFunc(VolunteersCreate)
	router.GetSecureRouter().Methods("GET").Path("/volunteers/id_adherent/{id_adherent}").Name("Show").HandlerFunc(VolunteersShowByIdAdherent)
	router.GetSecureRouter().Methods("PUT").Path("/volunteers/id_adherent/{id_adherent}").Name("Update").HandlerFunc(VolunteersUpdate)
	router.GetSecureRouter().Methods("DELETE").Path("/volunteers/{id}").Name("DELETE").HandlerFunc(VolunteersDelete)
}
