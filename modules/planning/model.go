package planning

// swagger:model Planning
type Planning struct {
	// Id of the planning
	// in: int64
	// read only: true
	Id int `json:"id"`
	// Name of the planning
	// in: string
	// example: Permanence
	Name string `json:"name"`
	// Type of the planning
	// in: string
	// example: permanence
	Type string `json:"type"`
	// Location of the planning
	// in: string
	// example: Labo Prof Chen
	Location string `json:"location"`
	// Date_begins of the planning
	// in: string
	// example: 2023-05-23
	Date_begins string `json:"date_begins"`
	// Date_end of the planning
	// in: string
	// example: 2023-05-23
	Date_end string `json:"date_end"`
	// Hour_begins of the planning
	// in: string
	// example: 9:00:00
	Hour_begins string `json:"hour_begins"`
	// Hour_end of the planning
	// in: string
	// example: 18:00:00
	Hour_end string `json:"hour_end"`
}

type Plannings []Planning

// swagger:model Planning_Attendee
type Attendee struct {
	// Id of the attendee
	// in: int64
	// read only: true
	Id int `json:"id"`
	// Id of the planning
	// in: int
	// example: 2
	// required: true
	Id_planning int `json:"id_planning"`
	// Id of the volunteer
	// in: int
	// example: 1
	// required: true
	Id_volunteer int `json:"id_volunteer"`
	// Job of the shift
	// in: string
	// example: staff
	// required: true
	Job string `json:"job"`
	// Date of the shift
	// in: string
	// example: 2023-05-23
	// required: true
	Date string `json:"date"`
	// Hour_begins of the shift
	// in: string
	// example: 10:00:00
	// required: true
	Hour_begins string `json:"hour_begins"`
	// Hour_end of the shift
	// in: string
	// example: 12:00:00
	// required: true
	Hour_end string `json:"hour_end"`
}

type Attendees []Attendee

func NewPlanning(pla *Planning) {

	stmt, _ := TheDb().Prepare(
		`INSERT INTO plannings
			(name, type, location, date_begins, date_end, hour_begins, hour_end)
		VALUES (?,?,?,?,?,?,?);`,
	)
	_, err := stmt.Exec(
		pla.Name,
		pla.Type,
		pla.Location,
		pla.Date_begins,
		pla.Date_end,
		pla.Hour_begins,
		pla.Hour_end,
	)
	if err != nil {
		TheLogger().LogError("planning", "can't create new planning.", err)
	}
}

func FindPlanningById(id int) *Planning {
	var pla Planning

	row := TheDb().QueryRow("SELECT * FROM plannings WHERE id = ?;", id)
	err := row.Scan(
		&pla.Id,
		&pla.Name,
		&pla.Type,
		&pla.Location,
		&pla.Date_begins,
		&pla.Date_end,
		&pla.Hour_begins,
		&pla.Hour_end,
	)

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

		err := rows.Scan(
			&pla.Id,
			&pla.Name,
			&pla.Type,
			&pla.Location,
			&pla.Date_begins,
			&pla.Date_end,
			&pla.Hour_begins,
			&pla.Hour_end,
		)

		if err != nil {
			TheLogger().LogError("planning", "planning not found.", err)
		}

		plas = append(plas, pla)
	}

	return &plas
}

func UpdatePlanning(pla *Planning) {
	stmt, err := TheDb().Prepare(
		`UPDATE plannings SET
			name=?, type=?, location=?, date_begins=?, date_end=?, hour_begins=?, hour_end=?
		WHERE id=?;`,
	)

	if err != nil {
		TheLogger().LogError("planning", "problem with the db.", err)
	}

	_, err = stmt.Exec(
		pla.Name,
		pla.Type,
		pla.Location,
		pla.Date_begins,
		pla.Date_end,
		pla.Hour_begins,
		pla.Hour_end,
		pla.Id,
	)

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
		TheLogger().LogError("planning_attendee", "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var att Attendee

		err := rows.Scan(
			&att.Id,
			&att.Id_planning,
			&att.Id_volunteer,
			&att.Job,
			&att.Date,
			&att.Hour_begins,
			&att.Hour_end,
		)

		if err != nil {
			TheLogger().LogError("planning_attendee", "attendee not found.", err)
		}

		atts = append(atts, att)
	}

	return &atts
}

func NewAttendee(att *Attendee) {
	stmt, _ := TheDb().Prepare(
		`INSERT INTO planning_attendees
			(id_planning, id_volunteer, job, date, hour_begins, hour_end)
		VALUES (?,?,?,?,?,?);`,
	)
	_, err := stmt.Exec(
		att.Id_planning,
		att.Id_volunteer,
		att.Job,
		att.Date,
		att.Hour_begins,
		att.Hour_end,
	)
	if err != nil {
		TheLogger().LogError("planning_attendee", "can't create new attendee.", err)
	}
}

func FindAttendeeById(id int) *Attendee {
	var att Attendee

	row := TheDb().QueryRow("SELECT * FROM planning_attendees WHERE id = ?;", id)
	err := row.Scan(
		&att.Id,
		&att.Id_planning,
		&att.Id_volunteer,
		&att.Job,
		&att.Date,
		&att.Hour_begins,
		&att.Hour_end,
	)

	if err != nil {
		TheLogger().LogWarning("planning_attendee", "attendee not found.", err)
	}

	return &att
}

func FindAttendeeByPlanningId(id_planning int) *Attendees {
	var atts Attendees

	rows, err := TheDb().Query(
		"SELECT * FROM planning_attendees WHERE id_planning = ?;",
		id_planning,
	)

	if err != nil {
		TheLogger().LogWarning("planning_attendee", "attendees with id planning not found.", err)
	}

	for rows.Next() {
		var att Attendee

		err := rows.Scan(
			&att.Id,
			&att.Id_planning,
			&att.Id_volunteer,
			&att.Job,
			&att.Date,
			&att.Hour_begins,
			&att.Hour_end,
		)

		if err != nil {
			TheLogger().LogError("planning_attendee", "attendee not found.", err)
		}

		atts = append(atts, att)
	}

	return &atts
}

func FindAttendeeByVolunteerId(id_volunteer int) *Attendees {
	var atts Attendees

	rows, err := TheDb().Query(
		"SELECT * FROM planning_attendees WHERE id_volunteer = ?;",
		id_volunteer,
	)

	if err != nil {
		TheLogger().LogWarning("planning_attendee", "attendees with id volunteer not found.", err)
	}

	for rows.Next() {
		var att Attendee

		err := rows.Scan(
			&att.Id,
			&att.Id_planning,
			&att.Id_volunteer,
			&att.Job,
			&att.Date,
			&att.Hour_begins,
			&att.Hour_end,
		)

		if err != nil {
			TheLogger().LogError("planning_attendee", "attendee not found.", err)
		}

		atts = append(atts, att)
	}

	return &atts
}

func UpdateAttendee(att *Attendee) {
	stmt, err := TheDb().Prepare(
		`UPDATE planning_attendees SET
			id_planning=?, id_volunteer=?, job=?, date=?, hour_begins=?, hour_end=?
		WHERE id=?;`,
	)

	if err != nil {
		TheLogger().LogError("planning_attendee", "problem with the db.", err)
	}

	_, err = stmt.Exec(
		att.Id_planning,
		att.Id_volunteer,
		att.Job,
		att.Date,
		att.Hour_begins,
		att.Hour_end,
		att.Id,
	)

	if err != nil {
		TheLogger().LogError("planning_attendee", "attendee can't be updated.", err)
	}
}

func DeleteAttendeeById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM planning_attendees WHERE id=?;")

	if err != nil {
		TheLogger().LogError("planning_attendee", "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError("planning_attendee", "attendee can't be deleted.", err)
	}

	return err
}
