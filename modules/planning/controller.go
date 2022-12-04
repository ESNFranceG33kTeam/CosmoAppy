package planning

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func PlanningsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(AllPlannings())
	if err != nil {
		TheLogger().LogError("planning", "problem with indexation.", err)
	} else {
		TheLogger().LogInfo("planning", "request GET : "+r.RequestURI)
	}
}

func PlanningsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("planning", "problem with create.", err)
	}

	var pla Planning

	err = json.Unmarshal(body, &pla)
	if err != nil {
		TheLogger().LogError("planning", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", pla.Date_begins)
	if err != nil {
		TheLogger().LogInfo("planning", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}
	_, err = time.Parse("2006-01-02", pla.Date_end)
	if err != nil {
		TheLogger().LogInfo("planning", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse("15:04:00", pla.Hour_begins)
	if err != nil {
		TheLogger().LogInfo("planning", "Hour format wrong.")
		http.Error(w, "Hour format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse("15:04:00", pla.Hour_end)
	if err != nil {
		TheLogger().LogInfo("planning", "Hour format wrong.")
		http.Error(w, "Hour format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	NewPlanning(&pla)

	err = json.NewEncoder(w).Encode(pla)
	if err != nil {
		TheLogger().LogError("planning", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("planning", "request POST : "+r.RequestURI)
	}
}

func PlanningsShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("planning", "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	adh := FindPlanningById(id)

	err = json.NewEncoder(w).Encode(adh)
	if err != nil {
		TheLogger().LogError("planning", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("planning", "request GET : "+r.RequestURI)
	}
}

func PlanningsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("planning", "unable to get id.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("planning", "problem with update.", err)
	}

	pla := FindPlanningById(id)

	err = json.Unmarshal(body, &pla)
	if err != nil {
		TheLogger().LogError("planning", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", pla.Date_begins)
	if err != nil {
		TheLogger().LogInfo("planning", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}
	_, err = time.Parse("2006-01-02", pla.Date_end)
	if err != nil {
		TheLogger().LogInfo("planning", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse("15:04:00", pla.Hour_begins)
	if err != nil {
		TheLogger().LogInfo("planning", "Hour format wrong.")
		http.Error(w, "Hour format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse("15:04:00", pla.Hour_end)
	if err != nil {
		TheLogger().LogInfo("planning", "Hour format wrong.")
		http.Error(w, "Hour format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	UpdatePlanning(pla)

	err = json.NewEncoder(w).Encode(pla)
	if err != nil {
		TheLogger().LogError("planning", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("planning", "request PUT : "+r.RequestURI)
	}
}

func PlanningsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("planning", "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	err = DeletePlanningById(id)
	if err != nil {
		TheLogger().LogError("planning", "unable to delete planning.", err)
	} else {
		TheLogger().LogInfo("planning", "request DELETE : "+r.RequestURI)
	}
}

func AttendeesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(AllAttendees())
	if err != nil {
		TheLogger().LogError("attendee", "problem with indexation.", err)
	} else {
		TheLogger().LogInfo("attendee", "request GET : "+r.RequestURI)
	}
}

func AttendeesCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("attendee", "problem with create.", err)
	}

	var att Attendee

	err = json.Unmarshal(body, &att)
	if err != nil {
		TheLogger().LogError("attendee", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", att.Date)
	if err != nil {
		TheLogger().LogInfo("attendee", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse("15:04:00", att.Hour_end)
	if err != nil {
		TheLogger().LogInfo("attendee", "Hour format wrong.")
		http.Error(w, "Hour format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse("15:04:00", att.Hour_begins)
	if err != nil {
		TheLogger().LogInfo("attendee", "Hour format wrong.")
		http.Error(w, "Hour format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	NewAttendee(&att)

	err = json.NewEncoder(w).Encode(att)
	if err != nil {
		TheLogger().LogError("attendee", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("attendee", "request POST : "+r.RequestURI)
	}
}

func AttendeesShowByIdPlanning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_planning, err := strconv.Atoi(vars["id_planning"])
	if err != nil {
		TheLogger().LogError("attendee", "unable to get id_planning.", err)
	}

	w.WriteHeader(http.StatusOK)
	vlt := FindAttendeeByPlanningId(id_planning)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError("attendee", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("attendee", "request GET : "+r.RequestURI)
	}
}

func AttendeesShowByIdAdherent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_adherent, err := strconv.Atoi(vars["id_adherent"])
	if err != nil {
		TheLogger().LogError("attendee", "unable to get id_adherent.", err)
	}

	w.WriteHeader(http.StatusOK)
	vlt := FindAttendeeByAdherentId(id_adherent)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError("attendee", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("attendee", "request GET : "+r.RequestURI)
	}
}

func AttendeesUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("attendee", "unable to get id_adherent.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("attendee", "problem with update.", err)
	}

	att := FindAttendeeById(id)

	err = json.Unmarshal(body, &att)
	if err != nil {
		TheLogger().LogError("attendee", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", att.Date)
	if err != nil {
		TheLogger().LogInfo("attendee", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse("15:04:00", att.Hour_end)
	if err != nil {
		TheLogger().LogInfo("attendee", "Hour format wrong.")
		http.Error(w, "Hour format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse("15:04:00", att.Hour_begins)
	if err != nil {
		TheLogger().LogInfo("attendee", "Hour format wrong.")
		http.Error(w, "Hour format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	UpdateAttendee(att)

	err = json.NewEncoder(w).Encode(att)
	if err != nil {
		TheLogger().LogError("attendee", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("attendee", "request PUT : "+r.RequestURI)
	}
}

func AttendeesDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("attendee", "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	err = DeleteAttendeeById(id)
	if err != nil {
		TheLogger().LogError("attendee", "unable to delete attendee.", err)
	} else {
		TheLogger().LogInfo("attendee", "request DELETE : "+r.RequestURI)
	}
}
