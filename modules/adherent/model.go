package adherent

import (
	"time"
)

// swagger:model Adherent
type Adherent struct {
	// Id of the adherent
	// in: int
	// read only: true
	Id int `json:"id"`
	// Firstname of the adherent
	// in: string
	// required: true
	// example: Pichu
	Firstname string `json:"firstname"`
	// Lastname of the adherent
	// in: string
	// required: true
	// example: Ketchum
	Lastname string `json:"lastname"`
	// Email of the adherent
	// in: string
	// required: true
	// example: pichu@kanto.kt
	Email string `json:"email"`
	// Dateofbirth of the adherent
	// in: string
	// required: true
	// example: 2000-05-22
	Dateofbirth string `json:"dateofbirth"`
	// Situation status of the adherent
	// in: string
	// required: true
	// example: student
	Situation string `json:"situation"`
	// University of the adherent
	// in: string
	// required: true
	// example: Indigo
	University string `json:"university"`
	// Homeland of the adherent
	// in: string
	// required: true
	// example: Kanto
	Homeland string `json:"homeland"`
	// Speakabout of the adherent
	// in: string
	// required: true
	// example: Twitter
	Speakabout string `json:"speakabout"`
	// Newsletter status of the adherent
	// in: bool
	// required: true
	// example: true
	Newsletter bool `json:"newsletter"`
	// AdhesionDate date of the adhesion
	// in: string
	// required: true
	// example: 2022-05-22
	AdhesionDate string `json:"adhesion_date"`
	// CreatedAt date of the adherent
	// in: time.Time
	// read only: true
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt date of the adherent
	// in: time.Time
	// read only: true
	UpdatedAt time.Time `json:"updated_at"`
}

type Adherents []Adherent

func NewAdherent(adh *Adherent) {
	adh.CreatedAt = time.Now()
	adh.UpdatedAt = time.Now()

	stmt, _ := TheDb().Prepare(
		`INSERT INTO adherents
			(firstname, lastname, email, dateofbirth, situation, university, homeland, speakabout,
			newsletter, adhesion_date, created_at, updated_at)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?);`,
	)
	_, err := stmt.Exec(
		adh.Firstname,
		adh.Lastname,
		adh.Email,
		adh.Dateofbirth,
		adh.Situation,
		adh.University,
		adh.Homeland,
		adh.Speakabout,
		adh.Newsletter,
		adh.AdhesionDate,
		adh.CreatedAt,
		adh.UpdatedAt,
	)
	if err != nil {
		TheLogger().LogError("adherent", "can't create new adherent.", err)
	}
}

func FindAdherentById(id int) *Adherent {
	var adh Adherent

	row := TheDb().QueryRow("SELECT * FROM adherents WHERE id = ?;", id)
	err := row.Scan(
		&adh.Id,
		&adh.Firstname,
		&adh.Lastname,
		&adh.Email,
		&adh.Dateofbirth,
		&adh.Situation,
		&adh.University,
		&adh.Homeland,
		&adh.Speakabout,
		&adh.Newsletter,
		&adh.AdhesionDate,
		&adh.CreatedAt,
		&adh.UpdatedAt,
	)

	if err != nil {
		TheLogger().LogWarning("adherent", "adherent not found.", err)
	}

	return &adh
}

func FindAdherentByName(firstname string, lastname string) *Adherent {
	var adh Adherent

	row := TheDb().QueryRow("SELECT * FROM adherents WHERE firstname = ? AND lastname = ?;",
		firstname, lastname)
	err := row.Scan(
		&adh.Id,
		&adh.Firstname,
		&adh.Lastname,
		&adh.Email,
		&adh.Dateofbirth,
		&adh.Situation,
		&adh.University,
		&adh.Homeland,
		&adh.Speakabout,
		&adh.Newsletter,
		&adh.AdhesionDate,
		&adh.CreatedAt,
		&adh.UpdatedAt,
	)

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

		err := rows.Scan(
			&adh.Id,
			&adh.Firstname,
			&adh.Lastname,
			&adh.Email,
			&adh.Dateofbirth,
			&adh.Situation,
			&adh.University,
			&adh.Homeland,
			&adh.Speakabout,
			&adh.Newsletter,
			&adh.AdhesionDate,
			&adh.CreatedAt,
			&adh.UpdatedAt,
		)

		if err != nil {
			TheLogger().LogError("adherent", "adherents not found.", err)
		}

		adhs = append(adhs, adh)
	}

	return &adhs
}

func UpdateAdherent(adh *Adherent) {
	adh.UpdatedAt = time.Now()

	stmt, err := TheDb().Prepare(
		`UPDATE adherents SET
			firstname=?, lastname=?, email=?, dateofbirth=?, situation=?, university=?,
			homeland=?, speakabout=?, newsletter=?, adhesion_date=?, updated_at=?
		WHERE id=?;`,
	)

	if err != nil {
		TheLogger().LogError("adherent", "problem with the db.", err)
	}

	_, err = stmt.Exec(
		adh.Firstname,
		adh.Lastname,
		adh.Email,
		adh.Dateofbirth,
		adh.Situation,
		adh.University,
		adh.Homeland,
		adh.Speakabout,
		adh.Newsletter,
		adh.AdhesionDate,
		adh.UpdatedAt,
		adh.Id,
	)

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
