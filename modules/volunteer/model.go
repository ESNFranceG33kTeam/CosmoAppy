package volunteer

// swagger:model Volunteer
type Volunteer struct {
	// Id of the volunteer
	// in: int64
	// read only: true
	Id int `json:"id"`
	// Firstname of the volunteer
	// in: string
	// required: true
	// example: Ash
	Firstname string `json:"firstname"`
	// Lastname of the volunteer
	// in: string
	// required: true
	// example: Ketchum
	Lastname string `json:"lastname"`
	// Email of the volunteer
	// in: string
	// required: true
	// example: ash.ketchum@indigo.kt
	Email string `json:"email"`
	// Discord pseudo
	// in: string
	// required: true
	// example: TheVeryBest
	Discord string `json:"discord"`
	// Phone number
	// in: string
	// required: true
	// example: 0123456789
	Phone string `json:"phone"`
	// University name
	// in: string
	// required: true
	// example: Indigo
	University string `json:"university"`
	// Postal address
	// in: string
	// required: true
	// example: 01 bis house, 00001 Pallet Town, Kanto
	PostalAddress string `json:"postal_address"`
	// Status of the volunteer
	// in: bool
	// required: true
	Actif bool `json:"actif"`
	// Does the volunteer has a bureau role
	// in: bool
	// required: true
	// example: false
	Bureau bool `json:"bureau"`
	// Status of the volunteer
	// in: string
	// required: true
	// example: volunteer
	HrStatus string `json:"hr_status"`
	// Started of date of volunteering
	// in: string
	// required: true
	// example: 1997-04-01
	StartedDate string `json:"started_date"`
}

type Volunteers []Volunteer

func NewVolunteer(vlt *Volunteer) {
	stmt, _ := TheDb().Prepare(
		`INSERT INTO ` + db_name + `
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
		TheLogger().LogError(log_name, "can't create new volunteer.", err)
	}
}

func FindVolunteerById(id int) *Volunteer {
	var vlt Volunteer

	row := TheDb().QueryRow("SELECT * FROM "+db_name+" WHERE id = ?;", id)
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
		TheLogger().LogWarning(log_name, "volunteer id not found.", err)
	}

	return &vlt
}

func AllVolunteers() *Volunteers {
	var vlts Volunteers

	rows, err := TheDb().Query("SELECT * FROM " + db_name)

	if err != nil {
		TheLogger().LogError(log_name, "problem with the db.", err)
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
			TheLogger().LogError(log_name, "volunteers not found.", err)
		}

		vlts = append(vlts, vlt)
	}

	return &vlts
}

func UpdateVolunteer(vlt *Volunteer) {
	stmt, err := TheDb().Prepare(
		`UPDATE ` + db_name + ` SET
			firstname=?, lastname=?, email=?, discord=?, phone=?, university=?, postal_address=?,
			actif=?, bureau=?, hr_status=?, started_date=?
		WHERE id=?;`,
	)

	if err != nil {
		TheLogger().LogError(log_name, "problem with the db.", err)
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
		TheLogger().LogError(log_name, "volunteer can't be updated.", err)
	}
}

func DeleteVolunteerById(id int) error {
	stmt, err := TheDb().Prepare("DELETE FROM " + db_name + " WHERE id=?;")

	if err != nil {
		TheLogger().LogError(log_name, "problem with the db.", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		TheLogger().LogError(log_name, "volunteer can't be deleted.", err)
	}

	return err
}
