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
	Confpathflag = "../../test/conf_local.yaml"
	InitFile()
}

func TestInitFile(t *testing.T) {
	InitFile()
	_, err := os.Stat(TheAppConfig().ConfPath)
	if err != nil {
		log.Fatalf("Conf path doesn't exist !")
	}
}

func TestReadconfig(t *testing.T) {
	ReadConfig()

	userdb := TheAppConfig().Userdb
	passdb := TheAppConfig().Passdb
	ipdb := TheAppConfig().Ipdb
	portdb := TheAppConfig().Portdb
	namedb := TheAppConfig().Namedb
	extradb := TheAppConfig().Extradb
	loglevel := TheAppConfig().Loglevel

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
