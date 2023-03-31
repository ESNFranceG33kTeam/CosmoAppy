package volunteer

func InitRoutes() {
	// routes /volunteers
	TheSecureRouter().Methods("GET").Path("/volunteers").Name("Index").HandlerFunc(VolunteersIndex)
	TheSecureRouter().Methods("POST").
		Path("/volunteers").
		Name("Create").
		HandlerFunc(VolunteersCreate)
	TheSecureRouter().Methods("GET").
		Path("/volunteers/{id}").
		Name("Show").
		HandlerFunc(VolunteersShowById)
	TheSecureRouter().Methods("PUT").
		Path("/volunteers/{id}").
		Name("Update").
		HandlerFunc(VolunteersUpdate)
	TheSecureRouter().Methods("DELETE").
		Path("/volunteers/{id}").
		Name("DELETE").
		HandlerFunc(VolunteersDelete)

	// stats
	TheSecureRouter().Methods("GET").
		Path("/volunteers/stats/monthly").
		Name("Index").
		HandlerFunc(MonthlyStatsIndex)
	TheSecureRouter().Methods("GET").
		Path("/volunteers/stats/monthly/{archive_date}").
		Name("Show").
		HandlerFunc(MonthlyStatShowByDate)
	TheSecureRouter().Methods("POST").
		Path("/volunteers/stats/monthly/create").
		Name("Create").
		HandlerFunc(AutoMonthlyStatCreate)
}
