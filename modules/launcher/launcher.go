package launcher

import (
	"github.com/ESNFranceG33kTeam/CosmoAppy/logger"
	"github.com/ESNFranceG33kTeam/CosmoAppy/modules/adherent"
	"github.com/ESNFranceG33kTeam/CosmoAppy/modules/esncard"
	"github.com/ESNFranceG33kTeam/CosmoAppy/modules/event"
	"github.com/ESNFranceG33kTeam/CosmoAppy/modules/health"
	"github.com/ESNFranceG33kTeam/CosmoAppy/modules/money"
	"github.com/ESNFranceG33kTeam/CosmoAppy/modules/planning"
	"github.com/ESNFranceG33kTeam/CosmoAppy/modules/volunteer"
)

// Modules to launch, the order can be important!
func LauncherModules() {
	// Adherent dependant
	adherent.Launcher()
	volunteer.Launcher()
	esncard.Launcher()
	event.Launcher()
	planning.Launcher()

	money.Launcher()
	//stock.Launcher()

	// Always last
	health.Launcher()
}

// Populate database, the order can be important!
func SpecimenModules() {
	// Adherent dependant
	adherent.Specimen()
	volunteer.Specimen()
	esncard.Specimen()
	event.Specimen()
	planning.Specimen()

	money.Specimen()
	//stock.Specimen()

	logger.GetLogger().LogInfo("Specimen", "Specimen data charged up.")
}
