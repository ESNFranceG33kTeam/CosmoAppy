package models

import (
	"time"

	"github.com/ESNFranceG33kTeam/sAPI/config"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
)

type Adherent struct {
	Id          int       `json:"id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Email       string    `json:"email"`
	Dateofbirth string    `json:"dateofbirth"`
	ESNcard     string    `json:"esncard"`
	Student     bool      `json:"student"`
	University  string    `json:"university"`
	Homeland    string    `json:"homeland"`
	Speakabout  string    `json:"speakabout"`
	Newsletter  bool      `json:"newsletter"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Adherents []Adherent

func NewAdherent(adh *Adherent) {
	adh.CreatedAt = time.Now()
	adh.UpdatedAt = time.Now()

	stmt, _ := config.Db().Prepare("INSERT INTO adherents (firstname, lastname, email, dateofbirth, esncard, student, university, homeland, speakabout, newsletter, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?);")
	_, err := stmt.Exec(adh.Firstname, adh.Lastname, adh.Email, adh.Dateofbirth, adh.ESNcard, adh.Student, adh.University, adh.Homeland, adh.Speakabout, adh.Newsletter, adh.CreatedAt, adh.UpdatedAt)
	if err != nil {
		logger.LogCritical("adherent", "can't create new adherent.", err)
	}

	if err != nil {
		logger.LogCritical("adherent", "can't create new adherent.", err)
	}
}

func FindAdherentById(id int) *Adherent {
	var adh Adherent

	row := config.Db().QueryRow("SELECT * FROM adherents WHERE id = ?;", id)
	err := row.Scan(&adh.Id, &adh.Firstname, &adh.Lastname, &adh.Email, &adh.Dateofbirth, &adh.ESNcard, &adh.Student, &adh.University, &adh.Homeland, &adh.Speakabout, &adh.Newsletter, &adh.CreatedAt, &adh.UpdatedAt)

	if err != nil {
		logger.LogWarning("adherent", "adherent not found.", err)
	}

	return &adh
}

func FindAdherentByName(firstname string, lastname string) *Adherent {
	var adh Adherent

	row := config.Db().QueryRow("SELECT * FROM adherents WHERE firstname = ? AND lastname = ?;", firstname, lastname)
	err := row.Scan(&adh.Id, &adh.Firstname, &adh.Lastname, &adh.Email, &adh.Dateofbirth, &adh.ESNcard, &adh.Student, &adh.University, &adh.Homeland, &adh.Speakabout, &adh.Newsletter, &adh.CreatedAt, &adh.UpdatedAt)

	if err != nil {
		logger.LogWarning("adherent", "adherent not found.", err)
	}

	return &adh
}

func AllAdherents() *Adherents {
	var adhs Adherents

	rows, err := config.Db().Query("SELECT * FROM adherents")

	if err != nil {
		logger.LogError("adherent", "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var adh Adherent

		err := rows.Scan(&adh.Id, &adh.Firstname, &adh.Lastname, &adh.Email, &adh.Dateofbirth, &adh.ESNcard, &adh.Student, &adh.University, &adh.Homeland, &adh.Speakabout, &adh.Newsletter, &adh.CreatedAt, &adh.UpdatedAt)

		if err != nil {
			logger.LogError("adherent", "adherents not found.", err)
		}

		adhs = append(adhs, adh)
	}

	return &adhs
}

func UpdateAdherent(adh *Adherent) {
	adh.UpdatedAt = time.Now()

	stmt, err := config.Db().Prepare("UPDATE adherents SET firstname=?, lastname=?, email=?, dateofbirth=?, esncard=?, student=?, university=?, homeland=?, speakabout=?, newsletter=?, updated_at=? WHERE id=?;")

	if err != nil {
		logger.LogError("adherent", "problem with the db.", err)
	}

	_, err = stmt.Exec(adh.Firstname, adh.Lastname, adh.Email, adh.Dateofbirth, adh.ESNcard, adh.Student, adh.University, adh.Homeland, adh.Speakabout, adh.Newsletter, adh.UpdatedAt, adh.Id)

	if err != nil {
		logger.LogError("adherent", "adherent can't be updated.", err)
	}
}

func DeleteAdherentById(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM adherents WHERE id=?;")

	if err != nil {
		logger.LogError("adherent", "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		logger.LogError("adherent", "adherent can't be deleted.", err)
	}

	return err
}
