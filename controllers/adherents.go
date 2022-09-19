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

// swagger:route GET /adherents adherent listAdherents
// Get adherents list
//
// security:
// - apiKey: []
// responses:
//
//	401: CommonError
//	200: GetAdherents
func AdherentsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(models.AllAdherents())
	if err != nil {
		logger.LogError("adherent", "problem with indexation.", err)
	} else {
		logger.LogInfo("adherent", "request GET : "+r.RequestURI)
	}
}

// swagger:route POST /adherents adherent addAdherent
// Create a new adherent
//
// security:
// - apiKey: []
// responses:
//
//	401: CommonError
//	200: GetAdherent
func AdherentsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.LogError("adherent", "problem with create.", err)
	} else {
		logger.LogInfo("adherent", "request POST : "+r.RequestURI)
	}

	var adh models.Adherent

	err = json.Unmarshal(body, &adh)
	if err != nil {
		logger.LogError("adherent", "problem with unmarshal.", err)
	}

	models.NewAdherent(&adh)

	err = json.NewEncoder(w).Encode(adh)
	if err != nil {
		logger.LogError("adherent", "problem with encoder.", err)
	}
}

// swagger:route  GET /adherents/{id} adherent showAdherent
// Show an adherent
//
// consumes:
//   - application/x-www-form-urlencoded
//
// security:
// - apiKey: []
// responses:
//
//	401: CommonError
//	200: GetAdherent
func AdherentsShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.LogError("adherent", "unable to get id.", err)
	} else {
		logger.LogInfo("adherent", "request GET : "+r.RequestURI)
	}

	adh := models.FindAdherentById(id)

	err = json.NewEncoder(w).Encode(adh)
	if err != nil {
		logger.LogError("adherent", "problem with encoder.", err)
	}
}

func AdherentsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.LogError("adherent", "unable to get id.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.LogError("adherent", "problem with update.", err)
	} else {
		logger.LogInfo("adherent", "request PUT : "+r.RequestURI)
	}

	adh := models.FindAdherentById(id)

	err = json.Unmarshal(body, &adh)
	if err != nil {
		logger.LogError("adherent", "problem with unmarshal.", err)
	}

	models.UpdateAdherent(adh)

	err = json.NewEncoder(w).Encode(adh)
	if err != nil {
		logger.LogError("adherent", "problem with encoder.", err)
	}
}

func AdherentsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.LogError("adherent", "unable to get id.", err)
	}

	err = models.DeleteAdherentById(id)
	if err != nil {
		logger.LogError("adherent", "unable to delete adherent.", err)
	}
}
