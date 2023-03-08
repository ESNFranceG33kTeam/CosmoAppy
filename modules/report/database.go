package report

func CreateReportsTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS reports (
			id INT NOT NULL AUTO_INCREMENT,
			type VARCHAR(45) NOT NULL,
			ref_ext INT NOT NULL,
			name VARCHAR(45) NOT NULL,
			date VARCHAR(45) NOT NULL,
			comment LONGTEXT NOT NULL,
			nb_reel_attendees INT NOT NULL,
			nb_subscribe_attendees INT NOT NULL,
			staffs_list LONGTEXT NOT NULL,
			nb_hours_prepa DECIMAL(6,2) NOT NULL,
			nb_hours DECIMAL(6,2) NOT NULL,
			nb_staffs_vlt DECIMAL(6,2) NOT NULL,
			nb_staffs_emp DECIMAL(6,2) NOT NULL,
			nb_staffs_scv DECIMAL(6,2) NOT NULL,
			taux_valorisation_vlt DECIMAL(6,2) NOT NULL,
			taux_valorisation_emp DECIMAL(6,2) NOT NULL,
			taux_valorisation_scv DECIMAL(6,2) NOT NULL,
			code_public VARCHAR(45) NOT NULL,
			code_project VARCHAR(45) NOT NULL,
			PRIMARY KEY (id),
			UNIQUE INDEX id_UNIQUE (id ASC)
		);
	`)
	if err != nil {
		TheLogger().LogCritical("report", "create table reports got a problem.", err)
	} else {
		TheLogger().LogInfo("report", "reports table successfully created.")
	}
}
