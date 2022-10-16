package controllers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/models"
)

func setUpVolunteer() {
	models.NewVolunteer(&models.Volunteer{Id_adherent: 2, Actif: true, Bureau: false})
	models.NewVolunteer(&models.Volunteer{Id_adherent: 3, Actif: false, Bureau: false})
	models.NewVolunteer(&models.Volunteer{Id_adherent: 4, Actif: true, Bureau: true})
}

func TestVolunteersIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/volunteers", nil)
	w := httptest.NewRecorder()
	VolunteersIndex(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestVolunteersCreate(t *testing.T) {
	var jsonStr = []byte(`{"id_adherent": 1, "actif": true, "bureau": false}`)

	req := httptest.NewRequest("POST", "/volunteers", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	VolunteersCreate(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}
