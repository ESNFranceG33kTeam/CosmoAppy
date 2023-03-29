package event

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func MonthlyStatsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	stats := AllMonthlyStats()

	err := json.NewEncoder(w).Encode(stats)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with indexation.", err)
	} else {
		TheLogger().LogInfo(log_name_monthly_stat, "request GET : "+r.RequestURI)
	}
}

func MonthlyStatShowByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	archive_date := vars["archive_date"]

	w.WriteHeader(http.StatusOK)
	stats := FindMonthlyStatByDate(archive_date)

	err := json.NewEncoder(w).Encode(stats)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	} else {
		TheLogger().LogInfo(log_name_monthly_stat, "request GET : "+r.RequestURI)
	}
}

func AutoMonthlyStatCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	AutoNewMonthlyStat()
	TheLogger().LogInfo(log_name_monthly_stat, "request GET : "+r.RequestURI)
}

func ForceMonthlyStatCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	archive_date := vars["archive_date"]
	w.WriteHeader(http.StatusOK)

	AutoNewMonthlyStat(archive_date)
	TheLogger().LogInfo(log_name_monthly_stat, "request GET : "+r.RequestURI)
}
