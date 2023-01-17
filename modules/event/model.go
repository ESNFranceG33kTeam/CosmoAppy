package event

// swagger:model Event
type Event struct {
	// Id of the event
	// in: int64
	Id int `json:"id"`
	// Name of the event
	// in: string
	Name string `json:"name"`
	// Date of the event
	// in: string
	Date string `json:"date"`
	// Location of the event
	// in: string
	Location string `json:"location"`
	// Number Spots Max of the event
	// in: int
	NbSpotsMax int `json:"nb_spots_max"`
	// Number Spots Taken of the event
	// in: int
	NbSpotsTaken int `json:"nb_spots_taken"`
	// Type of the event
	// in: string
	Type string `json:"type"`
	// Price of the event
	// in: float
	Price float64 `json:"price"`
	// Url of the event
	// in: string
	Url string `json:"url_facebook"`
	// Status of the event
	// in: bool
	Actif bool `json:"actif"`
}

type Events []Event

// swagger:model Event_Attendee
type Attendee struct {
	// Id of the attendee
	// in: int64
	Id int `json:"id"`
	// Id of the event
	// in: int64
	Id_event int `json:"id_event"`
	// Id of the adherent
	// in: int64
	Id_adherent int `json:"id_adherent"`
}

type Attendees []Attendee

// swagger:model Event_Staff
type Staff struct {
	// Id of the staff
	// in: int64
	Id int `json:"id"`
	// Id of the event
	// in: int64
	Id_event int `json:"id_event"`
	// Id of the volunteer
	// in: int64
	Id_volunteer int `json:"id_volunteer"`
}

type Staffs []Staff

func NewEvent(eve *Event) {

	stmt, _ := TheDb().Prepare("INSERT INTO events (name, date, location, nb_spots_max, nb_spots_taken, type, price, url_facebook, actif) VALUES (?,?,?,?,?,?,?,?,?);")
	_, err := stmt.Exec(eve.Name, eve.Date, eve.Location, eve.NbSpotsMax, eve.NbSpotsTaken, eve.Type, eve.Price, eve.Url, eve.Actif)
	if err != nil {
		TheLogger().LogError("event", "can't create new event.", err)
	}
}

func FindEventById(id int) *Event {
	var eve Event

	row := TheDb().QueryRow("SELECT * FROM events WHERE id = ?;", id)
	err := row.Scan(&eve.Id, &eve.Name, &eve.Date, &eve.Location, &eve.NbSpotsMax, &eve.NbSpotsTaken, &eve.Type, &eve.Price, &eve.Url, &eve.Actif)

	if err != nil {
		TheLogger().LogWarning("event", "event not found.", err)
	}

	return &eve
}

func AllEvents() *Events {
	var eves Events

	rows, err := TheDb().Query("SELECT * FROM events")

	if err != nil {
		TheLogger().LogError("event", "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var eve Event

		err := rows.Scan(&eve.Id, &eve.Name, &eve.Date, &eve.Location, &eve.NbSpotsMax, &eve.NbSpotsTaken, &eve.Type, &eve.Price, &eve.Url, &eve.Actif)

		if err != nil {
			TheLogger().LogError("event", "event not found.", err)
		}

		eves = append(eves, eve)
	}

	return &eves
}

func UpdateSpotsTakenEvent(eve *Event) {
	stmt, err := TheDb().Prepare("UPDATE events SET nb_spots_taken=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError("event", "problem with the db.", err)
	}

	_, err = stmt.Exec(eve.NbSpotsTaken, eve.Id)

	if err != nil {
		TheLogger().LogError("event", "nb_spots_taken can't be updated.", err)
	}
}

func UpdateEvent(eve *Event) {
	stmt, err := TheDb().Prepare("UPDATE events SET name=?, date=?, location=?, nb_spots_max=?, nb_spots_taken=?, type=?, price=?, url_facebook=?, actif=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError("event", "problem with the db.", err)
	}

	_, err = stmt.Exec(eve.Name, eve.Date, eve.Location, eve.NbSpotsMax, eve.NbSpotsTaken, eve.Type, eve.Price, eve.Url, eve.Actif, eve.Id)

	if err != nil {
		TheLogger().LogError("event", "event can't be updated.", err)
	}
}

func DeleteEventById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM events WHERE id=?;")

	if err != nil {
		TheLogger().LogError("event", "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError("event", "event can't be deleted.", err)
	}

	return err
}

func AllAttendees() *Attendees {
	var atts Attendees

	rows, err := TheDb().Query("SELECT * FROM event_attendees;")

	if err != nil {
		TheLogger().LogError("attendee", "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var att Attendee

		err := rows.Scan(&att.Id, &att.Id_event, &att.Id_adherent)

		if err != nil {
			TheLogger().LogError("attendee", "attendee not found.", err)
		}

		atts = append(atts, att)
	}

	return &atts
}

func NewAttendee(att *Attendee) {
	stmt, _ := TheDb().Prepare("INSERT INTO event_attendees (id_event, id_adherent) VALUES (?,?);")
	_, err := stmt.Exec(att.Id_event, att.Id_adherent)
	if err != nil {
		TheLogger().LogError("attendee", "can't create new attendee.", err)
	}
}

func FindAttendeeById(id int) *Attendee {
	var att Attendee

	row := TheDb().QueryRow("SELECT * FROM event_attendees WHERE id = ?;", id)
	err := row.Scan(&att.Id, &att.Id_event, &att.Id_adherent)

	if err != nil {
		TheLogger().LogWarning("attendee", "attendee not found.", err)
	}

	return &att
}

func FindAttendeeByEventId(id_event int) *Attendees {
	var atts Attendees

	rows, err := TheDb().Query("SELECT * FROM event_attendees WHERE id_event = ?;", id_event)

	if err != nil {
		TheLogger().LogWarning("attendee", "attendees with id event not found.", err)
	}

	for rows.Next() {
		var att Attendee

		err := rows.Scan(&att.Id, &att.Id_event, &att.Id_adherent)

		if err != nil {
			TheLogger().LogError("attendee", "attendee not found.", err)
		}

		atts = append(atts, att)
	}

	return &atts
}

func FindAttendeeByAdherentId(id_adherent int) *Attendees {
	var atts Attendees

	rows, err := TheDb().Query("SELECT * FROM event_attendees WHERE id_adherent = ?;", id_adherent)

	if err != nil {
		TheLogger().LogWarning("attendee", "attendees with id adherent not found.", err)
	}

	for rows.Next() {
		var att Attendee

		err := rows.Scan(&att.Id, &att.Id_event, &att.Id_adherent)

		if err != nil {
			TheLogger().LogError("attendee", "attendee not found.", err)
		}

		atts = append(atts, att)
	}

	return &atts
}

func UpdateAttendee(att *Attendee) {
	stmt, err := TheDb().Prepare("UPDATE event_attendees SET id_event=?, id_adherent=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError("attendee", "problem with the db.", err)
	}

	_, err = stmt.Exec(att.Id_event, att.Id_adherent, att.Id)

	if err != nil {
		TheLogger().LogError("attendee", "attendee can't be updated.", err)
	}
}

func DeleteAttendeeById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM event_attendees WHERE id=?;")

	if err != nil {
		TheLogger().LogError("attendee", "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError("attendee", "attendee can't be deleted.", err)
	}

	return err
}

func AllStaffs() *Staffs {
	var stas Staffs

	rows, err := TheDb().Query("SELECT * FROM event_staffs;")

	if err != nil {
		TheLogger().LogError("staff", "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var sta Staff

		err := rows.Scan(&sta.Id, &sta.Id_event, &sta.Id_volunteer)

		if err != nil {
			TheLogger().LogError("staff", "staff not found.", err)
		}

		stas = append(stas, sta)
	}

	return &stas
}

func NewStaff(sta *Staff) {
	stmt, _ := TheDb().Prepare("INSERT INTO event_staffs (id_event, id_volunteer) VALUES (?,?);")
	_, err := stmt.Exec(sta.Id_event, sta.Id_volunteer)
	if err != nil {
		TheLogger().LogError("staff", "can't create new attendee.", err)
	}
}

func FindStaffById(id int) *Staff {
	var sta Staff

	row := TheDb().QueryRow("SELECT * FROM event_staffs WHERE id = ?;", id)
	err := row.Scan(&sta.Id, &sta.Id_event, &sta.Id_volunteer)

	if err != nil {
		TheLogger().LogWarning("attendee", "attendee not found.", err)
	}

	return &sta
}

func FindStaffByEventId(id_event int) *Staffs {
	var stas Staffs

	rows, err := TheDb().Query("SELECT * FROM event_staffs WHERE id_event = ?;", id_event)

	if err != nil {
		TheLogger().LogWarning("staff", "staffs with id event not found.", err)
	}

	for rows.Next() {
		var sta Staff

		err := rows.Scan(&sta.Id, &sta.Id_event, &sta.Id_volunteer)

		if err != nil {
			TheLogger().LogError("attendee", "attendee not found.", err)
		}

		stas = append(stas, sta)
	}

	return &stas
}

func FindStaffByVolunteerId(id_volunteer int) *Staffs {
	var stas Staffs

	rows, err := TheDb().Query("SELECT * FROM event_staffs WHERE id_volunteer = ?;", id_volunteer)

	if err != nil {
		TheLogger().LogWarning("staff", "staffs with id adherent not found.", err)
	}

	for rows.Next() {
		var sta Staff

		err := rows.Scan(&sta.Id, &sta.Id_event, &sta.Id_volunteer)

		if err != nil {
			TheLogger().LogError("staff", "staff not found.", err)
		}

		stas = append(stas, sta)
	}

	return &stas
}

func UpdateStaff(sta *Staff) {
	stmt, err := TheDb().Prepare("UPDATE event_staffs SET id_event=?, id_volunteer=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError("staff", "problem with the db.", err)
	}

	_, err = stmt.Exec(sta.Id_event, sta.Id_volunteer, sta.Id)

	if err != nil {
		TheLogger().LogError("staff", "staff can't be updated.", err)
	}
}

func DeleteStaffById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM event_staffs WHERE id=?;")

	if err != nil {
		TheLogger().LogError("staff", "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError("staff", "staff can't be deleted.", err)
	}

	return err
}
