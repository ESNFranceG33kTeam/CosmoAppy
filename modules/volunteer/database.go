package volunteer

func CreateVolunteersTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS ` + db_name + ` (
			id INT NOT NULL AUTO_INCREMENT,
			firstname VARCHAR(45) NOT NULL,
			lastname VARCHAR(45) NOT NULL,
			email VARCHAR(45) NOT NULL,
			discord VARCHAR(45) NOT NULL,
			phone VARCHAR(45) NOT NULL,
			university VARCHAR(45) NOT NULL,
			postal_address VARCHAR(95) NOT NULL,
			actif TINYINT NOT NULL,
			bureau TINYINT NOT NULL,
			hr_status VARCHAR(45) NOT NULL,
			started_date VARCHAR(45) NOT NULL,
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
			nb_active_per_university LONGTEXT NOT NULL,
			nb_active_per_hrstatus LONGTEXT NOT NULL,
			nb_new_vlt_this_month INT NOT NULL,
			nb_alumnus INT NOT NULL,
			nb_bureau INT NOT NULL,
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
