package report

func CreateReportsTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS reports (
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(45) NOT NULL,
			date VARCHAR(45) NOT NULL,
			comment LONGTEXT NOT NULL,
			nb_hours DECIMAL(6,2) NOT NULL,
			nb_staffs DECIMAL(6,2) NOT NULL,
			taux_valorisation DECIMAL(6,2) NOT NULL,
			PRIMARY KEY (id),
			UNIQUE INDEX id_UNIQUE (id ASC)
		);
	`)
	if err != nil {
		TheLogger().LogCritical("database", "create table reports got a problem.", err)
	} else {
		TheLogger().LogInfo("database", "reports table successfully created.")
	}
}
