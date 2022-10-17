package volunteer

import (
	"github.com/ESNFranceG33kTeam/sAPI/database"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
)

func CreateVolunteerTable() {
	_, err := database.Db().Exec(`
		CREATE TABLE IF NOT EXISTS volunteers (
			id INT NOT NULL AUTO_INCREMENT,
			id_adherent INT NOT NULL,
			actif TINYINT NOT NULL,
			bureau TINYINT NOT NULL,
			PRIMARY KEY (id),
			INDEX idadherents_idx (id_adherent ASC),
			UNIQUE INDEX id_UNIQUE (id ASC),
			UNIQUE INDEX id_adherent_UNIQUE (id_adherent ASC),
			CONSTRAINT id_volunteer_adherent
				FOREIGN KEY (id_adherent)
				REFERENCES adherents (id)
				ON DELETE CASCADE
				ON UPDATE NO ACTION
		);
	`)
	if err != nil {
		logger.LogCritical("database", "create table volunteers got a problem.", err)
	} else {
		logger.LogInfo("database", "volunteers table successfully created.")
	}
}
