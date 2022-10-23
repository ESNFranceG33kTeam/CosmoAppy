package volunteer

func InitRoutes() {
	// routes /volunteers
	TheSecureRouter().Methods("GET").Path("/volunteers").Name("Index").HandlerFunc(VolunteersIndex)
	TheSecureRouter().Methods("POST").Path("/volunteers").Name("Create").HandlerFunc(VolunteersCreate)
	TheSecureRouter().Methods("GET").Path("/volunteers/id_adherent/{id_adherent}").Name("Show").HandlerFunc(VolunteersShowByIdAdherent)
	TheSecureRouter().Methods("PUT").Path("/volunteers/id_adherent/{id_adherent}").Name("Update").HandlerFunc(VolunteersUpdate)
	TheSecureRouter().Methods("DELETE").Path("/volunteers/{id}").Name("DELETE").HandlerFunc(VolunteersDelete)
}
