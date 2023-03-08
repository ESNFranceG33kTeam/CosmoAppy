package adherent

func InitRoutes() {
	// routes /adherents
	TheSecureRouter().Methods("GET").
		Path("/adherents").
		Name("Index").
		HandlerFunc(AdherentsIndex)

	TheSecureRouter().Methods("POST").
		Path("/adherents").
		Name("Create").
		HandlerFunc(AdherentsCreate)

	TheSecureRouter().Methods("GET").
		Path("/adherents/{id}").
		Name("Show").
		HandlerFunc(AdherentsShow)

	TheSecureRouter().Methods("PUT").
		Path("/adherents/{id}").
		Name("Update").
		HandlerFunc(AdherentsUpdate)

	TheSecureRouter().Methods("DELETE").
		Path("/adherents/{id}").
		Name("DELETE").
		HandlerFunc(AdherentsDelete)
}
