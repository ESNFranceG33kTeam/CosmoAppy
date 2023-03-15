package esncard

import (
	"time"
)

// swagger:model ESNcard
type ESNcard struct {
	// Id of the esncard
	// in: int
	// read only: true
	Id int `json:"id"`
	// Id of the adherent
	// in: int64
	// required: true
	// example: 2
	Id_adherent int `json:"id_adherent"`
	// Code of the esncard
	// in: string
	// required: true
	// example: aVeryTooLongCode
	Esncard string `json:"esncard"`
	// CreatedAt date of the esncard
	// in: time.Time
	// read only: true
	CreatedAt time.Time `json:"created_at"`
}

type ESNcards []ESNcard

func NewESNcard(card *ESNcard) {
	card.CreatedAt = time.Now()

	stmt, _ := TheDb().Prepare(`INSERT INTO ` + db_name + `
		(id_adherent, esncard, created_at) VALUES (?,?,?);`)
	_, err := stmt.Exec(card.Id_adherent, card.Esncard, card.CreatedAt)
	if err != nil {
		TheLogger().LogError(log_name, "can't create new esncard.", err)
	}
}

func FindESNcardByIdadherent(id_adherent int) *ESNcard {
	var card ESNcard

	row := TheDb().QueryRow("SELECT * FROM "+db_name+" WHERE id_adherent = ?;", id_adherent)
	err := row.Scan(&card.Id, &card.Id_adherent, &card.Esncard, &card.CreatedAt)

	if err != nil {
		TheLogger().LogWarning(log_name, "esncard id_adherent not found.", err)
	}

	return &card
}

func FindESNcardByESNcard(esncard string) *ESNcard {
	var card ESNcard

	row := TheDb().QueryRow("SELECT * FROM "+db_name+" WHERE esncard = ?;", esncard)
	err := row.Scan(&card.Id, &card.Id_adherent, &card.Esncard, &card.CreatedAt)

	if err != nil {
		TheLogger().LogWarning(log_name, "esncard number not found.", err)
	}

	return &card
}

func FindESNcardsByDate(created_at string) *ESNcards {
	var cards ESNcards

	rows, err := TheDb().Query("SELECT * FROM "+db_name+" WHERE created_at like ?;", created_at+"%")

	if err != nil {
		TheLogger().LogWarning(log_name, "cards with created_at not found.", err)
	}

	for rows.Next() {
		var card ESNcard

		err := rows.Scan(
			&card.Id,
			&card.Id_adherent,
			&card.Esncard,
			&card.CreatedAt,
		)

		if err != nil {
			TheLogger().LogInfo(log_name, "cards not found.")
		}

		cards = append(cards, card)
	}

	return &cards
}

func AllESNcards() *ESNcards {
	var cards ESNcards

	rows, err := TheDb().Query("SELECT * FROM " + db_name)

	if err != nil {
		TheLogger().LogError(log_name, "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var card ESNcard

		err := rows.Scan(&card.Id, &card.Id_adherent, &card.Esncard, &card.CreatedAt)

		if err != nil {
			TheLogger().LogError(log_name, "esncards not found.", err)
		}

		cards = append(cards, card)
	}

	return &cards
}

func DeleteESNcardById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM " + db_name + " WHERE id=?;")

	if err != nil {
		TheLogger().LogError(log_name, "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError(log_name, "esncard can't be deleted.", err)
	}

	return err
}
