package money

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setUpController() {
	NewMoney(
		&Money{
			Label:       "ESNcard",
			Price:       5,
			PaymentType: "Carte bancaire",
			PaymentDate: time.Now().AddDate(0, -1, 0).Format("2006-01-02"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Event",
			Price:       2,
			PaymentType: "Carte bancaire",
			PaymentDate: time.Now().AddDate(0, -1, 0).Format("2006-01-02"),
		},
	)
	NewMoney(
		&Money{
			Label:       "Travel",
			Price:       15,
			PaymentType: "Carte bancaire",
			PaymentDate: time.Now().AddDate(0, 0, 0).Format("2006-01-02"),
		},
	)
}

func TestMoneysIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/moneys", nil)
	w := httptest.NewRecorder()
	MoneysIndex(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestMoneysCreate(t *testing.T) {
	var jsonStr = []byte(`{"label": "ESNcard", "price": 5}`)

	req := httptest.NewRequest("POST", "/moneys", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	MoneysCreate(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestAutoMonthlyStatsIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/moneys/stats/monthly", nil)
	w := httptest.NewRecorder()
	MonthlyStatsIndex(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}
