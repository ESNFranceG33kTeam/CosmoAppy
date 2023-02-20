package report

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUpController() {
	NewReport(&Report{Type: "event", RefExt: 1, Name: "Report League", Date: "2016-04-14", Comment: "It was amazing !", NbReelAtt: 20, NbSubsAtt: 30, StaffsList: "Toto Titi, Tata Titi", NbHoursPrepa: 4, NbHours: 12, NbStaffs: 10, TauxValorisation: 20, CodePublic: "TST", CodeProject: "ALL"})
	NewReport(&Report{Type: "custom", RefExt: 0, Name: "Report Reu", Date: "2020-04-14", Comment: "It was amazing !", NbReelAtt: 2, NbSubsAtt: 0, StaffsList: "Toto Titi, Tata Titi", NbHoursPrepa: 2, NbHours: 12, NbStaffs: 1, TauxValorisation: 14, CodePublic: "TST", CodeProject: "ALL"})
	NewReport(&Report{Type: "event", RefExt: 2, Name: "Report NEM", Date: "2021-04-14", Comment: "It was amazing !", NbReelAtt: 35, NbSubsAtt: 30, StaffsList: "Toto Titi, Tata Titi", NbHoursPrepa: 4, NbHours: 18, NbStaffs: 10, TauxValorisation: 18, CodePublic: "TST", CodeProject: "ALL"})
}

func TestReportsIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/reports", nil)
	w := httptest.NewRecorder()
	ReportsIndex(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestReportsCreate(t *testing.T) {
	var jsonStr = []byte(`{"type": "planning", "ref_ext": 4, "name": "Report Meeting", "date": "2022-04-14", "comment": "Nice", "nb_reel_attendees": 12, "nb_subscribe_attendees": 12, "staffs_list": "Toto Titi", "nb_hours_prepa": 2, "nb_hours": 2, "nb_staffs": 1, "taux_valorisation": 18, "code_public": "TST", "code_project": "ALL"}`)

	req := httptest.NewRequest("POST", "/reports", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	ReportsCreate(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}
