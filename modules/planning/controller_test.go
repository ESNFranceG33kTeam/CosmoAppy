package planning

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUpController() {
	NewPlanning(
		&Planning{
			Name:        "Permanence",
			Type:        "Permanence",
			Location:    "MDE",
			Date_begins: "2023-04-23",
			Date_end:    "2023-04-25",
			Hour_begins: "9:00:00",
			Hour_end:    "18:00:00",
		},
	)
	NewPlanning(
		&Planning{
			Name:        "NEM",
			Type:        "Event",
			Location:    "Transborder",
			Date_begins: "2023-04-23",
			Date_end:    "2023-04-23",
			Hour_begins: "18:00:00",
			Hour_end:    "23:00:00",
		},
	)
	NewAttendee(
		&Attendee{
			Id_planning:  1,
			Id_volunteer: 2,
			Job:          "Photographer",
			Date:         "2023-04-23",
			Hour_begins:  "10:00:00",
			Hour_end:     "12:00:00",
		},
	)
	NewAttendee(
		&Attendee{
			Id_planning:  1,
			Id_volunteer: 3,
			Job:          "Staff",
			Date:         "2023-04-24",
			Hour_begins:  "16:00:00",
			Hour_end:     "18:00:00",
		},
	)
}

func TestPlanningsIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/plannings", nil)
	w := httptest.NewRecorder()
	PlanningsIndex(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestPlanningsCreate(t *testing.T) {
	var jsonStr = []byte(`
		{"name": "Permanence Spe", "location": "MDE", "date_begins": "2024-12-25",
		"date_end": "2024-12-25", "hour_begins": "9:00", "hour_end": "12:00"}`,
	)

	req := httptest.NewRequest("POST", "/plannings", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	PlanningsCreate(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestAttendeesIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/planning_attendees", nil)
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
	var jsonStr = []byte(`
		{"id_planning": 1, "id_volunteer": 2, "job": "Staff", "date": "2023-04-25",
		"hour_begins": "12:00", "hour_end": "14:00"}`,
	)

	req := httptest.NewRequest("POST", "/planning_attendees", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	AttendeesCreate(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}
