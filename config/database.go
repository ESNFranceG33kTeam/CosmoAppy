package config

import (
	"database/sql"
	"fmt"

	"github.com/ESNFranceG33kTeam/sAPI/logger"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	var userdb = "root"
	var passdb = "root"
	var ipdb = "127.0.0.1"
	var portbd = "3306"
	var namedb = "sapi"
	var extra = "?parseTime=true"
	infodb := fmt.Sprintf(userdb + ":" + passdb + "@tcp(" + ipdb + ":" + portbd + ")/" + namedb + extra)

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
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS adherents(id serial, firstname varchar(20), lastname varchar(20), email varchar(20), dateofbirth varchar(20), esncard varchar(20), student bool, university varchar(20), homeland varchar(20), speakabout varchar(20), newsletter bool, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")
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
