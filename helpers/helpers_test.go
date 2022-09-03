package helpers

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setUp()
	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")
	//tearDown()
	os.Exit(exitVal)
}

func setUp() {
	InitFile("../conf.yml")
}

func TestInitFile(t *testing.T) {
	InitFile("../conf.yml")
	_, err := os.Stat(AppConfig.ConfPath)
	if err != nil {
		log.Fatalf("Conf path doesn't exist !")
	}
}

func TestReadconfig(t *testing.T) {
	ReadConfig()

	userdb := AppConfig.Userdb
	passdb := AppConfig.Passdb
	ipdb := AppConfig.Ipdb
	portdb := AppConfig.Portdb
	namedb := AppConfig.Namedb
	extradb := AppConfig.Extradb
	loglevel := AppConfig.Loglevel

	assert.Equal(t, "root", userdb)
	assert.Equal(t, "root", passdb)
	assert.Equal(t, "127.0.0.1", ipdb)
	assert.Equal(t, "3306", portdb)
	assert.Equal(t, "sapi", namedb)
	assert.Equal(t, "?parseTime=true", extradb)
	assert.Equal(t, 0, loglevel)
}

func TestTheAppConfig(t *testing.T) {
	var got interface{} = TheAppConfig()

	_, ok := got.(*Cfg)
	if ok == false {
		log.Fatalf("got is not what i want !")
	}
}
