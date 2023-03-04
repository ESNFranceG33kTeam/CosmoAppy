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
		TheLogger().LogError("planning_attendee", "problem with indexation.", err)
	} else {
		TheLogger().LogInfo("planning_attendee", "request GET : "+r.RequestURI)
	}
}

func AttendeesCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("planning_attendee", "problem with create.", err)
	}

	var att Attendee

	err = json.Unmarshal(body, &att)
	if err != nil {
		TheLogger().LogError("planning_attendee", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", att.Date)
	if err != nil {
		TheLogger().LogInfo("planning_attendee", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse("15:04:00", att.Hour_end)
	if err != nil {
		TheLogger().LogInfo("planning_attendee", "Hour format wrong.")
		http.Error(w, "Hour format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse("15:04:00", att.Hour_begins)
	if err != nil {
		TheLogger().LogInfo("planning_attendee", "Hour format wrong.")
		http.Error(w, "Hour format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	NewAttendee(&att)

	err = json.NewEncoder(w).Encode(att)
	if err != nil {
		TheLogger().LogError("planning_attendee", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("planning_attendee", "request POST : "+r.RequestURI)
	}
}

func AttendeesShowByIdPlanning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_planning, err := strconv.Atoi(vars["id_planning"])
	if err != nil {
		TheLogger().LogError("planning_attendee", "unable to get id_planning.", err)
	}

	w.WriteHeader(http.StatusOK)
	vlt := FindAttendeeByPlanningId(id_planning)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError("planning_attendee", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("planning_attendee", "request GET : "+r.RequestURI)
	}
}

func AttendeesShowByIdVolunteer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_volunteer, err := strconv.Atoi(vars["id_volunteer"])
	if err != nil {
		TheLogger().LogError("planning_attendee", "unable to get id_volunteer.", err)
	}

	w.WriteHeader(http.StatusOK)
	vlt := FindAttendeeByVolunteerId(id_volunteer)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError("planning_attendee", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("planning_attendee", "request GET : "+r.RequestURI)
	}
}

func AttendeesUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("planning_attendee", "unable to get id_volunteer.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("planning_attendee", "problem with update.", err)
	}

	att := FindAttendeeById(id)

	err = json.Unmarshal(body, &att)
	if err != nil {
		TheLogger().LogError("planning_attendee", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", att.Date)
	if err != nil {
		TheLogger().LogInfo("planning_attendee", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse("15:04:00", att.Hour_end)
	if err != nil {
		TheLogger().LogInfo("planning_attendee", "Hour format wrong.")
		http.Error(w, "Hour format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	_, err = time.Parse("15:04:00", att.Hour_begins)
	if err != nil {
		TheLogger().LogInfo("planning_attendee", "Hour format wrong.")
		http.Error(w, "Hour format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	UpdateAttendee(att)

	err = json.NewEncoder(w).Encode(att)
	if err != nil {
		TheLogger().LogError("planning_attendee", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("planning_attendee", "request PUT : "+r.RequestURI)
	}
}

func AttendeesDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("planning_attendee", "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	err = DeleteAttendeeById(id)
	if err != nil {
		TheLogger().LogError("planning_attendee", "unable to delete attendee.", err)
	} else {
		TheLogger().LogInfo("planning_attendee", "request DELETE : "+r.RequestURI)
	}
}
