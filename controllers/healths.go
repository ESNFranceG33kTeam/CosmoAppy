package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/models"
)

func HealthsCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	health := models.GetHealth()

	err := json.NewEncoder(w).Encode(health)
	if err != nil {
		logger.LogError("health", "problem with encoder.", err)
	} else {
		logger.LogInfo("health", "request GET : "+r.RequestURI)
	}
}
