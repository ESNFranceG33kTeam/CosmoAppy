package volunteer

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func VolunteersIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(AllVolunteers())
	if err != nil {
		TheLogger().LogError("volunteer", "problem with indexation.", err)
	} else {
		TheLogger().LogInfo("volunteer", "request GET : "+r.RequestURI)
	}
}

func VolunteersCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("volunteer", "problem with create.", err)
	}

	var vlt Volunteer

	err = json.Unmarshal(body, &vlt)
	if err != nil {
		TheLogger().LogError("volunteer", "problem with unmarshal.", err)
	}

	w.WriteHeader(http.StatusOK)
	NewVolunteer(&vlt)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError("volunteer", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("volunteer", "request POST : "+r.RequestURI)
	}
}

func VolunteersShowByIdAdherent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_adherent, err := strconv.Atoi(vars["id_adherent"])
	if err != nil {
		TheLogger().LogError("volunteer", "unable to get id_adherent.", err)
	}

	w.WriteHeader(http.StatusOK)
	vlt := FindVolunteerByIdadherent(id_adherent)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError("volunteer", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("volunteer", "request GET : "+r.RequestURI)
	}
}

func VolunteersUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_adherent, err := strconv.Atoi(vars["id_adherent"])
	if err != nil {
		TheLogger().LogError("volunteer", "unable to get id_adherent.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("volunteer", "problem with update.", err)
	}

	vlt := FindVolunteerByIdadherent(id_adherent)

	err = json.Unmarshal(body, &vlt)
	if err != nil {
		TheLogger().LogError("volunteer", "problem with unmarshal.", err)
	}

	w.WriteHeader(http.StatusOK)
	UpdateVolunteer(vlt)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError("volunteer", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("volunteer", "request PUT : "+r.RequestURI)
	}
}

func VolunteersDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("volunteer", "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	err = DeleteVolunteerById(id)
	if err != nil {
		TheLogger().LogError("volunteer", "unable to delete volunteer.", err)
	} else {
		TheLogger().LogInfo("volunteer", "request DELETE : "+r.RequestURI)
	}
}
