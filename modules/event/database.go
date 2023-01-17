package event

func CreateEventsTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS events (
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
		TheLogger().LogCritical("database", "create table events got a problem.", err)
	} else {
		TheLogger().LogInfo("database", "events table successfully created.")
	}

	_, err = TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS event_attendees (
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
				REFERENCES events (id)
				ON DELETE CASCADE
				ON UPDATE NO ACTION
		);
	`)
	if err != nil {
		TheLogger().LogCritical("database", "create table event_attendees got a problem.", err)
	} else {
		TheLogger().LogInfo("database", "event_attendees table successfully created.")
	}

	_, err = TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS event_staffs (
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
				REFERENCES events (id)
				ON DELETE CASCADE
				ON UPDATE NO ACTION
		);
	`)
	if err != nil {
		TheLogger().LogCritical("database", "create table event_staffs got a problem.", err)
	} else {
		TheLogger().LogInfo("database", "event_staffs table successfully created.")
	}
}
