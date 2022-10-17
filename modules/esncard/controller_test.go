package esncard

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUpController() {
	NewESNcard(&ESNcard{Id_adherent: 2, Esncard: "Luigi"})
	NewESNcard(&ESNcard{Id_adherent: 3, Esncard: "Salameche"})
	NewESNcard(&ESNcard{Id_adherent: 4, Esncard: "AveryLongCode"})
}

func TestESNcardsIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/esncards", nil)
	w := httptest.NewRecorder()
	ESNcardsIndex(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestESNcardsCreate(t *testing.T) {
	var jsonStr = []byte(`{"id_adherent": 2, "esncard": "Brawl"}`)

	req := httptest.NewRequest("POST", "/esncards", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	ESNcardsCreate(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}
