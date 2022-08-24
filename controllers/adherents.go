package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ESNFranceG33kTeam/sAPI/models"
	"github.com/gorilla/mux"
)

func AdherentsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(models.AllAdherents())
	if err != nil {
		log.Fatal(err)
	}
}

func AdherentsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var adh models.Adherent

	err = json.Unmarshal(body, &adh)
	if err != nil {
		log.Fatal(err)
	}

	models.NewAdherent(&adh)

	err = json.NewEncoder(w).Encode(adh)
	if err != nil {
		log.Fatal(err)
	}
}

func AdherentsShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	adh := models.FindAdherentById(id)

	err = json.NewEncoder(w).Encode(adh)
	if err != nil {
		log.Fatal(err)
	}
}

func AdherentsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	adh := models.FindAdherentById(id)

	err = json.Unmarshal(body, &adh)
	if err != nil {
		log.Fatal(err)
	}

	models.UpdateAdherent(adh)

	err = json.NewEncoder(w).Encode(adh)
	if err != nil {
		log.Fatal(err)
	}
}

func AdherentsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	err = models.DeleteAdherentById(id)
	if err != nil {
		log.Fatal(err)
	}
}
