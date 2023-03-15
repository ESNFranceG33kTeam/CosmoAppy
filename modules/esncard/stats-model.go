package esncard

import (
	"errors"
	"time"
)

// swagger:model ESNcard_Stat
type ESNcard_Stat struct {
	// Id of the entry
	// in: int
	// read only: true
	Id int `json:"id"`
	// Archive date of the stat
	// in: string
	// example: 2023-01
	// required: true
	ArchiveDate string `json:"archive_date"`
	// Number of esncard sold
	// in: json
	// example: 12
	// read only: true
	NbESNcardSold int `json:"nb_esncard_sold"`
	// CreatedAt date of the stat
	// in: time.Time
	// read only: true
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt date of the stat
	// in: time.Time
	// read only: true
	UpdatedAt time.Time `json:"updated_at"`
}

type ESNcard_Stats []ESNcard_Stat

func NewMonthlyStat(stat *ESNcard_Stat) {
	var err error
	stmt, _ := TheDb().Prepare(
		"INSERT INTO " + db_name_monthly_stat + `
			(archive_date, nb_esncard_sold, created_at, updated_at)
		VALUES (?,?,?,?);`,
	)
	_, err = stmt.Exec(
		stat.ArchiveDate,
		stat.NbESNcardSold,
		stat.CreatedAt,
		stat.UpdatedAt,
	)
	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "can't create new card.")
	}
}

func UpdateMonthlyStat(stat *ESNcard_Stat) {
	stmt, err := TheDb().Prepare(
		"UPDATE " + db_name_monthly_stat + ` SET
			archive_date=?, nb_esncard_sold=?, updated_at=?
		WHERE id=?;`,
	)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with the db.", err)
	}

	_, err = stmt.Exec(
		stat.ArchiveDate,
		stat.NbESNcardSold,
		stat.UpdatedAt,
		stat.Id,
	)

	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "stat can't be updated.", err)
	}
}

func AutoNewMonthlyStat(archive_date ...string) {
	var stat ESNcard_Stat
	stat.CreatedAt = time.Now()
	stat.UpdatedAt = time.Now()

	// Get data of the past
	if archive_date != nil {
		stat.ArchiveDate = archive_date[0]
	} else {
		stat.ArchiveDate = stat.CreatedAt.AddDate(0, -1, 0).Format("2006-01")
	}
	month_stats := FindESNcardsByDate(stat.ArchiveDate)
	if len(*month_stats) == 0 {
		TheLogger().LogWarning(log_name_monthly_stat, "no data for this date.", errors.New("no data"))
		return
	}

	// Merge data
	stat.NbESNcardSold = len(*month_stats)

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

func FindMonthlyStatByDate(archive_date string) *ESNcard_Stat {
	var stat ESNcard_Stat

	rows := TheDb().QueryRow(
		"SELECT * FROM "+db_name_monthly_stat+" WHERE archive_date = ?;",
		archive_date,
	)
	err := rows.Scan(
		&stat.Id,
		&stat.ArchiveDate,
		&stat.NbESNcardSold,
		&stat.CreatedAt,
		&stat.UpdatedAt,
	)

	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "cards not found.")
	}

	return &stat
}

func AllMonthlyStats() *ESNcard_Stats {
	var stats ESNcard_Stats

	rows, err := TheDb().Query("SELECT * FROM " + db_name_monthly_stat)

	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var stat ESNcard_Stat

		err := rows.Scan(
			&stat.Id,
			&stat.ArchiveDate,
			&stat.NbESNcardSold,
			&stat.CreatedAt,
			&stat.UpdatedAt,
		)

		if err != nil {
			TheLogger().LogInfo(log_name_monthly_stat, "cards not found.")
		}

		stats = append(stats, stat)
	}

	return &stats
}
