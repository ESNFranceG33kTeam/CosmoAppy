package money

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func MoneysIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(AllMoneys())
	if err != nil {
		TheLogger().LogError("money", "problem with indexation.", err)
	} else {
		TheLogger().LogInfo("money", "request GET : "+r.RequestURI)
	}
}

func MoneysCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("money", "problem with create.", err)
	}

	var mon Money

	err = json.Unmarshal(body, &mon)
	if err != nil {
		TheLogger().LogError("money", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", mon.PaymentDate)

	if err != nil {
		TheLogger().LogInfo("money", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	NewMoney(&mon)

	err = json.NewEncoder(w).Encode(mon)
	if err != nil {
		TheLogger().LogError("money", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("money", "request POST : "+r.RequestURI)
	}
}

func MoneysShowByLabel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	label := vars["label"]

	w.WriteHeader(http.StatusOK)
	mons := FindMoneysByLabel(label)

	err := json.NewEncoder(w).Encode(mons)
	if err != nil {
		TheLogger().LogError("money", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("money", "request GET : "+r.RequestURI)
	}
}
