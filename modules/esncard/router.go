package esncard

func InitRoutes() {
	// routes /esncards
	TheSecureRouter().Methods("GET").
		Path("/esncards").
		Name("Index").
		HandlerFunc(ESNcardsIndex)

	TheSecureRouter().Methods("POST").
		Path("/esncards").
		Name("Create").
		HandlerFunc(ESNcardsCreate)

	TheSecureRouter().Methods("GET").
		Path("/esncards/id_adherent/{id_adherent}").
		Name("Show").
		HandlerFunc(ESNcardsShowByIdAdherent)

	TheSecureRouter().Methods("GET").
		Path("/esncards/esncard/{esncard}").
		Name("Show").
		HandlerFunc(ESNcardsShowByESNcard)

	TheSecureRouter().Methods("DELETE").
		Path("/esncards/{id}").
		Name("DELETE").
		HandlerFunc(ESNcardsDelete)
}
