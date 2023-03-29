package report

func CreateReportsTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS ` + db_name + ` (
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
		TheLogger().LogCritical(log_name, "create table "+db_name+" got a problem.", err)
	} else {
		TheLogger().LogInfo(log_name, db_name+" table successfully created.")
	}

	_, err = TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS ` + db_name_monthly_stat + ` (
			id INT NOT NULL AUTO_INCREMENT,
			archive_date VARCHAR(45) NOT NULL,
			nb_per_codepublic LONGTEXT NOT NULL,
			nb_per_codeproject LONGTEXT NOT NULL,
			nb_per_type LONGTEXT NOT NULL,
			nb_total INT NOT NULL,
			hours_vlt_per_codes LONGTEXT NOT NULL,
			hours_emp_per_codes LONGTEXT NOT NULL,
			hours_scv_per_codes LONGTEXT NOT NULL,
			valo_vlt_per_codes LONGTEXT NOT NULL,
			valo_emp_per_codes LONGTEXT NOT NULL,
			valo_scv_per_codes LONGTEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NULL DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE INDEX id_UNIQUE (id ASC)
		);
	`)
	if err != nil {
		TheLogger().LogCritical(log_name_monthly_stat, "create table "+db_name_monthly_stat+" got a problem.", err)
	} else {
		TheLogger().LogInfo(log_name_monthly_stat, db_name_monthly_stat+" table successfully created.")
	}
}
