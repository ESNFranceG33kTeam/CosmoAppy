package event

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func EventsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(AllEvents())
	if err != nil {
		TheLogger().LogError("event", "problem with indexation.", err)
	} else {
		TheLogger().LogInfo("event", "request GET : "+r.RequestURI)
	}
}

func EventsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("event", "problem with create.", err)
	}

	var eve Event

	err = json.Unmarshal(body, &eve)
	if err != nil {
		TheLogger().LogError("event", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", eve.Date)

	if err != nil {
		TheLogger().LogInfo("event", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	NewEvent(&eve)

	err = json.NewEncoder(w).Encode(eve)
	if err != nil {
		TheLogger().LogError("event", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("event", "request POST : "+r.RequestURI)
	}
}

func EventsShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("event", "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	adh := FindEventById(id)

	err = json.NewEncoder(w).Encode(adh)
	if err != nil {
		TheLogger().LogError("event", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("event", "request GET : "+r.RequestURI)
	}
}

func EventsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("event", "unable to get id.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("event", "problem with update.", err)
	}

	eve := FindEventById(id)

	err = json.Unmarshal(body, &eve)
	if err != nil {
		TheLogger().LogError("event", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", eve.Date)

	if err != nil {
		TheLogger().LogInfo("event", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	UpdateEvent(eve)

	err = json.NewEncoder(w).Encode(eve)
	if err != nil {
		TheLogger().LogError("event", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("event", "request PUT : "+r.RequestURI)
	}
}

func EventsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("event", "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	err = DeleteEventById(id)
	if err != nil {
		TheLogger().LogError("event", "unable to delete event.", err)
	} else {
		TheLogger().LogInfo("event", "request DELETE : "+r.RequestURI)
	}
}
