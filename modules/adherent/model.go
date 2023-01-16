package adherent

import (
	"time"
)

// swagger:model Adherent
type Adherent struct {
	// Id of the adherent
	// in: int64
	Id int `json:"id"`
	// Firstname of the adherent
	// in: string
	Firstname string `json:"firstname"`
	// Lastname of the adherent
	// in: string
	Lastname string `json:"lastname"`
	// Email of the adherent
	// in: string
	Email string `json:"email"`
	// Dateofbirth of the adherent
	// in: string
	Dateofbirth string `json:"dateofbirth"`
	// Situation status of the adherent
	// in: string
	Situation string `json:"situation"`
	// University of the adherent
	// in: string
	University string `json:"university"`
	// Homeland of the adherent
	// in: string
	Homeland string `json:"homeland"`
	// Speakabout of the adherent
	// in: string
	Speakabout string `json:"speakabout"`
	// Newsletter status of the adherent
	// in: bool
	Newsletter bool `json:"newsletter"`
	// AdhesionDate date of the adhesion
	// in: string
	AdhesionDate string `json:"adhesion_date"`
	// CreatedAt date of the adherent
	// in: time.Time
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt date of the adherent
	// in: time.Time
	UpdatedAt time.Time `json:"updated_at"`
}

type Adherents []Adherent

func NewAdherent(adh *Adherent) {
	adh.CreatedAt = time.Now()
	adh.UpdatedAt = time.Now()

	stmt, _ := TheDb().Prepare("INSERT INTO adherents (firstname, lastname, email, dateofbirth, situation, university, homeland, speakabout, newsletter, adhesion_date, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?);")
	_, err := stmt.Exec(adh.Firstname, adh.Lastname, adh.Email, adh.Dateofbirth, adh.Situation, adh.University, adh.Homeland, adh.Speakabout, adh.Newsletter, adh.AdhesionDate, adh.CreatedAt, adh.UpdatedAt)
	if err != nil {
		TheLogger().LogError("adherent", "can't create new adherent.", err)
	}
}

func FindAdherentById(id int) *Adherent {
	var adh Adherent

	row := TheDb().QueryRow("SELECT * FROM adherents WHERE id = ?;", id)
	err := row.Scan(&adh.Id, &adh.Firstname, &adh.Lastname, &adh.Email, &adh.Dateofbirth, &adh.Situation, &adh.University, &adh.Homeland, &adh.Speakabout, &adh.Newsletter, &adh.AdhesionDate, &adh.CreatedAt, &adh.UpdatedAt)

	if err != nil {
		TheLogger().LogWarning("adherent", "adherent not found.", err)
	}

	return &adh
}

func FindAdherentByName(firstname string, lastname string) *Adherent {
	var adh Adherent

	row := TheDb().QueryRow("SELECT * FROM adherents WHERE firstname = ? AND lastname = ?;", firstname, lastname)
	err := row.Scan(&adh.Id, &adh.Firstname, &adh.Lastname, &adh.Email, &adh.Dateofbirth, &adh.Situation, &adh.University, &adh.Homeland, &adh.Speakabout, &adh.Newsletter, &adh.AdhesionDate, &adh.CreatedAt, &adh.UpdatedAt)

	if err != nil {
		TheLogger().LogWarning("adherent", "adherent not found.", err)
	}

	return &adh
}

func AllAdherents() *Adherents {
	var adhs Adherents

	rows, err := TheDb().Query("SELECT * FROM adherents")

	if err != nil {
		TheLogger().LogError("adherent", "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var adh Adherent

		err := rows.Scan(&adh.Id, &adh.Firstname, &adh.Lastname, &adh.Email, &adh.Dateofbirth, &adh.Situation, &adh.University, &adh.Homeland, &adh.Speakabout, &adh.Newsletter, &adh.AdhesionDate, &adh.CreatedAt, &adh.UpdatedAt)

		if err != nil {
			TheLogger().LogError("adherent", "adherents not found.", err)
		}

		adhs = append(adhs, adh)
	}

	return &adhs
}

func UpdateAdherent(adh *Adherent) {
	adh.UpdatedAt = time.Now()

	stmt, err := TheDb().Prepare("UPDATE adherents SET firstname=?, lastname=?, email=?, dateofbirth=?, situation=?, university=?, homeland=?, speakabout=?, newsletter=?, adhesion_date=?, updated_at=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError("adherent", "problem with the db.", err)
	}

	_, err = stmt.Exec(adh.Firstname, adh.Lastname, adh.Email, adh.Dateofbirth, adh.Situation, adh.University, adh.Homeland, adh.Speakabout, adh.Newsletter, adh.AdhesionDate, adh.UpdatedAt, adh.Id)

	if err != nil {
		TheLogger().LogError("adherent", "adherent can't be updated.", err)
	}
}

func DeleteAdherentById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM adherents WHERE id=?;")

	if err != nil {
		TheLogger().LogError("adherent", "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError("adherent", "adherent can't be deleted.", err)
	}

	return err
}
