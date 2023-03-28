package adherent

func CreateAdherentsTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS ` + db_name + ` (
			id INT NOT NULL AUTO_INCREMENT,
			firstname VARCHAR(45) NOT NULL,
			lastname VARCHAR(45) NOT NULL,
			email VARCHAR(45) NOT NULL,
			dateofbirth VARCHAR(45) NOT NULL,
			situation VARCHAR(45) NOT NULL,
			university VARCHAR(45) NULL DEFAULT NULL,
			homeland VARCHAR(45) NOT NULL,
			speakabout VARCHAR(45) NULL DEFAULT NULL,
			newsletter TINYINT NOT NULL,
			adhesion_date VARCHAR(45) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NULL DEFAULT NULL,
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
			nb_per_country LONGTEXT NOT NULL,
			nb_per_university LONGTEXT NOT NULL,
			nb_per_situation LONGTEXT NOT NULL,
			nb_total INT NOT NULL,
			aboutus_per_type LONGTEXT NOT NULL,
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
