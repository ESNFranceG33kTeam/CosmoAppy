package event

// swagger:model Event
type Event struct {
	// Id of the event
	// in: int64
	// read only: true
	Id int `json:"id"`
	// Name of the event
	// in: string
	// required: true
	// example: Conseil des 4
	Name string `json:"name"`
	// Date of the event
	// in: string
	// required: true
	// example: 2023-05-23
	Date string `json:"date"`
	// Location of the event
	// in: string
	// required: true
	// example: Plateau indigo, Kanto
	Location string `json:"location"`
	// Number Spots Max of the event
	// in: int
	// required: true
	// example: 30
	NbSpotsMax int `json:"nb_spots_max"`
	// Number Spots Taken of the event
	// in: int
	// required: true
	// example: 30
	NbSpotsTaken int `json:"nb_spots_taken"`
	// Type of the event
	// in: string
	// required: true
	// example: sport
	Type string `json:"type"`
	// Price of the event
	// in: float
	// required: true
	// example: 0
	Price float64 `json:"price"`
	// Url of the event
	// in: string
	// required: true
	// example: https://facebook.com/ligue
	Url string `json:"url_facebook"`
	// Status of the event
	// in: bool
	// required: true
	// example: true
	Actif bool `json:"actif"`
}

type Events []Event

// swagger:model Event_Attendee
type Attendee struct {
	// Id of the attendee
	// in: int64
	// read only: true
	Id int `json:"id"`
	// Id of the event
	// in: int64
	// required: true
	// example: 2
	Id_event int `json:"id_event"`
	// Id of the adherent
	// in: int64
	// required: true
	// example: 1
	Id_adherent int `json:"id_adherent"`
}

type Attendees []Attendee

// swagger:model Event_Staff
type Staff struct {
	// Id of the staff
	// in: int64
	// read only: true
	Id int `json:"id"`
	// Id of the event
	// in: int
	// required: true
	// example: 2
	Id_event int `json:"id_event"`
	// Id of the volunteer
	// in: int
	// required: true
	// example: 1
	Id_volunteer int `json:"id_volunteer"`
}

type Staffs []Staff

func NewEvent(eve *Event) {

	stmt, _ := TheDb().Prepare(
		`INSERT INTO ` + db_name + `
			(name, date, location, nb_spots_max, nb_spots_taken, type, price, url_facebook, actif)
		VALUES (?,?,?,?,?,?,?,?,?);`,
	)
	_, err := stmt.Exec(
		eve.Name,
		eve.Date,
		eve.Location,
		eve.NbSpotsMax,
		eve.NbSpotsTaken,
		eve.Type,
		eve.Price,
		eve.Url,
		eve.Actif,
	)
	if err != nil {
		TheLogger().LogError(log_name, "can't create new event.", err)
	}
}

func FindEventById(id int) *Event {
	var eve Event

	row := TheDb().QueryRow("SELECT * FROM "+db_name+" WHERE id = ?;", id)
	err := row.Scan(
		&eve.Id,
		&eve.Name,
		&eve.Date,
		&eve.Location,
		&eve.NbSpotsMax,
		&eve.NbSpotsTaken,
		&eve.Type,
		&eve.Price,
		&eve.Url,
		&eve.Actif,
	)

	if err != nil {
		TheLogger().LogWarning(log_name, "event not found.", err)
	}

	return &eve
}

func FindEventsByDate(date string) *Events {
	var eves Events

	rows, err := TheDb().Query("SELECT * FROM "+db_name+" WHERE date like ?;", date+"%")

	if err != nil {
		TheLogger().LogWarning(log_name, "events with date_begins not found.", err)
	}

	for rows.Next() {
		var eve Event

		err := rows.Scan(
			&eve.Id,
			&eve.Name,
			&eve.Date,
			&eve.Location,
			&eve.NbSpotsMax,
			&eve.NbSpotsTaken,
			&eve.Type,
			&eve.Price,
			&eve.Url,
			&eve.Actif,
		)

		if err != nil {
			TheLogger().LogInfo(log_name, "eves not found.")
		}

		eves = append(eves, eve)
	}

	return &eves
}

func AllEvents() *Events {
	var eves Events

	rows, err := TheDb().Query("SELECT * FROM " + db_name)

	if err != nil {
		TheLogger().LogError(log_name, "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var eve Event

		err := rows.Scan(
			&eve.Id,
			&eve.Name,
			&eve.Date,
			&eve.Location,
			&eve.NbSpotsMax,
			&eve.NbSpotsTaken,
			&eve.Type,
			&eve.Price,
			&eve.Url,
			&eve.Actif,
		)

		if err != nil {
			TheLogger().LogError(log_name, "event not found.", err)
		}

		eves = append(eves, eve)
	}

	return &eves
}

func UpdateSpotsTakenEvent(eve *Event) {
	stmt, err := TheDb().Prepare("UPDATE " + db_name + " SET nb_spots_taken=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError(log_name, "problem with the db.", err)
	}

	_, err = stmt.Exec(eve.NbSpotsTaken, eve.Id)

	if err != nil {
		TheLogger().LogError(log_name, "nb_spots_taken can't be updated.", err)
	}
}

func UpdateEvent(eve *Event) {
	stmt, err := TheDb().Prepare(
		`UPDATE ` + db_name + ` SET
			name=?, date=?, location=?, nb_spots_max=?, nb_spots_taken=?, type=?, price=?,
			url_facebook=?, actif=?
		WHERE id=?;`,
	)

	if err != nil {
		TheLogger().LogError(log_name, "problem with the db.", err)
	}

	_, err = stmt.Exec(
		eve.Name,
		eve.Date,
		eve.Location,
		eve.NbSpotsMax,
		eve.NbSpotsTaken,
		eve.Type,
		eve.Price,
		eve.Url,
		eve.Actif,
		eve.Id,
	)

	if err != nil {
		TheLogger().LogError(log_name, "event can't be updated.", err)
	}
}

func DeleteEventById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM " + db_name + " WHERE id=?;")

	if err != nil {
		TheLogger().LogError(log_name, "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError(log_name, "event can't be deleted.", err)
	}

	return err
}

func AllAttendees() *Attendees {
	var atts Attendees

	rows, err := TheDb().Query("SELECT * FROM " + db_name_att + ";")

	if err != nil {
		TheLogger().LogError(log_name_att, "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var att Attendee

		err := rows.Scan(&att.Id, &att.Id_event, &att.Id_adherent)

		if err != nil {
			TheLogger().LogError(log_name_att, "attendee not found.", err)
		}

		atts = append(atts, att)
	}

	return &atts
}

func NewAttendee(att *Attendee) {
	stmt, _ := TheDb().Prepare("INSERT INTO " + db_name_att + " (id_event, id_adherent) VALUES (?,?);")
	_, err := stmt.Exec(att.Id_event, att.Id_adherent)
	if err != nil {
		TheLogger().LogError(log_name_att, "can't create new attendee.", err)
	}
}

func FindAttendeeById(id int) *Attendee {
	var att Attendee

	row := TheDb().QueryRow("SELECT * FROM "+db_name_att+" WHERE id = ?;", id)
	err := row.Scan(&att.Id, &att.Id_event, &att.Id_adherent)

	if err != nil {
		TheLogger().LogWarning(log_name_att, "attendee not found.", err)
	}

	return &att
}

func FindAttendeeByEventId(id_event int) *Attendees {
	var atts Attendees

	rows, err := TheDb().Query("SELECT * FROM "+db_name_att+" WHERE id_event = ?;", id_event)

	if err != nil {
		TheLogger().LogWarning(log_name_att, "attendees with id event not found.", err)
	}

	for rows.Next() {
		var att Attendee

		err := rows.Scan(&att.Id, &att.Id_event, &att.Id_adherent)

		if err != nil {
			TheLogger().LogError(log_name_att, "attendee not found.", err)
		}

		atts = append(atts, att)
	}

	return &atts
}

func FindAttendeeByAdherentId(id_adherent int) *Attendees {
	var atts Attendees

	rows, err := TheDb().Query("SELECT * FROM "+db_name_att+" WHERE id_adherent = ?;", id_adherent)

	if err != nil {
		TheLogger().LogWarning(log_name_att, "attendees with id adherent not found.", err)
	}

	for rows.Next() {
		var att Attendee

		err := rows.Scan(&att.Id, &att.Id_event, &att.Id_adherent)

		if err != nil {
			TheLogger().LogError(log_name_att, "attendee not found.", err)
		}

		atts = append(atts, att)
	}

	return &atts
}

func UpdateAttendee(att *Attendee) {
	stmt, err := TheDb().Prepare(
		`UPDATE ` + db_name_att + ` SET
			id_event=?, id_adherent=?
		WHERE id=?;`,
	)

	if err != nil {
		TheLogger().LogError(log_name_att, "problem with the db.", err)
	}

	_, err = stmt.Exec(att.Id_event, att.Id_adherent, att.Id)

	if err != nil {
		TheLogger().LogError(log_name_att, "attendee can't be updated.", err)
	}
}

func DeleteAttendeeById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM " + db_name_att + " WHERE id=?;")

	if err != nil {
		TheLogger().LogError(log_name_att, "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError(log_name_att, "attendee can't be deleted.", err)
	}

	return err
}

func AllStaffs() *Staffs {
	var stas Staffs

	rows, err := TheDb().Query("SELECT * FROM " + db_name_sta + ";")

	if err != nil {
		TheLogger().LogError(log_name_sta, "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var sta Staff

		err := rows.Scan(&sta.Id, &sta.Id_event, &sta.Id_volunteer)

		if err != nil {
			TheLogger().LogError(log_name_sta, "staff not found.", err)
		}

		stas = append(stas, sta)
	}

	return &stas
}

func NewStaff(sta *Staff) {
	stmt, _ := TheDb().Prepare("INSERT INTO " + db_name_sta + " (id_event, id_volunteer) VALUES (?,?);")
	_, err := stmt.Exec(sta.Id_event, sta.Id_volunteer)
	if err != nil {
		TheLogger().LogError(log_name_sta, "can't create new attendee.", err)
	}
}

func FindStaffById(id int) *Staff {
	var sta Staff

	row := TheDb().QueryRow("SELECT * FROM "+db_name_sta+" WHERE id = ?;", id)
	err := row.Scan(&sta.Id, &sta.Id_event, &sta.Id_volunteer)

	if err != nil {
		TheLogger().LogWarning(log_name_sta, "attendee not found.", err)
	}

	return &sta
}

func FindStaffByEventId(id_event int) *Staffs {
	var stas Staffs

	rows, err := TheDb().Query("SELECT * FROM "+db_name_sta+" WHERE id_event = ?;", id_event)

	if err != nil {
		TheLogger().LogWarning(log_name_sta, "staffs with id event not found.", err)
	}

	for rows.Next() {
		var sta Staff

		err := rows.Scan(&sta.Id, &sta.Id_event, &sta.Id_volunteer)

		if err != nil {
			TheLogger().LogError(log_name_sta, "attendee not found.", err)
		}

		stas = append(stas, sta)
	}

	return &stas
}

func FindStaffByVolunteerId(id_volunteer int) *Staffs {
	var stas Staffs

	rows, err := TheDb().Query("SELECT * FROM "+db_name_sta+" WHERE id_volunteer = ?;", id_volunteer)

	if err != nil {
		TheLogger().LogWarning(log_name_sta, "staffs with id adherent not found.", err)
	}

	for rows.Next() {
		var sta Staff

		err := rows.Scan(&sta.Id, &sta.Id_event, &sta.Id_volunteer)

		if err != nil {
			TheLogger().LogError(log_name_sta, "staff not found.", err)
		}

		stas = append(stas, sta)
	}

	return &stas
}

func UpdateStaff(sta *Staff) {
	stmt, err := TheDb().Prepare("UPDATE " + db_name_sta + " SET id_event=?, id_volunteer=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError(log_name_sta, "problem with the db.", err)
	}

	_, err = stmt.Exec(sta.Id_event, sta.Id_volunteer, sta.Id)

	if err != nil {
		TheLogger().LogError(log_name_sta, "staff can't be updated.", err)
	}
}

func DeleteStaffById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM " + db_name_sta + " WHERE id=?;")

	if err != nil {
		TheLogger().LogError(log_name_sta, "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError(log_name_sta, "staff can't be deleted.", err)
	}

	return err
}
