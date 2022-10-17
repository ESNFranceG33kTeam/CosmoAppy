package money

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/gorilla/mux"
)

func MoneysIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(AllMoneys())
	if err != nil {
		logger.LogError("money", "problem with indexation.", err)
	} else {
		logger.LogInfo("money", "request GET : "+r.RequestURI)
	}
}

func MoneysCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.LogError("money", "problem with create.", err)
	}

	var mon Money

	err = json.Unmarshal(body, &mon)
	if err != nil {
		logger.LogError("money", "problem with unmarshal.", err)
	}

	w.WriteHeader(http.StatusOK)
	NewMoney(&mon)

	err = json.NewEncoder(w).Encode(mon)
	if err != nil {
		logger.LogError("money", "problem with encoder.", err)
	} else {
		logger.LogInfo("money", "request POST : "+r.RequestURI)
	}
}

func MoneysShowByLabel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	label := vars["label"]

	w.WriteHeader(http.StatusOK)
	mons := FindMoneyByLabel(label)

	err := json.NewEncoder(w).Encode(mons)
	if err != nil {
		logger.LogError("money", "problem with encoder.", err)
	} else {
		logger.LogInfo("money", "request GET : "+r.RequestURI)
	}
}
