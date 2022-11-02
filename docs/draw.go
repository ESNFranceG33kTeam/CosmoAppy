package docs

import (
	"github.com/ESNFranceG33kTeam/CosmoAppy/logger"
)

const draw = `
   _____
  / ____|                           /\
 | |     ___  ___ _ __ ___   ___   /  \   _ __  _ __  _   _
 | |    / _ \/ __| '_ ' _ \ / _ \ / /\ \ | '_ \| '_ \| | | |
 | |___| (_) \__ \ | | | | | (_) / ____ \| |_) | |_) | |_| |
  \_____\___/|___/_| |_| |_|\___/_/    \_\ .__/| .__/ \__, |
                                         | |   | |     __/ |
                                         |_|   |_|    |___/
by the French ESN G33k Team.
`

func DrawStart() {

	logger.GetLogger().LogDraw(draw)
}
