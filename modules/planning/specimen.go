package planning

func Specimen() {
	NewPlanning(&Planning{Name: "Permanence", Location: "MDE", Date_begins: "2023-04-23", Date_end: "2023-04-25", Hour_begins: "9-00", Hour_end: "18-00"})
	NewPlanning(&Planning{Name: "NEM", Location: "Transborder", Date_begins: "2023-04-23", Date_end: "2023-04-23", Hour_begins: "18-00", Hour_end: "23-00"})
	NewAttendee(&Attendee{Id_planning: 1, Id_adherent: 2, Date: "2023-04-23", Hour_begins: "10-00", Hour_end: "12-00"})
	NewAttendee(&Attendee{Id_planning: 1, Id_adherent: 3, Date: "2023-04-24", Hour_begins: "16-00", Hour_end: "18-00"})
}
