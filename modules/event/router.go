package event

func InitRoutes() {
	// routes /events
	TheSecureRouter().Methods("GET").Path("/events").Name("Index").HandlerFunc(EventsIndex)
	TheSecureRouter().Methods("POST").Path("/events").Name("Create").HandlerFunc(EventsCreate)
	TheSecureRouter().Methods("GET").Path("/events/{id}").Name("Show").HandlerFunc(EventsShow)
	TheSecureRouter().Methods("PUT").Path("/events/{id}").Name("Update").HandlerFunc(EventsUpdate)
	TheSecureRouter().Methods("DELETE").
		Path("/events/{id}").
		Name("DELETE").
		HandlerFunc(EventsDelete)

	// /event_attendees
	TheSecureRouter().Methods("GET").
		Path("/event_attendees").
		Name("Index").
		HandlerFunc(AttendeesIndex)
	TheSecureRouter().Methods("POST").
		Path("/event_attendees").
		Name("Create").
		HandlerFunc(AttendeesCreate)
	TheSecureRouter().Methods("GET").
		Path("/event_attendees/id_event/{id_event}").
		Name("Show").
		HandlerFunc(AttendeesShowByIdEvent)
	TheSecureRouter().Methods("GET").
		Path("/event_attendees/id_adherent/{id_adherent}").
		Name("Show").
		HandlerFunc(AttendeesShowByIdAdherent)
	TheSecureRouter().Methods("PUT").
		Path("/event_attendees/{id}").
		Name("Update").
		HandlerFunc(AttendeesUpdate)
	TheSecureRouter().Methods("DELETE").
		Path("/event_attendees/{id}").
		Name("DELETE").
		HandlerFunc(AttendeesDelete)

	// /event_staffs
	TheSecureRouter().Methods("GET").Path("/event_staffs").Name("Index").HandlerFunc(StaffsIndex)
	TheSecureRouter().Methods("POST").Path("/event_staffs").Name("Create").HandlerFunc(StaffsCreate)
	TheSecureRouter().Methods("GET").
		Path("/event_staffs/id_event/{id_event}").
		Name("Show").
		HandlerFunc(StaffsShowByIdEvent)
	TheSecureRouter().Methods("GET").
		Path("/event_staffs/id_volunteer/{id_volunteer}").
		Name("Show").
		HandlerFunc(StaffsShowByIdVolunteer)
	TheSecureRouter().Methods("PUT").
		Path("/event_staffs/{id}").
		Name("Update").
		HandlerFunc(StaffsUpdate)
	TheSecureRouter().Methods("DELETE").
		Path("/event_staffs/{id}").
		Name("DELETE").
		HandlerFunc(StaffsDelete)

	// stats
	TheSecureRouter().Methods("GET").
		Path("/events/stats/monthly").
		Name("Index").
		HandlerFunc(MonthlyStatsIndex)
	TheSecureRouter().Methods("GET").
		Path("/events/stats/monthly/{archive_date}").
		Name("Show").
		HandlerFunc(MonthlyStatShowByDate)
	TheSecureRouter().Methods("POST").
		Path("/events/stats/monthly/create").
		Name("Create").
		HandlerFunc(AutoMonthlyStatCreate)
	TheSecureRouter().Methods("POST").
		Path("/events/stats/monthly/force/{archive_date}").
		Name("Create").
		HandlerFunc(ForceMonthlyStatCreate)
}
