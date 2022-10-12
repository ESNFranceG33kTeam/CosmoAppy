package models

import (
	"time"

	"github.com/ESNFranceG33kTeam/sAPI/config"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
)

// swagger:model ESNcard
type ESNcard struct {
	// Id of the esncard
	// in: int64
	Id int `json:"id"`
	// Id of the adherent
	// in: int64
	Id_adherent int `json:"id_adherent"`
	// Code of the esncard
	// in: string
	Esncard string `json:"esncard"`
	// CreatedAt date of the esncard
	// in: time.Time
	CreatedAt time.Time `json:"created_at"`
}

type ESNcards []ESNcard

func NewESNcard(card *ESNcard) {
	card.CreatedAt = time.Now()

	stmt, _ := config.Db().Prepare("INSERT INTO esncards (id_adherent, esncard, created_at) VALUES (?,?,?);")
	_, err := stmt.Exec(card.Id_adherent, card.Esncard, card.CreatedAt)
	if err != nil {
		logger.LogError("esncard", "can't create new esncard.", err)
	}
}

func FindESNcardByIdadherent(id_adherent int) *ESNcard {
	var card ESNcard

	row := config.Db().QueryRow("SELECT * FROM esncards WHERE id_adherent = ?;", id_adherent)
	err := row.Scan(&card.Id, &card.Id_adherent, &card.Esncard, &card.CreatedAt)

	if err != nil {
		logger.LogWarning("esncard", "esncard id_adherent not found.", err)
	}

	return &card
}

func FindESNcardByESNcard(esncard string) *ESNcard {
	var card ESNcard

	row := config.Db().QueryRow("SELECT * FROM esncards WHERE esncard = ?;", esncard)
	err := row.Scan(&card.Id, &card.Id_adherent, &card.Esncard, &card.CreatedAt)

	if err != nil {
		logger.LogWarning("esncard", "esncard number not found.", err)
	}

	return &card
}

func AllESNcards() *ESNcards {
	var cards ESNcards

	rows, err := config.Db().Query("SELECT * FROM esncards")

	if err != nil {
		logger.LogError("esncard", "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var card ESNcard

		err := rows.Scan(&card.Id, &card.Id_adherent, &card.Esncard, &card.CreatedAt)

		if err != nil {
			logger.LogError("esncard", "esncards not found.", err)
		}

		cards = append(cards, card)
	}

	return &cards
}

func DeleteESNcardById(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM esncards WHERE id=?;")

	if err != nil {
		logger.LogError("esncard", "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		logger.LogError("esncard", "esncard can't be deleted.", err)
	}

	return err
}
