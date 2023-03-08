package cocas

import (
	"encoding/json"
	"net/http"

	"github.com/ESNFranceG33kTeam/CosmoAppy/logger"
	"gopkg.in/cas.v2"
)

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	if !cas.IsAuthenticated(r) {
		cas.RedirectToLogin(w, r)
		return
	}

	if r.URL.Path == "/logout" {
		cas.RedirectToLogout(w, r)
		return
	}

	esner := GetProfile(r)

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(esner)
	if err != nil {
		logger.GetLogger().LogError("cocas", "problem with encoder.", err)
	} else {
		logger.GetLogger().LogInfo("cocas", "request GET : "+r.RequestURI)
	}
}
