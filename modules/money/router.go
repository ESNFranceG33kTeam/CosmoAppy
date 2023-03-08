package money

func InitRoutes() {
	// routes /moneys
	TheSecureRouter().Methods("GET").Path("/moneys").Name("Index").HandlerFunc(MoneysIndex)
	TheSecureRouter().Methods("POST").Path("/moneys").Name("Create").HandlerFunc(MoneysCreate)
	TheSecureRouter().Methods("GET").
		Path("/moneys/label/{label}").
		Name("Show").
		HandlerFunc(MoneysShowByLabel)
}
