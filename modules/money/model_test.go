package money

import (
	"log"
	"testing"
	"time"
)

func setUpModel() {
	NewMoney(&Money{Label: "ESNcard", Price: 5, PaymentType: "cb", PaymentDate: time.Now().AddDate(0, -1, 0).Format("2006-01-02")})
	NewMoney(&Money{Label: "Event", Price: 2, PaymentType: "cash", PaymentDate: time.Now().AddDate(0, -1, 0).Format("2006-01-02")})
	NewMoney(&Money{Label: "Travel", Price: 15, PaymentType: "cb", PaymentDate: time.Now().AddDate(0, 0, 0).Format("2006-01-02")})
}

func TestNewMoney(t *testing.T) {
	NewMoney(&Money{Label: "ESNcard", Price: 5, PaymentType: "cb", PaymentDate: "2023-01-10"})
}

func TestFindMoneybyLabel(t *testing.T) {
	mons := FindMoneysByLabel("ESNcard")

	if len(*mons) == 0 {
		log.Fatal("Label is empty")
	}
}

func TestFindMoneysbyDate(t *testing.T) {
	mons := FindMoneysByDate(time.Now().AddDate(0, 0, 0).Format("2006-01"))

	if len(*mons) == 0 {
		log.Fatal("Date is empty")
	}
}

func TestAllMoneys(t *testing.T) {
	mons := AllMoneys()

	if len(*mons) == 0 {
		log.Fatal("Money is empty")
	}
}

func TestAutoNewMonthlyStat(t *testing.T) {
	AutoNewMonthlyStat()
	AutoNewMonthlyStat(time.Now().AddDate(0, 0, 0).Format("2006-01"))
}

func TestFindMonthlyStatByDate(t *testing.T) {
	stat := FindMonthlyStatByDate(time.Now().AddDate(0, -1, 0).Format("2006-01"))

	if stat.Id == 0 {
		log.Fatal("Monthly stat of date is empty")
	}
}

func TestAllMonthlyStats(t *testing.T) {
	stats := AllMonthlyStats()

	if len(*stats) == 0 {
		log.Fatal("Monthly stat is empty")
	}
}
