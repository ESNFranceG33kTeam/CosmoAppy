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
	var dbserver *sql.DB

	// Connect to server and create database
	infodbserver := fmt.Sprintf(userdb + ":" + passdb + "@tcp(" + ipdb + ":" + portdb + ")/")
	dbserver, _ = sql.Open("mysql", infodbserver)
	initSchema(dbserver, namedb)

	// Connect to database
	infodb := fmt.Sprintf(userdb + ":" + passdb + "@tcp(" + ipdb + ":" + portdb + ")/" + namedb + extra)
	db, err = sql.Open("mysql", infodb)
	if err != nil {
		logger.LogCritical("database", "connection with database got a problem.", err)
	} else {
		logger.LogInfo("database", "Database is connected.")
	}

	// Create Tables if not exists
	createAdherentTable()
	createESNcardTable()
}

func initSchema(dbserver *sql.DB, namedb string) {
	_, err := dbserver.Exec("CREATE DATABASE IF NOT EXISTS " + namedb + ";")
	if err != nil {
		logger.LogCritical("database", "create database got a problem.", err)
	} else {
		logger.LogInfo("database", "database successfully created.")
	}
}

func createAdherentTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS adherents (
			id INT NOT NULL AUTO_INCREMENT,
			firstname VARCHAR(45) NOT NULL,
			lastname VARCHAR(45) NOT NULL,
			email VARCHAR(45) NOT NULL,
			dateofbirth VARCHAR(45) NOT NULL,
			student TINYINT(1) NOT NULL,
			university VARCHAR(45) NULL DEFAULT NULL,
			homeland VARCHAR(45) NOT NULL,
			speakabout VARCHAR(45) NULL DEFAULT NULL,
			newsletter TINYINT(1) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NULL DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE INDEX id_UNIQUE (id ASC)
		);
	`)
	if err != nil {
		logger.LogCritical("database", "create table adherents got a problem.", err)
	} else {
		logger.LogInfo("database", "adherents table successfully created.")
	}
}

func createESNcardTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS esncards(
			id INT NOT NULL AUTO_INCREMENT,
			id_adherent INT NOT NULL,
			esncard VARCHAR(45) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			PRIMARY KEY (id),
			INDEX id_adherent_idx (id_adherent ASC),
			UNIQUE INDEX id_UNIQUE (id ASC),
			UNIQUE INDEX esncard_UNIQUE (esncard ASC),
			CONSTRAINT id_adherent
			FOREIGN KEY (id_adherent)
			REFERENCES adherents (id)
			ON DELETE CASCADE
			ON UPDATE NO ACTION
		);
	`)
	if err != nil {
		logger.LogCritical("database", "create table esncards got a problem.", err)
	} else {
		logger.LogInfo("database", "esncards table successfully created.")
	}
}

// Getter for db var
func Db() *sql.DB {
	return db
}
