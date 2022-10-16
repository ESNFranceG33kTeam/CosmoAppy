package money

import (
	"time"

	"github.com/ESNFranceG33kTeam/sAPI/config"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
)

// swagger:model Money
type Money struct {
	// Id of the entry
	// in: int64
	Id int `json:"id"`
	// Label of the operation
	// in: string
	Label string `json:"label"`
	// Price of the operation
	// in: float
	Price float64 `json:"price"`
	// CreatedAt date of the operation
	// in: time.Time
	CreatedAt time.Time `json:"created_at"`
}

type Moneys []Money

func NewMoney(mon *Money) {
	mon.CreatedAt = time.Now()

	stmt, _ := config.Db().Prepare("INSERT INTO moneys (label, price, created_at) VALUES (?,?,?);")
	_, err := stmt.Exec(mon.Label, mon.Price, mon.CreatedAt)
	if err != nil {
		logger.LogError("money", "can't create new operation.", err)
	}
}

func FindMoneyByLabel(label string) *Moneys {
	var mons Moneys

	rows, err := config.Db().Query("SELECT * FROM moneys WHERE label = ?;", label)

	if err != nil {
		logger.LogWarning("money", "operations with label not found.", err)
	}

	for rows.Next() {
		var mon Money

		err := rows.Scan(&mon.Id, &mon.Label, &mon.Price, &mon.CreatedAt)

		if err != nil {
			logger.LogError("money", "operations not found.", err)
		}

		mons = append(mons, mon)
	}

	return &mons
}

func AllMoneys() *Moneys {
	var mons Moneys

	rows, err := config.Db().Query("SELECT * FROM moneys")

	if err != nil {
		logger.LogError("money", "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var mon Money

		err := rows.Scan(&mon.Id, &mon.Label, &mon.Price, &mon.CreatedAt)

		if err != nil {
			logger.LogError("money", "operations not found.", err)
		}

		mons = append(mons, mon)
	}

	return &mons
}
