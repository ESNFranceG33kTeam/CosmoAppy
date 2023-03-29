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

	// stats
	TheSecureRouter().Methods("GET").
		Path("/reports/stats/monthly").
		Name("Index").
		HandlerFunc(MonthlyStatsIndex)
	TheSecureRouter().Methods("GET").
		Path("/reports/stats/monthly/{archive_date}").
		Name("Show").
		HandlerFunc(MonthlyStatShowByDate)
	TheSecureRouter().Methods("POST").
		Path("/reports/stats/monthly/create").
		Name("Create").
		HandlerFunc(AutoMonthlyStatCreate)
	TheSecureRouter().Methods("POST").
		Path("/reports/stats/monthly/force/{archive_date}").
		Name("Create").
		HandlerFunc(ForceMonthlyStatCreate)
}
