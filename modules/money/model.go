package money

import (
	"time"
)

// swagger:model Money
type Money struct {
	// Id of the entry
	// in: int
	// read only: true
	Id int `json:"id"`
	// Label of the operation
	// in: string
	// example: event
	// required: true
	Label string `json:"label"`
	// Price of the operation
	// in: float
	// example: 1.5
	// required: true
	Price float64 `json:"price"`
	// Payment type of the operation
	// in: string
	// example: cash
	// required: true
	PaymentType string `json:"payment_type"`
	// Payment date of the operation
	// in: string
	// example: 2023-01-23
	// required: true
	PaymentDate string `json:"payment_date"`
	// CreatedAt date of the operation
	// in: time.Time
	// read only: true
	CreatedAt time.Time `json:"created_at"`
}

type Moneys []Money

func NewMoney(mon *Money) {
	mon.CreatedAt = time.Now()

	stmt, _ := TheDb().Prepare(
		`INSERT INTO ` + db_name + `
		(label, price, payment_type, payment_date, created_at)
		VALUES (?,?,?,?,?);`,
	)
	_, err := stmt.Exec(mon.Label, mon.Price, mon.PaymentType, mon.PaymentDate, mon.CreatedAt)
	if err != nil {
		TheLogger().LogError(log_name, "can't create new operation.", err)
	}
}

func FindMoneysByLabel(label string) *Moneys {
	var mons Moneys

	rows, err := TheDb().Query("SELECT * FROM "+db_name+" WHERE label = ?;", label)

	if err != nil {
		TheLogger().LogWarning(log_name, "operations with label not found.", err)
	}

	for rows.Next() {
		var mon Money

		err := rows.Scan(
			&mon.Id,
			&mon.Label,
			&mon.Price,
			&mon.PaymentType,
			&mon.PaymentDate,
			&mon.CreatedAt,
		)

		if err != nil {
			TheLogger().LogError(log_name, "operations not found.", err)
		}

		mons = append(mons, mon)
	}

	return &mons
}

func FindMoneysByDate(payment_date string) *Moneys {
	var mons Moneys

	rows, err := TheDb().Query("SELECT * FROM "+db_name+" WHERE payment_date like ?;", payment_date+"%")

	if err != nil {
		TheLogger().LogWarning(log_name, "operations with payment_date not found.", err)
	}

	for rows.Next() {
		var mon Money

		err := rows.Scan(
			&mon.Id,
			&mon.Label,
			&mon.Price,
			&mon.PaymentType,
			&mon.PaymentDate,
			&mon.CreatedAt,
		)

		if err != nil {
			TheLogger().LogInfo(log_name, "operations not found.")
		}

		mons = append(mons, mon)
	}

	return &mons
}

func AllMoneys() *Moneys {
	var mons Moneys

	rows, err := TheDb().Query("SELECT * FROM " + db_name)

	if err != nil {
		TheLogger().LogError(log_name, "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var mon Money

		err := rows.Scan(
			&mon.Id,
			&mon.Label,
			&mon.Price,
			&mon.PaymentType,
			&mon.PaymentDate,
			&mon.CreatedAt,
		)

		if err != nil {
			TheLogger().LogError(log_name, "operations not found.", err)
		}

		mons = append(mons, mon)
	}

	return &mons
}
