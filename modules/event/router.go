package event

func InitRoutes() {
	// routes /events
	TheSecureRouter().Methods("GET").Path("/events").Name("Index").HandlerFunc(EventsIndex)
	TheSecureRouter().Methods("POST").Path("/events").Name("Create").HandlerFunc(EventsCreate)
	TheSecureRouter().Methods("GET").Path("/events/{id}").Name("Show").HandlerFunc(EventsShow)
	TheSecureRouter().Methods("PUT").Path("/events/take_spot/{id}").Name("Update").HandlerFunc(EventsUpdateSpotsAvai)
	TheSecureRouter().Methods("PUT").Path("/events/update/{id}").Name("Update").HandlerFunc(EventsUpdate)
	TheSecureRouter().Methods("DELETE").Path("/events/{id}").Name("DELETE").HandlerFunc(EventsDelete)

	TheSecureRouter().Methods("GET").Path("/event_attendees").Name("Index").HandlerFunc(AttendeesIndex)
	TheSecureRouter().Methods("POST").Path("/event_attendees").Name("Create").HandlerFunc(AttendeesCreate)
	TheSecureRouter().Methods("GET").Path("/event_attendees/id_event/{id_event}").Name("Show").HandlerFunc(AttendeesShowByIdEvent)
	TheSecureRouter().Methods("GET").Path("/event_attendees/id_adherent/{id_adherent}").Name("Show").HandlerFunc(AttendeesShowByIdAdherent)
	TheSecureRouter().Methods("PUT").Path("/event_attendees/{id}").Name("Update").HandlerFunc(AttendeesUpdate)
	TheSecureRouter().Methods("DELETE").Path("/event_attendees/{id}").Name("DELETE").HandlerFunc(AttendeesDelete)
}
