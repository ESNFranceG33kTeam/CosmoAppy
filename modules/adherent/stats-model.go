package adherent

import (
	"encoding/json"
	"errors"
	"time"
)

// swagger:model Adherent_Stat
type Adherent_Stat struct {
	// Id of the entry
	// in: int
	// read only: true
	Id int `json:"id"`
	// Archive date of the stat
	// in: string
	// example: 2023-01
	// required: true
	ArchiveDate string `json:"archive_date"`
	// Number of adherent by country
	// in: json
	// example: {"Poland": 12, "Belgium": 4}
	// read only: true
	NbPerCountry json.RawMessage `json:"nb_per_country"`
	// Number of adherent by university
	// in: json
	// example: {"Lumiere": 1, "SciencePo": 4}
	// read only: true
	NbPerUniversity json.RawMessage `json:"nb_per_university"`
	// Number of adherent by situation
	// in: json
	// example: {"student": 12, "worker": 4}
	// read only: true
	NbPerSituation json.RawMessage `json:"nb_per_situation"`
	// Number of adherent total
	// in: int
	// example: 12
	// read only: true
	NbTotal int `json:"nb_total"`
	// Number of how adherent heared about us per type
	// in: json
	// example: {"twitter": 12, "facebook": 4}
	// read only: true
	AboutusPerType json.RawMessage `json:"aboutus_per_type"`
	// CreatedAt date of the stat
	// in: time.Time
	// read only: true
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt date of the stat
	// in: time.Time
	// read only: true
	UpdatedAt time.Time `json:"updated_at"`
}

type Adherent_Stats []Adherent_Stat

func NewMonthlyStat(stat *Adherent_Stat) {
	var err error
	stmt, _ := TheDb().Prepare(
		"INSERT INTO " + db_name_monthly_stat + `
			(archive_date, nb_per_country, nb_per_university, nb_per_situation, nb_total,
			aboutus_per_type, created_at, updated_at)
		VALUES (?,?,?,?,?,?,?,?);`,
	)
	_, err = stmt.Exec(
		stat.ArchiveDate,
		stat.NbPerCountry,
		stat.NbPerUniversity,
		stat.NbPerSituation,
		stat.NbTotal,
		stat.AboutusPerType,
		stat.CreatedAt,
		stat.UpdatedAt,
	)
	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "can't create new stat.")
	}
}

func UpdateMonthlyStat(stat *Adherent_Stat) {
	stmt, err := TheDb().Prepare(
		"UPDATE " + db_name_monthly_stat + ` SET
			archive_date=?, nb_per_country=?, nb_per_university=?, nb_per_situation=?,
			nb_total=?, aboutus_per_type=?, updated_at=?
		WHERE id=?;`,
	)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with the db.", err)
	}

	_, err = stmt.Exec(
		stat.ArchiveDate,
		stat.NbPerCountry,
		stat.NbPerUniversity,
		stat.NbPerSituation,
		stat.NbTotal,
		stat.AboutusPerType,
		stat.UpdatedAt,
		stat.Id,
	)

	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "stat can't be updated.", err)
	}
}

func AutoNewMonthlyStat(archive_date ...string) {
	var stat Adherent_Stat
	var err error
	stat.CreatedAt = time.Now()
	stat.UpdatedAt = time.Now()

	// Get data of the past
	if archive_date != nil {
		stat.ArchiveDate = archive_date[0]
	} else {
		stat.ArchiveDate = stat.CreatedAt.AddDate(0, -1, 0).Format("2006-01")
	}
	month_stats := FindAdherentsByDate(stat.ArchiveDate)
	if len(*month_stats) == 0 {
		TheLogger().LogWarning(log_name_monthly_stat, "no data for this date.", errors.New("no data"))
		return
	}

	// Merge data
	countCountry := make(map[string]int)
	countUniversity := make(map[string]int)
	countSituation := make(map[string]int)
	countAboutus := make(map[string]int)
	for _, adherent := range *month_stats {
		countCountry[adherent.Homeland]++
		countUniversity[adherent.University]++
		countSituation[adherent.Situation]++
		countAboutus[adherent.Speakabout]++
	}

	stat.NbPerCountry, err = json.Marshal(countCountry)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbPerUniversity, err = json.Marshal(countUniversity)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbPerSituation, err = json.Marshal(countSituation)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbTotal = len(*month_stats)
	stat.AboutusPerType, err = json.Marshal(countAboutus)
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

func FindMonthlyStatByDate(archive_date string) *Adherent_Stat {
	var stat Adherent_Stat

	rows := TheDb().QueryRow(
		"SELECT * FROM "+db_name_monthly_stat+" WHERE archive_date = ?;",
		archive_date,
	)
	err := rows.Scan(
		&stat.Id,
		&stat.ArchiveDate,
		&stat.NbPerCountry,
		&stat.NbPerUniversity,
		&stat.NbPerSituation,
		&stat.NbTotal,
		&stat.AboutusPerType,
		&stat.CreatedAt,
		&stat.UpdatedAt,
	)

	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "cards not found.")
	}

	return &stat
}

func AllMonthlyStats() *Adherent_Stats {
	var stats Adherent_Stats

	rows, err := TheDb().Query("SELECT * FROM " + db_name_monthly_stat)

	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var stat Adherent_Stat

		err := rows.Scan(
			&stat.Id,
			&stat.ArchiveDate,
			&stat.NbPerCountry,
			&stat.NbPerUniversity,
			&stat.NbPerSituation,
			&stat.NbTotal,
			&stat.AboutusPerType,
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
