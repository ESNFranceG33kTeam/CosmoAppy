package controllers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ESNFranceG33kTeam/sAPI/config"
	"github.com/ESNFranceG33kTeam/sAPI/helpers"
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/models"
)

func TestMain(m *testing.M) {
	setUp()

	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")
	tearDown()
	os.Exit(exitVal)

}

func setUp() {
	helpers.InitFile("../test/conf_local.yaml")
	helpers.ReadConfig()
	logger.LogInit(helpers.AppConfig.Loglevel)
	config.DatabaseInit(helpers.AppConfig.Userdb, helpers.AppConfig.Passdb, helpers.AppConfig.Ipdb, helpers.AppConfig.Portdb, helpers.AppConfig.Namedb, helpers.AppConfig.Extradb)

	models.NewAdherent(&models.Adherent{Firstname: "Test1", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", ESNcard: "grgerrbrbreht", Student: false})
	models.NewAdherent(&models.Adherent{Firstname: "Test2", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", ESNcard: "grgerrbrbreht", Student: false})
	models.NewAdherent(&models.Adherent{Firstname: "Test3", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", ESNcard: "grgerrbrbreht", Student: false})
	models.NewAdherent(&models.Adherent{Firstname: "Test4", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", ESNcard: "grgerrbrbreht", Student: false})
}

func tearDown() {
	_, err := config.Db().Exec("DROP TABLE adherents;")
	if err != nil {
		log.Fatal(err)
	}
}

func TestAdherentsIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/adherents", nil)
	w := httptest.NewRecorder()
	AdherentsIndex(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

func TestAdherentsCreate(t *testing.T) {
	var jsonStr = []byte(`{"firstname": "Luigi", "lastname": "Bros", "dateofbirth": "24-04-1995"}`)

	req := httptest.NewRequest("POST", "/adherents", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	AdherentsCreate(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}

// func TestAdherentsShow(t *testing.T) {
// 	req, _ := http.NewRequest("Get", "/adherents/2", nil)

// 	w := httptest.NewRecorder()
// 	AdherentsShow(w, req)

// 	res := w.Result()
// 	defer res.Body.Close()
// 	_, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Errorf("expected error to be nil got %v", err)
// 	}
// }

// func TestAdherentsUpdate(t *testing.T) {

// }

// func TestAdherentsDelete(t *testing.T) {

// }
