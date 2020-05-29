package civsession

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/civ"
)

// CivConfig holds a config for a CivSession.
type CivConfig struct {
	NumBans        int
	NumPicks       int
	NumRepicks     int
	UseFilthyTiers bool
}

// CivSession holds data for a single civ-bot session.
type CivSession struct {
	Config           *CivConfig
	Players          map[string]*discordgo.User
	Civs             []*civ.Civ
	Bans             map[string]*civ.Civ
	Picks            map[string][]*civ.Civ
	PickTime         time.Time
	RePickVotes      int
	RePicksRemaining int
}

// NewCivSession returns a new CivSession, note map fields are initialized to
// empty zero lengtrh maps.
func NewCivSession() *CivSession {
	cfg := &CivConfig{
		NumBans:        1,
		NumPicks:       3,
		NumRepicks:     3,
		UseFilthyTiers: false,
	}

	return &CivSession{
		Config:           cfg,
		Players:          make(map[string]*discordgo.User),
		Civs:             civ.GenCivs(),
		Bans:             make(map[string]*civ.Civ),
		Picks:            make(map[string][]*civ.Civ),
		RePicksRemaining: 3,
	}
}

// Reset clears the CivSession referenced by the pointer receiver to the func.
func (cs *CivSession) Reset() {
	cs.Players = make(map[string]*discordgo.User)
	cs.Civs = civ.GenCivs()
	cs.Bans = make(map[string]*civ.Civ)
	cs.Picks = make(map[string][]*civ.Civ)
	cs.PickTime = time.Time{}
	cs.RePickVotes = 0
	cs.RePicksRemaining = 3
}
