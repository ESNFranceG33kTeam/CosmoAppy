package health

func InitRoutes() {
	// route /health
	TheRouter().Methods("GET").Path("/health").Name("Index").HandlerFunc(HealthsCheck)
	TheSecureRouter().Methods("GET").Path("/status").Name("Index").HandlerFunc(StatusCheck)
}
