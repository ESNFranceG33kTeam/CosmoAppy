package planning

func Specimen() {
	NewPlanning(&Planning{Name: "Permanence", Type: "permanence", Location: "Conseil des 4", Date_begins: "2023-04-23", Date_end: "2023-04-24", Hour_begins: "9:00:00", Hour_end: "18:00:00"})
	NewPlanning(&Planning{Name: "PokePong", Type: "event", Location: "Transborder", Date_begins: "2023-04-23", Date_end: "2023-04-23", Hour_begins: "18:00:00", Hour_end: "23:00:00"})
	NewAttendee(&Attendee{Id_planning: 1, Id_volunteer: 3, Job: "staff", Date: "2023-04-23", Hour_begins: "10:00:00", Hour_end: "12:00:00"})
	NewAttendee(&Attendee{Id_planning: 1, Id_volunteer: 4, Job: "master", Date: "2023-04-24", Hour_begins: "16:00:00", Hour_end: "18:00:00"})
}
