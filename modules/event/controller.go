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
		TheLogger().LogError(log_name, "problem with indexation.", err)
	} else {
		TheLogger().LogInfo(log_name, "request GET : "+r.RequestURI)
	}
}

func EventsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError(log_name, "problem with create.", err)
	}

	var eve Event

	err = json.Unmarshal(body, &eve)
	if err != nil {
		TheLogger().LogError(log_name, "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", eve.Date)

	if err != nil {
		TheLogger().LogInfo(log_name, "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	NewEvent(&eve)

	err = json.NewEncoder(w).Encode(eve)
	if err != nil {
		TheLogger().LogError(log_name, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name, "request POST : "+r.RequestURI)
	}
}

func EventsShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError(log_name, "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	adh := FindEventById(id)

	err = json.NewEncoder(w).Encode(adh)
	if err != nil {
		TheLogger().LogError(log_name, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name, "request GET : "+r.RequestURI)
	}
}

func EventsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError(log_name, "unable to get id.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError(log_name, "problem with update.", err)
	}

	eve := FindEventById(id)

	err = json.Unmarshal(body, &eve)
	if err != nil {
		TheLogger().LogError(log_name, "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", eve.Date)

	if err != nil {
		TheLogger().LogInfo(log_name, "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	UpdateEvent(eve)

	err = json.NewEncoder(w).Encode(eve)
	if err != nil {
		TheLogger().LogError(log_name, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name, "request PUT : "+r.RequestURI)
	}
}

func EventsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError(log_name, "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	err = DeleteEventById(id)
	if err != nil {
		TheLogger().LogError(log_name, "unable to delete event.", err)
	} else {
		TheLogger().LogInfo(log_name, "request DELETE : "+r.RequestURI)
	}
}

func AttendeesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(AllAttendees())
	if err != nil {
		TheLogger().LogError(log_name_att, "problem with indexation.", err)
	} else {
		TheLogger().LogInfo(log_name_att, "request GET : "+r.RequestURI)
	}
}

func AttendeesCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError(log_name_att, "problem with create.", err)
	}

	var att Attendee

	err = json.Unmarshal(body, &att)
	if err != nil {
		TheLogger().LogError(log_name_att, "problem with unmarshal.", err)
	}

	eve := FindEventById(att.Id_event)
	if eve.NbSpotsTaken >= eve.NbSpotsMax {
		TheLogger().LogInfo(log_name, "No spot available.")
		http.Error(w, "No spot available.", http.StatusBadRequest)

		return
	} else {
		eve.NbSpotsTaken += 1
	}

	w.WriteHeader(http.StatusOK)
	UpdateSpotsTakenEvent(eve)
	NewAttendee(&att)

	err = json.NewEncoder(w).Encode(att)
	if err != nil {
		TheLogger().LogError(log_name_att, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name_att, "request POST : "+r.RequestURI)
	}
}

func AttendeesShowByIdEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_event, err := strconv.Atoi(vars["id_event"])
	if err != nil {
		TheLogger().LogError(log_name_att, "unable to get id_event.", err)
	}

	w.WriteHeader(http.StatusOK)
	vlt := FindAttendeeByEventId(id_event)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError(log_name_att, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name_att, "request GET : "+r.RequestURI)
	}
}

func AttendeesShowByIdAdherent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_adherent, err := strconv.Atoi(vars["id_adherent"])
	if err != nil {
		TheLogger().LogError(log_name_att, "unable to get id_adherent.", err)
	}

	w.WriteHeader(http.StatusOK)
	vlt := FindAttendeeByAdherentId(id_adherent)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError(log_name_att, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name_att, "request GET : "+r.RequestURI)
	}
}

func AttendeesUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError(log_name_att, "unable to get id_adherent.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError(log_name_att, "problem with update.", err)
	}

	att := FindAttendeeById(id)

	err = json.Unmarshal(body, &att)
	if err != nil {
		TheLogger().LogError(log_name_att, "problem with unmarshal.", err)
	}

	w.WriteHeader(http.StatusOK)
	UpdateAttendee(att)

	err = json.NewEncoder(w).Encode(att)
	if err != nil {
		TheLogger().LogError(log_name_att, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name_att, "request PUT : "+r.RequestURI)
	}
}

func AttendeesDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError(log_name_att, "unable to get id.", err)
	}

	eve := FindEventById(id)
	eve.NbSpotsTaken -= 1
	UpdateSpotsTakenEvent(eve)

	w.WriteHeader(http.StatusOK)
	err = DeleteAttendeeById(id)
	if err != nil {
		TheLogger().LogError(log_name_att, "unable to delete attendee.", err)
	} else {
		TheLogger().LogInfo(log_name_att, "request DELETE : "+r.RequestURI)
	}
}

func StaffsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(AllStaffs())
	if err != nil {
		TheLogger().LogError(log_name_sta, "problem with indexation.", err)
	} else {
		TheLogger().LogInfo(log_name_sta, "request GET : "+r.RequestURI)
	}
}

func StaffsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError(log_name_sta, "problem with create.", err)
	}

	var sta Staff

	err = json.Unmarshal(body, &sta)
	if err != nil {
		TheLogger().LogError(log_name_sta, "problem with unmarshal.", err)
	}

	eve := FindEventById(sta.Id_event)
	if eve.NbSpotsTaken >= eve.NbSpotsMax {
		TheLogger().LogInfo(log_name, "No spot available.")
		http.Error(w, "No spot available.", http.StatusBadRequest)

		return
	} else {
		eve.NbSpotsTaken += 1
	}

	w.WriteHeader(http.StatusOK)
	UpdateSpotsTakenEvent(eve)
	NewStaff(&sta)

	err = json.NewEncoder(w).Encode(sta)
	if err != nil {
		TheLogger().LogError(log_name_sta, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name_sta, "request POST : "+r.RequestURI)
	}
}

func StaffsShowByIdEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_event, err := strconv.Atoi(vars["id_event"])
	if err != nil {
		TheLogger().LogError(log_name_sta, "unable to get id_event.", err)
	}

	w.WriteHeader(http.StatusOK)
	sta := FindStaffByEventId(id_event)

	err = json.NewEncoder(w).Encode(sta)
	if err != nil {
		TheLogger().LogError(log_name_sta, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name_sta, "request GET : "+r.RequestURI)
	}
}

func StaffsShowByIdVolunteer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_volunteer, err := strconv.Atoi(vars["id_volunteer"])
	if err != nil {
		TheLogger().LogError(log_name_sta, "unable to get id_volunteer.", err)
	}

	w.WriteHeader(http.StatusOK)
	vlt := FindStaffByVolunteerId(id_volunteer)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError(log_name_sta, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name_sta, "request GET : "+r.RequestURI)
	}
}

func StaffsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError(log_name_sta, "unable to get id_adherent.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError(log_name_sta, "problem with update.", err)
	}

	sta := FindStaffById(id)

	err = json.Unmarshal(body, &sta)
	if err != nil {
		TheLogger().LogError(log_name_sta, "problem with unmarshal.", err)
	}

	w.WriteHeader(http.StatusOK)
	UpdateStaff(sta)

	err = json.NewEncoder(w).Encode(sta)
	if err != nil {
		TheLogger().LogError(log_name_sta, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name_sta, "request PUT : "+r.RequestURI)
	}
}

func StaffsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError(log_name_sta, "unable to get id.", err)
	}

	eve := FindEventById(id)
	eve.NbSpotsTaken -= 1
	UpdateSpotsTakenEvent(eve)

	w.WriteHeader(http.StatusOK)
	err = DeleteStaffById(id)
	if err != nil {
		TheLogger().LogError(log_name_sta, "unable to delete attendee.", err)
	} else {
		TheLogger().LogInfo(log_name_sta, "request DELETE : "+r.RequestURI)
	}
}
