package event

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUpController() {
	NewEvent(&Event{Name: "Voyage a Hawai", Date: "2023-04-23", Location: "Hawai", Type: "voyage", Price: 120.42, Url: "facebook.com/voyageHawai", Actif: true})
	NewEvent(&Event{Name: "Saturday Night Fever", Date: "2023-05-23", Location: "3 rue Albert 1er, 69000 Lyon", Type: "soiree", Price: 20, Url: "facebook.com/sturdayfever", Actif: false})
	NewAttendee(&Attendee{Id_event: 1, Id_adherent: 2, Staff: true})
	NewAttendee(&Attendee{Id_event: 1, Id_adherent: 3, Staff: false})
}

func TestEventsIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/events", nil)
	w := httptest.NewRecorder()
	EventsIndex(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestEventsCreate(t *testing.T) {
	var jsonStr = []byte(`{"name": "Visite au Musée", "date": "2024-12-25", "location": "Musée des miniatures, 69000 Lyon", "type": "culture", "price": 0, "url_facebook": "facebook.com/eventmusee", "actif": true}`)

	req := httptest.NewRequest("POST", "/events", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	EventsCreate(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestAttendeesIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/event_attendees", nil)
	w := httptest.NewRecorder()
	AttendeesIndex(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestAttendeesCreate(t *testing.T) {
	var jsonStr = []byte(`{"id_event": 1, "id_adherent": 2, "staff": true}`)

	req := httptest.NewRequest("POST", "/event_attendees", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	AttendeesCreate(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}
