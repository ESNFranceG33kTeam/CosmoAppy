package event

// swagger:model Event
type Event struct {
	// Id of the event
	// in: int64
	Id int `json:"id"`
	// Label of the event
	// in: string
	Name string `json:"name"`
	// Date of the event
	// in: string
	Date string `json:"date"`
	// Location of the event
	// in: string
	Location string `json:"location"`
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
type Attendees struct {
	// Id of the attendee
	// in: int64
	Id int `json:"id"`
	// Id of the event
	// in: int64
	IdEvent int `json:"id_event"`
	// Id of the adherent
	// in: int64
	IdAdherent int `json:"id_adherent"`
	// Status of the event
	// in: bool
	Staff bool `json:"staff"`
}

type Attendee []Attendees

func NewEvent(eve *Event) {

	stmt, _ := TheDb().Prepare("INSERT INTO events (name, date, location, type, price, url_facebook, actif) VALUES (?,?,?,?,?,?,?);")
	_, err := stmt.Exec(eve.Name, eve.Date, eve.Location, eve.Type, eve.Price, eve.Url, eve.Actif)
	if err != nil {
		TheLogger().LogError("event", "can't create new event.", err)
	}
}

func FindEventById(id int) *Event {
	var eve Event

	row := TheDb().QueryRow("SELECT * FROM events WHERE id = ?;", id)
	err := row.Scan(&eve.Id, &eve.Name, &eve.Date, &eve.Location, &eve.Type, &eve.Price, &eve.Url, &eve.Actif)

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

		err := rows.Scan(&eve.Id, &eve.Name, &eve.Date, &eve.Location, &eve.Type, &eve.Price, &eve.Url, &eve.Actif)

		if err != nil {
			TheLogger().LogError("event", "event not found.", err)
		}

		eves = append(eves, eve)
	}

	return &eves
}

func UpdateEvent(eve *Event) {
	stmt, err := TheDb().Prepare("UPDATE events SET name=?, date=?, location=?, type=?, price=?, url_facebook=?, actif=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError("event", "problem with the db.", err)
	}

	_, err = stmt.Exec(eve.Name, eve.Date, eve.Location, eve.Type, eve.Price, eve.Url, eve.Actif, eve.Id)

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
