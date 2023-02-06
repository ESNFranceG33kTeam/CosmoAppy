package planning

func InitRoutes() {
	// routes /events
	TheSecureRouter().Methods("GET").Path("/plannings").Name("Index").HandlerFunc(PlanningsIndex)
	TheSecureRouter().Methods("POST").Path("/plannings").Name("Create").HandlerFunc(PlanningsCreate)
	TheSecureRouter().Methods("GET").Path("/plannings/{id}").Name("Show").HandlerFunc(PlanningsShow)
	TheSecureRouter().Methods("PUT").Path("/plannings/{id}").Name("Update").HandlerFunc(PlanningsUpdate)
	TheSecureRouter().Methods("DELETE").Path("/plannings/{id}").Name("DELETE").HandlerFunc(PlanningsDelete)

	TheSecureRouter().Methods("GET").Path("/planning_attendees").Name("Index").HandlerFunc(AttendeesIndex)
	TheSecureRouter().Methods("POST").Path("/planning_attendees").Name("Create").HandlerFunc(AttendeesCreate)
	TheSecureRouter().Methods("GET").Path("/planning_attendees/id_planning/{id_planning}").Name("Show").HandlerFunc(AttendeesShowByIdPlanning)
	TheSecureRouter().Methods("GET").Path("/planning_attendees/id_volunteer/{id_volunteer}").Name("Show").HandlerFunc(AttendeesShowByIdVolunteer)
	TheSecureRouter().Methods("PUT").Path("/planning_attendees/{id}").Name("Update").HandlerFunc(AttendeesUpdate)
	TheSecureRouter().Methods("DELETE").Path("/planning_attendees/{id}").Name("DELETE").HandlerFunc(AttendeesDelete)
}
