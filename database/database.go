package database

import (
	"database/sql"
	"fmt"

	"github.com/ESNFranceG33kTeam/CosmoAppy/helpers"
	"github.com/ESNFranceG33kTeam/CosmoAppy/logger"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DatabaseInit() {
	var err error
	var dbserver *sql.DB

	var userdb = helpers.TheAppConfig().Userdb
	var passdb = helpers.TheAppConfig().Passdb
	var ipdb = helpers.TheAppConfig().Ipdb
	var portdb = helpers.TheAppConfig().Portdb
	var namedb = helpers.TheAppConfig().Namedb
	var extradb = helpers.TheAppConfig().Extradb

	// Connect to server and create database
	infodbserver := fmt.Sprintf(userdb + ":" + passdb + "@tcp(" + ipdb + ":" + portdb + ")/")
	dbserver, _ = sql.Open("mysql", infodbserver)
	initSchema(dbserver, namedb)

	// Connect to database
	infodb := fmt.Sprintf(userdb + ":" + passdb + "@tcp(" + ipdb + ":" + portdb + ")/" + namedb + extradb)
	db, err = sql.Open("mysql", infodb)
	if err != nil {
		logger.GetLogger().LogCritical("database", "connection with database got a problem.", err)
	} else {
		logger.GetLogger().LogInfo("database", "Database is connected.")
	}
}

func initSchema(dbserver *sql.DB, namedb string) {
	_, err := dbserver.Exec("CREATE DATABASE IF NOT EXISTS " + namedb + ";")
	if err != nil {
		logger.GetLogger().LogCritical("database", "create database got a problem.", err)
	} else {
		logger.GetLogger().LogInfo("database", "database successfully created.")
	}
}

// Getter for db var
func Db() *sql.DB {
	return db
}
