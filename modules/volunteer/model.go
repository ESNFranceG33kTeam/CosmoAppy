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
	// Discord pseudo
	// in: string
	Discord string `json:"discord"`
	// Phone number
	// in: string
	Phone string `json:"phone"`
	// University name
	// in: string
	University string `json:"university"`
	// Postal address
	// in: string
	PostalAddress string `json:"postal_address"`
	// Status of the volunteer
	// in: bool
	Actif bool `json:"actif"`
	// Does the volunteer has a bureau role
	// in: bool
	Bureau bool `json:"bureau"`
	// Status of the volunteer
	// in: string
	HrStatus string `json:"hr_status"`
	// Started of date of volunteering
	// in: string
	StartedDate string `json:"started_date"`
}

type Volunteers []Volunteer

func NewVolunteer(vlt *Volunteer) {
	stmt, _ := TheDb().Prepare(
		`INSERT INTO volunteers
			(firstname, lastname, email, discord, phone, university, postal_address, actif, bureau,
			hr_status, started_date)
		VALUES (?,?,?,?,?,?,?,?,?,?,?);`,
	)
	_, err := stmt.Exec(
		vlt.Firstname,
		vlt.Lastname,
		vlt.Email,
		vlt.Discord,
		vlt.Phone,
		vlt.University,
		vlt.PostalAddress,
		vlt.Actif,
		vlt.Bureau,
		vlt.HrStatus,
		vlt.StartedDate,
	)
	if err != nil {
		TheLogger().LogError("volunteer", "can't create new volunteer.", err)
	}
}

func FindVolunteerById(id int) *Volunteer {
	var vlt Volunteer

	row := TheDb().QueryRow("SELECT * FROM volunteers WHERE id = ?;", id)
	err := row.Scan(
		&vlt.Id,
		&vlt.Firstname,
		&vlt.Lastname,
		&vlt.Email,
		&vlt.Discord,
		&vlt.Phone,
		&vlt.University,
		&vlt.PostalAddress,
		&vlt.Actif,
		&vlt.Bureau,
		&vlt.HrStatus,
		&vlt.StartedDate,
	)

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

		err := rows.Scan(
			&vlt.Id,
			&vlt.Firstname,
			&vlt.Lastname,
			&vlt.Email,
			&vlt.Discord,
			&vlt.Phone,
			&vlt.University,
			&vlt.PostalAddress,
			&vlt.Actif,
			&vlt.Bureau,
			&vlt.HrStatus,
			&vlt.StartedDate,
		)

		if err != nil {
			TheLogger().LogError("volunteer", "volunteers not found.", err)
		}

		vlts = append(vlts, vlt)
	}

	return &vlts
}

func UpdateVolunteer(vlt *Volunteer) {
	stmt, err := TheDb().Prepare(
		`UPDATE volunteers SET
			firstname=?, lastname=?, email=?, discord=?, phone=?, university=?, postal_address=?,
			actif=?, bureau=?, hr_status=?, started_date=?
		WHERE id=?;`,
	)

	if err != nil {
		TheLogger().LogError("volunteer", "problem with the db.", err)
	}

	_, err = stmt.Exec(
		vlt.Firstname,
		vlt.Lastname,
		vlt.Email,
		vlt.Discord,
		vlt.Phone,
		vlt.University,
		vlt.PostalAddress,
		vlt.Actif,
		vlt.Bureau,
		vlt.HrStatus,
		vlt.StartedDate,
		vlt.Id,
	)

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
