package planning

// swagger:model Planning
type Planning struct {
	// Id of the planning
	// in: int64
	Id int `json:"id"`
	// Label of the planning
	// in: string
	Name string `json:"name"`
	// Location of the planning
	// in: string
	Location string `json:"location"`
	// Date_begins of the planning
	// in: string
	Date_begins string `json:"date_begins"`
	// Date_end of the planning
	// in: string
	Date_end string `json:"date_end"`
	// Hour_begins of the planning
	// in: string
	Hour_begins string `json:"hour_begins"`
	// Hour_end of the planning
	// in: string
	Hour_end string `json:"hour_end"`
}

type Plannings []Planning

// swagger:model Planning_Attendee
type Attendee struct {
	// Id of the attendee
	// in: int64
	Id int `json:"id"`
	// Id of the planning
	// in: int64
	Id_planning int `json:"id_planning"`
	// Id of the adherent
	// in: int64
	Id_adherent int `json:"id_adherent"`
	// Date of the shift
	// in: string
	Date string `json:"date"`
	// Hour_begins of the shift
	// in: string
	Hour_begins string `json:"hour_begins"`
	// Hour_end of the shift
	// in: string
	Hour_end string `json:"hour_end"`
}

type Attendees []Attendee

func NewPlanning(pla *Planning) {

	stmt, _ := TheDb().Prepare("INSERT INTO plannings (name, location, date_begins, date_end, hour_begins, hour_end) VALUES (?,?,?,?,?,?);")
	_, err := stmt.Exec(pla.Name, pla.Location, pla.Date_begins, pla.Date_end, pla.Hour_begins, pla.Hour_end)
	if err != nil {
		TheLogger().LogError("planning", "can't create new planning.", err)
	}
}

func FindPlanningById(id int) *Planning {
	var pla Planning

	row := TheDb().QueryRow("SELECT * FROM plannings WHERE id = ?;", id)
	err := row.Scan(&pla.Id, &pla.Name, &pla.Location, &pla.Date_begins, &pla.Date_end, &pla.Hour_begins, &pla.Hour_end)

	if err != nil {
		TheLogger().LogWarning("planning", "event not found.", err)
	}

	return &pla
}

func AllPlannings() *Plannings {
	var plas Plannings

	rows, err := TheDb().Query("SELECT * FROM plannings")

	if err != nil {
		TheLogger().LogError("planning", "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var pla Planning

		err := rows.Scan(&pla.Id, &pla.Name, &pla.Location, &pla.Date_begins, &pla.Date_end, &pla.Hour_begins, &pla.Hour_end)

		if err != nil {
			TheLogger().LogError("planning", "planning not found.", err)
		}

		plas = append(plas, pla)
	}

	return &plas
}

func UpdatePlanning(pla *Planning) {
	stmt, err := TheDb().Prepare("UPDATE plannings SET name=?, location=?, date_begins=?, date_end=?, hour_begins=?, hour_end=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError("planning", "problem with the db.", err)
	}

	_, err = stmt.Exec(pla.Name, pla.Location, pla.Date_begins, pla.Date_end, pla.Hour_begins, pla.Hour_end, pla.Id)

	if err != nil {
		TheLogger().LogError("planning", "planning can't be updated.", err)
	}
}

func DeletePlanningById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM plannings WHERE id=?;")

	if err != nil {
		TheLogger().LogError("planning", "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError("planning", "planning can't be deleted.", err)
	}

	return err
}

func AllAttendees() *Attendees {
	var atts Attendees

	rows, err := TheDb().Query("SELECT * FROM planning_attendees;")

	if err != nil {
		TheLogger().LogError("planning", "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var att Attendee

		err := rows.Scan(&att.Id, &att.Id_planning, &att.Id_adherent, &att.Date, &att.Hour_begins, &att.Hour_end)

		if err != nil {
			TheLogger().LogError("planning", "attendee not found.", err)
		}

		atts = append(atts, att)
	}

	return &atts
}

func NewAttendee(att *Attendee) {
	stmt, _ := TheDb().Prepare("INSERT INTO planning_attendees (id_planning, id_adherent, date, hour_begins, hour_end) VALUES (?,?,?,?,?);")
	_, err := stmt.Exec(att.Id_planning, att.Id_adherent, att.Date, att.Hour_begins, att.Hour_end)
	if err != nil {
		TheLogger().LogError("planning", "can't create new attendee.", err)
	}
}

func FindAttendeeById(id int) *Attendee {
	var att Attendee

	row := TheDb().QueryRow("SELECT * FROM planning_attendees WHERE id = ?;", id)
	err := row.Scan(&att.Id, &att.Id_planning, &att.Id_adherent, &att.Date, &att.Hour_begins, &att.Hour_end)

	if err != nil {
		TheLogger().LogWarning("planning", "attendee not found.", err)
	}

	return &att
}

func FindAttendeeByPlanningId(id_planning int) *Attendees {
	var atts Attendees

	rows, err := TheDb().Query("SELECT * FROM planning_attendees WHERE id_planning = ?;", id_planning)

	if err != nil {
		TheLogger().LogWarning("planning", "attendees with id planning not found.", err)
	}

	for rows.Next() {
		var att Attendee

		err := rows.Scan(&att.Id, &att.Id_planning, &att.Id_adherent, &att.Date, &att.Hour_begins, &att.Hour_end)

		if err != nil {
			TheLogger().LogError("planning", "attendee not found.", err)
		}

		atts = append(atts, att)
	}

	return &atts
}

func FindAttendeeByAdherentId(id_adherent int) *Attendees {
	var atts Attendees

	rows, err := TheDb().Query("SELECT * FROM planning_attendees WHERE id_adherent = ?;", id_adherent)

	if err != nil {
		TheLogger().LogWarning("planning", "attendees with id adherent not found.", err)
	}

	for rows.Next() {
		var att Attendee

		err := rows.Scan(&att.Id, &att.Id_planning, &att.Id_adherent, &att.Date, &att.Hour_begins, &att.Hour_end)

		if err != nil {
			TheLogger().LogError("planning", "attendee not found.", err)
		}

		atts = append(atts, att)
	}

	return &atts
}

func UpdateAttendee(att *Attendee) {
	stmt, err := TheDb().Prepare("UPDATE planning_attendees SET id_planning=?, id_adherent=?, date=?, hour_begins=?, hour_end=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError("planning", "problem with the db.", err)
	}

	_, err = stmt.Exec(att.Id_planning, att.Id_adherent, att.Date, att.Hour_begins, att.Hour_end, att.Id)

	if err != nil {
		TheLogger().LogError("planning", "attendee can't be updated.", err)
	}
}

func DeleteAttendeeById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM planning_attendees WHERE id=?;")

	if err != nil {
		TheLogger().LogError("planning", "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError("planning", "attendee can't be deleted.", err)
	}

	return err
}
