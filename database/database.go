package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ESNFranceG33kTeam/CosmoAppy/helpers"
	"github.com/ESNFranceG33kTeam/CosmoAppy/logger"
	_ "github.com/go-sql-driver/mysql"
)

type ConfDb struct {
	Userdb  string
	Passdb  string
	Ipdb    string
	Portdb  string
	Namedb  string
	Extradb string
	Infodb  string
	SqlDb   *sql.DB
}

var myconf ConfDb

func DatabaseInit() {
	var err error
	var dbserver *sql.DB
	myconf.InitConfDb()

	// Connect to server and create database
	infodbserver := fmt.Sprintf(myconf.Userdb + ":" + myconf.Passdb + "@tcp(" + myconf.Ipdb + ":" + myconf.Portdb + ")/")
	dbserver, _ = sql.Open("mysql", infodbserver)
	initSchema(dbserver)
	dbserver.Close()

	// Connect to database
	myconf.SqlDb, err = sql.Open("mysql", myconf.Infodb)
	if err != nil {
		logger.GetLogger().LogCritical("database", "connection with database got a problem.", err)
	} else {
		logger.GetLogger().LogInfo("database", "Database is connected.")
	}
}

func initSchema(dbserver *sql.DB) {
	_, err := dbserver.Exec("CREATE DATABASE IF NOT EXISTS " + myconf.Namedb + ";")
	if err != nil {
		logger.GetLogger().LogCritical("database", "create database got a problem.", err)
	} else {
		logger.GetLogger().LogInfo("database", "database successfully init.")
	}
}

func (conf *ConfDb) InitConfDb() {
	conf.Userdb = helpers.TheAppConfig().Userdb
	conf.Passdb = os.Getenv(helpers.TheAppConfig().Passdb)
	conf.Ipdb = helpers.TheAppConfig().Ipdb
	conf.Portdb = helpers.TheAppConfig().Portdb
	conf.Namedb = helpers.TheAppConfig().Namedb
	conf.Extradb = helpers.TheAppConfig().Extradb
	conf.Infodb = fmt.Sprintf(conf.Userdb + ":" + conf.Passdb + "@tcp(" + conf.Ipdb + ":" + conf.Portdb + ")/" + conf.Namedb + conf.Extradb)
}

// Getter for db var
func Db() *sql.DB {
	err := myconf.SqlDb.Ping()
	if err != nil {
		logger.GetLogger().LogWarning("database", "reconnect to db.", err)
		myconf.SqlDb, err = sql.Open("mysql", myconf.Infodb)
		if err != nil {
			logger.GetLogger().LogCritical("database", "reconnection with database got a problem.", err)
		} else {
			logger.GetLogger().LogInfo("database", "database is reconnected.")
		}
	}

	return myconf.SqlDb
}
