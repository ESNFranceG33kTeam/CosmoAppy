package volunteer

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUpController() {
	NewVolunteer(
		&Volunteer{
			Firstname: "Toto",
			Lastname:  "Bob",
			Email:     "toto.toto@bob.com",
			Actif:     true,
			Bureau:    false,
		},
	)
	NewVolunteer(
		&Volunteer{
			Firstname: "Toto",
			Lastname:  "Bob",
			Email:     "toto.toto@bob.com",
			Actif:     false,
			Bureau:    false,
		},
	)
	NewVolunteer(
		&Volunteer{
			Firstname: "Toto",
			Lastname:  "Bob",
			Email:     "toto.toto@bob.com",
			Actif:     true,
			Bureau:    true,
		},
	)
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
	var jsonStr = []byte(`
		{"fistname": "Titi", "lastname": "titi", "email": "titi@titi.com", "actif": true,
		"bureau": false}`,
	)

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
