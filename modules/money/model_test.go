package money

import (
	"log"
	"testing"
)

func setUpModel() {
	NewMoney(&Money{Label: "ESNcard", Price: 5})
	NewMoney(&Money{Label: "Event", Price: 2})
	NewMoney(&Money{Label: "Travel", Price: 15})
}

func TestNewMoney(t *testing.T) {
	NewMoney(&Money{Label: "ESNcard", Price: 5})
}

func TestFindMoneybyLabel(t *testing.T) {
	mons := FindMoneyByLabel("ESNcard")

	if len(*mons) == 0 {
		log.Fatal("Label is empty")
	}
}

func TestAllMoneys(t *testing.T) {
	mons := AllMoneys()

	//log.Println(adhs)
	if len(*mons) == 0 {
		log.Fatal("Money is empty")
	}
}
