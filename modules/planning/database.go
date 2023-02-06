package planning

func CreatePlanningsTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS plannings (
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
		TheLogger().LogCritical("database", "create table plannings got a problem.", err)
	} else {
		TheLogger().LogInfo("database", "plannings table successfully created.")
	}

	_, err = TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS planning_attendees (
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
				REFERENCES plannings (id)
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
		TheLogger().LogCritical("database", "create table planning_attendees got a problem.", err)
	} else {
		TheLogger().LogInfo("database", "planning_attendees table successfully created.")
	}
}
