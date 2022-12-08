package volunteer

// swagger:model Volunteer
type Volunteer struct {
	// Id of the volunteer
	// in: int64
	Id int `json:"id"`
	// Firstname of the volunteer
	// in: string
	Firstname string `json:"firstname"`
	// Lastname of the volunteer
	// in: string
	Lastname string `json:"lastname"`
	// Email of the volunteer
	// in: string
	Email string `json:"email"`
	// Status of the volunteer
	// in: bool
	Actif bool `json:"actif"`
	// Does the volunteer has a bureau role
	// in: bool
	Bureau bool `json:"bureau"`
}

type Volunteers []Volunteer

func NewVolunteer(vlt *Volunteer) {
	stmt, _ := TheDb().Prepare("INSERT INTO volunteers (firstname, lastname, email, actif, bureau) VALUES (?,?,?,?,?);")
	_, err := stmt.Exec(vlt.Firstname, vlt.Lastname, vlt.Email, vlt.Actif, vlt.Bureau)
	if err != nil {
		TheLogger().LogError("volunteer", "can't create new volunteer.", err)
	}
}

func FindVolunteerById(id int) *Volunteer {
	var vlt Volunteer

	row := TheDb().QueryRow("SELECT * FROM volunteers WHERE id = ?;", id)
	err := row.Scan(&vlt.Id, &vlt.Firstname, &vlt.Lastname, &vlt.Email, &vlt.Actif, &vlt.Bureau)

	if err != nil {
		TheLogger().LogWarning("volunteer", "volunteer id not found.", err)
	}

	return &vlt
}

func AllVolunteers() *Volunteers {
	var vlts Volunteers

	rows, err := TheDb().Query("SELECT * FROM volunteers")

	if err != nil {
		TheLogger().LogError("volunteer", "problem with the db.", err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var vlt Volunteer

		err := rows.Scan(&vlt.Id, &vlt.Firstname, &vlt.Lastname, &vlt.Email, &vlt.Actif, &vlt.Bureau)

		if err != nil {
			TheLogger().LogError("volunteer", "volunteers not found.", err)
		}

		vlts = append(vlts, vlt)
	}

	return &vlts
}

func UpdateVolunteer(vlt *Volunteer) {
	stmt, err := TheDb().Prepare("UPDATE volunteers SET firstname=?, lastname=?, email=?, actif=?, bureau=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError("volunteer", "problem with the db.", err)
	}

	_, err = stmt.Exec(vlt.Firstname, vlt.Lastname, vlt.Email, vlt.Actif, vlt.Bureau, vlt.Id)

	if err != nil {
		TheLogger().LogError("volunteer", "volunteer can't be updated.", err)
	}
}

func DeleteVolunteerById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM volunteers WHERE id=?;")

	if err != nil {
		TheLogger().LogError("volunteer", "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError("volunteer", "volunteer can't be deleted.", err)
	}

	return err
}
