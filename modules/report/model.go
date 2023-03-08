package report

// swagger:model Report
type Report struct {
	// Id of the report
	// in: int
	Id int `json:"id"`
	// Type of the report
	// in: string
	Type string `json:"type"`
	// External reference of the report
	// in: int
	RefExt int `json:"ref_ext"`
	// Name of the report
	// in: string
	Name string `json:"name"`
	// Date of the event/planning/custom
	// in: string
	Date string `json:"date"`
	// Comment of the report
	// in: string
	Comment string `json:"comment"`
	// Number reel of attendees of the event/planning/custom
	// in: int64
	NbReelAtt int `json:"nb_reel_attendees"`
	// Number subscribes of attendees of the event/planning/custom
	// in: int
	NbSubsAtt int `json:"nb_subscribe_attendees"`
	// List of staffs of the event/planning/custom
	// in: string
	StaffsList string `json:"staffs_list"`
	// Number of hours of preparation
	// in: float
	NbHoursPrepa float64 `json:"nb_hours_prepa"`
	// Number of hours of the event/planning/custom
	// in: float
	NbHours float64 `json:"nb_hours"`
	// Number of volunteer staffs of the event/planning/custom
	// in: float
	NbStaffsVlt float64 `json:"nb_staffs_vlt"`
	// Number of employee staffs of the event/planning/custom
	// in: float
	NbStaffsEmp float64 `json:"nb_staffs_emp"`
	// Number of Civic service staffs of the event/planning/custom
	// in: float
	NbStaffsScv float64 `json:"nb_staffs_scv"`
	// Taux volunteer Valorisation of the event/planning/custom
	// in: float
	TauxValorisationVlt float64 `json:"taux_valorisation_vlt"`
	// Taux employee Valorisation of the event/planning/custom
	// in: float
	TauxValorisationEmp float64 `json:"taux_valorisation_emp"`
	// Taux civic service Valorisation of the event/planning/custom
	// in: float
	TauxValorisationScv float64 `json:"taux_valorisation_scv"`
	// Code public of the event/planning/custom
	// in: string
	CodePublic string `json:"code_public"`
	// Code projet of the event/planning/custom
	// in: string
	CodeProject string `json:"code_project"`
}

type Reports []Report

func NewReport(rpt *Report) {
	stmt, _ := TheDb().Prepare(`
		INSERT INTO reports
			(type, ref_ext, name, date, comment, nb_reel_attendees, nb_subscribe_attendees,
			staffs_list, nb_hours_prepa, nb_hours, nb_staffs_vlt, nb_staffs_emp, nb_staffs_scv,
			taux_valorisation_vlt, taux_valorisation_emp, taux_valorisation_scv, code_public,
			code_project)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`,
	)
	_, err := stmt.Exec(
		rpt.Type,
		rpt.RefExt,
		rpt.Name,
		rpt.Date,
		rpt.Comment,
		rpt.NbReelAtt,
		rpt.NbSubsAtt,
		rpt.StaffsList,
		rpt.NbHoursPrepa,
		rpt.NbHours,
		rpt.NbStaffsVlt,
		rpt.NbStaffsEmp,
		rpt.NbStaffsScv,
		rpt.TauxValorisationVlt,
		rpt.TauxValorisationEmp,
		rpt.TauxValorisationScv,
		rpt.CodePublic,
		rpt.CodeProject,
	)
	if err != nil {
		TheLogger().LogError("report", "can't create new report.", err)
	}
}

func FindReportById(id int) *Report {
	var rpt Report

	row := TheDb().QueryRow("SELECT * FROM reports WHERE id = ?;", id)
	err := row.Scan(
		&rpt.Id,
		&rpt.Type,
		&rpt.RefExt,
		&rpt.Name,
		&rpt.Date,
		&rpt.Comment,
		&rpt.NbReelAtt,
		&rpt.NbSubsAtt,
		&rpt.StaffsList,
		&rpt.NbHoursPrepa,
		&rpt.NbHours,
		&rpt.NbStaffsVlt,
		&rpt.NbStaffsEmp,
		&rpt.NbStaffsScv,
		&rpt.TauxValorisationVlt,
		&rpt.TauxValorisationEmp,
		&rpt.TauxValorisationScv,
		&rpt.CodePublic,
		&rpt.CodeProject,
	)

	if err != nil {
		TheLogger().LogWarning("report", "report id not found.", err)
	}

	return &rpt
}

func AllReports() *Reports {
	var rpts Reports

	rows, err := TheDb().Query("SELECT * FROM reports")

	if err != nil {
		TheLogger().LogError("report", "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var rpt Report

		err := rows.Scan(
			&rpt.Id,
			&rpt.Type,
			&rpt.RefExt,
			&rpt.Name,
			&rpt.Date,
			&rpt.Comment,
			&rpt.NbReelAtt,
			&rpt.NbSubsAtt,
			&rpt.StaffsList,
			&rpt.NbHoursPrepa,
			&rpt.NbHours,
			&rpt.NbStaffsVlt,
			&rpt.NbStaffsEmp,
			&rpt.NbStaffsScv,
			&rpt.TauxValorisationVlt,
			&rpt.TauxValorisationEmp,
			&rpt.TauxValorisationScv,
			&rpt.CodePublic,
			&rpt.CodeProject,
		)

		if err != nil {
			TheLogger().LogError("report", "reports not found.", err)
		}

		rpts = append(rpts, rpt)
	}

	return &rpts
}

func UpdateReport(rpt *Report) {
	stmt, err := TheDb().Prepare(
		`UPDATE reports SET
			type=?, ref_ext=?, name=?, date=?, comment=?, nb_reel_attendees=?,
			nb_subscribe_attendees=?, staffs_list=?, nb_hours_prepa=?, nb_hours=?, nb_staffs_vlt=?,
			nb_staffs_emp=?, nb_staffs_scv=?, taux_valorisation_vlt=?, taux_valorisation_emp=?,
			taux_valorisation_scv=?, code_public=?, code_project=?
		WHERE id=?;`,
	)

	if err != nil {
		TheLogger().LogError("report", "problem with the db.", err)
	}

	_, err = stmt.Exec(
		rpt.Type,
		rpt.RefExt,
		rpt.Name,
		rpt.Date,
		rpt.Comment,
		rpt.NbReelAtt,
		rpt.NbSubsAtt,
		rpt.StaffsList,
		rpt.NbHoursPrepa,
		rpt.NbHours,
		rpt.NbStaffsVlt,
		rpt.NbStaffsEmp,
		rpt.NbStaffsScv,
		rpt.TauxValorisationVlt,
		rpt.TauxValorisationEmp,
		rpt.TauxValorisationScv,
		rpt.CodePublic,
		rpt.CodeProject,
		rpt.Id,
	)

	if err != nil {
		TheLogger().LogError("report", "report can't be updated.", err)
	}
}

func DeleteReportById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM reports WHERE id=?;")

	if err != nil {
		TheLogger().LogError("report", "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError("report", "report can't be deleted.", err)
	}

	return err
}
