package volunteer

// swagger:model Volunteer
type Volunteer struct {
	// Id of the volunteer
	// in: int64
	Id int `json:"id"`
	// Id of the adherent
	// in: int64
	Id_adherent int `json:"id_adherent"`
	// Status of the volunteer
	// in: bool
	Actif bool `json:"actif"`
	// Does the volunteer has a bureau role
	// in: bool
	Bureau bool `json:"bureau"`
}

type Volunteers []Volunteer

func NewVolunteer(vlt *Volunteer) {
	stmt, _ := TheDb().Prepare("INSERT INTO volunteers (id_adherent, actif, bureau) VALUES (?,?,?);")
	_, err := stmt.Exec(vlt.Id_adherent, vlt.Actif, vlt.Bureau)
	if err != nil {
		TheLogger().LogError("volunteer", "can't create new volunteer.", err)
	}
}

func FindVolunteerByIdadherent(id_adherent int) *Volunteer {
	var vlt Volunteer

	row := TheDb().QueryRow("SELECT * FROM volunteers WHERE id_adherent = ?;", id_adherent)
	err := row.Scan(&vlt.Id, &vlt.Id_adherent, &vlt.Actif, &vlt.Bureau)

	if err != nil {
		TheLogger().LogWarning("volunteer", "volunteer id_adherent not found.", err)
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

		err := rows.Scan(&vlt.Id, &vlt.Id_adherent, &vlt.Actif, &vlt.Bureau)

		if err != nil {
			TheLogger().LogError("volunteer", "volunteers not found.", err)
		}

		vlts = append(vlts, vlt)
	}

	return &vlts
}

func UpdateVolunteer(vlt *Volunteer) {
	stmt, err := TheDb().Prepare("UPDATE volunteers SET id_adherent=?, actif=?, bureau=? WHERE id=?;")

	if err != nil {
		TheLogger().LogError("volunteer", "problem with the db.", err)
	}

	_, err = stmt.Exec(vlt.Id_adherent, vlt.Actif, vlt.Bureau, vlt.Id)

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
		TheLogger().LogError("volunteer", "esncard can't be deleted.", err)
	}

	return err
}
