package report

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUpController() {
	NewReport(&Report{Name: "Report League", Date: "2016-04-14", Comment: "It was amazing !", NbHours: 12, NbStaffs: 10, TauxValorisation: 20})
	NewReport(&Report{Name: "Report Reu", Date: "2020-04-14", Comment: "It was amazing !", NbHours: 12, NbStaffs: 20, TauxValorisation: 14})
	NewReport(&Report{Name: "Report NEM", Date: "2021-04-14", Comment: "It was amazing !", NbHours: 18, NbStaffs: 10, TauxValorisation: 18})
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
	var jsonStr = []byte(`{"name": "Report Meeting", "date": "2022-04-14", "comment": "Nice", "nb_hours": 2, "nb_staffs": 1, "taux_valorisation": 18}`)

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
