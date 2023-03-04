package report

func InitRoutes() {
	// routes /reports
	TheSecureRouter().Methods("GET").Path("/reports").Name("Index").HandlerFunc(ReportsIndex)
	TheSecureRouter().Methods("POST").Path("/reports").Name("Create").HandlerFunc(ReportsCreate)
	TheSecureRouter().Methods("GET").
		Path("/reports/{id}").
		Name("Show").
		HandlerFunc(ReportsShowById)
	TheSecureRouter().Methods("PUT").
		Path("/reports/{id}").
		Name("Update").
		HandlerFunc(ReportsUpdate)
	TheSecureRouter().Methods("DELETE").
		Path("/reports/{id}").
		Name("DELETE").
		HandlerFunc(ReportsDelete)
}
