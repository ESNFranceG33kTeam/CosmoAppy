package money

import (
	"encoding/json"
	"errors"
	"time"
)

// swagger:model Money_Stats
type Money_Stat struct {
	// Id of the entry
	// in: int
	// read only: true
	Id int `json:"id"`
	// Archive date of the stat
	// in: string
	// example: 2023-01
	// required: true
	ArchiveDate string `json:"archive_date"`
	// Number of each label
	// in: json
	// example: {"event": 12, "esncard": 4}
	// read only: true
	NbPerLabel json.RawMessage `json:"nb_per_label"`
	// Average prices per type of label
	// in: json
	// example: {"adhesion": 12.5, "esncard": 5}
	// read only: true
	AvgPrices json.RawMessage `json:"avg_prices"`
	// Min prices per type of label
	// in: json
	// example: {"adhesion": 1, "esncard": 5}
	// read only: true
	MinPrices json.RawMessage `json:"min_prices"`
	// Max prices per type of label
	// in: json
	// example: {"adhesion": 20, "esncard": 5}
	// read only: true
	MaxPrices json.RawMessage `json:"max_prices"`
	// Number of Payments type
	// in: json
	// example: {"cash": 12, "cb": 5}
	// read only: true
	NbPaymentTypes json.RawMessage `json:"nb_payments_type"`
	// CreatedAt date of the operation
	// in: time.Time
	// read only: true
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt date of the operation
	// in: time.Time
	// read only: true
	UpdatedAt time.Time `json:"updated_at"`
}

type Money_Stats []Money_Stat

func NewMonthlyStat(stat *Money_Stat) {
	var err error
	stmt, _ := TheDb().Prepare(
		`INSERT INTO ` + db_name_monthly_stat + `
			(archive_date, nb_per_label, avg_prices, min_prices, max_prices, nb_payments_type,
			created_at, updated_at)
		VALUES (?,?,?,?,?,?,?,?);`,
	)
	_, err = stmt.Exec(
		stat.ArchiveDate,
		stat.NbPerLabel,
		stat.AvgPrices,
		stat.MinPrices,
		stat.MaxPrices,
		stat.NbPaymentTypes,
		stat.CreatedAt,
		stat.UpdatedAt,
	)
	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "can't create new stat.")
	}
}

func UpdateMonthlyStat(stat *Money_Stat) {
	stmt, err := TheDb().Prepare(
		`UPDATE ` + db_name_monthly_stat + ` SET
			archive_date=?, nb_per_label=?, avg_prices=?, min_prices=?, max_prices=?,
			nb_payments_type=?, updated_at=?
		WHERE id=?;`,
	)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with the db.", err)
	}

	_, err = stmt.Exec(
		stat.ArchiveDate,
		stat.NbPerLabel,
		stat.AvgPrices,
		stat.MinPrices,
		stat.MaxPrices,
		stat.NbPaymentTypes,
		stat.UpdatedAt,
		stat.Id,
	)

	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "stat can't be updated.")
	}
}

func AutoNewMonthlyStat(archive_date ...string) {
	var stat Money_Stat
	var err error
	stat.CreatedAt = time.Now()
	stat.UpdatedAt = time.Now()

	// Get data of the past
	if archive_date != nil {
		stat.ArchiveDate = archive_date[0]
	} else {
		stat.ArchiveDate = stat.CreatedAt.AddDate(0, -1, 0).Format("2006-01")
	}
	month_stats := FindMoneysByDate(stat.ArchiveDate)
	if len(*month_stats) == 0 {
		TheLogger().LogWarning(log_name_monthly_stat, "no data for this date.", errors.New("no data"))
		return
	}

	// Merge data
	countLabels := make(map[string]int)
	sumPrices := make(map[string]int)
	avgPrices := make(map[string]float64)
	minPrices := make(map[string]float64)
	maxPrices := make(map[string]float64)
	countPaymentsType := make(map[string]int)
	for _, money := range *month_stats {
		countLabels[money.Label]++
		sumPrices[money.Label] += int(money.Price)
		if minPrices[money.Label] == 0 || money.Price < minPrices[money.Label] {
			minPrices[money.Label] = money.Price
		}
		if money.Price > maxPrices[money.Label] {
			maxPrices[money.Label] = money.Price
		}
		countPaymentsType[money.PaymentType]++
	}
	for key, value := range countLabels {
		avgPrices[key] = float64(sumPrices[key]) / float64(value)
	}

	stat.NbPerLabel, err = json.Marshal(countLabels)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.AvgPrices, err = json.Marshal(avgPrices)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.MinPrices, err = json.Marshal(minPrices)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.MaxPrices, err = json.Marshal(maxPrices)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbPaymentTypes, err = json.Marshal(countPaymentsType)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}

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

func FindMonthlyStatByDate(archive_date string) *Money_Stat {
	var stat Money_Stat

	rows := TheDb().QueryRow(
		"SELECT * FROM "+db_name_monthly_stat+" WHERE archive_date = ?;",
		archive_date,
	)
	err := rows.Scan(
		&stat.Id,
		&stat.ArchiveDate,
		&stat.NbPerLabel,
		&stat.AvgPrices,
		&stat.MinPrices,
		&stat.MaxPrices,
		&stat.NbPaymentTypes,
		&stat.CreatedAt,
		&stat.UpdatedAt,
	)

	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "operations not found.")
	}

	return &stat
}

func AllMonthlyStats() *Money_Stats {
	var stats Money_Stats

	rows, err := TheDb().Query("SELECT * FROM " + db_name_monthly_stat)

	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var stat Money_Stat

		err := rows.Scan(
			&stat.Id,
			&stat.ArchiveDate,
			&stat.NbPerLabel,
			&stat.AvgPrices,
			&stat.MinPrices,
			&stat.MaxPrices,
			&stat.NbPaymentTypes,
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
