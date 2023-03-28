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

	// stats
	TheSecureRouter().Methods("GET").
		Path("/adherents/stats/monthly").
		Name("Index").
		HandlerFunc(MonthlyStatsIndex)
	TheSecureRouter().Methods("GET").
		Path("/adherents/stats/monthly/{archive_date}").
		Name("Show").
		HandlerFunc(MonthlyStatShowByDate)
	TheSecureRouter().Methods("POST").
		Path("/adherents/stats/monthly/create").
		Name("Create").
		HandlerFunc(AutoMonthlyStatCreate)
	TheSecureRouter().Methods("POST").
		Path("/adherents/stats/monthly/force/{archive_date}").
		Name("Create").
		HandlerFunc(ForceMonthlyStatCreate)
}
