package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/models"
	"github.com/gorilla/mux"
)

func ESNcardsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(models.AllESNcards())
	if err != nil {
		logger.LogError("esncard", "problem with indexation.", err)
	} else {
		logger.LogInfo("esncard", "request GET : "+r.RequestURI)
	}
}

func ESNcardsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.LogError("esncard", "problem with create.", err)
	}

	var card models.ESNcard

	err = json.Unmarshal(body, &card)
	if err != nil {
		logger.LogError("esncard", "problem with unmarshal.", err)
	}

	w.WriteHeader(http.StatusOK)
	models.NewESNcard(&card)

	err = json.NewEncoder(w).Encode(card)
	if err != nil {
		logger.LogError("esncard", "problem with encoder.", err)
	} else {
		logger.LogInfo("esncard", "request POST : "+r.RequestURI)
	}
}

func ESNcardsShowByIdAdherent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_adherent, err := strconv.Atoi(vars["id_adherent"])
	if err != nil {
		logger.LogError("esncard", "unable to get id_adherent.", err)
	}

	w.WriteHeader(http.StatusOK)
	card := models.FindESNcardByIdadherent(id_adherent)

	err = json.NewEncoder(w).Encode(card)
	if err != nil {
		logger.LogError("esncard", "problem with encoder.", err)
	} else {
		logger.LogInfo("esncard", "request GET : "+r.RequestURI)
	}
}

func ESNcardsShowByESNcard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	esncard := vars["esncard"]

	w.WriteHeader(http.StatusOK)
	adh := models.FindESNcardByESNcard(esncard)

	err := json.NewEncoder(w).Encode(adh)
	if err != nil {
		logger.LogError("esncard", "problem with encoder.", err)
	} else {
		logger.LogInfo("esncard", "request GET : "+r.RequestURI)
	}
}

func ESNcardsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.LogError("esncard", "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	err = models.DeleteESNcardById(id)
	if err != nil {
		logger.LogError("esncard", "unable to delete esncard.", err)
	} else {
		logger.LogInfo("esncard", "request DELETE : "+r.RequestURI)
	}
}
