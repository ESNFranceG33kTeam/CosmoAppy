package esncard

func CreateESNcardsTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS esncards (
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
		TheLogger().LogCritical("database", "create table esncards got a problem.", err)
	} else {
		TheLogger().LogInfo("database", "esncards table successfully created.")
	}
}
