package money

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
		TheLogger().LogError("money_stat_monthly", "problem with indexation.", err)
	} else {
		TheLogger().LogInfo("money_stat_monthly", "request GET : "+r.RequestURI)
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
		TheLogger().LogError("money_stat_monthly", "problem with encoder.", err)
	} else {
		TheLogger().LogInfo("money_stat_monthly", "request GET : "+r.RequestURI)
	}
}

func MonthlyStatCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	AutoNewMonthlyStat()
}
