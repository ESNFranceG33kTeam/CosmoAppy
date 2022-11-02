package launcher

import (
	"github.com/ESNFranceG33kTeam/sAPI/logger"
	"github.com/ESNFranceG33kTeam/sAPI/modules/adherent"
	"github.com/ESNFranceG33kTeam/sAPI/modules/esncard"
	"github.com/ESNFranceG33kTeam/sAPI/modules/event"
	"github.com/ESNFranceG33kTeam/sAPI/modules/health"
	"github.com/ESNFranceG33kTeam/sAPI/modules/money"
	"github.com/ESNFranceG33kTeam/sAPI/modules/planning"
	"github.com/ESNFranceG33kTeam/sAPI/modules/volunteer"
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
