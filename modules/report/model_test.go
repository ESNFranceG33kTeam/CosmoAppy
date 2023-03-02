package report

import (
	"log"
	"testing"
)

func setUpModel() {
	NewReport(&Report{Type: "event", RefExt: 1, Name: "Report League", Date: "2016-04-14", Comment: "It was amazing !", NbReelAtt: 20, NbSubsAtt: 30, StaffsList: "Toto Titi, Tata Titi", NbHoursPrepa: 4, NbHours: 12, NbStaffsVlt: 2, NbStaffsEmp: 0, NbStaffsScv: 0, TauxValorisationVlt: 14, TauxValorisationEmp: 14, TauxValorisationScv: 14, CodePublic: "TST", CodeProject: "ALL"})
	NewReport(&Report{Type: "custom", RefExt: 0, Name: "Report Reu", Date: "2020-04-14", Comment: "It was amazing !", NbReelAtt: 2, NbSubsAtt: 0, StaffsList: "Toto Titi, Tata Titi", NbHoursPrepa: 2, NbHours: 12, NbStaffsVlt: 2, NbStaffsEmp: 0, NbStaffsScv: 0, TauxValorisationVlt: 14, TauxValorisationEmp: 14, TauxValorisationScv: 14, CodePublic: "TST", CodeProject: "ALL"})
	NewReport(&Report{Type: "event", RefExt: 2, Name: "Report NEM", Date: "2021-04-14", Comment: "It was amazing !", NbReelAtt: 35, NbSubsAtt: 30, StaffsList: "Toto Titi, Tata Titi", NbHoursPrepa: 4, NbHours: 18, NbStaffsVlt: 2, NbStaffsEmp: 0, NbStaffsScv: 0, TauxValorisationVlt: 14, TauxValorisationEmp: 14, TauxValorisationScv: 14, CodePublic: "TST", CodeProject: "ALL"})
}

func TestNewReport(t *testing.T) {
	NewReport(&Report{Type: "event", RefExt: 3, Name: "Report BeerPong", Date: "2021-04-14", Comment: "It was amazing !", NbReelAtt: 35, NbSubsAtt: 30, StaffsList: "Toto Titi, Tata Titi", NbHoursPrepa: 4, NbHours: 18, NbStaffsVlt: 2, NbStaffsEmp: 0, NbStaffsScv: 0, TauxValorisationVlt: 14, TauxValorisationEmp: 14, TauxValorisationScv: 14, CodePublic: "TST", CodeProject: "ALL"})
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
	rpt := Report{Id: 3, Type: "planning", RefExt: 1, Name: "Report NEEEEM", Date: "2021-04-14", Comment: "It was amazing !", NbReelAtt: 35, NbSubsAtt: 30, StaffsList: "Toto Titi, Tata Titi", NbHoursPrepa: 4, NbHours: 18, NbStaffsVlt: 2, NbStaffsEmp: 0, NbStaffsScv: 0, TauxValorisationVlt: 14, TauxValorisationEmp: 14, TauxValorisationScv: 14, CodePublic: "TST", CodeProject: "ALL"}
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
