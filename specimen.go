package main

import (
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/modules/adherent"
	"github.com/ESNFranceG33kTeam/sAPI/modules/esncard"
	"github.com/ESNFranceG33kTeam/sAPI/modules/event"
	"github.com/ESNFranceG33kTeam/sAPI/modules/money"
	"github.com/ESNFranceG33kTeam/sAPI/modules/volunteer"
)

func PopulateDb() {

	// Populate database
	// Adherents
	adherent.Specimen()
	// Volunteers
	volunteer.Specimen()
	// ESNcards
	esncard.Specimen()
	// Events
	event.Specimen()
	// Plannings
	//plannings.Specimen()

	// Moneys
	money.Specimen()
	// Stocks
	//stocks.Specimen()

	logger.GetLogger().LogInfo("Specimen", "Specimen data charged up.")
}
