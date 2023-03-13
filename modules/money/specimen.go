package money

import "time"

func payment_date(paymentdate string) string {
	datedate, _ := time.Parse("2006-01-02", paymentdate)
	payment_date := datedate.Format("2006-01-02")

	return payment_date
}

func Specimen() {
	NewMoney(
		&Money{
			Label:       "Adhesion",
			Price:       1,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2021-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Adhesion",
			Price:       1,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Adhesion",
			Price:       1,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Adhesion",
			Price:       1,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Adhesion",
			Price:       1,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Adhesion",
			Price:       1,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Adhesion",
			Price:       1,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Adhesion",
			Price:       1,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Adhesion",
			Price:       1,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Adhesion",
			Price:       1,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)

	NewMoney(
		&Money{
			Label:       "ESNcard",
			Price:       5,
			PaymentType: "Cash",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "ESNcard",
			Price:       5,
			PaymentType: "Cash",
			PaymentDate: payment_date("2022-04-24"),
		},
	)

	NewMoney(
		&Money{
			Label:       "Event",
			Price:       120,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Event",
			Price:       120,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Event",
			Price:       120,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Event",
			Price:       80,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Event",
			Price:       80,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2022-04-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Event",
			Price:       80,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2023-02-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Event",
			Price:       20,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2023-02-24"),
		},
	)
	NewMoney(
		&Money{
			Label:       "ESNcard",
			Price:       80,
			PaymentType: "Carte bancaire",
			PaymentDate: payment_date("2023-02-01"),
		},
	)
}
