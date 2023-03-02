package report

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func ReportsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(AllReports())
	if err != nil {
		TheLogger().LogError("report", "problem with indexation.", err)
	} else {
		TheLogger().LogInfo("report", "request GET : "+r.RequestURI)
	}
}

func ReportsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("report", "problem with create.", err)
	}

	var rpt Report

	err = json.Unmarshal(body, &rpt)
	if err != nil {
		TheLogger().LogError("report", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", rpt.Date)

	if err != nil {
		TheLogger().LogInfo("report", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	NewReport(&rpt)

	err = json.NewEncoder(w).Encode(rpt)
	if err != nil {
		TheLogger().LogError("report", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("report", "request POST : "+r.RequestURI)
	}
}

func ReportsShowById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("report", "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	rpt := FindReportById(id)

	err = json.NewEncoder(w).Encode(rpt)
	if err != nil {
		TheLogger().LogError("report", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("report", "request GET : "+r.RequestURI)
	}
}

func ReportsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("report", "unable to get id.", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		TheLogger().LogError("report", "problem with update.", err)
	}

	rpt := FindReportById(id)

	err = json.Unmarshal(body, &rpt)
	if err != nil {
		TheLogger().LogError("report", "problem with unmarshal.", err)
	}

	_, err = time.Parse("2006-01-02", rpt.Date)

	if err != nil {
		TheLogger().LogInfo("report", "Date format wrong.")
		http.Error(w, "Date format wrong : "+err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	UpdateReport(rpt)

	err = json.NewEncoder(w).Encode(rpt)
	if err != nil {
		TheLogger().LogError("report", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("report", "request PUT : "+r.RequestURI)
	}
}

func ReportsDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		TheLogger().LogError("report", "unable to get id.", err)
	}

	w.WriteHeader(http.StatusOK)
	err = DeleteReportById(id)
	if err != nil {
		TheLogger().LogError("report", "unable to delete report.", err)
	} else {
		TheLogger().LogInfo("report", "request DELETE : "+r.RequestURI)
	}
}
