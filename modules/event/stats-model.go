package event

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"time"
)

// swagger:model Event_Stat
type Event_Stat struct {
	// Id of the entry
	// in: int
	// read only: true
	Id int `json:"id"`
	// Archive date of the stat
	// in: string
	// example: 2023-01
	// required: true
	ArchiveDate string `json:"archive_date"`
	// Number of event by location
	// in: json
	// example: {"MDE": 12, "MDEL": 4}
	// read only: true
	NbPerLocation json.RawMessage `json:"nb_per_location"`
	// Number of event by type
	// in: json
	// example: {"NEM": 12, "Gala": 4}
	// read only: true
	NbPerType json.RawMessage `json:"nb_per_type"`
	// Number of event by price
	// in: json
	// example: {"10": 2, "5": 4, "0": 1}
	// read only: true
	NbPerPrice json.RawMessage `json:"nb_per_price"`
	// Number of event cancel
	// in: int
	// example: 12
	// read only: true
	NbCancel int `json:"nb_cancel"`
	// Number of event total
	// in: int
	// example: 12
	// read only: true
	NbTotal int `json:"nb_total"`
	// Average fulfill per type of event
	// in: json
	// example: {"NEM": 99.9, "gala": 80}
	// read only: true
	FulfillAvgPerType json.RawMessage `json:"tx_avg_fulfill_per_type"`
	// Average attendee per type of event
	// in: json
	// example: {"NEM": 12.5, "gala": 5}
	// read only: true
	AvgAttPerType json.RawMessage `json:"nb_avg_att_per_type"`
	// Average staff per type of event
	// in: json
	// example: {"NEM": 12.5, "gala": 5}
	// read only: true
	AvgStaPerType json.RawMessage `json:"nb_avg_sta_per_type"`
	// Number of event attendees total
	// in: int
	// example: 12
	// read only: true
	NbAttTotal int `json:"nb_att_total"`
	// Number of event staffs total
	// in: int
	// example: 12
	// read only: true
	NbStaTotal int `json:"nb_sta_total"`
	// CreatedAt date of the stat
	// in: time.Time
	// read only: true
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt date of the stat
	// in: time.Time
	// read only: true
	UpdatedAt time.Time `json:"updated_at"`
}

type Event_Stats []Event_Stat

func NewMonthlyStat(stat *Event_Stat) {
	var err error
	stmt, _ := TheDb().Prepare(
		"INSERT INTO " + db_name_monthly_stat + `
			(archive_date, nb_per_location, nb_per_type, nb_per_price, nb_cancel, nb_total,
			tx_avg_fulfill_per_type, nb_avg_att_per_type, nb_avg_sta_per_type, nb_att_total, nb_sta_total,
			created_at, updated_at)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`,
	)
	_, err = stmt.Exec(
		stat.ArchiveDate,
		stat.NbPerLocation,
		stat.NbPerType,
		stat.NbPerPrice,
		stat.NbCancel,
		stat.NbTotal,
		stat.FulfillAvgPerType,
		stat.AvgAttPerType,
		stat.AvgStaPerType,
		stat.NbAttTotal,
		stat.NbStaTotal,
		stat.CreatedAt,
		stat.UpdatedAt,
	)
	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "can't create new stat.")
	}
}

func UpdateMonthlyStat(stat *Event_Stat) {
	stmt, err := TheDb().Prepare(
		"UPDATE " + db_name_monthly_stat + ` SET
			archive_date=?, nb_per_location=?, nb_per_type=?, nb_per_price=?, nb_cancel=?,
			nb_total=?, tx_avg_fulfill_per_type=?, nb_avg_att_per_type=?, nb_avg_sta_per_type=?,
			nb_att_total=?, nb_sta_total=?, updated_at=?
		WHERE id=?;`,
	)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with the db.", err)
	}

	_, err = stmt.Exec(
		stat.ArchiveDate,
		stat.NbPerLocation,
		stat.NbPerType,
		stat.NbPerPrice,
		stat.NbCancel,
		stat.NbTotal,
		stat.FulfillAvgPerType,
		stat.AvgAttPerType,
		stat.AvgStaPerType,
		stat.NbAttTotal,
		stat.NbStaTotal,
		stat.UpdatedAt,
		stat.Id,
	)

	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "stat can't be updated.", err)
	}
}

func AutoNewMonthlyStat(archive_date ...string) {
	var stat Event_Stat
	var err error
	stat.CreatedAt = time.Now()
	stat.UpdatedAt = time.Now()

	// Get data of the past
	if archive_date != nil {
		stat.ArchiveDate = archive_date[0]
	} else {
		stat.ArchiveDate = stat.CreatedAt.AddDate(0, -1, 0).Format("2006-01")
	}
	month_stats := FindEventsByDate(stat.ArchiveDate)
	if len(*month_stats) == 0 {
		TheLogger().LogWarning(log_name_monthly_stat, "no data for this date.", errors.New("no data"))
		return
	}

	// Merge data
	countLocation := make(map[string]int)
	countType := make(map[string]int)
	countPrice := make(map[string]int)
	countCancel := 0
	avgAttPercent := make(map[string]float64)
	fulfillAvgType := make(map[string]float64)
	avgAtt := make(map[string]float64)
	avgAttType := make(map[string]int)
	avgSta := make(map[string]float64)
	avgStaType := make(map[string]int)
	countAttTotal := 0
	countStaTotal := 0
	for _, event := range *month_stats {
		staffs := FindStaffByEventId(event.Id)
		countLocation[event.Location]++
		countType[event.Type]++
		countPrice[fmt.Sprintf("%.2f", event.Price)]++
		if !event.Actif {
			countCancel++
		}
		avgAttPercent[event.Type] += float64(
			float64(event.NbSpotsTaken)/float64(event.NbSpotsMax),
		) * 100
		avgAtt[event.Type] += float64(event.NbSpotsTaken)
		avgSta[event.Type] += float64(len(*staffs))
		countAttTotal += event.NbSpotsTaken
		countStaTotal += len(*staffs)
	}
	for key, value := range countType {
		fulfillAvgType[key] = math.Round(float64(avgAttPercent[key])/float64(value)*100) / 100
		avgAttType[key] = int(math.RoundToEven(avgAtt[key] / float64(value)))
		avgStaType[key] = int(math.RoundToEven(avgSta[key] / float64(value)))
	}

	stat.NbPerLocation, err = json.Marshal(countLocation)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbPerType, err = json.Marshal(countType)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbPerPrice, err = json.Marshal(countPrice)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbTotal = len(*month_stats)
	stat.NbCancel = countCancel
	stat.FulfillAvgPerType, err = json.Marshal(fulfillAvgType)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.AvgAttPerType, err = json.Marshal(avgAttType)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.AvgStaPerType, err = json.Marshal(avgStaType)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbAttTotal = countAttTotal
	stat.NbStaTotal = countStaTotal

	// Add data into db
	previous_gen := FindMonthlyStatByDate(stat.ArchiveDate)
	if previous_gen.Id != 0 {
		TheLogger().LogInfo(log_name_monthly_stat, "this date has already been generated.")
		stat.Id = previous_gen.Id
		UpdateMonthlyStat(&stat)
	} else {
		NewMonthlyStat(&stat)
	}
}

func FindMonthlyStatByDate(archive_date string) *Event_Stat {
	var stat Event_Stat

	rows := TheDb().QueryRow(
		"SELECT * FROM "+db_name_monthly_stat+" WHERE archive_date = ?;",
		archive_date,
	)
	err := rows.Scan(
		&stat.Id,
		&stat.ArchiveDate,
		&stat.NbPerLocation,
		&stat.NbPerType,
		&stat.NbPerPrice,
		&stat.NbCancel,
		&stat.NbTotal,
		&stat.FulfillAvgPerType,
		&stat.AvgAttPerType,
		&stat.AvgStaPerType,
		&stat.NbAttTotal,
		&stat.NbStaTotal,
		&stat.CreatedAt,
		&stat.UpdatedAt,
	)

	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "cards not found.")
	}

	return &stat
}

func AllMonthlyStats() *Event_Stats {
	var stats Event_Stats

	rows, err := TheDb().Query("SELECT * FROM " + db_name_monthly_stat)

	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var stat Event_Stat

		err := rows.Scan(
			&stat.Id,
			&stat.ArchiveDate,
			&stat.NbPerLocation,
			&stat.NbPerType,
			&stat.NbPerPrice,
			&stat.NbCancel,
			&stat.NbTotal,
			&stat.FulfillAvgPerType,
			&stat.AvgAttPerType,
			&stat.AvgStaPerType,
			&stat.NbAttTotal,
			&stat.NbStaTotal,
			&stat.CreatedAt,
			&stat.UpdatedAt,
		)

		if err != nil {
			TheLogger().LogInfo(log_name_monthly_stat, "stats not found.")
		}

		stats = append(stats, stat)
	}

	return &stats
}
