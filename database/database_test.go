package database

import (
	"database/sql"
	"log"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/helpers"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
)

func TestDatabaseInit(t *testing.T) {
	helpers.InitFile()
	helpers.ReadConfig()
	logger.LogInit()
	DatabaseInit()
}

func TestDb(t *testing.T) {
	var got interface{} = Db()

	_, ok := got.(*sql.DB)

	if ok == false {
		log.Fatalln("got is not what i want !")
	}

}
