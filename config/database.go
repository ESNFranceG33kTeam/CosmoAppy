package config

import (
	"database/sql"
	"fmt"

	"github.com/ESNFranceG33kTeam/sAPI/logger"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DatabaseInit(userdb string, passdb string, ipdb string, portdb string, namedb string, extra string) {
	var err error

	infodb := fmt.Sprintf(userdb + ":" + passdb + "@tcp(" + ipdb + ":" + portdb + ")/" + namedb + extra)

	db, err = sql.Open("mysql", infodb)
	if err != nil {
		logger.LogCritical("database", "connection with data got a problem.", err)
	} else {
		logger.LogInfo("database", "Database is connected.")
	}

	// Create Table endpoint if not exists
	createAdherentTable()
}

func createAdherentTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS adherents(id serial, firstname varchar(20) NOT NULL, lastname varchar(20) NOT NULL, email varchar(20) NOT NULL, dateofbirth varchar(20) NOT NULL, esncard varchar(20), student bool, university varchar(20), homeland varchar(20) NOT NULL, speakabout varchar(20), newsletter bool, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")
	if err != nil {
		logger.LogCritical("database", "create table adherent got a problem.", err)
	} else {
		logger.LogInfo("database", "adherents table successfully created.")
	}
}

// Getter for db var
func Db() *sql.DB {
	return db
}
