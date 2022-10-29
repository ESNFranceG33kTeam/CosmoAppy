package event

func InitRoutes() {
	// routes /events
	TheSecureRouter().Methods("GET").Path("/events").Name("Index").HandlerFunc(EventsIndex)
	TheSecureRouter().Methods("POST").Path("/events").Name("Create").HandlerFunc(EventsCreate)
	TheSecureRouter().Methods("GET").Path("/events/{id}").Name("Show").HandlerFunc(EventsShow)
	TheSecureRouter().Methods("PUT").Path("/events/{id}").Name("Update").HandlerFunc(EventsUpdate)
	TheSecureRouter().Methods("DELETE").Path("/events/{id}").Name("DELETE").HandlerFunc(EventsDelete)
}
