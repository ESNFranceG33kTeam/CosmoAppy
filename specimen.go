package main

import (
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/modules/adherent"
	"github.com/ESNFranceG33kTeam/sAPI/modules/esncard"
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

	// Moneys
	money.Specimen()

	logger.GetLogger().LogInfo("Specimen", "Specimen data charged up.")
}
