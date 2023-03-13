package money

func CreateMoneysTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS moneys (
			id INT NOT NULL AUTO_INCREMENT,
			label VARCHAR(45) NOT NULL,
			price DECIMAL(6,2) NOT NULL,
			payment_type VARCHAR(45) NOT NULL,
			payment_date VARCHAR(45) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			PRIMARY KEY (id),
			UNIQUE INDEX id_UNIQUE (id ASC)
		);
	`)
	if err != nil {
		TheLogger().LogCritical("money", "create table moneys got a problem.", err)
	} else {
		TheLogger().LogInfo("money", "moneys table successfully created.")
	}

	_, err = TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS money_stats_monthly (
			id INT NOT NULL AUTO_INCREMENT,
			archive_date VARCHAR(45) NOT NULL,
			nb_per_label LONGTEXT NOT NULL,
			avg_prices LONGTEXT NOT NULL,
			min_prices LONGTEXT NOT NULL,
			max_prices LONGTEXT NOT NULL,
			nb_payments_type LONGTEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NULL DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE INDEX id_UNIQUE (id ASC)
		);
	`)
	if err != nil {
		TheLogger().LogCritical("money_stat_monthly", "create table money_stats_monthly got a problem.", err)
	} else {
		TheLogger().LogInfo("money_stat_monthly", "money_stats_monthly table successfully created.")
	}
}
