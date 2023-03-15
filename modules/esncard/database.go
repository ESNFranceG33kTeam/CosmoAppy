package esncard

func CreateESNcardsTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS ` + db_name + ` (
			id INT NOT NULL AUTO_INCREMENT,
			id_adherent INT NOT NULL,
			esncard VARCHAR(45) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			PRIMARY KEY (id),
			INDEX id_adherent_idx (id_adherent ASC),
			UNIQUE INDEX id_UNIQUE (id ASC),
			UNIQUE INDEX esncard_UNIQUE (esncard ASC),
			CONSTRAINT id_esncard_adherent
				FOREIGN KEY (id_adherent)
				REFERENCES adherents (id)
				ON DELETE CASCADE
				ON UPDATE NO ACTION
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
			nb_esncard_sold INT NOT NULL,
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
