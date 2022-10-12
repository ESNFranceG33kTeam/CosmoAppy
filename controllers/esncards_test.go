package controllers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/models"
)

func setUpESNcard() {
	models.NewESNcard(&models.ESNcard{Id_adherent: 2, Esncard: "Mario"})
	models.NewESNcard(&models.ESNcard{Id_adherent: 3, Esncard: "Pikachu"})
	models.NewESNcard(&models.ESNcard{Id_adherent: 4, Esncard: "AveryLittleCode"})
}

func TestESNcardsIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/esncards", nil)
	w := httptest.NewRecorder()
	AdherentsIndex(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestESNcardsCreate(t *testing.T) {
	var jsonStr = []byte(`{"id_adherent": 2, "esncard": "Bros"}`)

	req := httptest.NewRequest("POST", "/esncards", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	AdherentsCreate(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}
