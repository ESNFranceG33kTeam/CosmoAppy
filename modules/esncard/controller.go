package esncard

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ESNcardsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(AllESNcards())
	if err != nil {
		TheLogger().LogError(log_name, "problem with indexation.", err)
	} else {
		TheLogger().LogInfo(log_name, "request GET : "+r.RequestURI)
	}
}

func ESNcardsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError(log_name, "problem with create.", err)
	}

	var card ESNcard

	err = json.Unmarshal(body, &card)
	if err != nil {
		TheLogger().LogError(log_name, "problem with unmarshal.", err)
	}

	w.WriteHeader(http.StatusOK)
	NewESNcard(&card)

	err = json.NewEncoder(w).Encode(card)
	if err != nil {
		TheLogger().LogError(log_name, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name, "request POST : "+r.RequestURI)
	}
}

func ESNcardsShowByIdAdherent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_adherent, err := strconv.Atoi(vars["id_adherent"])
	if err != nil {
		TheLogger().LogError(log_name, "unable to get id_adherent.", err)
	}

	w.WriteHeader(http.StatusOK)
	card := FindESNcardByIdadherent(id_adherent)

	err = json.NewEncoder(w).Encode(card)
	if err != nil {
		TheLogger().LogError(log_name, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name, "request GET : "+r.RequestURI)
	}
}

func ESNcardsShowByESNcard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	esncard := vars["esncard"]

	w.WriteHeader(http.StatusOK)
	card := FindESNcardByESNcard(esncard)

	err := json.NewEncoder(w).Encode(card)
	if err != nil {
		TheLogger().LogError(log_name, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name, "request GET : "+r.RequestURI)
	}
}

func ESNcardsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError(log_name, "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	err = DeleteESNcardById(id)
	if err != nil {
		TheLogger().LogError(log_name, "unable to delete esncard.", err)
	} else {
		TheLogger().LogInfo(log_name, "request DELETE : "+r.RequestURI)
	}
}
