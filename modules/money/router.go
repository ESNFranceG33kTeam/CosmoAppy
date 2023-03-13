package money

func InitRoutes() {
	// routes /moneys
	TheSecureRouter().Methods("GET").Path("/moneys").Name("Index").HandlerFunc(MoneysIndex)
	TheSecureRouter().Methods("POST").Path("/moneys").Name("Create").HandlerFunc(MoneysCreate)
	TheSecureRouter().Methods("GET").
		Path("/moneys/label/{label}").
		Name("Show").
		HandlerFunc(MoneysShowByLabel)

	// stats
	TheSecureRouter().Methods("GET").
		Path("/moneys/stats/monthly").
		Name("Index").
		HandlerFunc(MonthlyStatsIndex)

	TheSecureRouter().Methods("GET").
		Path("/moneys/stats/monthly/{archive_date}").
		Name("Show").
		HandlerFunc(MonthlyStatShowByDate)

	TheSecureRouter().Methods("POST").
		Path("/moneys/stats/monthly/create").
		Name("Create").
		HandlerFunc(AutoMonthlyStatCreate)

	TheSecureRouter().Methods("POST").
		Path("/moneys/stats/monthly/force/{archive_date}").
		Name("Create").
		HandlerFunc(ForceMonthlyStatCreate)
}
