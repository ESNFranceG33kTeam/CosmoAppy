package volunteer

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// swagger:model Volunteer_Stat
type Volunteer_Stat struct {
	// Id of the entry
	// in: int
	// read only: true
	Id int `json:"id"`
	// Archive date of the stat
	// in: string
	// example: 2023-01
	// required: true
	ArchiveDate string `json:"archive_date"`
	// Number of active volunteer by university
	// in: json
	// example: {"Lumiere": 1, "SciencePo": 4}
	// read only: true
	NbActivePerUniversity json.RawMessage `json:"nb_active_per_university"`
	// Number of active volunteer by hr status
	// in: json
	// example: {"volunteer": 1, "employee": 4}
	// read only: true
	NbActivePerHrstatus json.RawMessage `json:"nb_active_per_hrstatus"`
	// Number of new volunteer this month
	// in: int
	// example: 12
	// read only: true
	NbNewVlt int `json:"nb_new_vlt_this_month"`
	// Number of alumnus
	// in: int
	// example: 12
	// read only: true
	NbAlumnus int `json:"nb_alumnus"`
	// Number of volunteer in the bureau
	// in: int
	// example: 12
	// read only: true
	NbBureau int `json:"nb_bureau"`
	// CreatedAt date of the stat
	// in: time.Time
	// read only: true
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt date of the stat
	// in: time.Time
	// read only: true
	UpdatedAt time.Time `json:"updated_at"`
}

type Volunteer_Stats []Volunteer_Stat

func NewMonthlyStat(stat *Volunteer_Stat) {
	var err error
	stmt, _ := TheDb().Prepare(
		"INSERT INTO " + db_name_monthly_stat + `
			(archive_date, nb_active_per_university, nb_active_per_hrstatus, nb_new_vlt_this_month,
			nb_alumnus, nb_bureau, created_at, updated_at)
		VALUES (?,?,?,?,?,?,?,?);`,
	)
	_, err = stmt.Exec(
		stat.ArchiveDate,
		stat.NbActivePerUniversity,
		stat.NbActivePerHrstatus,
		stat.NbNewVlt,
		stat.NbAlumnus,
		stat.NbBureau,
		stat.CreatedAt,
		stat.UpdatedAt,
	)
	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "can't create new stat.")
	}
}

func UpdateMonthlyStat(stat *Volunteer_Stat) {
	stmt, err := TheDb().Prepare(
		"UPDATE " + db_name_monthly_stat + ` SET
			archive_date=?, nb_active_per_university=?, nb_active_per_hrstatus=?,
			nb_new_vlt_this_month=?, nb_alumnus=?, nb_bureau=?, updated_at=?
		WHERE id=?;`,
	)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with the db.", err)
	}

	_, err = stmt.Exec(
		stat.ArchiveDate,
		stat.NbActivePerUniversity,
		stat.NbActivePerHrstatus,
		stat.NbNewVlt,
		stat.NbAlumnus,
		stat.NbBureau,
		stat.UpdatedAt,
		stat.Id,
	)

	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "stat can't be updated.", err)
	}
}

func AutoNewMonthlyStat() {
	var stat Volunteer_Stat
	var err error
	stat.CreatedAt = time.Now()
	stat.UpdatedAt = time.Now()

	// Get data of the past
	stat.ArchiveDate = stat.CreatedAt.AddDate(0, -1, 0).Format("2006-01")

	month_stats := AllVolunteers()
	if len(*month_stats) == 0 {
		TheLogger().LogWarning(log_name_monthly_stat, "no data for this date.", errors.New("no data"))
		return
	}

	// Merge data
	countActiveUniversity := make(map[string]int)
	countActiveStatus := make(map[string]int)
	countNewVlt := 0
	countAlumnus := 0
	countBureau := 0
	for _, volunteer := range *month_stats {
		if volunteer.Actif {
			countActiveUniversity[volunteer.University]++
			countActiveStatus[volunteer.HrStatus]++
		}
		if volunteer.HrStatus == "volunteer" {
			if strings.Contains(volunteer.StartedDate, stat.ArchiveDate) {
				countNewVlt++
			}
			if !volunteer.Actif {
				countAlumnus++
			}
			if volunteer.Bureau {
				countBureau++
			}
		}
	}

	stat.NbActivePerUniversity, err = json.Marshal(countActiveUniversity)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbActivePerHrstatus, err = json.Marshal(countActiveStatus)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbNewVlt = countNewVlt
	stat.NbAlumnus = countAlumnus
	stat.NbBureau = countBureau

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

func FindMonthlyStatByDate(archive_date string) *Volunteer_Stat {
	var stat Volunteer_Stat

	rows := TheDb().QueryRow(
		"SELECT * FROM "+db_name_monthly_stat+" WHERE archive_date = ?;",
		archive_date,
	)
	err := rows.Scan(
		&stat.Id,
		&stat.ArchiveDate,
		&stat.NbActivePerUniversity,
		&stat.NbActivePerHrstatus,
		&stat.NbNewVlt,
		&stat.NbAlumnus,
		&stat.NbBureau,
		&stat.CreatedAt,
		&stat.UpdatedAt,
	)

	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "cards not found.")
	}

	return &stat
}

func AllMonthlyStats() *Volunteer_Stats {
	var stats Volunteer_Stats

	rows, err := TheDb().Query("SELECT * FROM " + db_name_monthly_stat)

	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var stat Volunteer_Stat

		err := rows.Scan(
			&stat.Id,
			&stat.ArchiveDate,
			&stat.NbActivePerUniversity,
			&stat.NbActivePerHrstatus,
			&stat.NbNewVlt,
			&stat.NbAlumnus,
			&stat.NbBureau,
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
