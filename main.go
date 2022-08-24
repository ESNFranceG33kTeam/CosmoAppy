package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ESNFranceG33kTeam/sAPI/config"
	"github.com/ESNFranceG33kTeam/sAPI/models"
)

func main() {
	config.DatabaseInit()
	router := InitializeRouter()

	// Populate database
	birth_time, _ := time.Parse("2006-01-02", "1992-09-23")
	models.NewAdherent(&models.Adherent{Firstname: "Titi", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: birth_time, ESNcard: "grgerrbrbreht", Student: false})

	fmt.Println("Finish add to db")

	log.Fatal(http.ListenAndServe(":8080", router))

}
