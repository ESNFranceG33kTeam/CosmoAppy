package models

import (
	"log"
	"time"

	"github.com/ESNFranceG33kTeam/sAPI/config"
)

type Adherent struct {
	Id          int       `json:"id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Email       string    `json:"email"`
	Dateofbirth time.Time `json:"dateofbirth"`
	ESNcard     string    `json:"esncard"`
	Student     bool      `json:"student"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Adherents []Adherent

func NewAdherent(adh *Adherent) {
	if adh == nil {
		log.Fatal(adh)
	}
	adh.CreatedAt = time.Now()
	adh.UpdatedAt = time.Now()

	stmt, _ := config.Db().Prepare("INSERT INTO adherents (firstname, lastname, email, dateofbirth, esncard, student, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?);")
	res, err := stmt.Exec(adh.Firstname, adh.Lastname, adh.Email, adh.Dateofbirth, adh.ESNcard, adh.Student, adh.CreatedAt, adh.UpdatedAt)
	if err != nil {
		log.Fatalf("Fatal error : %s", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	//glog.Infoln(lastId)
	log.Println(lastId)
}

func FindAdherentById(id int) *Adherent {
	var adh Adherent

	row := config.Db().QueryRow("SELECT * FROM adherents WHERE id = ?;", id)
	err := row.Scan(&adh.Id, &adh.Firstname, &adh.Lastname, &adh.Email, &adh.Dateofbirth, adh.ESNcard, adh.Student, &adh.CreatedAt, &adh.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &adh
}

func FindAdherentByName(firstname string, lastname string) *Adherent {
	var adh Adherent

	row := config.Db().QueryRow("SELECT * FROM adherents WHERE firstname = ? AND lastname = ?;", firstname, lastname)
	err := row.Scan(&adh.Id, &adh.Firstname, &adh.Lastname, &adh.Email, &adh.Dateofbirth, adh.ESNcard, adh.Student, &adh.CreatedAt, &adh.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &adh
}

func AllAdherents() *Adherents {
	var adhs Adherents

	rows, err := config.Db().Query("SELECT * FROM adherents")

	if err != nil {
		log.Fatal(err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var adh Adherent

		err := rows.Scan(&adh.Id, &adh.Firstname, &adh.Lastname, &adh.Email, &adh.Dateofbirth, &adh.ESNcard, &adh.Student, &adh.CreatedAt, &adh.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}

		adhs = append(adhs, adh)
	}

	return &adhs
}

func UpdateAdherent(adh *Adherent) {
	adh.UpdatedAt = time.Now()

	stmt, err := config.Db().Prepare("UPDATE adherents SET firstname=?, lastname=?, email=?, dateofbirth=?, esncard=?, student=?, updated_at=? WHERE id=?;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(adh.Firstname, adh.Lastname, adh.Email, adh.Dateofbirth, adh.ESNcard, adh.Student, adh.UpdatedAt, adh.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteAdherentById(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM adherents WHERE id=?;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
}
