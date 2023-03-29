package event

func CreateEventsTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS ` + db_name + ` (
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(45) NOT NULL,
			date VARCHAR(45) NOT NULL,
			location VARCHAR(45) NOT NULL,
			nb_spots_max INT NOT NULL,
			nb_spots_taken INT NOT NULL,
			type VARCHAR(45) NOT NULL,
			price DECIMAL NOT NULL,
			url_facebook VARCHAR(45) NULL DEFAULT NULL,
			actif TINYINT NOT NULL,
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
			id_event INT NOT NULL,
			id_adherent INT NOT NULL,
			PRIMARY KEY (id),
			INDEX idadherent_idx (id_adherent ASC),
			INDEX idevent_idx (id_event ASC),
			UNIQUE INDEX id_UNIQUE (id ASC),
			CONSTRAINT id_event_att_adherent
				FOREIGN KEY (id_adherent)
				REFERENCES adherents (id)
				ON DELETE CASCADE
				ON UPDATE NO ACTION,
			CONSTRAINT id_att_event
				FOREIGN KEY (id_event)
				REFERENCES ` + db_name + ` (id)
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
		CREATE TABLE IF NOT EXISTS ` + db_name_sta + ` (
			id INT NOT NULL AUTO_INCREMENT,
			id_event INT NOT NULL,
			id_volunteer INT NOT NULL,
			PRIMARY KEY (id),
			INDEX idvolunteer_idx (id_volunteer ASC),
			INDEX idevent_idx (id_event ASC),
			UNIQUE INDEX id_UNIQUE (id ASC),
			CONSTRAINT id_event_sta_volunteer
				FOREIGN KEY (id_volunteer)
				REFERENCES volunteers (id)
				ON DELETE CASCADE
				ON UPDATE NO ACTION,
			CONSTRAINT id_sta_event
				FOREIGN KEY (id_event)
				REFERENCES ` + db_name + ` (id)
				ON DELETE CASCADE
				ON UPDATE NO ACTION
		);
	`)
	if err != nil {
		TheLogger().LogCritical(log_name_sta, "create table "+db_name_sta+" got a problem.", err)
	} else {
		TheLogger().LogInfo(log_name_sta, db_name_sta+" table successfully created.")
	}

	_, err = TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS ` + db_name_monthly_stat + ` (
			id INT NOT NULL AUTO_INCREMENT,
			archive_date VARCHAR(45) NOT NULL,
			nb_per_location LONGTEXT NOT NULL,
			nb_per_type LONGTEXT NOT NULL,
			nb_per_price LONGTEXT NOT NULL,
			nb_cancel INT NOT NULL,
			nb_total INT NOT NULL,
			tx_avg_fulfill_per_type LONGTEXT NOT NULL,
			nb_avg_att_per_type LONGTEXT NOT NULL,
			nb_avg_sta_per_type LONGTEXT NOT NULL,
			nb_att_total INT NOT NULL,
			nb_sta_total INT NOT NULL,
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
