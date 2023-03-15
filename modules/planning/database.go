package planning

func CreatePlanningsTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS ` + db_name + ` (
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(45) NOT NULL,
			type VARCHAR(45) NOT NULL,
			location VARCHAR(45) NOT NULL,
			date_begins VARCHAR(45) NOT NULL,
			date_end VARCHAR(45) NOT NULL,
			hour_begins VARCHAR(45) NOT NULL,
			hour_end VARCHAR(45) NOT NULL,
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
		CREATE TABLE IF NOT EXISTS ` + db_name_att + ` (
			id INT NOT NULL AUTO_INCREMENT,
			id_planning INT NOT NULL,
			id_volunteer INT NOT NULL,
			job VARCHAR(45) NOT NULL,
			date VARCHAR(45) NOT NULL,
			hour_begins VARCHAR(45) NOT NULL,
			hour_end VARCHAR(45) NOT NULL,
			PRIMARY KEY (id),
			INDEX id_planning_idx (id_planning ASC),
			INDEX id_volunteer_idx (id_volunteer ASC),
			UNIQUE INDEX id_UNIQUE (id ASC),
			CONSTRAINT id_att_planning
				FOREIGN KEY (id_planning)
				REFERENCES ` + db_name + ` (id)
				ON DELETE CASCADE
				ON UPDATE NO ACTION,
			CONSTRAINT id_planning_att_volunteer
				FOREIGN KEY (id_volunteer)
				REFERENCES volunteers (id)
				ON DELETE CASCADE
				ON UPDATE NO ACTION
		);
	`)
	if err != nil {
		TheLogger().LogCritical(log_name_att, "create table "+db_name_att+" got a problem.", err)
	} else {
		TheLogger().LogInfo(log_name_att, db_name_att+" table successfully created.")
	}

	_, err = TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS ` + db_name_monthly_stat + ` (
			id INT NOT NULL AUTO_INCREMENT,
			archive_date VARCHAR(45) NOT NULL,
			nb_per_location LONGTEXT NOT NULL,
			nb_per_type LONGTEXT NOT NULL,
			nb_total INT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NULL DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE INDEX id_UNIQUE (id ASC)
		);
	`)
	if err != nil {
		TheLogger().LogCritical(db_name_monthly_stat, "create table "+db_name_monthly_stat+" got a problem.", err)
	} else {
		TheLogger().LogInfo(db_name_monthly_stat, db_name_monthly_stat+" table successfully created.")
	}
}
