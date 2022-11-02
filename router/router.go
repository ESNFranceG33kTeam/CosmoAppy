package router

import (
	"net/http"

	"github.com/ESNFranceG33kTeam/CosmoAppy/helpers"
	"github.com/ESNFranceG33kTeam/CosmoAppy/logger"
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
			logger.GetLogger().LogInfo("auth", "user is connected : "+user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
			logger.GetLogger().LogInfo("auth", "auth failed")
		}
	})
}

func InitializeRouter() {
	// StrictSlash is true => redirect /adherents/ to /adherents
	router = mux.NewRouter().StrictSlash(true)

	// swagger
	router.Handle(helpers.Swaggerpathflag, http.FileServer(http.Dir(".")))

	// documentation for developers
	opts := middleware.SwaggerUIOpts{SpecURL: helpers.Swaggerpathflag, Path: "swagger"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/swagger", sh)

	// documentation for share
	opts1 := middleware.RedocOpts{SpecURL: helpers.Swaggerpathflag, Path: "docs"}
	sh1 := middleware.Redoc(opts1, nil)
	router.Handle("/docs", sh1)
}

func GetRouter() *mux.Router {
	return router
}

func GetSecureRouter() *mux.Router {
	// authentication into /auth
	amw := authenticationMiddleware{tokenUsers: make(map[string]string)}
	amw.Populate(helpers.TheAppConfig().Usersapi, helpers.TheAppConfig().Tokensapi)
	secure := router.PathPrefix("/auth").Subrouter()
	secure.Use(amw.Middleware)
	return secure
}
