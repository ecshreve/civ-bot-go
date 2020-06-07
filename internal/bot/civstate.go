package bot

import (
	"time"

	"github.com/benbjohnson/clock"

	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/constants"
)

// PickState state information specific to making picks for an instance of the Bot.
type PickState struct {
	PickTime         time.Time
	RePickVotes      int
	RePicksRemaining int
}

// NewPickState returns a new PickState based on the given CivConfig.
func NewPickState(config *CivConfig) *PickState {
	return &PickState{
		RePicksRemaining: config.RePicks,
	}
}

// CivState stores state information for an instance of the Bot.
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

// NewCivState returns a CivState based on the DefaultCivConfig.
func NewCivState() *CivState {
	return &CivState{
		Clk:       clock.New(),
		Civs:      civ.GenCivs(),
		CivMap:    civ.GenCivMap(civ.GenCivs()),
		PickState: NewPickState(DefaultCivConfig),
	}
}
