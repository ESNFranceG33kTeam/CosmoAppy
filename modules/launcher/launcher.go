package launcher

import (
	"github.com/ESNFranceG33kTeam/sAPI/modules/adherent"
	"github.com/ESNFranceG33kTeam/sAPI/modules/esncard"
	"github.com/ESNFranceG33kTeam/sAPI/modules/health"
	"github.com/ESNFranceG33kTeam/sAPI/modules/money"
	"github.com/ESNFranceG33kTeam/sAPI/modules/volunteer"
)

func LauncherModules() {
	// Adherent dependant
	adherent.Launcher()
	volunteer.Launcher()
	esncard.Launcher()

	money.Launcher()

	// Always last
	health.Launcher()
}
