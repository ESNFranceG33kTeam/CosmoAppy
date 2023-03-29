package report

import (
	"encoding/json"
	"errors"
	"time"
)

// swagger:model Report_Stat
type Report_Stat struct {
	// Id of the entry
	// in: int
	// read only: true
	Id int `json:"id"`
	// Archive date of the stat
	// in: string
	// example: 2023-01
	// required: true
	ArchiveDate string `json:"archive_date"`
	// Number of report by code public
	// in: json
	// example: {"CHAMP": 12, "MAST": 4}
	// read only: true
	NbPerCodePublic json.RawMessage `json:"nb_per_codepublic"`
	// Number of report by code project
	// in: json
	// example: {"PKM": 1, "SMI": 4}
	// read only: true
	NbPerCodeProject json.RawMessage `json:"nb_per_codeproject"`
	// Number of report by type
	// in: json
	// example: {"student": 12, "worker": 4}
	// read only: true
	NbPerType json.RawMessage `json:"nb_per_type"`
	// Number of report total
	// in: int
	// example: 12
	// read only: true
	NbTotal int `json:"nb_total"`
	// Number of hours for a volunteer by code
	// in: json
	// example: {"PKM": 12, "SMI": 4}
	// read only: true
	HoursVltCodes json.RawMessage `json:"hours_vlt_per_codes"`
	// Number of hours for an employee by code
	// in: json
	// example: {"PKM": 12, "SMI": 4}
	// read only: true
	HoursEmpCodes json.RawMessage `json:"hours_emp_per_codes"`
	// Number of hours for a scv by code
	// in: json
	// example: {"PKM": 12, "SMI": 4}
	// read only: true
	HoursScvCodes json.RawMessage `json:"hours_scv_per_codes"`
	// Valorisation for a volunteer by code
	// in: json
	// example: {"PKM": 12, "SMI": 4}
	// read only: true
	ValoVltCodes json.RawMessage `json:"valo_vlt_per_codes"`
	// Valorisation for an employee by code
	// in: json
	// example: {"PKM": 12, "SMI": 4}
	// read only: true
	ValoEmpCodes json.RawMessage `json:"valo_emp_per_codes"`
	// Valorisation for a scv by code
	// in: json
	// example: {"PKM": 12, "SMI": 4}
	// read only: true
	ValoScvCodes json.RawMessage `json:"valo_scv_per_codes"`
	// CreatedAt date of the stat
	// in: time.Time
	// read only: true
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt date of the stat
	// in: time.Time
	// read only: true
	UpdatedAt time.Time `json:"updated_at"`
}

type Report_Stats []Report_Stat

func NewMonthlyStat(stat *Report_Stat) {
	var err error
	stmt, _ := TheDb().Prepare(
		"INSERT INTO " + db_name_monthly_stat + `
			(archive_date, nb_per_codepublic, nb_per_codeproject, nb_per_type, nb_total,
			hours_vlt_per_codes, hours_emp_per_codes, hours_scv_per_codes, valo_vlt_per_codes,
			valo_emp_per_codes, valo_scv_per_codes, created_at, updated_at)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`,
	)
	_, err = stmt.Exec(
		stat.ArchiveDate,
		stat.NbPerCodePublic,
		stat.NbPerCodeProject,
		stat.NbPerType,
		stat.NbTotal,
		stat.HoursVltCodes,
		stat.HoursEmpCodes,
		stat.HoursScvCodes,
		stat.ValoVltCodes,
		stat.ValoEmpCodes,
		stat.ValoScvCodes,
		stat.CreatedAt,
		stat.UpdatedAt,
	)
	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "can't create new stat.")
	}
}

func UpdateMonthlyStat(stat *Report_Stat) {
	stmt, err := TheDb().Prepare(
		"UPDATE " + db_name_monthly_stat + ` SET
			archive_date=?, nb_per_codepublic=?, nb_per_codeproject=?, nb_per_type=?, nb_total=?,
			hours_vlt_per_codes=?, hours_emp_per_codes=?, hours_scv_per_codes=?,
			valo_vlt_per_codes=?, valo_emp_per_codes=?, valo_scv_per_codes=?, updated_at=?
		WHERE id=?;`,
	)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with the db.", err)
	}

	_, err = stmt.Exec(
		stat.ArchiveDate,
		stat.NbPerCodePublic,
		stat.NbPerCodeProject,
		stat.NbPerType,
		stat.NbTotal,
		stat.HoursVltCodes,
		stat.HoursEmpCodes,
		stat.HoursScvCodes,
		stat.ValoVltCodes,
		stat.ValoEmpCodes,
		stat.ValoScvCodes,
		stat.UpdatedAt,
		stat.Id,
	)

	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "stat can't be updated.", err)
	}
}

func AutoNewMonthlyStat(archive_date ...string) {
	var stat Report_Stat
	var err error
	stat.CreatedAt = time.Now()
	stat.UpdatedAt = time.Now()

	// Get data of the past
	if archive_date != nil {
		stat.ArchiveDate = archive_date[0]
	} else {
		stat.ArchiveDate = stat.CreatedAt.AddDate(0, -1, 0).Format("2006-01")
	}
	month_stats := FindReportsByDate(stat.ArchiveDate)
	if len(*month_stats) == 0 {
		TheLogger().LogWarning(log_name_monthly_stat, "no data for this date.", errors.New("no data"))
		return
	}

	// Merge data
	countCodePublic := make(map[string]int)
	countCodeProject := make(map[string]int)
	countType := make(map[string]int)
	hoursVltCodes := make(map[string]map[string]float64)
	hoursVltCodes["code_public"] = make(map[string]float64)
	hoursVltCodes["code_project"] = make(map[string]float64)

	hoursEmpCodes := make(map[string]map[string]float64)
	hoursEmpCodes["code_public"] = make(map[string]float64)
	hoursEmpCodes["code_project"] = make(map[string]float64)

	hoursScvCodes := make(map[string]map[string]float64)
	hoursScvCodes["code_public"] = make(map[string]float64)
	hoursScvCodes["code_project"] = make(map[string]float64)

	valoVltCodes := make(map[string]map[string]float64)
	valoVltCodes["code_public"] = make(map[string]float64)
	valoVltCodes["code_project"] = make(map[string]float64)

	valoEmpCodes := make(map[string]map[string]float64)
	valoEmpCodes["code_public"] = make(map[string]float64)
	valoEmpCodes["code_project"] = make(map[string]float64)

	valoScvCodes := make(map[string]map[string]float64)
	valoScvCodes["code_public"] = make(map[string]float64)
	valoScvCodes["code_project"] = make(map[string]float64)

	for _, report := range *month_stats {
		countCodePublic[report.CodePublic]++
		countCodeProject[report.CodeProject]++
		countType[report.Type]++

		hoursVltCodes["code_public"][report.CodePublic] += report.NbHours * report.NbStaffsVlt
		hoursVltCodes["code_project"][report.CodeProject] += report.NbHours * report.NbStaffsVlt
		hoursEmpCodes["code_public"][report.CodePublic] += report.NbHours * report.NbStaffsEmp
		hoursEmpCodes["code_project"][report.CodeProject] += report.NbHours * report.NbStaffsEmp
		hoursScvCodes["code_public"][report.CodePublic] += report.NbHours * report.NbStaffsScv
		hoursScvCodes["code_project"][report.CodeProject] += report.NbHours * report.NbStaffsScv

		valoVltCodes["code_public"][report.CodePublic] += report.TauxValorisationVlt * report.NbStaffsVlt
		valoVltCodes["code_project"][report.CodeProject] += report.TauxValorisationVlt * report.NbStaffsVlt
		valoEmpCodes["code_public"][report.CodePublic] += report.TauxValorisationEmp * report.NbStaffsEmp
		valoEmpCodes["code_project"][report.CodeProject] += report.TauxValorisationEmp * report.NbStaffsEmp
		valoScvCodes["code_public"][report.CodePublic] += report.TauxValorisationScv * report.NbStaffsScv
		valoScvCodes["code_project"][report.CodeProject] += report.TauxValorisationScv * report.NbStaffsScv

	}

	stat.NbPerCodePublic, err = json.Marshal(countCodePublic)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbPerCodeProject, err = json.Marshal(countCodeProject)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbPerType, err = json.Marshal(countType)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.NbTotal = len(*month_stats)
	stat.HoursVltCodes, err = json.Marshal(hoursVltCodes)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.HoursEmpCodes, err = json.Marshal(hoursEmpCodes)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.HoursScvCodes, err = json.Marshal(hoursScvCodes)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.ValoVltCodes, err = json.Marshal(valoVltCodes)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.ValoEmpCodes, err = json.Marshal(valoEmpCodes)
	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with encoder.", err)
	}
	stat.ValoScvCodes, err = json.Marshal(valoScvCodes)
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

func FindMonthlyStatByDate(archive_date string) *Report_Stat {
	var stat Report_Stat

	rows := TheDb().QueryRow(
		"SELECT * FROM "+db_name_monthly_stat+" WHERE archive_date = ?;",
		archive_date,
	)
	err := rows.Scan(
		&stat.Id,
		&stat.ArchiveDate,
		&stat.NbPerCodePublic,
		&stat.NbPerCodeProject,
		&stat.NbPerType,
		&stat.NbTotal,
		&stat.HoursVltCodes,
		&stat.HoursEmpCodes,
		&stat.HoursScvCodes,
		&stat.ValoVltCodes,
		&stat.ValoEmpCodes,
		&stat.ValoScvCodes,
		&stat.CreatedAt,
		&stat.UpdatedAt,
	)

	if err != nil {
		TheLogger().LogInfo(log_name_monthly_stat, "cards not found.")
	}

	return &stat
}

func AllMonthlyStats() *Report_Stats {
	var stats Report_Stats

	rows, err := TheDb().Query("SELECT * FROM " + db_name_monthly_stat)

	if err != nil {
		TheLogger().LogError(log_name_monthly_stat, "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var stat Report_Stat

		err := rows.Scan(
			&stat.Id,
			&stat.ArchiveDate,
			&stat.NbPerCodePublic,
			&stat.NbPerCodeProject,
			&stat.NbPerType,
			&stat.NbTotal,
			&stat.HoursVltCodes,
			&stat.HoursEmpCodes,
			&stat.HoursScvCodes,
			&stat.ValoVltCodes,
			&stat.ValoEmpCodes,
			&stat.ValoScvCodes,
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
