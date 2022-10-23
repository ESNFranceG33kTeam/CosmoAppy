package health

func InitRoutes() {
	// route /health
	TheRouter().Methods("GET").Path("/health").Name("Index").HandlerFunc(HealthsCheck)
}
