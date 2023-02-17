package report

import (
	"log"
	"testing"
)

func setUpModel() {
	NewReport(&Report{Name: "Report League", Date: "2016-04-14", Comment: "It was amazing !", NbHours: 12, NbStaffs: 10, TauxValorisation: 20})
	NewReport(&Report{Name: "Report Reu", Date: "2020-04-14", Comment: "It was amazing !", NbHours: 12, NbStaffs: 20, TauxValorisation: 14})
	NewReport(&Report{Name: "Report NEM", Date: "2021-04-14", Comment: "It was amazing !", NbHours: 18, NbStaffs: 10, TauxValorisation: 18})
}

func TestNewReport(t *testing.T) {
	NewReport(&Report{Name: "Report BeerPong", Date: "2021-04-14", Comment: "It was amazing !", NbHours: 18, NbStaffs: 10, TauxValorisation: 20})
}

func TestFindReportById(t *testing.T) {
	rpt := FindReportById(2)

	if rpt.Name != "Report Reu" {
		log.Fatal("Report is not the good one.")
	}
}

func TestAllReports(t *testing.T) {
	rpts := AllReports()

	if len(*rpts) == 0 {
		log.Fatal("Report is empty")
	}
}

func TestUpdateReport(t *testing.T) {
	rpt := Report{Id: 3, Name: "Report NEEEEM", Date: "2021-04-14", Comment: "It was amazing !", NbHours: 18, NbStaffs: 10, TauxValorisation: 18}
	UpdateReport(&rpt)

	rpt_3 := FindReportById(3)
	if rpt_3.Name != "Report NEEEEM" {
		log.Fatal("Report_3 didn't updated !")
	}
}

func TestDeleteReportById(t *testing.T) {
	err := DeleteReportById(3)

	if err != nil {
		log.Fatal("Delete had a problem : ", err)
	}

	vlts := AllReports()

	for _, vlt := range *vlts {
		if vlt.Id == 3 {
			log.Fatal("Report_3 didn't be removed !")
		}
	}
}
