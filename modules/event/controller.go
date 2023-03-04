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

func AttendeesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(AllAttendees())
	if err != nil {
		TheLogger().LogError("event_attendee", "problem with indexation.", err)
	} else {
		TheLogger().LogInfo("event_attendee", "request GET : "+r.RequestURI)
	}
}

func AttendeesCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("event_attendee", "problem with create.", err)
	}

	var att Attendee

	err = json.Unmarshal(body, &att)
	if err != nil {
		TheLogger().LogError("event_attendee", "problem with unmarshal.", err)
	}

	eve := FindEventById(att.Id_event)
	if eve.NbSpotsTaken >= eve.NbSpotsMax {
		TheLogger().LogInfo("event", "No spot available.")
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
		TheLogger().LogError("event_attendee", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("event_attendee", "request POST : "+r.RequestURI)
	}
}

func AttendeesShowByIdEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_event, err := strconv.Atoi(vars["id_event"])
	if err != nil {
		TheLogger().LogError("event_attendee", "unable to get id_event.", err)
	}

	w.WriteHeader(http.StatusOK)
	vlt := FindAttendeeByEventId(id_event)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError("event_attendee", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("event_attendee", "request GET : "+r.RequestURI)
	}
}

func AttendeesShowByIdAdherent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_adherent, err := strconv.Atoi(vars["id_adherent"])
	if err != nil {
		TheLogger().LogError("event_attendee", "unable to get id_adherent.", err)
	}

	w.WriteHeader(http.StatusOK)
	vlt := FindAttendeeByAdherentId(id_adherent)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError("event_attendee", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("event_attendee", "request GET : "+r.RequestURI)
	}
}

func AttendeesUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("event_attendee", "unable to get id_adherent.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("event_attendee", "problem with update.", err)
	}

	att := FindAttendeeById(id)

	err = json.Unmarshal(body, &att)
	if err != nil {
		TheLogger().LogError("event_attendee", "problem with unmarshal.", err)
	}

	w.WriteHeader(http.StatusOK)
	UpdateAttendee(att)

	err = json.NewEncoder(w).Encode(att)
	if err != nil {
		TheLogger().LogError("event_attendee", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("event_attendee", "request PUT : "+r.RequestURI)
	}
}

func AttendeesDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("event_attendee", "unable to get id.", err)
	}

	eve := FindEventById(id)
	eve.NbSpotsTaken -= 1
	UpdateSpotsTakenEvent(eve)

	w.WriteHeader(http.StatusOK)
	err = DeleteAttendeeById(id)
	if err != nil {
		TheLogger().LogError("event_attendee", "unable to delete attendee.", err)
	} else {
		TheLogger().LogInfo("event_attendee", "request DELETE : "+r.RequestURI)
	}
}

func StaffsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(AllStaffs())
	if err != nil {
		TheLogger().LogError("event_staff", "problem with indexation.", err)
	} else {
		TheLogger().LogInfo("event_staff", "request GET : "+r.RequestURI)
	}
}

func StaffsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("event_staff", "problem with create.", err)
	}

	var sta Staff

	err = json.Unmarshal(body, &sta)
	if err != nil {
		TheLogger().LogError("event_staff", "problem with unmarshal.", err)
	}

	eve := FindEventById(sta.Id_event)
	if eve.NbSpotsTaken >= eve.NbSpotsMax {
		TheLogger().LogInfo("event", "No spot available.")
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
		TheLogger().LogError("event_staff", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("event_staff", "request POST : "+r.RequestURI)
	}
}

func StaffsShowByIdEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_event, err := strconv.Atoi(vars["id_event"])
	if err != nil {
		TheLogger().LogError("event_staff", "unable to get id_event.", err)
	}

	w.WriteHeader(http.StatusOK)
	sta := FindStaffByEventId(id_event)

	err = json.NewEncoder(w).Encode(sta)
	if err != nil {
		TheLogger().LogError("event_staff", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("event_staff", "request GET : "+r.RequestURI)
	}
}

func StaffsShowByIdVolunteer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id_volunteer, err := strconv.Atoi(vars["id_volunteer"])
	if err != nil {
		TheLogger().LogError("event_staff", "unable to get id_volunteer.", err)
	}

	w.WriteHeader(http.StatusOK)
	vlt := FindStaffByVolunteerId(id_volunteer)

	err = json.NewEncoder(w).Encode(vlt)
	if err != nil {
		TheLogger().LogError("event_staff", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("event_staff", "request GET : "+r.RequestURI)
	}
}

func StaffsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("event_staff", "unable to get id_adherent.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("event_staff", "problem with update.", err)
	}

	sta := FindStaffById(id)

	err = json.Unmarshal(body, &sta)
	if err != nil {
		TheLogger().LogError("event_staff", "problem with unmarshal.", err)
	}

	w.WriteHeader(http.StatusOK)
	UpdateStaff(sta)

	err = json.NewEncoder(w).Encode(sta)
	if err != nil {
		TheLogger().LogError("event_staff", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("event_staff", "request PUT : "+r.RequestURI)
	}
}

func StaffsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("event_staff", "unable to get id.", err)
	}

	eve := FindEventById(id)
	eve.NbSpotsTaken -= 1
	UpdateSpotsTakenEvent(eve)

	w.WriteHeader(http.StatusOK)
	err = DeleteStaffById(id)
	if err != nil {
		TheLogger().LogError("event_staff", "unable to delete attendee.", err)
	} else {
		TheLogger().LogInfo("event_staff", "request DELETE : "+r.RequestURI)
	}
}
