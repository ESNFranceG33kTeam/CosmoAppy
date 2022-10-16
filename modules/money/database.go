package money

import (
	"github.com/ESNFranceG33kTeam/sAPI/config"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
)

func CreateMoneysTable() {
	_, err := config.Db().Exec(`
		CREATE TABLE IF NOT EXISTS money (
			id INT NOT NULL AUTO_INCREMENT,
			label VARCHAR(45) NOT NULL,
			price FLOAT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			PRIMARY KEY (id),
			UNIQUE INDEX id_UNIQUE (id ASC)
		);
	`)
	if err != nil {
		logger.LogCritical("database", "create table moneys got a problem.", err)
	} else {
		logger.LogInfo("database", "moneys table successfully created.")
	}
}
