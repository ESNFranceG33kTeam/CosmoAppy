package launcher

import (
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/modules/adherent"
	"github.com/ESNFranceG33kTeam/sAPI/modules/esncard"
	"github.com/ESNFranceG33kTeam/sAPI/modules/event"
	"github.com/ESNFranceG33kTeam/sAPI/modules/health"
	"github.com/ESNFranceG33kTeam/sAPI/modules/money"
	"github.com/ESNFranceG33kTeam/sAPI/modules/volunteer"
)

func LauncherModules() {
	// Adherent dependant
	adherent.Launcher()
	volunteer.Launcher()
	esncard.Launcher()
	event.Launcher()

	money.Launcher()

	// Always last
	health.Launcher()
}

func SpecimenModules() {
	// Populate database
	// Adherent dependant
	adherent.Specimen()
	volunteer.Specimen()
	esncard.Specimen()
	event.Specimen()
	//plannings.Specimen()

	// Moneys
	money.Specimen()
	//stocks.Specimen()

	logger.GetLogger().LogInfo("Specimen", "Specimen data charged up.")
}
