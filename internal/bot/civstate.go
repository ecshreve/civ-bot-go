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
		PlayerMap: make(map[PlayerID]*Player),
		Civs:      civ.GenCivs(),
		CivMap:    civ.GenCivMap(civ.GenCivs()),
		Bans:      make(map[PlayerID][]*civ.Civ),
		Picks:     make(map[PlayerID][]*civ.Civ),
		PickState: NewPickState(&DefaultCivConfig),
	}
}

// NewCivStateWithConfig returns a CivState based on the given CivConfig.
func NewCivStateWithConfig(cfg *CivConfig) *CivState {
	cs := NewCivState()
	cs.PickState = NewPickState(cfg)
	return cs
}

// Reset clears the Bot's CivState conditionally maintains the CivConfig based on
// the keepConfig argument.
func (b *Bot) Reset(keepConfig bool) {
	cfg := b.CivConfig

	if keepConfig {
		b.CivState = NewCivStateWithConfig(cfg)
		return
	}

	b.CivConfig = NewCivConfig()
	b.CivState = NewCivState()
}

// ReadyToPick returns a boolean based on the Bot's CivState indicating if the Bot
// should proceed to picking Civs. True if all players have entered the number of bans
// defined in the CivConfig then, otherwise false.
func (b *Bot) ReadyToPick() bool {
	bans := b.CivState.Bans
	if len(bans) != len(b.CivState.Players) {
		return false
	}

	for _, bansForPlayer := range bans {
		if len(bansForPlayer) < b.CivConfig.Bans {
			return false
		}
	}

	return true
}
