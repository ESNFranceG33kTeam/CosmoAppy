package report

// swagger:model Report
type Report struct {
	// Id of the report
	// in: int64
	Id int `json:"id"`
	// Name of the report
	// in: string
	Name string `json:"name"`
	// Date of the event/planning/custom
	// in: string
	Date string `json:"date"`
	// Comment of the report
	// in: string
	Comment string `json:"comment"`
	// Number of hours of the event/planning/custom
	// in: string
	NbHours float64 `json:"nb_hours"`
	// Number of staffs of the event/planning/custom
	// in: string
	NbStaffs float64 `json:"nb_staffs"`
	// Taux Valorisation of the event/planning/custom
	// in: string
	TauxValorisation float64 `json:"taux_valorisation"`
}

type Reports []Report

func NewReport(rpt *Report) {
	stmt, _ := TheDb().Prepare("INSERT INTO reports (name, date, comment, nb_hours, nb_staffs, taux_valorisation) VALUES (?,?,?,?,?,?);")
	_, err := stmt.Exec(rpt.Name, rpt.Date, rpt.Comment, rpt.NbHours, rpt.NbStaffs, rpt.TauxValorisation)
	if err != nil {
		TheLogger().LogError("report", "can't create new report.", err)
	}
}

func FindReportById(id int) *Report {
	var rpt Report

	row := TheDb().QueryRow("SELECT * FROM reports WHERE id = ?;", id)
	err := row.Scan(&rpt.Id, &rpt.Name, &rpt.Date, &rpt.Comment, &rpt.NbHours, &rpt.NbStaffs, &rpt.TauxValorisation)

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

		err := rows.Scan(&rpt.Id, &rpt.Name, &rpt.Date, &rpt.Comment, &rpt.NbHours, &rpt.NbStaffs, &rpt.TauxValorisation)

		if err != nil {
			TheLogger().LogError("report", "reports not found.", err)
		}

		rpts = append(rpts, rpt)
	}

	return &rpts
}

func UpdateReport(rpt *Report) {
	stmt, err := TheDb().Prepare("UPDATE reports SET name=?, date=?, comment=?, nb_hours=?, nb_staffs=?, taux_valorisation=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError("report", "problem with the db.", err)
	}

	_, err = stmt.Exec(rpt.Name, rpt.Date, rpt.Comment, rpt.NbHours, rpt.NbStaffs, rpt.TauxValorisation, rpt.Id)

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
