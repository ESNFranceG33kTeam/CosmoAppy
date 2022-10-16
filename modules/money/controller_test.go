package money

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUpController() {
	NewMoney(&Money{Label: "ESNcard", Price: 5})
	NewMoney(&Money{Label: "Event", Price: 2})
	NewMoney(&Money{Label: "Travel", Price: 15})
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
