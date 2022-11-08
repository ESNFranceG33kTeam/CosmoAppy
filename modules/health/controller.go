package health

import (
	"encoding/json"
	"net/http"
)

func HealthsCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	health := GetHealth()

	err := json.NewEncoder(w).Encode(health)
	if err != nil {
		TheLogger().LogError("health", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("health", "request GET : "+r.RequestURI)
	}
}

func StatusCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	status := GetStatus()

	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		TheLogger().LogError("health", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("health", "request GET : "+r.RequestURI)
	}
}
