package planning

import (
	"log"
	"testing"
)

func setUpModel() {
	NewPlanning(&Planning{Name: "Permanence", Type: "Permanence", Location: "MDE", Date_begins: "2023-04-23", Date_end: "2023-04-25", Hour_begins: "9-00", Hour_end: "18-00"})
	NewPlanning(&Planning{Name: "NEM", Type: "Event", Location: "Transborder", Date_begins: "2023-04-23", Date_end: "2023-04-23", Hour_begins: "18-00", Hour_end: "23-00"})
	NewAttendee(&Attendee{Id_planning: 1, Id_volunteer: 2, Job: "Photographer", Date: "2023-04-23", Hour_begins: "10-00", Hour_end: "12-00"})
	NewAttendee(&Attendee{Id_planning: 1, Id_volunteer: 3, Job: "Staff", Date: "2023-04-24", Hour_begins: "16-00", Hour_end: "18-00"})
}

func TestNewPlanning(t *testing.T) {
	NewPlanning(&Planning{Name: "Permanence Special", Location: "MDE", Date_begins: "2023-04-22", Date_end: "2023-04-22", Hour_begins: "9-00", Hour_end: "12-00"})
}

func TestFindPlanningById(t *testing.T) {
	pla := FindPlanningById(2)

	if pla.Id != 2 {
		log.Fatal("Planning is not him")
	}
}

func TestAllPlannings(t *testing.T) {
	plas := AllPlannings()

	//log.Println(adhs)
	if len(*plas) == 0 {
		log.Fatal("Planning is empty")
	}
}

func TestUpdatePlanning(t *testing.T) {
	pla := Planning{Id: 2, Name: "NEM", Location: "Transborder", Date_begins: "2023-04-23", Date_end: "2023-04-23", Hour_begins: "18-00", Hour_end: "23-59"}
	UpdatePlanning(&pla)

	pla_2 := FindPlanningById(2)
	if pla_2.Hour_end != "23-59" {
		log.Fatal("Pla_2 didn't updated !")
	}
}

func TestDeletePlanningById(t *testing.T) {
	err := DeletePlanningById(3)

	if err != nil {
		log.Fatal("Delete had a problem : ", err)
	}

	plas := AllPlannings()

	for _, pla := range *plas {
		if pla.Id == 3 {
			log.Fatal("Eve_3 didn't be removed !")
		}
	}
}
func TestNewAttendee(t *testing.T) {
	NewAttendee(&Attendee{Id_planning: 2, Id_volunteer: 3, Job: "Staff", Date: "2023-04-24", Hour_begins: "16-00", Hour_end: "18-00"})
}

func TestAllAttendees(t *testing.T) {
	atts := AllAttendees()

	//log.Println(adhs)
	if len(*atts) == 0 {
		log.Fatal("Attendee is empty")
	}
}

func TestFindAttendeeById(t *testing.T) {
	att := FindPlanningById(2)

	if att.Id != 2 {
		log.Fatal("Attendee is not him")
	}
}

func TestFindAttendeeByPlanningId(t *testing.T) {
	atts := FindAttendeeByPlanningId(1)

	if len(*atts) == 0 {
		log.Fatal("Attendee is empty")
	}
}

func TestFindAttendeeByVolunteerId(t *testing.T) {
	atts := FindAttendeeByVolunteerId(3)

	if len(*atts) == 0 {
		log.Fatal("Attendee is empty")
	}
}

func TestUpdateAttendee(t *testing.T) {
	att := Attendee{Id: 1, Id_planning: 1, Id_volunteer: 2, Job: "Staff", Date: "2023-04-23", Hour_begins: "10-00", Hour_end: "14-00"}
	UpdateAttendee(&att)

	att_1 := FindAttendeeById(1)
	if att_1.Hour_end != "14-00" {
		log.Fatal("att_1 didn't updated !")
	}
}

func TestDeleteAttendeeById(t *testing.T) {
	err := DeleteAttendeeById(3)

	if err != nil {
		log.Fatal("Delete had a problem : ", err)
	}

	atts := AllAttendees()

	for _, att := range *atts {
		if att.Id == 3 {
			log.Fatal("Att_3 didn't be removed !")
		}
	}
}
