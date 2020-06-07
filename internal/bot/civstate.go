package bot

import (
	"time"

	"github.com/benbjohnson/clock"
	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/constants"
)

type PickState struct {
	PickTime         time.Time
	RePickVotes      int
	RePicksRemaining int
}

func NewPickState(config *CivConfig) *PickState {
	return &PickState{
		RePicksRemaining: config.RePicks,
	}
}

type CivState struct {
	Clk       clock.Clock
	Players   []*Player
	PlayerMap map[PlayerID]*Player
	Civs      []*civ.Civ
	CivMap    map[constants.CivKey]*civ.Civ
	Bans      map[PlayerID][]*civ.Civ
	Picks     map[PlayerID][]*civ.Civ
	*PickState
}

func NewCivState() *CivState {
	return &CivState{
		Clk:       clock.New(),
		Civs:      civ.GenCivs(),
		CivMap:    civ.GenCivMap(civ.GenCivs()),
		PickState: NewPickState(DefaultCivConfig),
	}
}
