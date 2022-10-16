package router

import (
	"net/http"

	"github.com/ESNFranceG33kTeam/sAPI/controllers"
	"github.com/ESNFranceG33kTeam/sAPI/helpers"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

var router *mux.Router

// Define our struct
type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// Initialize it somewhere
func (amw *authenticationMiddleware) Populate(usersapi []string, tokensapi []string) {
	for induser, valuser := range usersapi {
		amw.tokenUsers[tokensapi[induser]] = valuser
	}
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := amw.tokenUsers[token]; found {
			// We found the token in our map
			logger.LogInfo("auth", "user is connected : "+user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
			logger.LogInfo("auth", "auth failed")
		}
	})
}

func InitializeRouter() {
	router = mux.NewRouter().StrictSlash(true)

	// healthcheck
	router.Methods("GET").Path("/health").Name("Index").HandlerFunc(controllers.HealthsCheck)

	// swagger
	router.Handle(helpers.Swaggerpathflag, http.FileServer(http.Dir("./")))

	// documentation for developers
	opts := middleware.SwaggerUIOpts{SpecURL: helpers.Swaggerpathflag, Path: "swagger"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/swagger", sh)

	// documentation for share
	opts1 := middleware.RedocOpts{SpecURL: helpers.Swaggerpathflag, Path: "docs"}
	sh1 := middleware.Redoc(opts1, nil)
	router.Handle("/docs", sh1)

	// authentication into /auth
	// amw := authenticationMiddleware{tokenUsers: make(map[string]string)}
	// amw.Populate(helpers.AppConfig.Usersapi, helpers.AppConfig.Tokensapi)
	// secure := router.PathPrefix("/auth").Subrouter()
	// secure.Use(amw.Middleware)

	// StrictSlash is true => redirect /adherents/ to /adherents
	// route /adherents
	// secure.Methods("GET").Path("/adherents").Name("Index").HandlerFunc(controllers.AdherentsIndex)
	// secure.Methods("POST").Path("/adherents").Name("Create").HandlerFunc(controllers.AdherentsCreate)
	// secure.Methods("GET").Path("/adherents/{id}").Name("Show").HandlerFunc(controllers.AdherentsShow)
	// secure.Methods("PUT").Path("/adherents/{id}").Name("Update").HandlerFunc(controllers.AdherentsUpdate)
	// secure.Methods("DELETE").Path("/adherents/{id}").Name("DELETE").HandlerFunc(controllers.AdherentsDelete)

	// // route /esncards
	// secure.Methods("GET").Path("/esncards").Name("Index").HandlerFunc(controllers.ESNcardsIndex)
	// secure.Methods("POST").Path("/esncards").Name("Create").HandlerFunc(controllers.ESNcardsCreate)
	// secure.Methods("GET").Path("/esncards/id_adherent/{id_adherent}").Name("Show").HandlerFunc(controllers.ESNcardsShowByIdAdherent)
	// secure.Methods("GET").Path("/esncards/esncard/{esncard}").Name("Show").HandlerFunc(controllers.ESNcardsShowByESNcard)
	// secure.Methods("DELETE").Path("/esncards/{id}").Name("DELETE").HandlerFunc(controllers.ESNcardsDelete)

	// // route /volunteers
	// secure.Methods("GET").Path("/volunteers").Name("Index").HandlerFunc(controllers.VolunteersIndex)
	// secure.Methods("POST").Path("/volunteers").Name("Create").HandlerFunc(controllers.VolunteersCreate)
	// secure.Methods("GET").Path("/volunteers/id_adherent/{id_adherent}").Name("Show").HandlerFunc(controllers.VolunteersShowByIdAdherent)
	// secure.Methods("PUT").Path("/volunteers/id_adherent/{id_adherent}").Name("Update").HandlerFunc(controllers.VolunteersUpdate)
	// secure.Methods("DELETE").Path("/volunteers/{id}").Name("DELETE").HandlerFunc(controllers.VolunteersDelete)
}

func GetRouter() *mux.Router {
	return router
}

func GetSecureRouter() *mux.Router {
	// authentication into /auth
	amw := authenticationMiddleware{tokenUsers: make(map[string]string)}
	amw.Populate(helpers.AppConfig.Usersapi, helpers.AppConfig.Tokensapi)
	secure := router.PathPrefix("/auth").Subrouter()
	secure.Use(amw.Middleware)
	return secure
}
