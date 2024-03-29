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

	// stats
	TheSecureRouter().Methods("GET").
		Path("/esncards/stats/monthly").
		Name("Index").
		HandlerFunc(MonthlyStatsIndex)

	TheSecureRouter().Methods("GET").
		Path("/esncards/stats/monthly/{archive_date}").
		Name("Show").
		HandlerFunc(MonthlyStatShowByDate)

	TheSecureRouter().Methods("POST").
		Path("/esncards/stats/monthly/create").
		Name("Create").
		HandlerFunc(AutoMonthlyStatCreate)

	TheSecureRouter().Methods("POST").
		Path("/esncards/stats/monthly/force/{archive_date}").
		Name("Create").
		HandlerFunc(ForceMonthlyStatCreate)
}
