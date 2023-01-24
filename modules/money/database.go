package money

func CreateMoneysTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS moneys (
			id INT NOT NULL AUTO_INCREMENT,
			label VARCHAR(45) NOT NULL,
			price DECIMAL(6,2) NOT NULL,
			payment_date VARCHAR(45) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			PRIMARY KEY (id),
			UNIQUE INDEX id_UNIQUE (id ASC)
		);
	`)
	if err != nil {
		TheLogger().LogCritical("database", "create table moneys got a problem.", err)
	} else {
		TheLogger().LogInfo("database", "moneys table successfully created.")
	}
}
