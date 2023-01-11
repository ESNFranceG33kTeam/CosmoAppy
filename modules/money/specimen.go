package money

func Specimen() {
	NewMoney(&Money{Label: "ESNcard", Price: 5})
	NewMoney(&Money{Label: "ESNcard", Price: 5})

	NewMoney(&Money{Label: "Adhesion", Price: 1})
	NewMoney(&Money{Label: "Adhesion", Price: 1})
	NewMoney(&Money{Label: "Adhesion", Price: 1})
	NewMoney(&Money{Label: "Adhesion", Price: 1})
	NewMoney(&Money{Label: "Adhesion", Price: 1})
	NewMoney(&Money{Label: "Adhesion", Price: 1})
	NewMoney(&Money{Label: "Adhesion", Price: 1})
	NewMoney(&Money{Label: "Adhesion", Price: 1})
	NewMoney(&Money{Label: "Adhesion", Price: 1})

	NewMoney(&Money{Label: "Event", Price: 120})
	NewMoney(&Money{Label: "Event", Price: 0})
	NewMoney(&Money{Label: "Event", Price: 120})
	NewMoney(&Money{Label: "Event", Price: 120})
	NewMoney(&Money{Label: "Event", Price: 80})
	NewMoney(&Money{Label: "Event", Price: 80})
	NewMoney(&Money{Label: "Event", Price: 0})
	NewMoney(&Money{Label: "Event", Price: 0})
}
