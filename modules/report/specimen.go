package report

func Specimen() {
	NewReport(&Report{Type: "event", RefExt: 1, Name: "Report League Indigo", Date: "2016-04-14", Comment: "It was amazing !", NbReelAtt: 25, NbSubsAtt: 30, StaffsList: "Cynthia Shirona, Peter Wataru", NbHoursPrepa: 8, NbHours: 20, NbStaffs: 2, TauxValorisation: 14, CodePublic: "CHAMP", CodeProject: "PKM"})
	NewReport(&Report{Type: "event", RefExt: 2, Name: "Report Masters", Date: "2020-04-14", Comment: "FIIIIIRE !", NbReelAtt: 25, NbSubsAtt: 30, StaffsList: "Cynthia Shirona", NbHoursPrepa: 8, NbHours: 30, NbStaffs: 1, TauxValorisation: 18, CodePublic: "MAST", CodeProject: "PKM"})
	NewReport(&Report{Type: "planning", RefExt: 1, Name: "Report Permanence", Date: "2020-04-14", Comment: "Not lot of people !", NbReelAtt: 8, NbSubsAtt: 0, StaffsList: "Cynthia Shirona", NbHoursPrepa: 0, NbHours: 8, NbStaffs: 1, TauxValorisation: 18, CodePublic: "CHAMP", CodeProject: "PKM"})
	NewReport(&Report{Type: "custom", RefExt: 0, Name: "Report Meeting with Mayor", Date: "2020-04-14", Comment: "Good meeting !", NbReelAtt: 2, NbSubsAtt: 2, StaffsList: "Cynthia Shirona", NbHoursPrepa: 1, NbHours: 2, NbStaffs: 1, TauxValorisation: 18, CodePublic: "CHAMP", CodeProject: "PKM"})
}
