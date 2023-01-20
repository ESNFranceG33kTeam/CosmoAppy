package volunteer

func CreateVolunteersTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS volunteers (
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
			started_date VARCHAR(45) NOT NULL,
			PRIMARY KEY (id),
			UNIQUE INDEX id_UNIQUE (id ASC)
		);
	`)
	if err != nil {
		TheLogger().LogCritical("database", "create table volunteers got a problem.", err)
	} else {
		TheLogger().LogInfo("database", "volunteers table successfully created.")
	}
}
